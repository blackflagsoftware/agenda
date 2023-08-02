// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package roleuser is a generated GoMock package.
package roleuser

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataRoleUserAdapter is a mock of DataRoleUserAdapter interface.
type MockDataRoleUserAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockDataRoleUserAdapterMockRecorder
}

// MockDataRoleUserAdapterMockRecorder is the mock recorder for MockDataRoleUserAdapter.
type MockDataRoleUserAdapterMockRecorder struct {
	mock *MockDataRoleUserAdapter
}

// NewMockDataRoleUserAdapter creates a new mock instance.
func NewMockDataRoleUserAdapter(ctrl *gomock.Controller) *MockDataRoleUserAdapter {
	mock := &MockDataRoleUserAdapter{ctrl: ctrl}
	mock.recorder = &MockDataRoleUserAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataRoleUserAdapter) EXPECT() *MockDataRoleUserAdapterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDataRoleUserAdapter) Create(arg0 *RoleUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDataRoleUserAdapterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDataRoleUserAdapter)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockDataRoleUserAdapter) Delete(arg0 *RoleUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDataRoleUserAdapterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataRoleUserAdapter)(nil).Delete), arg0)
}

// Login mocks base method.
func (m *MockDataRoleUserAdapter) Login(arg0 RoleUser) (RoleLogin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(RoleLogin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockDataRoleUserAdapterMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockDataRoleUserAdapter)(nil).Login), arg0)
}

// Read mocks base method.
func (m *MockDataRoleUserAdapter) Read(arg0 *RoleUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockDataRoleUserAdapterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDataRoleUserAdapter)(nil).Read), arg0)
}

// ReadAll mocks base method.
func (m *MockDataRoleUserAdapter) ReadAll(arg0 *[]RoleUser, arg1 RoleUserParam) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockDataRoleUserAdapterMockRecorder) ReadAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockDataRoleUserAdapter)(nil).ReadAll), arg0, arg1)
}

// Update mocks base method.
func (m *MockDataRoleUserAdapter) Update(arg0 RoleUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDataRoleUserAdapterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataRoleUserAdapter)(nil).Update), arg0)
}
