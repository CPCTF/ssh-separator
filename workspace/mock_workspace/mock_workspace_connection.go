// Code generated by MockGen. DO NOT EDIT.
// Source: workspace_connection.go

// Package mock_workspace is a generated GoMock package.
package mock_workspace

import (
	context "context"
	reflect "reflect"

	domain "github.com/CPCTF2022/ssh-separator/domain"
	values "github.com/CPCTF2022/ssh-separator/domain/values"
	gomock "github.com/golang/mock/gomock"
)

// MockIWorkspaceConnection is a mock of IWorkspaceConnection interface.
type MockIWorkspaceConnection struct {
	ctrl     *gomock.Controller
	recorder *MockIWorkspaceConnectionMockRecorder
}

// MockIWorkspaceConnectionMockRecorder is the mock recorder for MockIWorkspaceConnection.
type MockIWorkspaceConnectionMockRecorder struct {
	mock *MockIWorkspaceConnection
}

// NewMockIWorkspaceConnection creates a new mock instance.
func NewMockIWorkspaceConnection(ctrl *gomock.Controller) *MockIWorkspaceConnection {
	mock := &MockIWorkspaceConnection{ctrl: ctrl}
	mock.recorder = &MockIWorkspaceConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWorkspaceConnection) EXPECT() *MockIWorkspaceConnectionMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockIWorkspaceConnection) Connect(ctx context.Context, workspace *domain.Workspace) (*domain.WorkspaceConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx, workspace)
	ret0, _ := ret[0].(*domain.WorkspaceConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockIWorkspaceConnectionMockRecorder) Connect(ctx, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockIWorkspaceConnection)(nil).Connect), ctx, workspace)
}

// Disconnect mocks base method.
func (m *MockIWorkspaceConnection) Disconnect(ctx context.Context, connection *domain.WorkspaceConnection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect", ctx, connection)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockIWorkspaceConnectionMockRecorder) Disconnect(ctx, connection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockIWorkspaceConnection)(nil).Disconnect), ctx, connection)
}

// Resize mocks base method.
func (m *MockIWorkspaceConnection) Resize(ctx context.Context, connection *domain.WorkspaceConnection, window *values.Window) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", ctx, connection, window)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resize indicates an expected call of Resize.
func (mr *MockIWorkspaceConnectionMockRecorder) Resize(ctx, connection, window interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockIWorkspaceConnection)(nil).Resize), ctx, connection, window)
}
