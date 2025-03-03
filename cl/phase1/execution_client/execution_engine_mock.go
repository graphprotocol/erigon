// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go
//
// Generated by this command:
//
//	mockgen -typed=true -source=./interface.go -destination=./execution_engine_mock.go -package=execution_client . ExecutionEngine
//

// Package execution_client is a generated GoMock package.
package execution_client

import (
	context "context"
	big "math/big"
	reflect "reflect"

	common "github.com/erigontech/erigon-lib/common"
	hexutil "github.com/erigontech/erigon-lib/common/hexutil"
	cltypes "github.com/erigontech/erigon/cl/cltypes"
	types "github.com/erigontech/erigon/core/types"
	engine_types "github.com/erigontech/erigon/turbo/engineapi/engine_types"
	gomock "go.uber.org/mock/gomock"
)

// MockExecutionEngine is a mock of ExecutionEngine interface.
type MockExecutionEngine struct {
	ctrl     *gomock.Controller
	recorder *MockExecutionEngineMockRecorder
	isgomock struct{}
}

// MockExecutionEngineMockRecorder is the mock recorder for MockExecutionEngine.
type MockExecutionEngineMockRecorder struct {
	mock *MockExecutionEngine
}

// NewMockExecutionEngine creates a new mock instance.
func NewMockExecutionEngine(ctrl *gomock.Controller) *MockExecutionEngine {
	mock := &MockExecutionEngine{ctrl: ctrl}
	mock.recorder = &MockExecutionEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutionEngine) EXPECT() *MockExecutionEngineMockRecorder {
	return m.recorder
}

// CurrentHeader mocks base method.
func (m *MockExecutionEngine) CurrentHeader(ctx context.Context) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentHeader", ctx)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentHeader indicates an expected call of CurrentHeader.
func (mr *MockExecutionEngineMockRecorder) CurrentHeader(ctx any) *MockExecutionEngineCurrentHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentHeader", reflect.TypeOf((*MockExecutionEngine)(nil).CurrentHeader), ctx)
	return &MockExecutionEngineCurrentHeaderCall{Call: call}
}

// MockExecutionEngineCurrentHeaderCall wrap *gomock.Call
type MockExecutionEngineCurrentHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineCurrentHeaderCall) Return(arg0 *types.Header, arg1 error) *MockExecutionEngineCurrentHeaderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineCurrentHeaderCall) Do(f func(context.Context) (*types.Header, error)) *MockExecutionEngineCurrentHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineCurrentHeaderCall) DoAndReturn(f func(context.Context) (*types.Header, error)) *MockExecutionEngineCurrentHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ForkChoiceUpdate mocks base method.
func (m *MockExecutionEngine) ForkChoiceUpdate(ctx context.Context, finalized, head common.Hash, attributes *engine_types.PayloadAttributes) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForkChoiceUpdate", ctx, finalized, head, attributes)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForkChoiceUpdate indicates an expected call of ForkChoiceUpdate.
func (mr *MockExecutionEngineMockRecorder) ForkChoiceUpdate(ctx, finalized, head, attributes any) *MockExecutionEngineForkChoiceUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForkChoiceUpdate", reflect.TypeOf((*MockExecutionEngine)(nil).ForkChoiceUpdate), ctx, finalized, head, attributes)
	return &MockExecutionEngineForkChoiceUpdateCall{Call: call}
}

// MockExecutionEngineForkChoiceUpdateCall wrap *gomock.Call
type MockExecutionEngineForkChoiceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineForkChoiceUpdateCall) Return(arg0 []byte, arg1 error) *MockExecutionEngineForkChoiceUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineForkChoiceUpdateCall) Do(f func(context.Context, common.Hash, common.Hash, *engine_types.PayloadAttributes) ([]byte, error)) *MockExecutionEngineForkChoiceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineForkChoiceUpdateCall) DoAndReturn(f func(context.Context, common.Hash, common.Hash, *engine_types.PayloadAttributes) ([]byte, error)) *MockExecutionEngineForkChoiceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FrozenBlocks mocks base method.
func (m *MockExecutionEngine) FrozenBlocks(ctx context.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FrozenBlocks", ctx)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// FrozenBlocks indicates an expected call of FrozenBlocks.
func (mr *MockExecutionEngineMockRecorder) FrozenBlocks(ctx any) *MockExecutionEngineFrozenBlocksCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FrozenBlocks", reflect.TypeOf((*MockExecutionEngine)(nil).FrozenBlocks), ctx)
	return &MockExecutionEngineFrozenBlocksCall{Call: call}
}

