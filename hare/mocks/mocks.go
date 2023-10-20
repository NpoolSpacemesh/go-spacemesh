// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go
//
// Generated by this command:
//
//	mockgen -typed -package=mocks -destination=./mocks/mocks.go -source=./interfaces.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/spacemeshos/go-spacemesh/common/types"
	datastore "github.com/spacemeshos/go-spacemesh/datastore"
	gomock "go.uber.org/mock/gomock"
)

// MocklayerPatrol is a mock of layerPatrol interface.
type MocklayerPatrol struct {
	ctrl     *gomock.Controller
	recorder *MocklayerPatrolMockRecorder
}

// MocklayerPatrolMockRecorder is the mock recorder for MocklayerPatrol.
type MocklayerPatrolMockRecorder struct {
	mock *MocklayerPatrol
}

// NewMocklayerPatrol creates a new mock instance.
func NewMocklayerPatrol(ctrl *gomock.Controller) *MocklayerPatrol {
	mock := &MocklayerPatrol{ctrl: ctrl}
	mock.recorder = &MocklayerPatrolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocklayerPatrol) EXPECT() *MocklayerPatrolMockRecorder {
	return m.recorder
}

// SetHareInCharge mocks base method.
func (m *MocklayerPatrol) SetHareInCharge(arg0 types.LayerID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHareInCharge", arg0)
}

// SetHareInCharge indicates an expected call of SetHareInCharge.
func (mr *MocklayerPatrolMockRecorder) SetHareInCharge(arg0 any) *layerPatrolSetHareInChargeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHareInCharge", reflect.TypeOf((*MocklayerPatrol)(nil).SetHareInCharge), arg0)
	return &layerPatrolSetHareInChargeCall{Call: call}
}

// layerPatrolSetHareInChargeCall wrap *gomock.Call
type layerPatrolSetHareInChargeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *layerPatrolSetHareInChargeCall) Return() *layerPatrolSetHareInChargeCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *layerPatrolSetHareInChargeCall) Do(f func(types.LayerID)) *layerPatrolSetHareInChargeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *layerPatrolSetHareInChargeCall) DoAndReturn(f func(types.LayerID)) *layerPatrolSetHareInChargeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockRolacle is a mock of Rolacle interface.
type MockRolacle struct {
	ctrl     *gomock.Controller
	recorder *MockRolacleMockRecorder
}

// MockRolacleMockRecorder is the mock recorder for MockRolacle.
type MockRolacleMockRecorder struct {
	mock *MockRolacle
}

// NewMockRolacle creates a new mock instance.
func NewMockRolacle(ctrl *gomock.Controller) *MockRolacle {
	mock := &MockRolacle{ctrl: ctrl}
	mock.recorder = &MockRolacleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRolacle) EXPECT() *MockRolacleMockRecorder {
	return m.recorder
}

// CalcEligibility mocks base method.
func (m *MockRolacle) CalcEligibility(arg0 context.Context, arg1 types.LayerID, arg2 uint32, arg3 int, arg4 types.NodeID, arg5 types.VrfSignature) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcEligibility", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalcEligibility indicates an expected call of CalcEligibility.
func (mr *MockRolacleMockRecorder) CalcEligibility(arg0, arg1, arg2, arg3, arg4, arg5 any) *RolacleCalcEligibilityCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcEligibility", reflect.TypeOf((*MockRolacle)(nil).CalcEligibility), arg0, arg1, arg2, arg3, arg4, arg5)
	return &RolacleCalcEligibilityCall{Call: call}
}

// RolacleCalcEligibilityCall wrap *gomock.Call
type RolacleCalcEligibilityCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RolacleCalcEligibilityCall) Return(arg0 uint16, arg1 error) *RolacleCalcEligibilityCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RolacleCalcEligibilityCall) Do(f func(context.Context, types.LayerID, uint32, int, types.NodeID, types.VrfSignature) (uint16, error)) *RolacleCalcEligibilityCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RolacleCalcEligibilityCall) DoAndReturn(f func(context.Context, types.LayerID, uint32, int, types.NodeID, types.VrfSignature) (uint16, error)) *RolacleCalcEligibilityCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsIdentityActiveOnConsensusView mocks base method.
func (m *MockRolacle) IsIdentityActiveOnConsensusView(arg0 context.Context, arg1 types.NodeID, arg2 types.LayerID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIdentityActiveOnConsensusView", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsIdentityActiveOnConsensusView indicates an expected call of IsIdentityActiveOnConsensusView.
func (mr *MockRolacleMockRecorder) IsIdentityActiveOnConsensusView(arg0, arg1, arg2 any) *RolacleIsIdentityActiveOnConsensusViewCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIdentityActiveOnConsensusView", reflect.TypeOf((*MockRolacle)(nil).IsIdentityActiveOnConsensusView), arg0, arg1, arg2)
	return &RolacleIsIdentityActiveOnConsensusViewCall{Call: call}
}

