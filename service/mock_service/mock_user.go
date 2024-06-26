// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	values "github.com/mazrean/separated-webshell/domain/values"
)

// MockIUser is a mock of IUser interface.
type MockIUser struct {
	ctrl     *gomock.Controller
	recorder *MockIUserMockRecorder
}

// MockIUserMockRecorder is the mock recorder for MockIUser.
type MockIUserMockRecorder struct {
	mock *MockIUser
}

// NewMockIUser creates a new mock instance.
func NewMockIUser(ctrl *gomock.Controller) *MockIUser {
	mock := &MockIUser{ctrl: ctrl}
	mock.recorder = &MockIUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUser) EXPECT() *MockIUserMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockIUser) Auth(ctx context.Context, name values.UserName, password values.Password) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", ctx, name, password)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth.
func (mr *MockIUserMockRecorder) Auth(ctx, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockIUser)(nil).Auth), ctx, name, password)
}

// New mocks base method.
func (m *MockIUser) New(ctx context.Context, name values.UserName, password values.Password) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", ctx, name, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockIUserMockRecorder) New(ctx, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockIUser)(nil).New), ctx, name, password)
}

// ResetContainer mocks base method.
func (m *MockIUser) ResetContainer(ctx context.Context, userName values.UserName) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetContainer", ctx, userName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetContainer indicates an expected call of ResetContainer.
func (mr *MockIUserMockRecorder) ResetContainer(ctx, userName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetContainer", reflect.TypeOf((*MockIUser)(nil).ResetContainer), ctx, userName)
}