// MockExecutionEngineFrozenBlocksCall wrap *gomock.Call
type MockExecutionEngineFrozenBlocksCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineFrozenBlocksCall) Return(arg0 uint64) *MockExecutionEngineFrozenBlocksCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineFrozenBlocksCall) Do(f func(context.Context) uint64) *MockExecutionEngineFrozenBlocksCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineFrozenBlocksCall) DoAndReturn(f func(context.Context) uint64) *MockExecutionEngineFrozenBlocksCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAssembledBlock mocks base method.
func (m *MockExecutionEngine) GetAssembledBlock(ctx context.Context, id []byte) (*cltypes.Eth1Block, *engine_types.BlobsBundleV1, *big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssembledBlock", ctx, id)
	ret0, _ := ret[0].(*cltypes.Eth1Block)
	ret1, _ := ret[1].(*engine_types.BlobsBundleV1)
	ret2, _ := ret[2].(*big.Int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetAssembledBlock indicates an expected call of GetAssembledBlock.
func (mr *MockExecutionEngineMockRecorder) GetAssembledBlock(ctx, id any) *MockExecutionEngineGetAssembledBlockCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssembledBlock", reflect.TypeOf((*MockExecutionEngine)(nil).GetAssembledBlock), ctx, id)
	return &MockExecutionEngineGetAssembledBlockCall{Call: call}
}

// MockExecutionEngineGetAssembledBlockCall wrap *gomock.Call
type MockExecutionEngineGetAssembledBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineGetAssembledBlockCall) Return(arg0 *cltypes.Eth1Block, arg1 *engine_types.BlobsBundleV1, arg2 *big.Int, arg3 error) *MockExecutionEngineGetAssembledBlockCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineGetAssembledBlockCall) Do(f func(context.Context, []byte) (*cltypes.Eth1Block, *engine_types.BlobsBundleV1, *big.Int, error)) *MockExecutionEngineGetAssembledBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineGetAssembledBlockCall) DoAndReturn(f func(context.Context, []byte) (*cltypes.Eth1Block, *engine_types.BlobsBundleV1, *big.Int, error)) *MockExecutionEngineGetAssembledBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetBodiesByHashes mocks base method.
func (m *MockExecutionEngine) GetBodiesByHashes(ctx context.Context, hashes []common.Hash) ([]*types.RawBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBodiesByHashes", ctx, hashes)
	ret0, _ := ret[0].([]*types.RawBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBodiesByHashes indicates an expected call of GetBodiesByHashes.
func (mr *MockExecutionEngineMockRecorder) GetBodiesByHashes(ctx, hashes any) *MockExecutionEngineGetBodiesByHashesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBodiesByHashes", reflect.TypeOf((*MockExecutionEngine)(nil).GetBodiesByHashes), ctx, hashes)
	return &MockExecutionEngineGetBodiesByHashesCall{Call: call}
}

// MockExecutionEngineGetBodiesByHashesCall wrap *gomock.Call
type MockExecutionEngineGetBodiesByHashesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineGetBodiesByHashesCall) Return(arg0 []*types.RawBody, arg1 error) *MockExecutionEngineGetBodiesByHashesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineGetBodiesByHashesCall) Do(f func(context.Context, []common.Hash) ([]*types.RawBody, error)) *MockExecutionEngineGetBodiesByHashesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineGetBodiesByHashesCall) DoAndReturn(f func(context.Context, []common.Hash) ([]*types.RawBody, error)) *MockExecutionEngineGetBodiesByHashesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetBodiesByRange mocks base method.
func (m *MockExecutionEngine) GetBodiesByRange(ctx context.Context, start, count uint64) ([]*types.RawBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBodiesByRange", ctx, start, count)
	ret0, _ := ret[0].([]*types.RawBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBodiesByRange indicates an expected call of GetBodiesByRange.
func (mr *MockExecutionEngineMockRecorder) GetBodiesByRange(ctx, start, count any) *MockExecutionEngineGetBodiesByRangeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBodiesByRange", reflect.TypeOf((*MockExecutionEngine)(nil).GetBodiesByRange), ctx, start, count)
	return &MockExecutionEngineGetBodiesByRangeCall{Call: call}
}

// MockExecutionEngineGetBodiesByRangeCall wrap *gomock.Call
type MockExecutionEngineGetBodiesByRangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineGetBodiesByRangeCall) Return(arg0 []*types.RawBody, arg1 error) *MockExecutionEngineGetBodiesByRangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineGetBodiesByRangeCall) Do(f func(context.Context, uint64, uint64) ([]*types.RawBody, error)) *MockExecutionEngineGetBodiesByRangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineGetBodiesByRangeCall) DoAndReturn(f func(context.Context, uint64, uint64) ([]*types.RawBody, error)) *MockExecutionEngineGetBodiesByRangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HasBlock mocks base method.
func (m *MockExecutionEngine) HasBlock(ctx context.Context, hash common.Hash) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBlock", ctx, hash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasBlock indicates an expected call of HasBlock.
func (mr *MockExecutionEngineMockRecorder) HasBlock(ctx, hash any) *MockExecutionEngineHasBlockCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBlock", reflect.TypeOf((*MockExecutionEngine)(nil).HasBlock), ctx, hash)
	return &MockExecutionEngineHasBlockCall{Call: call}
}

