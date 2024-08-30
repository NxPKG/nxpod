// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nxpkg/nxpod/image-builder/api (interfaces: ImageBuilderClient,ImageBuilder_BuildClient,ImageBuilder_LogsClient,ImageBuilderServer,ImageBuilder_BuildServer,ImageBuilder_LogsServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	api "github.com/nxpkg/nxpod/image-builder/api"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockImageBuilderClient is a mock of ImageBuilderClient interface.
type MockImageBuilderClient struct {
	ctrl     *gomock.Controller
	recorder *MockImageBuilderClientMockRecorder
}

// MockImageBuilderClientMockRecorder is the mock recorder for MockImageBuilderClient.
type MockImageBuilderClientMockRecorder struct {
	mock *MockImageBuilderClient
}

// NewMockImageBuilderClient creates a new mock instance.
func NewMockImageBuilderClient(ctrl *gomock.Controller) *MockImageBuilderClient {
	mock := &MockImageBuilderClient{ctrl: ctrl}
	mock.recorder = &MockImageBuilderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageBuilderClient) EXPECT() *MockImageBuilderClientMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockImageBuilderClient) Build(arg0 context.Context, arg1 *api.BuildRequest, arg2 ...grpc.CallOption) (api.ImageBuilder_BuildClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Build", varargs...)
	ret0, _ := ret[0].(api.ImageBuilder_BuildClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockImageBuilderClientMockRecorder) Build(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockImageBuilderClient)(nil).Build), varargs...)
}

// ListBuilds mocks base method.
func (m *MockImageBuilderClient) ListBuilds(arg0 context.Context, arg1 *api.ListBuildsRequest, arg2 ...grpc.CallOption) (*api.ListBuildsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuilds", varargs...)
	ret0, _ := ret[0].(*api.ListBuildsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuilds indicates an expected call of ListBuilds.
func (mr *MockImageBuilderClientMockRecorder) ListBuilds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuilds", reflect.TypeOf((*MockImageBuilderClient)(nil).ListBuilds), varargs...)
}

// Logs mocks base method.
func (m *MockImageBuilderClient) Logs(arg0 context.Context, arg1 *api.LogsRequest, arg2 ...grpc.CallOption) (api.ImageBuilder_LogsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Logs", varargs...)
	ret0, _ := ret[0].(api.ImageBuilder_LogsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockImageBuilderClientMockRecorder) Logs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockImageBuilderClient)(nil).Logs), varargs...)
}

// ResolveBaseImage mocks base method.
func (m *MockImageBuilderClient) ResolveBaseImage(arg0 context.Context, arg1 *api.ResolveBaseImageRequest, arg2 ...grpc.CallOption) (*api.ResolveBaseImageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResolveBaseImage", varargs...)
	ret0, _ := ret[0].(*api.ResolveBaseImageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveBaseImage indicates an expected call of ResolveBaseImage.
func (mr *MockImageBuilderClientMockRecorder) ResolveBaseImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBaseImage", reflect.TypeOf((*MockImageBuilderClient)(nil).ResolveBaseImage), varargs...)
}

// ResolveWorkspaceImage mocks base method.
func (m *MockImageBuilderClient) ResolveWorkspaceImage(arg0 context.Context, arg1 *api.ResolveWorkspaceImageRequest, arg2 ...grpc.CallOption) (*api.ResolveWorkspaceImageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResolveWorkspaceImage", varargs...)
	ret0, _ := ret[0].(*api.ResolveWorkspaceImageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveWorkspaceImage indicates an expected call of ResolveWorkspaceImage.
func (mr *MockImageBuilderClientMockRecorder) ResolveWorkspaceImage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveWorkspaceImage", reflect.TypeOf((*MockImageBuilderClient)(nil).ResolveWorkspaceImage), varargs...)
}

// MockImageBuilder_BuildClient is a mock of ImageBuilder_BuildClient interface.
type MockImageBuilder_BuildClient struct {
	ctrl     *gomock.Controller
	recorder *MockImageBuilder_BuildClientMockRecorder
}

// MockImageBuilder_BuildClientMockRecorder is the mock recorder for MockImageBuilder_BuildClient.
type MockImageBuilder_BuildClientMockRecorder struct {
	mock *MockImageBuilder_BuildClient
}

// NewMockImageBuilder_BuildClient creates a new mock instance.
func NewMockImageBuilder_BuildClient(ctrl *gomock.Controller) *MockImageBuilder_BuildClient {
	mock := &MockImageBuilder_BuildClient{ctrl: ctrl}
	mock.recorder = &MockImageBuilder_BuildClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageBuilder_BuildClient) EXPECT() *MockImageBuilder_BuildClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockImageBuilder_BuildClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockImageBuilder_BuildClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockImageBuilder_BuildClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockImageBuilder_BuildClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockImageBuilder_BuildClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockImageBuilder_BuildClient)(nil).Context))
}