// RolacleIsIdentityActiveOnConsensusViewCall wrap *gomock.Call
type RolacleIsIdentityActiveOnConsensusViewCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RolacleIsIdentityActiveOnConsensusViewCall) Return(arg0 bool, arg1 error) *RolacleIsIdentityActiveOnConsensusViewCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RolacleIsIdentityActiveOnConsensusViewCall) Do(f func(context.Context, types.NodeID, types.LayerID) (bool, error)) *RolacleIsIdentityActiveOnConsensusViewCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RolacleIsIdentityActiveOnConsensusViewCall) DoAndReturn(f func(context.Context, types.NodeID, types.LayerID) (bool, error)) *RolacleIsIdentityActiveOnConsensusViewCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Proof mocks base method.
func (m *MockRolacle) Proof(arg0 context.Context, arg1 types.LayerID, arg2 uint32) (types.VrfSignature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proof", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.VrfSignature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Proof indicates an expected call of Proof.
func (mr *MockRolacleMockRecorder) Proof(arg0, arg1, arg2 any) *RolacleProofCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proof", reflect.TypeOf((*MockRolacle)(nil).Proof), arg0, arg1, arg2)
	return &RolacleProofCall{Call: call}
}

// RolacleProofCall wrap *gomock.Call
type RolacleProofCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RolacleProofCall) Return(arg0 types.VrfSignature, arg1 error) *RolacleProofCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RolacleProofCall) Do(f func(context.Context, types.LayerID, uint32) (types.VrfSignature, error)) *RolacleProofCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RolacleProofCall) DoAndReturn(f func(context.Context, types.LayerID, uint32) (types.VrfSignature, error)) *RolacleProofCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Validate mocks base method.
func (m *MockRolacle) Validate(arg0 context.Context, arg1 types.LayerID, arg2 uint32, arg3 int, arg4 types.NodeID, arg5 types.VrfSignature, arg6 uint16) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockRolacleMockRecorder) Validate(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *RolacleValidateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockRolacle)(nil).Validate), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	return &RolacleValidateCall{Call: call}
}

// RolacleValidateCall wrap *gomock.Call
type RolacleValidateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RolacleValidateCall) Return(arg0 bool, arg1 error) *RolacleValidateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RolacleValidateCall) Do(f func(context.Context, types.LayerID, uint32, int, types.NodeID, types.VrfSignature, uint16) (bool, error)) *RolacleValidateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RolacleValidateCall) DoAndReturn(f func(context.Context, types.LayerID, uint32, int, types.NodeID, types.VrfSignature, uint16) (bool, error)) *RolacleValidateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockstateQuerier is a mock of stateQuerier interface.
type MockstateQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockstateQuerierMockRecorder
}

// MockstateQuerierMockRecorder is the mock recorder for MockstateQuerier.
type MockstateQuerierMockRecorder struct {
	mock *MockstateQuerier
}

// NewMockstateQuerier creates a new mock instance.
func NewMockstateQuerier(ctrl *gomock.Controller) *MockstateQuerier {
	mock := &MockstateQuerier{ctrl: ctrl}
	mock.recorder = &MockstateQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstateQuerier) EXPECT() *MockstateQuerierMockRecorder {
	return m.recorder
}

// IsIdentityActiveOnConsensusView mocks base method.
func (m *MockstateQuerier) IsIdentityActiveOnConsensusView(arg0 context.Context, arg1 types.NodeID, arg2 types.LayerID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIdentityActiveOnConsensusView", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsIdentityActiveOnConsensusView indicates an expected call of IsIdentityActiveOnConsensusView.
func (mr *MockstateQuerierMockRecorder) IsIdentityActiveOnConsensusView(arg0, arg1, arg2 any) *stateQuerierIsIdentityActiveOnConsensusViewCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIdentityActiveOnConsensusView", reflect.TypeOf((*MockstateQuerier)(nil).IsIdentityActiveOnConsensusView), arg0, arg1, arg2)
	return &stateQuerierIsIdentityActiveOnConsensusViewCall{Call: call}
}

