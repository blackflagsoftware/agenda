// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package announcement is a generated GoMock package.
package announcement

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataAnnouncementAdapter is a mock of DataAnnouncementAdapter interface.
type MockDataAnnouncementAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockDataAnnouncementAdapterMockRecorder
}

// MockDataAnnouncementAdapterMockRecorder is the mock recorder for MockDataAnnouncementAdapter.
type MockDataAnnouncementAdapterMockRecorder struct {
	mock *MockDataAnnouncementAdapter
}

// NewMockDataAnnouncementAdapter creates a new mock instance.
func NewMockDataAnnouncementAdapter(ctrl *gomock.Controller) *MockDataAnnouncementAdapter {
	mock := &MockDataAnnouncementAdapter{ctrl: ctrl}
	mock.recorder = &MockDataAnnouncementAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataAnnouncementAdapter) EXPECT() *MockDataAnnouncementAdapterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDataAnnouncementAdapter) Create(arg0 *Announcement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDataAnnouncementAdapterMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDataAnnouncementAdapter)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockDataAnnouncementAdapter) Delete(arg0 *Announcement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDataAnnouncementAdapterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataAnnouncementAdapter)(nil).Delete), arg0)
}

// Read mocks base method.
func (m *MockDataAnnouncementAdapter) Read(arg0 *Announcement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockDataAnnouncementAdapterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDataAnnouncementAdapter)(nil).Read), arg0)
}

// ReadAll mocks base method.
func (m *MockDataAnnouncementAdapter) ReadAll(arg0 *[]Announcement, arg1 AnnouncementParam) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockDataAnnouncementAdapterMockRecorder) ReadAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockDataAnnouncementAdapter)(nil).ReadAll), arg0, arg1)
}

// Update mocks base method.
func (m *MockDataAnnouncementAdapter) Update(arg0 Announcement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDataAnnouncementAdapterMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataAnnouncementAdapter)(nil).Update), arg0)
}