// Header mocks base method.
func (m *MockImageBuilder_BuildClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockImageBuilder_BuildClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockImageBuilder_BuildClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockImageBuilder_BuildClient) Recv() (*api.BuildResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*api.BuildResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockImageBuilder_BuildClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockImageBuilder_BuildClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockImageBuilder_BuildClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockImageBuilder_BuildClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockImageBuilder_BuildClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockImageBuilder_BuildClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockImageBuilder_BuildClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockImageBuilder_BuildClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockImageBuilder_BuildClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockImageBuilder_BuildClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockImageBuilder_BuildClient)(nil).Trailer))
}

// MockImageBuilder_LogsClient is a mock of ImageBuilder_LogsClient interface.
type MockImageBuilder_LogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockImageBuilder_LogsClientMockRecorder
}

// MockImageBuilder_LogsClientMockRecorder is the mock recorder for MockImageBuilder_LogsClient.
type MockImageBuilder_LogsClientMockRecorder struct {
	mock *MockImageBuilder_LogsClient
}

// NewMockImageBuilder_LogsClient creates a new mock instance.
func NewMockImageBuilder_LogsClient(ctrl *gomock.Controller) *MockImageBuilder_LogsClient {
	mock := &MockImageBuilder_LogsClient{ctrl: ctrl}
	mock.recorder = &MockImageBuilder_LogsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageBuilder_LogsClient) EXPECT() *MockImageBuilder_LogsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockImageBuilder_LogsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockImageBuilder_LogsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockImageBuilder_LogsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockImageBuilder_LogsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockImageBuilder_LogsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockImageBuilder_LogsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockImageBuilder_LogsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockImageBuilder_LogsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockImageBuilder_LogsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockImageBuilder_LogsClient) Recv() (*api.LogsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*api.LogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockImageBuilder_LogsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockImageBuilder_LogsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockImageBuilder_LogsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockImageBuilder_LogsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockImageBuilder_LogsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockImageBuilder_LogsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockImageBuilder_LogsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockImageBuilder_LogsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockImageBuilder_LogsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockImageBuilder_LogsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockImageBuilder_LogsClient)(nil).Trailer))
}

// MockImageBuilderServer is a mock of ImageBuilderServer interface.
type MockImageBuilderServer struct {
	ctrl     *gomock.Controller
	recorder *MockImageBuilderServerMockRecorder
}

// MockImageBuilderServerMockRecorder is the mock recorder for MockImageBuilderServer.
type MockImageBuilderServerMockRecorder struct {
	mock *MockImageBuilderServer
}

// NewMockImageBuilderServer creates a new mock instance.
func NewMockImageBuilderServer(ctrl *gomock.Controller) *MockImageBuilderServer {
	mock := &MockImageBuilderServer{ctrl: ctrl}
	mock.recorder = &MockImageBuilderServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageBuilderServer) EXPECT() *MockImageBuilderServerMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockImageBuilderServer) Build(arg0 *api.BuildRequest, arg1 api.ImageBuilder_BuildServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockImageBuilderServerMockRecorder) Build(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockImageBuilderServer)(nil).Build), arg0, arg1)
}

// ListBuilds mocks base method.
func (m *MockImageBuilderServer) ListBuilds(arg0 context.Context, arg1 *api.ListBuildsRequest) (*api.ListBuildsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuilds", arg0, arg1)
	ret0, _ := ret[0].(*api.ListBuildsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuilds indicates an expected call of ListBuilds.
func (mr *MockImageBuilderServerMockRecorder) ListBuilds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuilds", reflect.TypeOf((*MockImageBuilderServer)(nil).ListBuilds), arg0, arg1)
}

// Logs mocks base method.
func (m *MockImageBuilderServer) Logs(arg0 *api.LogsRequest, arg1 api.ImageBuilder_LogsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logs indicates an expected call of Logs.
func (mr *MockImageBuilderServerMockRecorder) Logs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockImageBuilderServer)(nil).Logs), arg0, arg1)
}