// stateQuerierIsIdentityActiveOnConsensusViewCall wrap *gomock.Call
type stateQuerierIsIdentityActiveOnConsensusViewCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *stateQuerierIsIdentityActiveOnConsensusViewCall) Return(arg0 bool, arg1 error) *stateQuerierIsIdentityActiveOnConsensusViewCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *stateQuerierIsIdentityActiveOnConsensusViewCall) Do(f func(context.Context, types.NodeID, types.LayerID) (bool, error)) *stateQuerierIsIdentityActiveOnConsensusViewCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *stateQuerierIsIdentityActiveOnConsensusViewCall) DoAndReturn(f func(context.Context, types.NodeID, types.LayerID) (bool, error)) *stateQuerierIsIdentityActiveOnConsensusViewCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Mockmesh is a mock of mesh interface.
type Mockmesh struct {
	ctrl     *gomock.Controller
	recorder *MockmeshMockRecorder
}

// MockmeshMockRecorder is the mock recorder for Mockmesh.
type MockmeshMockRecorder struct {
	mock *Mockmesh
}

// NewMockmesh creates a new mock instance.
func NewMockmesh(ctrl *gomock.Controller) *Mockmesh {
	mock := &Mockmesh{ctrl: ctrl}
	mock.recorder = &MockmeshMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockmesh) EXPECT() *MockmeshMockRecorder {
	return m.recorder
}

// Ballot mocks base method.
func (m *Mockmesh) Ballot(arg0 types.BallotID) (*types.Ballot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ballot", arg0)
	ret0, _ := ret[0].(*types.Ballot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ballot indicates an expected call of Ballot.
func (mr *MockmeshMockRecorder) Ballot(arg0 any) *meshBallotCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ballot", reflect.TypeOf((*Mockmesh)(nil).Ballot), arg0)
	return &meshBallotCall{Call: call}
}

// meshBallotCall wrap *gomock.Call
type meshBallotCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshBallotCall) Return(arg0 *types.Ballot, arg1 error) *meshBallotCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshBallotCall) Do(f func(types.BallotID) (*types.Ballot, error)) *meshBallotCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshBallotCall) DoAndReturn(f func(types.BallotID) (*types.Ballot, error)) *meshBallotCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Cache mocks base method.
func (m *Mockmesh) Cache() *datastore.CachedDB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cache")
	ret0, _ := ret[0].(*datastore.CachedDB)
	return ret0
}

// Cache indicates an expected call of Cache.
func (mr *MockmeshMockRecorder) Cache() *meshCacheCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cache", reflect.TypeOf((*Mockmesh)(nil).Cache))
	return &meshCacheCall{Call: call}
}

