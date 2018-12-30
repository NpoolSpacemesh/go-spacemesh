package sync

import (
	"fmt"
	"github.com/spacemeshos/go-spacemesh/mesh"
	"github.com/spacemeshos/go-spacemesh/p2p/service"
	"testing"
	"time"
)

func ListenerFactory(peers Peers, name string) *BlockListener {
	return NewBlockListener(peers, BlockValidatorMock{}, getMesh("TestBlockListener_"+name), 1*time.Second, 2)
}

func TestBlockListener(t *testing.T) {

	fmt.Println("test sync start")
	sim := service.NewSimulator()
	n1 := sim.NewNode()
	n2 := sim.NewNode()
	bl1 := ListenerFactory(PeersImpl{n1, func() []Peer { return []Peer{n2.PublicKey()} }}, "1")
	bl2 := ListenerFactory(PeersImpl{n1, func() []Peer { return []Peer{n1.PublicKey()} }}, "2")
	bl2.Start()

	block1 := mesh.NewExistingBlock(mesh.BlockID(123), 0, nil)
	block2 := mesh.NewExistingBlock(mesh.BlockID(321), 1, nil)
	block3 := mesh.NewExistingBlock(mesh.BlockID(222), 2, nil)

	block1.BlockVotes[block2.ID()] = true
	block1.BlockVotes[block3.ID()] = true

	bl1.AddBlock(block1)
	bl1.AddBlock(block2)
	bl1.AddBlock(block3)

	bl2.FetchBlock(block1.Id)
	timeout := time.After(30 * time.Second)
loop:
	for {
		select {
		// Got a timeout! fail with a timeout error
		case <-timeout:
			t.Error("timed out ")
		default:
			if b, err := bl2.GetBlock(block1.Id); err == nil {
				fmt.Println("  ", b)
				t.Log("done!")
				break loop
			}
		}
	}
}

func TestBlockListener2(t *testing.T) {

	fmt.Println("test sync start")
	sim := service.NewSimulator()
	n1 := sim.NewNode()
	n2 := sim.NewNode()

	bl1 := ListenerFactory(PeersImpl{n1, func() []Peer { return []Peer{n2.PublicKey()} }}, "3")
	bl2 := ListenerFactory(PeersImpl{n2, func() []Peer { return []Peer{n1.PublicKey()} }}, "4")

	bl2.Start()

	block1 := mesh.NewBlock(true, nil, time.Now(), 0)
	block2 := mesh.NewBlock(true, nil, time.Now(), 1)
	block3 := mesh.NewBlock(true, nil, time.Now(), 2)
	block4 := mesh.NewBlock(true, nil, time.Now(), 2)
	block5 := mesh.NewBlock(true, nil, time.Now(), 3)
	block6 := mesh.NewBlock(true, nil, time.Now(), 3)
	block7 := mesh.NewBlock(true, nil, time.Now(), 4)
	block8 := mesh.NewBlock(true, nil, time.Now(), 4)
	block9 := mesh.NewBlock(true, nil, time.Now(), 4)
	block10 := mesh.NewBlock(true, nil, time.Now(), 5)

	block2.BlockVotes[block1.ID()] = true
	block3.BlockVotes[block2.ID()] = true
	block4.BlockVotes[block2.ID()] = true
	block5.BlockVotes[block3.ID()] = true
	block5.BlockVotes[block4.ID()] = true
	block6.BlockVotes[block4.ID()] = true
	block7.BlockVotes[block6.ID()] = true
	block7.BlockVotes[block5.ID()] = true
	block8.BlockVotes[block6.ID()] = true
	block9.BlockVotes[block5.ID()] = true
	block10.BlockVotes[block8.ID()] = true
	block10.BlockVotes[block9.ID()] = true

	bl1.AddBlock(block1)
	bl1.AddBlock(block2)
	bl1.AddBlock(block3)
	bl1.AddBlock(block4)
	bl1.AddBlock(block5)
	bl1.AddBlock(block6)
	bl1.AddBlock(block7)
	bl1.AddBlock(block8)
	bl1.AddBlock(block9)
	bl1.AddBlock(block10)

	bl2.FetchBlock(block10.Id)

	timeout := time.After(10 * time.Second)
loop:
	for {
		select {
		// Got a timeout! fail with a timeout error
		case <-timeout:
			t.Error("timed out ")
		default:
			if b, err := bl2.GetBlock(block1.Id); err == nil {
				fmt.Println("  ", b)
				t.Log("done!")
				break loop
			}
		}
	}
}

//todo integration testing