// ResolveBaseImage mocks base method.
func (m *MockImageBuilderServer) ResolveBaseImage(arg0 context.Context, arg1 *api.ResolveBaseImageRequest) (*api.ResolveBaseImageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveBaseImage", arg0, arg1)
	ret0, _ := ret[0].(*api.ResolveBaseImageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveBaseImage indicates an expected call of ResolveBaseImage.
func (mr *MockImageBuilderServerMockRecorder) ResolveBaseImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBaseImage", reflect.TypeOf((*MockImageBuilderServer)(nil).ResolveBaseImage), arg0, arg1)
}

// ResolveWorkspaceImage mocks base method.
func (m *MockImageBuilderServer) ResolveWorkspaceImage(arg0 context.Context, arg1 *api.ResolveWorkspaceImageRequest) (*api.ResolveWorkspaceImageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveWorkspaceImage", arg0, arg1)
	ret0, _ := ret[0].(*api.ResolveWorkspaceImageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveWorkspaceImage indicates an expected call of ResolveWorkspaceImage.
func (mr *MockImageBuilderServerMockRecorder) ResolveWorkspaceImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveWorkspaceImage", reflect.TypeOf((*MockImageBuilderServer)(nil).ResolveWorkspaceImage), arg0, arg1)
}

// mustEmbedUnimplementedImageBuilderServer mocks base method.
func (m *MockImageBuilderServer) mustEmbedUnimplementedImageBuilderServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedImageBuilderServer")
}

// mustEmbedUnimplementedImageBuilderServer indicates an expected call of mustEmbedUnimplementedImageBuilderServer.
func (mr *MockImageBuilderServerMockRecorder) mustEmbedUnimplementedImageBuilderServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedImageBuilderServer", reflect.TypeOf((*MockImageBuilderServer)(nil).mustEmbedUnimplementedImageBuilderServer))
}

// MockImageBuilder_BuildServer is a mock of ImageBuilder_BuildServer interface.
type MockImageBuilder_BuildServer struct {
	ctrl     *gomock.Controller
	recorder *MockImageBuilder_BuildServerMockRecorder
}

// MockImageBuilder_BuildServerMockRecorder is the mock recorder for MockImageBuilder_BuildServer.
type MockImageBuilder_BuildServerMockRecorder struct {
	mock *MockImageBuilder_BuildServer
}

// NewMockImageBuilder_BuildServer creates a new mock instance.
func NewMockImageBuilder_BuildServer(ctrl *gomock.Controller) *MockImageBuilder_BuildServer {
	mock := &MockImageBuilder_BuildServer{ctrl: ctrl}
	mock.recorder = &MockImageBuilder_BuildServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageBuilder_BuildServer) EXPECT() *MockImageBuilder_BuildServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockImageBuilder_BuildServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockImageBuilder_BuildServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockImageBuilder_BuildServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockImageBuilder_BuildServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockImageBuilder_BuildServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockImageBuilder_BuildServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockImageBuilder_BuildServer) Send(arg0 *api.BuildResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockImageBuilder_BuildServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockImageBuilder_BuildServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockImageBuilder_BuildServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockImageBuilder_BuildServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockImageBuilder_BuildServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockImageBuilder_BuildServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockImageBuilder_BuildServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockImageBuilder_BuildServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockImageBuilder_BuildServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockImageBuilder_BuildServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockImageBuilder_BuildServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockImageBuilder_BuildServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockImageBuilder_BuildServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockImageBuilder_BuildServer)(nil).SetTrailer), arg0)
}

// MockImageBuilder_LogsServer is a mock of ImageBuilder_LogsServer interface.
type MockImageBuilder_LogsServer struct {
	ctrl     *gomock.Controller
	recorder *MockImageBuilder_LogsServerMockRecorder
}

// MockImageBuilder_LogsServerMockRecorder is the mock recorder for MockImageBuilder_LogsServer.
type MockImageBuilder_LogsServerMockRecorder struct {
	mock *MockImageBuilder_LogsServer
}

// NewMockImageBuilder_LogsServer creates a new mock instance.
func NewMockImageBuilder_LogsServer(ctrl *gomock.Controller) *MockImageBuilder_LogsServer {
	mock := &MockImageBuilder_LogsServer{ctrl: ctrl}
	mock.recorder = &MockImageBuilder_LogsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageBuilder_LogsServer) EXPECT() *MockImageBuilder_LogsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockImageBuilder_LogsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockImageBuilder_LogsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockImageBuilder_LogsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockImageBuilder_LogsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockImageBuilder_LogsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockImageBuilder_LogsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockImageBuilder_LogsServer) Send(arg0 *api.LogsResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockImageBuilder_LogsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockImageBuilder_LogsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockImageBuilder_LogsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockImageBuilder_LogsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockImageBuilder_LogsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockImageBuilder_LogsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockImageBuilder_LogsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockImageBuilder_LogsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockImageBuilder_LogsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockImageBuilder_LogsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockImageBuilder_LogsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockImageBuilder_LogsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockImageBuilder_LogsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockImageBuilder_LogsServer)(nil).SetTrailer), arg0)
}