// meshCacheCall wrap *gomock.Call
type meshCacheCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshCacheCall) Return(arg0 *datastore.CachedDB) *meshCacheCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshCacheCall) Do(f func() *datastore.CachedDB) *meshCacheCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshCacheCall) DoAndReturn(f func() *datastore.CachedDB) *meshCacheCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAtxHeader mocks base method.
func (m *Mockmesh) GetAtxHeader(arg0 types.ATXID) (*types.ActivationTxHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAtxHeader", arg0)
	ret0, _ := ret[0].(*types.ActivationTxHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAtxHeader indicates an expected call of GetAtxHeader.
func (mr *MockmeshMockRecorder) GetAtxHeader(arg0 any) *meshGetAtxHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAtxHeader", reflect.TypeOf((*Mockmesh)(nil).GetAtxHeader), arg0)
	return &meshGetAtxHeaderCall{Call: call}
}

// meshGetAtxHeaderCall wrap *gomock.Call
type meshGetAtxHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshGetAtxHeaderCall) Return(arg0 *types.ActivationTxHeader, arg1 error) *meshGetAtxHeaderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshGetAtxHeaderCall) Do(f func(types.ATXID) (*types.ActivationTxHeader, error)) *meshGetAtxHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshGetAtxHeaderCall) DoAndReturn(f func(types.ATXID) (*types.ActivationTxHeader, error)) *meshGetAtxHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetEpochAtx mocks base method.
func (m *Mockmesh) GetEpochAtx(arg0 types.EpochID, arg1 types.NodeID) (*types.ActivationTxHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpochAtx", arg0, arg1)
	ret0, _ := ret[0].(*types.ActivationTxHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEpochAtx indicates an expected call of GetEpochAtx.
func (mr *MockmeshMockRecorder) GetEpochAtx(arg0, arg1 any) *meshGetEpochAtxCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpochAtx", reflect.TypeOf((*Mockmesh)(nil).GetEpochAtx), arg0, arg1)
	return &meshGetEpochAtxCall{Call: call}
}

// meshGetEpochAtxCall wrap *gomock.Call
type meshGetEpochAtxCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshGetEpochAtxCall) Return(arg0 *types.ActivationTxHeader, arg1 error) *meshGetEpochAtxCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshGetEpochAtxCall) Do(f func(types.EpochID, types.NodeID) (*types.ActivationTxHeader, error)) *meshGetEpochAtxCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshGetEpochAtxCall) DoAndReturn(f func(types.EpochID, types.NodeID) (*types.ActivationTxHeader, error)) *meshGetEpochAtxCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMalfeasanceProof mocks base method.
func (m *Mockmesh) GetMalfeasanceProof(arg0 types.NodeID) (*types.MalfeasanceProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMalfeasanceProof", arg0)
	ret0, _ := ret[0].(*types.MalfeasanceProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMalfeasanceProof indicates an expected call of GetMalfeasanceProof.
func (mr *MockmeshMockRecorder) GetMalfeasanceProof(arg0 any) *meshGetMalfeasanceProofCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMalfeasanceProof", reflect.TypeOf((*Mockmesh)(nil).GetMalfeasanceProof), arg0)
	return &meshGetMalfeasanceProofCall{Call: call}
}

// meshGetMalfeasanceProofCall wrap *gomock.Call
type meshGetMalfeasanceProofCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshGetMalfeasanceProofCall) Return(arg0 *types.MalfeasanceProof, arg1 error) *meshGetMalfeasanceProofCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshGetMalfeasanceProofCall) Do(f func(types.NodeID) (*types.MalfeasanceProof, error)) *meshGetMalfeasanceProofCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshGetMalfeasanceProofCall) DoAndReturn(f func(types.NodeID) (*types.MalfeasanceProof, error)) *meshGetMalfeasanceProofCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Proposals mocks base method.
func (m *Mockmesh) Proposals(arg0 types.LayerID) ([]*types.Proposal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proposals", arg0)
	ret0, _ := ret[0].([]*types.Proposal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Proposals indicates an expected call of Proposals.
func (mr *MockmeshMockRecorder) Proposals(arg0 any) *meshProposalsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proposals", reflect.TypeOf((*Mockmesh)(nil).Proposals), arg0)
	return &meshProposalsCall{Call: call}
}

// meshProposalsCall wrap *gomock.Call
type meshProposalsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *meshProposalsCall) Return(arg0 []*types.Proposal, arg1 error) *meshProposalsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *meshProposalsCall) Do(f func(types.LayerID) ([]*types.Proposal, error)) *meshProposalsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *meshProposalsCall) DoAndReturn(f func(types.LayerID) ([]*types.Proposal, error)) *meshProposalsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockweakCoin is a mock of weakCoin interface.
type MockweakCoin struct {
	ctrl     *gomock.Controller
	recorder *MockweakCoinMockRecorder
}

// MockweakCoinMockRecorder is the mock recorder for MockweakCoin.
type MockweakCoinMockRecorder struct {
	mock *MockweakCoin
}

// NewMockweakCoin creates a new mock instance.
func NewMockweakCoin(ctrl *gomock.Controller) *MockweakCoin {
	mock := &MockweakCoin{ctrl: ctrl}
	mock.recorder = &MockweakCoinMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockweakCoin) EXPECT() *MockweakCoinMockRecorder {
	return m.recorder
}

// Set mocks base method.
func (m *MockweakCoin) Set(arg0 types.LayerID, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockweakCoinMockRecorder) Set(arg0, arg1 any) *weakCoinSetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockweakCoin)(nil).Set), arg0, arg1)
	return &weakCoinSetCall{Call: call}
}

// weakCoinSetCall wrap *gomock.Call
type weakCoinSetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *weakCoinSetCall) Return(arg0 error) *weakCoinSetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *weakCoinSetCall) Do(f func(types.LayerID, bool) error) *weakCoinSetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *weakCoinSetCall) DoAndReturn(f func(types.LayerID, bool) error) *weakCoinSetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
