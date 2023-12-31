// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package wardbusinesssus is a generated GoMock package.
package wardbusinesssus

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataWardBusinessSusAdapter is a mock of DataWardBusinessSusAdapter interface.
type MockDataWardBusinessSusAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockDataWardBusinessSusAdapterMockRecorder
}

// MockDataWardBusinessSusAdapterMockRecorder is the mock recorder for MockDataWardBusinessSusAdapter.
type MockDataWardBusinessSusAdapterMockRecorder struct {
	mock *MockDataWardBusinessSusAdapter
}

// NewMockDataWardBusinessSusAdapter creates a new mock instance.
func NewMockDataWardBusinessSusAdapter(ctrl *gomock.Controller) *MockDataWardBusinessSusAdapter {
	mock := &MockDataWardBusinessSusAdapter{ctrl: ctrl}
	mock.recorder = &MockDataWardBusinessSusAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataWardBusinessSusAdapter) EXPECT() *MockDataWardBusinessSusAdapterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDataWardBusinessSusAdapter) Create(arg0 *WardBusinessSus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDataWardBusinessSusAdapterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDataWardBusinessSusAdapter)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockDataWardBusinessSusAdapter) Delete(arg0 *WardBusinessSus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDataWardBusinessSusAdapterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataWardBusinessSusAdapter)(nil).Delete), arg0)
}

// Read mocks base method.
func (m *MockDataWardBusinessSusAdapter) Read(arg0 *WardBusinessSus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockDataWardBusinessSusAdapterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDataWardBusinessSusAdapter)(nil).Read), arg0)
}

// ReadAll mocks base method.
func (m *MockDataWardBusinessSusAdapter) ReadAll(arg0 *[]WardBusinessSus, arg1 WardBusinessSusParam) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockDataWardBusinessSusAdapterMockRecorder) ReadAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockDataWardBusinessSusAdapter)(nil).ReadAll), arg0, arg1)
}

// Update mocks base method.
func (m *MockDataWardBusinessSusAdapter) Update(arg0 WardBusinessSus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDataWardBusinessSusAdapterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataWardBusinessSusAdapter)(nil).Update), arg0)
}
