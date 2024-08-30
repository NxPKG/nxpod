// Copyright (c) 2023 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nxpkg/nxpod/ws-daemon/pkg/controller (interfaces: WorkspaceOperations)

// Package controller is a generated GoMock package.
package controller

import (
	context "context"
	reflect "reflect"

	api "github.com/nxpkg/nxpod/content-service/api"
	gomock "github.com/golang/mock/gomock"
)

// MockWorkspaceOperations is a mock of WorkspaceOperations interface.
type MockWorkspaceOperations struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceOperationsMockRecorder
}

// MockWorkspaceOperationsMockRecorder is the mock recorder for MockWorkspaceOperations.
type MockWorkspaceOperationsMockRecorder struct {
	mock *MockWorkspaceOperations
}

// NewMockWorkspaceOperations creates a new mock instance.
func NewMockWorkspaceOperations(ctrl *gomock.Controller) *MockWorkspaceOperations {
	mock := &MockWorkspaceOperations{ctrl: ctrl}
	mock.recorder = &MockWorkspaceOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceOperations) EXPECT() *MockWorkspaceOperationsMockRecorder {
	return m.recorder
}

// BackupWorkspace mocks base method.
func (m *MockWorkspaceOperations) BackupWorkspace(arg0 context.Context, arg1 BackupOptions) (*api.GitStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupWorkspace", arg0, arg1)
	ret0, _ := ret[0].(*api.GitStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackupWorkspace indicates an expected call of BackupWorkspace.
func (mr *MockWorkspaceOperationsMockRecorder) BackupWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupWorkspace", reflect.TypeOf((*MockWorkspaceOperations)(nil).BackupWorkspace), arg0, arg1)
}

// DeleteWorkspace mocks base method.
func (m *MockWorkspaceOperations) DeleteWorkspace(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkspace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkspace indicates an expected call of DeleteWorkspace.
func (mr *MockWorkspaceOperationsMockRecorder) DeleteWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkspace", reflect.TypeOf((*MockWorkspaceOperations)(nil).DeleteWorkspace), arg0, arg1)
}

// InitWorkspace mocks base method.
func (m *MockWorkspaceOperations) InitWorkspace(arg0 context.Context, arg1 InitOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitWorkspace", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitWorkspace indicates an expected call of InitWorkspace.
func (mr *MockWorkspaceOperationsMockRecorder) InitWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitWorkspace", reflect.TypeOf((*MockWorkspaceOperations)(nil).InitWorkspace), arg0, arg1)
}

// SetupWorkspace mocks base method.
func (m *MockWorkspaceOperations) SetupWorkspace(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupWorkspace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupWorkspace indicates an expected call of SetupWorkspace.
func (mr *MockWorkspaceOperationsMockRecorder) SetupWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupWorkspace", reflect.TypeOf((*MockWorkspaceOperations)(nil).SetupWorkspace), arg0, arg1)
}

// Snapshot mocks base method.
func (m *MockWorkspaceOperations) Snapshot(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockWorkspaceOperationsMockRecorder) Snapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockWorkspaceOperations)(nil).Snapshot), arg0, arg1, arg2)
}

// SnapshotIDs mocks base method.
func (m *MockWorkspaceOperations) SnapshotIDs(arg0 context.Context, arg1 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotIDs", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SnapshotIDs indicates an expected call of SnapshotIDs.
func (mr *MockWorkspaceOperationsMockRecorder) SnapshotIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotIDs", reflect.TypeOf((*MockWorkspaceOperations)(nil).SnapshotIDs), arg0, arg1)
}
