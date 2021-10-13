// Code generated by MockGen. DO NOT EDIT.
// Source: ethereum/client.go

// Package mock_ethereum is a generated GoMock package.
package mock_ethereum

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ethereum "github.com/ethereum/go-ethereum"
	types "github.com/ethereum/go-ethereum/core/types"
	domain "github.com/forta-protocol/forta-node/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockrpcClient is a mock of rpcClient interface.
type MockrpcClient struct {
	ctrl     *gomock.Controller
	recorder *MockrpcClientMockRecorder
}

// MockrpcClientMockRecorder is the mock recorder for MockrpcClient.
type MockrpcClientMockRecorder struct {
	mock *MockrpcClient
}

// NewMockrpcClient creates a new mock instance.
func NewMockrpcClient(ctrl *gomock.Controller) *MockrpcClient {
	mock := &MockrpcClient{ctrl: ctrl}
	mock.recorder = &MockrpcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockrpcClient) EXPECT() *MockrpcClientMockRecorder {
	return m.recorder
}

// CallContext mocks base method.
func (m *MockrpcClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, result, method}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CallContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CallContext indicates an expected call of CallContext.
func (mr *MockrpcClientMockRecorder) CallContext(ctx, result, method interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, result, method}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContext", reflect.TypeOf((*MockrpcClient)(nil).CallContext), varargs...)
}

// Close mocks base method.
func (m *MockrpcClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockrpcClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockrpcClient)(nil).Close))
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method.
func (m *MockClient) BlockByHash(ctx context.Context, hash string) (*domain.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash)
	ret0, _ := ret[0].(*domain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockClientMockRecorder) BlockByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockClient)(nil).BlockByHash), ctx, hash)
}

// BlockByNumber mocks base method.
func (m *MockClient) BlockByNumber(ctx context.Context, number *big.Int) (*domain.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number)
	ret0, _ := ret[0].(*domain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockClientMockRecorder) BlockByNumber(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockClient)(nil).BlockByNumber), ctx, number)
}

// BlockNumber mocks base method.
func (m *MockClient) BlockNumber(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockClientMockRecorder) BlockNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockClient)(nil).BlockNumber), ctx)
}

// ChainID mocks base method.
func (m *MockClient) ChainID(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainID indicates an expected call of ChainID.
func (mr *MockClientMockRecorder) ChainID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockClient)(nil).ChainID), ctx)
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// GetLogs mocks base method.
func (m *MockClient) GetLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", ctx, q)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockClientMockRecorder) GetLogs(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockClient)(nil).GetLogs), ctx, q)
}

// TraceBlock mocks base method.
func (m *MockClient) TraceBlock(ctx context.Context, number *big.Int) ([]domain.Trace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceBlock", ctx, number)
	ret0, _ := ret[0].([]domain.Trace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TraceBlock indicates an expected call of TraceBlock.
func (mr *MockClientMockRecorder) TraceBlock(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBlock", reflect.TypeOf((*MockClient)(nil).TraceBlock), ctx, number)
}

// TransactionReceipt mocks base method.
func (m *MockClient) TransactionReceipt(ctx context.Context, txHash string) (*domain.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*domain.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockClientMockRecorder) TransactionReceipt(ctx, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockClient)(nil).TransactionReceipt), ctx, txHash)
}
