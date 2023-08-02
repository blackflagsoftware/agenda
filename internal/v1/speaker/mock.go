// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package speaker is a generated GoMock package.
package speaker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataSpeakerAdapter is a mock of DataSpeakerAdapter interface.
type MockDataSpeakerAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockDataSpeakerAdapterMockRecorder
}

// MockDataSpeakerAdapterMockRecorder is the mock recorder for MockDataSpeakerAdapter.
type MockDataSpeakerAdapterMockRecorder struct {
	mock *MockDataSpeakerAdapter
}

// NewMockDataSpeakerAdapter creates a new mock instance.
func NewMockDataSpeakerAdapter(ctrl *gomock.Controller) *MockDataSpeakerAdapter {
	mock := &MockDataSpeakerAdapter{ctrl: ctrl}
	mock.recorder = &MockDataSpeakerAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSpeakerAdapter) EXPECT() *MockDataSpeakerAdapterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDataSpeakerAdapter) Create(arg0 *Speaker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDataSpeakerAdapterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDataSpeakerAdapter)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockDataSpeakerAdapter) Delete(arg0 *Speaker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDataSpeakerAdapterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataSpeakerAdapter)(nil).Delete), arg0)
}

// Read mocks base method.
func (m *MockDataSpeakerAdapter) Read(arg0 *Speaker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockDataSpeakerAdapterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDataSpeakerAdapter)(nil).Read), arg0)
}

// ReadAll mocks base method.
func (m *MockDataSpeakerAdapter) ReadAll(arg0 *[]Speaker, arg1 SpeakerParam) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockDataSpeakerAdapterMockRecorder) ReadAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockDataSpeakerAdapter)(nil).ReadAll), arg0, arg1)
}

// Update mocks base method.
func (m *MockDataSpeakerAdapter) Update(arg0 Speaker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDataSpeakerAdapterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataSpeakerAdapter)(nil).Update), arg0)
}
