// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/seagullbird/headr-common/mq/dispatch (interfaces: Dispatcher)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDispatcher is a mock of Dispatcher interface
type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

// MockDispatcherMockRecorder is the mock recorder for MockDispatcher
type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

// NewMockDispatcher creates a new mock instance
func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

// DispatchMessage mocks base method
func (m *MockDispatcher) DispatchMessage(arg0 string, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "DispatchMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchMessage indicates an expected call of DispatchMessage
func (mr *MockDispatcherMockRecorder) DispatchMessage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchMessage", reflect.TypeOf((*MockDispatcher)(nil).DispatchMessage), arg0, arg1)
}