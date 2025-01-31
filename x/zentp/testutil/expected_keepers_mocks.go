// Code generated by MockGen. DO NOT EDIT.
// Source: x/zentp/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	types "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(arg0 context.Context, arg1 types0.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), arg0, arg1)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// HasAccount mocks base method.
func (m *MockAccountKeeper) HasAccount(arg0 context.Context, arg1 types0.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAccount", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasAccount indicates an expected call of HasAccount.
func (mr *MockAccountKeeperMockRecorder) HasAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAccount", reflect.TypeOf((*MockAccountKeeper)(nil).HasAccount), arg0, arg1)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx context.Context, moduleAccount string, amounts types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleAccount, amounts)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleAccount, amounts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleAccount, amounts)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(ctx context.Context, addr types0.AccAddress, denom string) types0.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types0.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(ctx, addr, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), ctx, addr, denom)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(arg0 context.Context, arg1 types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", arg0, arg1)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), arg0, arg1)
}

// MockParamSubspace is a mock of ParamSubspace interface.
type MockParamSubspace struct {
	ctrl     *gomock.Controller
	recorder *MockParamSubspaceMockRecorder
}

// MockParamSubspaceMockRecorder is the mock recorder for MockParamSubspace.
type MockParamSubspaceMockRecorder struct {
	mock *MockParamSubspace
}

// NewMockParamSubspace creates a new mock instance.
func NewMockParamSubspace(ctrl *gomock.Controller) *MockParamSubspace {
	mock := &MockParamSubspace{ctrl: ctrl}
	mock.recorder = &MockParamSubspaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParamSubspace) EXPECT() *MockParamSubspaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockParamSubspace) Get(arg0 context.Context, arg1 []byte, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", arg0, arg1, arg2)
}

// Get indicates an expected call of Get.
func (mr *MockParamSubspaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockParamSubspace)(nil).Get), arg0, arg1, arg2)
}

// Set mocks base method.
func (m *MockParamSubspace) Set(arg0 context.Context, arg1 []byte, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1, arg2)
}

// Set indicates an expected call of Set.
func (mr *MockParamSubspaceMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockParamSubspace)(nil).Set), arg0, arg1, arg2)
}

// MockTreasuryKeeper is a mock of TreasuryKeeper interface.
type MockTreasuryKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockTreasuryKeeperMockRecorder
}

// MockTreasuryKeeperMockRecorder is the mock recorder for MockTreasuryKeeper.
type MockTreasuryKeeperMockRecorder struct {
	mock *MockTreasuryKeeper
}

// NewMockTreasuryKeeper creates a new mock instance.
func NewMockTreasuryKeeper(ctrl *gomock.Controller) *MockTreasuryKeeper {
	mock := &MockTreasuryKeeper{ctrl: ctrl}
	mock.recorder = &MockTreasuryKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTreasuryKeeper) EXPECT() *MockTreasuryKeeperMockRecorder {
	return m.recorder
}

// GetKey mocks base method.
func (m *MockTreasuryKeeper) GetKey(ctx types0.Context, keyID uint64) (*types.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", ctx, keyID)
	ret0, _ := ret[0].(*types.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKey indicates an expected call of GetKey.
func (mr *MockTreasuryKeeperMockRecorder) GetKey(ctx, keyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockTreasuryKeeper)(nil).GetKey), ctx, keyID)
}