// MockExecutionEngineHasBlockCall wrap *gomock.Call
type MockExecutionEngineHasBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineHasBlockCall) Return(arg0 bool, arg1 error) *MockExecutionEngineHasBlockCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineHasBlockCall) Do(f func(context.Context, common.Hash) (bool, error)) *MockExecutionEngineHasBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineHasBlockCall) DoAndReturn(f func(context.Context, common.Hash) (bool, error)) *MockExecutionEngineHasBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HasGapInSnapshots mocks base method.
func (m *MockExecutionEngine) HasGapInSnapshots(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasGapInSnapshots", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasGapInSnapshots indicates an expected call of HasGapInSnapshots.
func (mr *MockExecutionEngineMockRecorder) HasGapInSnapshots(ctx any) *MockExecutionEngineHasGapInSnapshotsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasGapInSnapshots", reflect.TypeOf((*MockExecutionEngine)(nil).HasGapInSnapshots), ctx)
	return &MockExecutionEngineHasGapInSnapshotsCall{Call: call}
}

// MockExecutionEngineHasGapInSnapshotsCall wrap *gomock.Call
type MockExecutionEngineHasGapInSnapshotsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineHasGapInSnapshotsCall) Return(arg0 bool) *MockExecutionEngineHasGapInSnapshotsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineHasGapInSnapshotsCall) Do(f func(context.Context) bool) *MockExecutionEngineHasGapInSnapshotsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineHasGapInSnapshotsCall) DoAndReturn(f func(context.Context) bool) *MockExecutionEngineHasGapInSnapshotsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InsertBlock mocks base method.
func (m *MockExecutionEngine) InsertBlock(ctx context.Context, block *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlock", ctx, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBlock indicates an expected call of InsertBlock.
func (mr *MockExecutionEngineMockRecorder) InsertBlock(ctx, block any) *MockExecutionEngineInsertBlockCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlock", reflect.TypeOf((*MockExecutionEngine)(nil).InsertBlock), ctx, block)
	return &MockExecutionEngineInsertBlockCall{Call: call}
}

// MockExecutionEngineInsertBlockCall wrap *gomock.Call
type MockExecutionEngineInsertBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineInsertBlockCall) Return(arg0 error) *MockExecutionEngineInsertBlockCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineInsertBlockCall) Do(f func(context.Context, *types.Block) error) *MockExecutionEngineInsertBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineInsertBlockCall) DoAndReturn(f func(context.Context, *types.Block) error) *MockExecutionEngineInsertBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InsertBlocks mocks base method.
func (m *MockExecutionEngine) InsertBlocks(ctx context.Context, blocks []*types.Block, wait bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlocks", ctx, blocks, wait)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBlocks indicates an expected call of InsertBlocks.
func (mr *MockExecutionEngineMockRecorder) InsertBlocks(ctx, blocks, wait any) *MockExecutionEngineInsertBlocksCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlocks", reflect.TypeOf((*MockExecutionEngine)(nil).InsertBlocks), ctx, blocks, wait)
	return &MockExecutionEngineInsertBlocksCall{Call: call}
}

