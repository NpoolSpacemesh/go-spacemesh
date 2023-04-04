package hare

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/libp2p/go-libp2p/core/peer"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/hare/config"
	"github.com/spacemeshos/go-spacemesh/hare/mocks"
	"github.com/spacemeshos/go-spacemesh/log/logtest"
	"github.com/spacemeshos/go-spacemesh/p2p/pubsub"
)

// Test - run multiple CPs simultaneously.
func Test_multipleCPs(t *testing.T) {
	logtest.SetupGlobal(t)

	r := require.New(t)
	totalCp := uint32(3)
	finalLyr := types.GetEffectiveGenesis().Add(totalCp)
	test := newHareWrapper(totalCp)
	totalNodes := 10
	// RoundDuration is not used because we override the newRoundClock
	// function, wakeupDelta controls whether a consensus process will skip a
	// layer, if the layer tick arrives after wakeup delta then the process
	// skips the layer.
	cfg := config.Config{N: totalNodes, WakeupDelta: 5 * time.Second, RoundDuration: 0, ExpectedLeaders: 5, LimitIterations: 1000, LimitConcurrent: 100, Hdist: 20}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mesh, err := mocknet.FullMeshLinked(totalNodes)
	require.NoError(t, err)

	test.initialSets = make([]*Set, totalNodes)

	pList := make(map[types.LayerID][]*types.Proposal)
	for j := types.GetEffectiveGenesis().Add(1); !j.After(finalLyr); j = j.Add(1) {
		for i := uint64(0); i < 20; i++ {
			p := genLayerProposal(j, []types.TransactionID{})
			p.EpochData = &types.EpochData{
				Beacon: types.EmptyBeacon,
			}
			pList[j] = append(pList[j], p)
		}
	}
	meshes := make([]*mocks.Mockmesh, 0, totalNodes)
	ctrl, ctx := gomock.WithContext(ctx, t)
	for i := 0; i < totalNodes; i++ {
		mockMesh := mocks.NewMockmesh(ctrl)
		mockMesh.EXPECT().GetEpochAtx(gomock.Any(), gomock.Any()).Return(&types.ActivationTxHeader{BaseTickHeight: 11, TickCount: 1}, nil).AnyTimes()
		mockMesh.EXPECT().VRFNonce(gomock.Any(), gomock.Any()).Return(types.VRFPostIndex(0), nil).AnyTimes()
		mockMesh.EXPECT().GetMalfeasanceProof(gomock.Any()).AnyTimes()
		mockMesh.EXPECT().SetWeakCoin(gomock.Any(), gomock.Any()).AnyTimes()
		for lid := types.GetEffectiveGenesis().Add(1); !lid.After(finalLyr); lid = lid.Add(1) {
			mockMesh.EXPECT().Proposals(lid).Return(pList[lid], nil)
			for _, p := range pList[lid] {
				mockMesh.EXPECT().GetAtxHeader(p.AtxID).Return(&types.ActivationTxHeader{BaseTickHeight: 11, TickCount: 1}, nil).AnyTimes()
				mockMesh.EXPECT().Ballot(p.Ballot.ID()).Return(&p.Ballot, nil).AnyTimes()
			}
		}
		mockMesh.EXPECT().Proposals(gomock.Any()).Return([]*types.Proposal{}, nil).AnyTimes()
		meshes = append(meshes, mockMesh)
	}

	// setup roundClocks to progress a layer only when all nodes have received messages from all nodes.
	roundClocks := newSharedRoundClocks(totalNodes*totalNodes, 200*time.Millisecond)
	var pubsubs []*pubsub.PubSub
	outputs := make([]map[types.LayerID]LayerOutput, totalNodes)
	var outputsWaitGroup sync.WaitGroup
	for i := 0; i < totalNodes; i++ {
		host := mesh.Hosts()[i]
		ps, err := pubsub.New(ctx, logtest.New(t), host, pubsub.DefaultConfig())
		require.NoError(t, err)
		pubsubs = append(pubsubs, ps)
		// We wrap the pubsub system to notify round clocks whenever a message
		// is received.
		testPs := &testPublisherSubscriber{
			inner: ps,
			register: func(protocol string, handler pubsub.GossipHandler) {
				ps.Register(protocol, func(ctx context.Context, peer peer.ID, message []byte) pubsub.ValidationResult {
					res := handler(ctx, peer, message)

					m, err := MessageFromBuffer(message)
					if err != nil {
						panic(err)
					}
					roundClocks.clock(m.Layer).incMessages(int(m.Eligibility.Count), m.Round)
					return res
				})
			},
		}
		h := createTestHare(t, meshes[i], cfg, test.clock, testPs, t.Name())
		// override the round clocks method to use our shared round clocks
		h.newRoundClock = roundClocks.roundClock
		h.mockRoracle.EXPECT().IsIdentityActiveOnConsensusView(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
		h.mockRoracle.EXPECT().Proof(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(make([]byte, 80), nil).AnyTimes()
		h.mockRoracle.EXPECT().CalcEligibility(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(uint16(1), nil).AnyTimes()
		h.mockRoracle.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
		outputsWaitGroup.Add(1)
		go func(idx int) {
			defer outputsWaitGroup.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case out, ok := <-h.blockGenCh:
					if !ok {
						return
					}
					if outputs[idx] == nil {
						outputs[idx] = make(map[types.LayerID]LayerOutput)
					}
					outputs[idx][out.Layer] = out
				}
			}
		}(i)
		test.hare = append(test.hare, h.Hare)
		e := h.Start(ctx)
		r.NoError(e)
	}
	require.NoError(t, mesh.ConnectAllButSelf())
	require.Eventually(t, func() bool {
		for _, ps := range pubsubs {
			if len(ps.ProtocolPeers(pubsub.HareProtocol)) != len(mesh.Hosts())-1 {
				return false
			}
		}
		return true
	}, 5*time.Second, 10*time.Millisecond)
	log := logtest.New(t)
	for j := types.GetEffectiveGenesis().Add(1); !j.After(finalLyr); j = j.Add(1) {
		test.clock.advanceLayer()
		time.Sleep(10 * time.Millisecond)
		log.Warning("advancing to layer %d", j.Uint32())
	}

	// There are 5 rounds per layer and totalCPs layers and we double for good measure.
	test.WaitForTimedTermination(t, time.Minute)
	for _, h := range test.hare {
		close(h.blockGenCh)
	}
	outputsWaitGroup.Wait()
	for _, out := range outputs {
		for lid := types.GetEffectiveGenesis().Add(1); !lid.After(finalLyr); lid = lid.Add(1) {
			require.NotNil(t, out[lid])
			require.ElementsMatch(t, types.ToProposalIDs(pList[lid]), out[lid].Proposals)
		}
	}
	t.Cleanup(func() {
		for _, h := range test.hare {
			h.Close()
		}
	})
}

// Test - run multiple CPs where one of them runs more than one iteration.
func Test_multipleCPsAndIterations(t *testing.T) {
	logtest.SetupGlobal(t)

	r := require.New(t)
	totalCp := uint32(1)
	finalLyr := types.GetEffectiveGenesis().Add(totalCp)
	test := newHareWrapper(totalCp)
	totalNodes := 10
	// RoundDuration is not used because we override the newRoundClock
	// function, wakeupDelta controls whether a consensus process will skip a
	// layer, if the layer tick arrives after wakeup delta then the process
	// skips the layer.
	cfg := config.Config{N: totalNodes, WakeupDelta: 5 * time.Second, RoundDuration: 0, ExpectedLeaders: 5, LimitIterations: 1000, LimitConcurrent: 100, Hdist: 20}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mesh, err := mocknet.FullMeshLinked(totalNodes)
	require.NoError(t, err)

	test.initialSets = make([]*Set, totalNodes)

	pList := make(map[types.LayerID][]*types.Proposal)
	for j := types.GetEffectiveGenesis().Add(1); !j.After(finalLyr); j = j.Add(1) {
		for i := uint64(0); i < 20; i++ {
			p := genLayerProposal(j, []types.TransactionID{})
			p.EpochData = &types.EpochData{
				Beacon: types.EmptyBeacon,
			}
			pList[j] = append(pList[j], p)
		}
	}

	meshes := make([]*mocks.Mockmesh, 0, totalNodes)
	ctrl, ctx := gomock.WithContext(ctx, t)
	for i := 0; i < totalNodes; i++ {
		mockMesh := mocks.NewMockmesh(ctrl)
		mockMesh.EXPECT().GetEpochAtx(gomock.Any(), gomock.Any()).Return(&types.ActivationTxHeader{BaseTickHeight: 11, TickCount: 1}, nil).AnyTimes()
		mockMesh.EXPECT().VRFNonce(gomock.Any(), gomock.Any()).Return(types.VRFPostIndex(0), nil).AnyTimes()
		mockMesh.EXPECT().GetMalfeasanceProof(gomock.Any()).AnyTimes()
		mockMesh.EXPECT().SetWeakCoin(gomock.Any(), gomock.Any()).AnyTimes()
		for lid := types.GetEffectiveGenesis().Add(1); !lid.After(finalLyr); lid = lid.Add(1) {
			mockMesh.EXPECT().Proposals(lid).Return(pList[lid], nil)
			for _, p := range pList[lid] {
				mockMesh.EXPECT().GetAtxHeader(p.AtxID).Return(&types.ActivationTxHeader{BaseTickHeight: 11, TickCount: 1}, nil).AnyTimes()
				mockMesh.EXPECT().Ballot(p.Ballot.ID()).Return(&p.Ballot, nil).AnyTimes()
			}
		}
		mockMesh.EXPECT().Proposals(gomock.Any()).Return([]*types.Proposal{}, nil).AnyTimes()
		meshes = append(meshes, mockMesh)
	}

	// setup roundClocks to progress a layer only when all nodes have received messages from all nodes.
	roundClocks := newSharedRoundClocks(totalNodes*totalNodes, 200*time.Millisecond)

	// The stalled layer drops messages from rounds 0-7 so to ensure that those
	// rounds progress without requiring notification via
	// testPublisherSubscriber we "pre close" the rounds by providing a closed
	// channel.
	// roundClocks.roundClock(stalledLayer)
	// stalledLayerRoundClock := roundClocks.clock(stalledLayer)
	// closedChan := make(chan struct{})
	// close(closedChan)
	// So this is tricky, ideally this would work but then layer 8 when advance round is called it tries to close the channel for layer 7 and that is already closed and we panic, so what to do?
	// Ok I could swithc the shared round clock to send values on channels but that is a bit tricky because I then have to have the right number of values for all the calls from all the cp instances.
	// A better approach could be to use my original approach and simply call AwaitEndOfRound before calling advacne round to make sure the value is inited. Then I don't need to sleep.
	// for r := 0; r < 8; r++ {
	// 	stalledLayerRoundClock.rounds[uint32(r)] = closedChan
	// }
	var pubsubs []*pubsub.PubSub
	outputs := make([]map[types.LayerID]LayerOutput, totalNodes)
	var outputsWaitGroup sync.WaitGroup
	var layer8Count int
	var notifyMutex sync.Mutex
	var targetLayerWg sync.WaitGroup
	targetLayerWg.Add(totalNodes)
	stalledLayer := types.GetEffectiveGenesis().Add(1)
	maxStalledRound := 7
	for i := 0; i < totalNodes; i++ {
		host := mesh.Hosts()[i]
		ps, err := pubsub.New(ctx, logtest.New(t), host, pubsub.DefaultConfig())
		require.NoError(t, err)

		testPs := &testPublisherSubscriber{
			inner: ps,
			publish: func(ctx context.Context, topic string, message []byte) error {
				// drop messages from rounds 0 - 7 for stalled layer but not for preRound.
				msg, _ := MessageFromBuffer(message)
				if msg.Layer == stalledLayer && int(msg.Round) <= maxStalledRound && msg.Round != preRound {
					return errors.New("fake error")
				}
				// We need to wait before publishing for the unstalled rounds
				// to make sure all nodes have reached the unstalled round
				// otherwise nodes that are still working through the stalled
				// rounds may interpret the message as contextually invalid.
				if msg.Layer == stalledLayer && int(msg.Round) == maxStalledRound+1 {
					targetLayerWg.Done()
					targetLayerWg.Wait()
				}
				return ps.Publish(ctx, topic, message)
			},
			register: func(topic string, handler pubsub.GossipHandler) {
				// Register a wrapped handler with the real pubsub
				ps.Register(topic, func(ctx context.Context, peer peer.ID, message []byte) pubsub.ValidationResult {
					// Call the handler
					res := handler(ctx, peer, message)

					// Now we will ensure the round clock is progressed properly.
					notifyMutex.Lock()
					defer notifyMutex.Unlock()
					m, err := MessageFromBuffer(message)
					if err != nil {
						panic(err)
					}

					// Incremente message count.
					roundClocks.clock(m.Layer).incMessages(int(m.Eligibility.Count), m.Round)

					// The stalled layer blocks messages from being published,
					// so after the preRound we need a different mechanism to
					// progress the rounds, because we will not be receiving
					// messages and therefore we will not be making calls to
					// incMessages.
					if m.Layer == stalledLayer && m.Round == preRound {
						// Keep track of received preRound messages
						layer8Count += int(m.Eligibility.Count)
						if layer8Count == totalNodes*totalNodes {
							// Once all preRound messages have been received wait for the preRound to complete.
							<-roundClocks.clock(m.Layer).AwaitEndOfRound(preRound)
							println("status round completed")
							// time.Sleep(time.Second)
							// Manually progress rounds to the round following the maxStalledRound.
							roundClocks.clock(m.Layer).advanceToRound(uint32(maxStalledRound + 1))
							// roundClocks.clock(m.Layer).m.Lock()
							// defer roundClocks.clock(m.Layer).m.Unlock()
							// for r := 0; r < 8; r++ {
							// 	println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaadd")
							// 	// roundClocks.clock(m.Layer).AwaitEndOfRound(uint32(r))
							// 	time.Sleep(time.Millisecond * 200)
							// 	roundClocks.clock(m.Layer).advanceRound()
							// }
						}
					}
					return res
				})
			},

			// notify: func(msg []byte) {
			// 	layer, eligibilityCount := extractInstanceID(msg)
			// 	roundClocks.clock(layer).IncMessages(int(eligibilityCount))
			// },
			// notify: func(msg []byte) {
			// 	notifyMutex.Lock()
			// 	defer notifyMutex.Unlock()
			// 	m, err := MessageFromBuffer(msg)
			// 	if err != nil {
			// 		panic(err)
			// 	}

			// 	roundClocks.clock(m.Layer).incMessages(int(m.Eligibility.Count), m.Round)
			// 	// Keep track of progress in the stalled layer
			// 	if m.Layer == stalledLayer && m.Round == preRound {
			// 		// println("layer8adding", m.Eligibility.Count)
			// 		layer8Count += int(m.Eligibility.Count)
			// 		// println("layer8count", layer8Count)
			// 		if layer8Count == totalNodes*totalNodes {
			// 			// Gotta sleep to make sure that advance round is called first by the round clock
			// 			// time.Sleep(4 * roundClocks.processingDelay)
			// 			<-roundClocks.clock(m.Layer).AwaitEndOfRound(preRound)
			// 			println("status round completed")
			// 			// time.Sleep(time.Second)
			// 			roundClocks.clock(m.Layer).advanceToRound(8)
			// 			// roundClocks.clock(m.Layer).m.Lock()
			// 			// defer roundClocks.clock(m.Layer).m.Unlock()
			// 			// for r := 0; r < 8; r++ {
			// 			// 	println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaadd")
			// 			// 	// roundClocks.clock(m.Layer).AwaitEndOfRound(uint32(r))
			// 			// 	time.Sleep(time.Millisecond * 200)
			// 			// 	roundClocks.clock(m.Layer).advanceRound()
			// 			// }
			// 		}
			// 	}
			// },
		}
		// the manipulatior drops meessages for the given layer on the first iteration this will result in a second iteration s
		mp2p := &p2pManipulator{PublishSubsciber: testPs, stalledLayer: stalledLayer, err: errors.New("fake err")}
		h := createTestHare(t, meshes[i], cfg, test.clock, mp2p, t.Name())
		// override the round clocks method to use our shared round clocks
		h.newRoundClock = roundClocks.roundClock
		h.mockRoracle.EXPECT().IsIdentityActiveOnConsensusView(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
		h.mockRoracle.EXPECT().Proof(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(make([]byte, 80), nil).AnyTimes()
		h.mockRoracle.EXPECT().CalcEligibility(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(uint16(1), nil).AnyTimes()
		h.mockRoracle.EXPECT().Validate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
		outputsWaitGroup.Add(1)
		go func(idx int) {
			defer outputsWaitGroup.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case out, ok := <-h.blockGenCh:
					if !ok {
						return
					}
					if outputs[idx] == nil {
						outputs[idx] = make(map[types.LayerID]LayerOutput)
					}
					outputs[idx][out.Layer] = out
				}
			}
		}(i)
		test.hare = append(test.hare, h.Hare)
		e := h.Start(ctx)
		r.NoError(e)
	}
	require.NoError(t, mesh.ConnectAllButSelf())
	require.Eventually(t, func() bool {
		for _, ps := range pubsubs {
			if len(ps.ProtocolPeers(pubsub.HareProtocol)) != len(mesh.Hosts())-1 {
				return false
			}
		}
		return true
	}, 5*time.Second, 10*time.Millisecond)
	layerDuration := 250 * time.Millisecond
	log := logtest.New(t)
	go func() {
		for j := types.GetEffectiveGenesis().Add(1); !j.After(finalLyr); j = j.Add(1) {
			test.clock.advanceLayer()
			log.Warning("advancing to layer %d", j.Uint32())
			time.Sleep(layerDuration)
		}
	}()

	// There are 5 rounds per layer and totalCPs layers and we double to allow
	// for the for good measure. Also one layer in this test will run 2
	// iterations so we increase the layer count by 1.
	test.WaitForTimedTermination(t, time.Minute)
	for _, h := range test.hare {
		close(h.blockGenCh)
	}
	outputsWaitGroup.Wait()
	for _, out := range outputs {
		for lid := types.GetEffectiveGenesis().Add(1); !lid.After(finalLyr); lid = lid.Add(1) {
			require.NotNil(t, out[lid])
			require.ElementsMatch(t, types.ToProposalIDs(pList[lid]), out[lid].Proposals)
		}
	}
	t.Cleanup(func() {
		for _, h := range test.hare {
			h.Close()
		}
	})
}
