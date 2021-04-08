// Code generated by MockGen. DO NOT EDIT.
// Source: workspace.go

// Package mock_workspace is a generated GoMock package.
package mock_workspace

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mazrean/separated-webshell/domain"
)

// MockIWorkspace is a mock of IWorkspace interface.
type MockIWorkspace struct {
	ctrl     *gomock.Controller
	recorder *MockIWorkspaceMockRecorder
}

// MockIWorkspaceMockRecorder is the mock recorder for MockIWorkspace.
type MockIWorkspaceMockRecorder struct {
	mock *MockIWorkspace
}

// NewMockIWorkspace creates a new mock instance.
func NewMockIWorkspace(ctrl *gomock.Controller) *MockIWorkspace {
	mock := &MockIWorkspace{ctrl: ctrl}
	mock.recorder = &MockIWorkspaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWorkspace) EXPECT() *MockIWorkspaceMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockIWorkspace) Connect(ctx context.Context, userName string, isTty bool, winCh <-chan *domain.Window, stdin io.Reader, stdout, stderr io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx, userName, isTty, winCh, stdin, stdout, stderr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockIWorkspaceMockRecorder) Connect(ctx, userName, isTty, winCh, stdin, stdout, stderr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockIWorkspace)(nil).Connect), ctx, userName, isTty, winCh, stdin, stdout, stderr)
}

// Create mocks base method.
func (m *MockIWorkspace) Create(ctx context.Context, userName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIWorkspaceMockRecorder) Create(ctx, userName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIWorkspace)(nil).Create), ctx, userName)
}

// Remove mocks base method.
func (m *MockIWorkspace) Remove(ctx context.Context, userName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, userName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockIWorkspaceMockRecorder) Remove(ctx, userName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockIWorkspace)(nil).Remove), ctx, userName)
}
