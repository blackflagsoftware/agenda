// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package agenda is a generated GoMock package.
package agenda

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataAgendaAdapter is a mock of DataAgendaAdapter interface.
type MockDataAgendaAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockDataAgendaAdapterMockRecorder
}

// MockDataAgendaAdapterMockRecorder is the mock recorder for MockDataAgendaAdapter.
type MockDataAgendaAdapterMockRecorder struct {
	mock *MockDataAgendaAdapter
}

// NewMockDataAgendaAdapter creates a new mock instance.
func NewMockDataAgendaAdapter(ctrl *gomock.Controller) *MockDataAgendaAdapter {
	mock := &MockDataAgendaAdapter{ctrl: ctrl}
	mock.recorder = &MockDataAgendaAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataAgendaAdapter) EXPECT() *MockDataAgendaAdapterMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockDataAgendaAdapter) Check(arg0 *Agenda) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockDataAgendaAdapterMockRecorder) Check(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockDataAgendaAdapter)(nil).Check), arg0)
}

// Create mocks base method.
func (m *MockDataAgendaAdapter) Create(arg0 *Agenda) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDataAgendaAdapterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDataAgendaAdapter)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockDataAgendaAdapter) Delete(arg0 *Agenda) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDataAgendaAdapterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataAgendaAdapter)(nil).Delete), arg0)
}

// Read mocks base method.
func (m *MockDataAgendaAdapter) Read(arg0 *Agenda) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockDataAgendaAdapterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDataAgendaAdapter)(nil).Read), arg0)
}

// ReadAll mocks base method.
func (m *MockDataAgendaAdapter) ReadAll(arg0 *[]Agenda, arg1 AgendaParam) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockDataAgendaAdapterMockRecorder) ReadAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockDataAgendaAdapter)(nil).ReadAll), arg0, arg1)
}

// Update mocks base method.
func (m *MockDataAgendaAdapter) Update(arg0 Agenda) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDataAgendaAdapterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataAgendaAdapter)(nil).Update), arg0)
}