// MockExecutionEngineInsertBlocksCall wrap *gomock.Call
type MockExecutionEngineInsertBlocksCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineInsertBlocksCall) Return(arg0 error) *MockExecutionEngineInsertBlocksCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineInsertBlocksCall) Do(f func(context.Context, []*types.Block, bool) error) *MockExecutionEngineInsertBlocksCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineInsertBlocksCall) DoAndReturn(f func(context.Context, []*types.Block, bool) error) *MockExecutionEngineInsertBlocksCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsCanonicalHash mocks base method.
func (m *MockExecutionEngine) IsCanonicalHash(ctx context.Context, hash common.Hash) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCanonicalHash", ctx, hash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCanonicalHash indicates an expected call of IsCanonicalHash.
func (mr *MockExecutionEngineMockRecorder) IsCanonicalHash(ctx, hash any) *MockExecutionEngineIsCanonicalHashCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCanonicalHash", reflect.TypeOf((*MockExecutionEngine)(nil).IsCanonicalHash), ctx, hash)
	return &MockExecutionEngineIsCanonicalHashCall{Call: call}
}

// MockExecutionEngineIsCanonicalHashCall wrap *gomock.Call
type MockExecutionEngineIsCanonicalHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineIsCanonicalHashCall) Return(arg0 bool, arg1 error) *MockExecutionEngineIsCanonicalHashCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineIsCanonicalHashCall) Do(f func(context.Context, common.Hash) (bool, error)) *MockExecutionEngineIsCanonicalHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineIsCanonicalHashCall) DoAndReturn(f func(context.Context, common.Hash) (bool, error)) *MockExecutionEngineIsCanonicalHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewPayload mocks base method.
func (m *MockExecutionEngine) NewPayload(ctx context.Context, payload *cltypes.Eth1Block, beaconParentRoot *common.Hash, versionedHashes []common.Hash, executionRequestsList []hexutil.Bytes) (PayloadStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPayload", ctx, payload, beaconParentRoot, versionedHashes, executionRequestsList)
	ret0, _ := ret[0].(PayloadStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPayload indicates an expected call of NewPayload.
func (mr *MockExecutionEngineMockRecorder) NewPayload(ctx, payload, beaconParentRoot, versionedHashes, executionRequestsList any) *MockExecutionEngineNewPayloadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPayload", reflect.TypeOf((*MockExecutionEngine)(nil).NewPayload), ctx, payload, beaconParentRoot, versionedHashes, executionRequestsList)
	return &MockExecutionEngineNewPayloadCall{Call: call}
}

// MockExecutionEngineNewPayloadCall wrap *gomock.Call
type MockExecutionEngineNewPayloadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineNewPayloadCall) Return(arg0 PayloadStatus, arg1 error) *MockExecutionEngineNewPayloadCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineNewPayloadCall) Do(f func(context.Context, *cltypes.Eth1Block, *common.Hash, []common.Hash, []hexutil.Bytes) (PayloadStatus, error)) *MockExecutionEngineNewPayloadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineNewPayloadCall) DoAndReturn(f func(context.Context, *cltypes.Eth1Block, *common.Hash, []common.Hash, []hexutil.Bytes) (PayloadStatus, error)) *MockExecutionEngineNewPayloadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Ready mocks base method.
func (m *MockExecutionEngine) Ready(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ready indicates an expected call of Ready.
func (mr *MockExecutionEngineMockRecorder) Ready(ctx any) *MockExecutionEngineReadyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockExecutionEngine)(nil).Ready), ctx)
	return &MockExecutionEngineReadyCall{Call: call}
}

// MockExecutionEngineReadyCall wrap *gomock.Call
type MockExecutionEngineReadyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineReadyCall) Return(arg0 bool, arg1 error) *MockExecutionEngineReadyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineReadyCall) Do(f func(context.Context) (bool, error)) *MockExecutionEngineReadyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineReadyCall) DoAndReturn(f func(context.Context) (bool, error)) *MockExecutionEngineReadyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SupportInsertion mocks base method.
func (m *MockExecutionEngine) SupportInsertion() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportInsertion")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SupportInsertion indicates an expected call of SupportInsertion.
func (mr *MockExecutionEngineMockRecorder) SupportInsertion() *MockExecutionEngineSupportInsertionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportInsertion", reflect.TypeOf((*MockExecutionEngine)(nil).SupportInsertion))
	return &MockExecutionEngineSupportInsertionCall{Call: call}
}

// MockExecutionEngineSupportInsertionCall wrap *gomock.Call
type MockExecutionEngineSupportInsertionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExecutionEngineSupportInsertionCall) Return(arg0 bool) *MockExecutionEngineSupportInsertionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExecutionEngineSupportInsertionCall) Do(f func() bool) *MockExecutionEngineSupportInsertionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExecutionEngineSupportInsertionCall) DoAndReturn(f func() bool) *MockExecutionEngineSupportInsertionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
