// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/v1/workspace.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/nxpkg/nxpod/components/public-api/go/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// WorkspaceServiceName is the fully-qualified name of the WorkspaceService service.
	WorkspaceServiceName = "gitpod.v1.WorkspaceService"
)

// WorkspaceServiceClient is a client for the gitpod.v1.WorkspaceService service.
type WorkspaceServiceClient interface {
	// GetWorkspace returns a single workspace.
	//
	// +return NOT_FOUND User does not have access to a workspace with the given
	// ID +return NOT_FOUND Workspace does not exist
	GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error)
	// WatchWorkspaceStatus watches the workspaces status changes
	//
	// workspace_id +return NOT_FOUND Workspace does not exist
	WatchWorkspaceStatus(context.Context, *connect_go.Request[v1.WatchWorkspaceStatusRequest]) (*connect_go.ServerStreamForClient[v1.WatchWorkspaceStatusResponse], error)
	// ListWorkspaces returns a list of workspaces that match the query.
	ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error)
	// ListWorkspaceSessions returns a list of workspace sessions that match the
	ListWorkspaceSessions(context.Context, *connect_go.Request[v1.ListWorkspaceSessionsRequest]) (*connect_go.Response[v1.ListWorkspaceSessionsResponse], error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error)
	// StartWorkspace starts an existing workspace.
	// If the specified workspace is not in stopped phase, this will return the
	// workspace as is.
	StartWorkspace(context.Context, *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error)
	// UpdateWorkspace updates the workspace.
	UpdateWorkspace(context.Context, *connect_go.Request[v1.UpdateWorkspaceRequest]) (*connect_go.Response[v1.UpdateWorkspaceResponse], error)
	// StopWorkspace stops a running workspace.
	StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.Response[v1.StopWorkspaceResponse], error)
	// DeleteWorkspace deletes a workspace.
	// When the workspace is running, it will be stopped as well.
	// Deleted workspaces cannot be started again.
	DeleteWorkspace(context.Context, *connect_go.Request[v1.DeleteWorkspaceRequest]) (*connect_go.Response[v1.DeleteWorkspaceResponse], error)
	// ListWorkspaceClasses enumerates all available workspace classes.
	ListWorkspaceClasses(context.Context, *connect_go.Request[v1.ListWorkspaceClassesRequest]) (*connect_go.Response[v1.ListWorkspaceClassesResponse], error)
	// ParseContextURL parses a context URL and returns the workspace metadata and
	// spec. Not implemented yet.
	ParseContextURL(context.Context, *connect_go.Request[v1.ParseContextURLRequest]) (*connect_go.Response[v1.ParseContextURLResponse], error)
	// GetWorkspaceDefaultImage returns the default workspace image of specified
	// workspace.
	GetWorkspaceDefaultImage(context.Context, *connect_go.Request[v1.GetWorkspaceDefaultImageRequest]) (*connect_go.Response[v1.GetWorkspaceDefaultImageResponse], error)
	// SendHeartBeat sends a heartbeat to activate the workspace
	SendHeartBeat(context.Context, *connect_go.Request[v1.SendHeartBeatRequest]) (*connect_go.Response[v1.SendHeartBeatResponse], error)
	// GetWorkspaceOwnerToken returns an owner token of workspace.
	GetWorkspaceOwnerToken(context.Context, *connect_go.Request[v1.GetWorkspaceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceOwnerTokenResponse], error)
	// GetWorkspaceEditorCredentials returns an credentials that is used in editor
	// to encrypt and decrypt secrets
	GetWorkspaceEditorCredentials(context.Context, *connect_go.Request[v1.GetWorkspaceEditorCredentialsRequest]) (*connect_go.Response[v1.GetWorkspaceEditorCredentialsResponse], error)
	// CreateWorkspaceSnapshot creates a snapshot of the workspace that can be
	// shared with others.
	CreateWorkspaceSnapshot(context.Context, *connect_go.Request[v1.CreateWorkspaceSnapshotRequest]) (*connect_go.Response[v1.CreateWorkspaceSnapshotResponse], error)
	// WaitWorkspaceSnapshot waits for the snapshot to be available or failed.
	WaitForWorkspaceSnapshot(context.Context, *connect_go.Request[v1.WaitForWorkspaceSnapshotRequest]) (*connect_go.Response[v1.WaitForWorkspaceSnapshotResponse], error)
	// UpdateWorkspacePort updates the port of workspace.
	UpdateWorkspacePort(context.Context, *connect_go.Request[v1.UpdateWorkspacePortRequest]) (*connect_go.Response[v1.UpdateWorkspacePortResponse], error)
}

// NewWorkspaceServiceClient constructs a client for the gitpod.v1.WorkspaceService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkspaceServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) WorkspaceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workspaceServiceClient{
		getWorkspace: connect_go.NewClient[v1.GetWorkspaceRequest, v1.GetWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/GetWorkspace",
			opts...,
		),
		watchWorkspaceStatus: connect_go.NewClient[v1.WatchWorkspaceStatusRequest, v1.WatchWorkspaceStatusResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/WatchWorkspaceStatus",
			opts...,
		),
		listWorkspaces: connect_go.NewClient[v1.ListWorkspacesRequest, v1.ListWorkspacesResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/ListWorkspaces",
			opts...,
		),
		listWorkspaceSessions: connect_go.NewClient[v1.ListWorkspaceSessionsRequest, v1.ListWorkspaceSessionsResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/ListWorkspaceSessions",
			opts...,
		),
		createAndStartWorkspace: connect_go.NewClient[v1.CreateAndStartWorkspaceRequest, v1.CreateAndStartWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/CreateAndStartWorkspace",
			opts...,
		),
		startWorkspace: connect_go.NewClient[v1.StartWorkspaceRequest, v1.StartWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/StartWorkspace",
			opts...,
		),
		updateWorkspace: connect_go.NewClient[v1.UpdateWorkspaceRequest, v1.UpdateWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/UpdateWorkspace",
			opts...,
		),
		stopWorkspace: connect_go.NewClient[v1.StopWorkspaceRequest, v1.StopWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/StopWorkspace",
			opts...,
		),
		deleteWorkspace: connect_go.NewClient[v1.DeleteWorkspaceRequest, v1.DeleteWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/DeleteWorkspace",
			opts...,
		),
		listWorkspaceClasses: connect_go.NewClient[v1.ListWorkspaceClassesRequest, v1.ListWorkspaceClassesResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/ListWorkspaceClasses",
			opts...,
		),
		parseContextURL: connect_go.NewClient[v1.ParseContextURLRequest, v1.ParseContextURLResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/ParseContextURL",
			opts...,
		),
		getWorkspaceDefaultImage: connect_go.NewClient[v1.GetWorkspaceDefaultImageRequest, v1.GetWorkspaceDefaultImageResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/GetWorkspaceDefaultImage",
			opts...,
		),
		sendHeartBeat: connect_go.NewClient[v1.SendHeartBeatRequest, v1.SendHeartBeatResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/SendHeartBeat",
			opts...,
		),
		getWorkspaceOwnerToken: connect_go.NewClient[v1.GetWorkspaceOwnerTokenRequest, v1.GetWorkspaceOwnerTokenResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/GetWorkspaceOwnerToken",
			opts...,
		),
		getWorkspaceEditorCredentials: connect_go.NewClient[v1.GetWorkspaceEditorCredentialsRequest, v1.GetWorkspaceEditorCredentialsResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/GetWorkspaceEditorCredentials",
			opts...,
		),
		createWorkspaceSnapshot: connect_go.NewClient[v1.CreateWorkspaceSnapshotRequest, v1.CreateWorkspaceSnapshotResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/CreateWorkspaceSnapshot",
			opts...,
		),
		waitForWorkspaceSnapshot: connect_go.NewClient[v1.WaitForWorkspaceSnapshotRequest, v1.WaitForWorkspaceSnapshotResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/WaitForWorkspaceSnapshot",
			opts...,
		),
		updateWorkspacePort: connect_go.NewClient[v1.UpdateWorkspacePortRequest, v1.UpdateWorkspacePortResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspaceService/UpdateWorkspacePort",
			opts...,
		),
	}
}

// workspaceServiceClient implements WorkspaceServiceClient.
type workspaceServiceClient struct {
	getWorkspace                  *connect_go.Client[v1.GetWorkspaceRequest, v1.GetWorkspaceResponse]
	watchWorkspaceStatus          *connect_go.Client[v1.WatchWorkspaceStatusRequest, v1.WatchWorkspaceStatusResponse]
	listWorkspaces                *connect_go.Client[v1.ListWorkspacesRequest, v1.ListWorkspacesResponse]
	listWorkspaceSessions         *connect_go.Client[v1.ListWorkspaceSessionsRequest, v1.ListWorkspaceSessionsResponse]
	createAndStartWorkspace       *connect_go.Client[v1.CreateAndStartWorkspaceRequest, v1.CreateAndStartWorkspaceResponse]
	startWorkspace                *connect_go.Client[v1.StartWorkspaceRequest, v1.StartWorkspaceResponse]
	updateWorkspace               *connect_go.Client[v1.UpdateWorkspaceRequest, v1.UpdateWorkspaceResponse]
	stopWorkspace                 *connect_go.Client[v1.StopWorkspaceRequest, v1.StopWorkspaceResponse]
	deleteWorkspace               *connect_go.Client[v1.DeleteWorkspaceRequest, v1.DeleteWorkspaceResponse]
	listWorkspaceClasses          *connect_go.Client[v1.ListWorkspaceClassesRequest, v1.ListWorkspaceClassesResponse]
	parseContextURL               *connect_go.Client[v1.ParseContextURLRequest, v1.ParseContextURLResponse]
	getWorkspaceDefaultImage      *connect_go.Client[v1.GetWorkspaceDefaultImageRequest, v1.GetWorkspaceDefaultImageResponse]
	sendHeartBeat                 *connect_go.Client[v1.SendHeartBeatRequest, v1.SendHeartBeatResponse]
	getWorkspaceOwnerToken        *connect_go.Client[v1.GetWorkspaceOwnerTokenRequest, v1.GetWorkspaceOwnerTokenResponse]
	getWorkspaceEditorCredentials *connect_go.Client[v1.GetWorkspaceEditorCredentialsRequest, v1.GetWorkspaceEditorCredentialsResponse]
	createWorkspaceSnapshot       *connect_go.Client[v1.CreateWorkspaceSnapshotRequest, v1.CreateWorkspaceSnapshotResponse]
	waitForWorkspaceSnapshot      *connect_go.Client[v1.WaitForWorkspaceSnapshotRequest, v1.WaitForWorkspaceSnapshotResponse]
	updateWorkspacePort           *connect_go.Client[v1.UpdateWorkspacePortRequest, v1.UpdateWorkspacePortResponse]
}

// GetWorkspace calls gitpod.v1.WorkspaceService.GetWorkspace.
func (c *workspaceServiceClient) GetWorkspace(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	return c.getWorkspace.CallUnary(ctx, req)
}

// WatchWorkspaceStatus calls gitpod.v1.WorkspaceService.WatchWorkspaceStatus.
func (c *workspaceServiceClient) WatchWorkspaceStatus(ctx context.Context, req *connect_go.Request[v1.WatchWorkspaceStatusRequest]) (*connect_go.ServerStreamForClient[v1.WatchWorkspaceStatusResponse], error) {
	return c.watchWorkspaceStatus.CallServerStream(ctx, req)
}

// ListWorkspaces calls gitpod.v1.WorkspaceService.ListWorkspaces.
func (c *workspaceServiceClient) ListWorkspaces(ctx context.Context, req *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	return c.listWorkspaces.CallUnary(ctx, req)
}

// ListWorkspaceSessions calls gitpod.v1.WorkspaceService.ListWorkspaceSessions.
func (c *workspaceServiceClient) ListWorkspaceSessions(ctx context.Context, req *connect_go.Request[v1.ListWorkspaceSessionsRequest]) (*connect_go.Response[v1.ListWorkspaceSessionsResponse], error) {
	return c.listWorkspaceSessions.CallUnary(ctx, req)
}

// CreateAndStartWorkspace calls gitpod.v1.WorkspaceService.CreateAndStartWorkspace.
func (c *workspaceServiceClient) CreateAndStartWorkspace(ctx context.Context, req *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	return c.createAndStartWorkspace.CallUnary(ctx, req)
}

// StartWorkspace calls gitpod.v1.WorkspaceService.StartWorkspace.
func (c *workspaceServiceClient) StartWorkspace(ctx context.Context, req *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error) {
	return c.startWorkspace.CallUnary(ctx, req)
}

// UpdateWorkspace calls gitpod.v1.WorkspaceService.UpdateWorkspace.
func (c *workspaceServiceClient) UpdateWorkspace(ctx context.Context, req *connect_go.Request[v1.UpdateWorkspaceRequest]) (*connect_go.Response[v1.UpdateWorkspaceResponse], error) {
	return c.updateWorkspace.CallUnary(ctx, req)
}

// StopWorkspace calls gitpod.v1.WorkspaceService.StopWorkspace.
func (c *workspaceServiceClient) StopWorkspace(ctx context.Context, req *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.Response[v1.StopWorkspaceResponse], error) {
	return c.stopWorkspace.CallUnary(ctx, req)
}

// DeleteWorkspace calls gitpod.v1.WorkspaceService.DeleteWorkspace.
func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, req *connect_go.Request[v1.DeleteWorkspaceRequest]) (*connect_go.Response[v1.DeleteWorkspaceResponse], error) {
	return c.deleteWorkspace.CallUnary(ctx, req)
}

// ListWorkspaceClasses calls gitpod.v1.WorkspaceService.ListWorkspaceClasses.
func (c *workspaceServiceClient) ListWorkspaceClasses(ctx context.Context, req *connect_go.Request[v1.ListWorkspaceClassesRequest]) (*connect_go.Response[v1.ListWorkspaceClassesResponse], error) {
	return c.listWorkspaceClasses.CallUnary(ctx, req)
}

// ParseContextURL calls gitpod.v1.WorkspaceService.ParseContextURL.
func (c *workspaceServiceClient) ParseContextURL(ctx context.Context, req *connect_go.Request[v1.ParseContextURLRequest]) (*connect_go.Response[v1.ParseContextURLResponse], error) {
	return c.parseContextURL.CallUnary(ctx, req)
}

// GetWorkspaceDefaultImage calls gitpod.v1.WorkspaceService.GetWorkspaceDefaultImage.
func (c *workspaceServiceClient) GetWorkspaceDefaultImage(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceDefaultImageRequest]) (*connect_go.Response[v1.GetWorkspaceDefaultImageResponse], error) {
	return c.getWorkspaceDefaultImage.CallUnary(ctx, req)
}

// SendHeartBeat calls gitpod.v1.WorkspaceService.SendHeartBeat.
func (c *workspaceServiceClient) SendHeartBeat(ctx context.Context, req *connect_go.Request[v1.SendHeartBeatRequest]) (*connect_go.Response[v1.SendHeartBeatResponse], error) {
	return c.sendHeartBeat.CallUnary(ctx, req)
}

// GetWorkspaceOwnerToken calls gitpod.v1.WorkspaceService.GetWorkspaceOwnerToken.
func (c *workspaceServiceClient) GetWorkspaceOwnerToken(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceOwnerTokenResponse], error) {
	return c.getWorkspaceOwnerToken.CallUnary(ctx, req)
}

// GetWorkspaceEditorCredentials calls gitpod.v1.WorkspaceService.GetWorkspaceEditorCredentials.
func (c *workspaceServiceClient) GetWorkspaceEditorCredentials(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceEditorCredentialsRequest]) (*connect_go.Response[v1.GetWorkspaceEditorCredentialsResponse], error) {
	return c.getWorkspaceEditorCredentials.CallUnary(ctx, req)
}

// CreateWorkspaceSnapshot calls gitpod.v1.WorkspaceService.CreateWorkspaceSnapshot.
func (c *workspaceServiceClient) CreateWorkspaceSnapshot(ctx context.Context, req *connect_go.Request[v1.CreateWorkspaceSnapshotRequest]) (*connect_go.Response[v1.CreateWorkspaceSnapshotResponse], error) {
	return c.createWorkspaceSnapshot.CallUnary(ctx, req)
}

// WaitForWorkspaceSnapshot calls gitpod.v1.WorkspaceService.WaitForWorkspaceSnapshot.
func (c *workspaceServiceClient) WaitForWorkspaceSnapshot(ctx context.Context, req *connect_go.Request[v1.WaitForWorkspaceSnapshotRequest]) (*connect_go.Response[v1.WaitForWorkspaceSnapshotResponse], error) {
	return c.waitForWorkspaceSnapshot.CallUnary(ctx, req)
}

// UpdateWorkspacePort calls gitpod.v1.WorkspaceService.UpdateWorkspacePort.
func (c *workspaceServiceClient) UpdateWorkspacePort(ctx context.Context, req *connect_go.Request[v1.UpdateWorkspacePortRequest]) (*connect_go.Response[v1.UpdateWorkspacePortResponse], error) {
	return c.updateWorkspacePort.CallUnary(ctx, req)
}

// WorkspaceServiceHandler is an implementation of the gitpod.v1.WorkspaceService service.
type WorkspaceServiceHandler interface {
	// GetWorkspace returns a single workspace.
	//
	// +return NOT_FOUND User does not have access to a workspace with the given
	// ID +return NOT_FOUND Workspace does not exist
	GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error)
	// WatchWorkspaceStatus watches the workspaces status changes
	//
	// workspace_id +return NOT_FOUND Workspace does not exist
	WatchWorkspaceStatus(context.Context, *connect_go.Request[v1.WatchWorkspaceStatusRequest], *connect_go.ServerStream[v1.WatchWorkspaceStatusResponse]) error
	// ListWorkspaces returns a list of workspaces that match the query.
	ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error)
	// ListWorkspaceSessions returns a list of workspace sessions that match the
	ListWorkspaceSessions(context.Context, *connect_go.Request[v1.ListWorkspaceSessionsRequest]) (*connect_go.Response[v1.ListWorkspaceSessionsResponse], error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error)
	// StartWorkspace starts an existing workspace.
	// If the specified workspace is not in stopped phase, this will return the
	// workspace as is.
	StartWorkspace(context.Context, *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error)
	// UpdateWorkspace updates the workspace.
	UpdateWorkspace(context.Context, *connect_go.Request[v1.UpdateWorkspaceRequest]) (*connect_go.Response[v1.UpdateWorkspaceResponse], error)
	// StopWorkspace stops a running workspace.
	StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.Response[v1.StopWorkspaceResponse], error)
	// DeleteWorkspace deletes a workspace.
	// When the workspace is running, it will be stopped as well.
	// Deleted workspaces cannot be started again.
	DeleteWorkspace(context.Context, *connect_go.Request[v1.DeleteWorkspaceRequest]) (*connect_go.Response[v1.DeleteWorkspaceResponse], error)
	// ListWorkspaceClasses enumerates all available workspace classes.
	ListWorkspaceClasses(context.Context, *connect_go.Request[v1.ListWorkspaceClassesRequest]) (*connect_go.Response[v1.ListWorkspaceClassesResponse], error)
	// ParseContextURL parses a context URL and returns the workspace metadata and
	// spec. Not implemented yet.
	ParseContextURL(context.Context, *connect_go.Request[v1.ParseContextURLRequest]) (*connect_go.Response[v1.ParseContextURLResponse], error)
	// GetWorkspaceDefaultImage returns the default workspace image of specified
	// workspace.
	GetWorkspaceDefaultImage(context.Context, *connect_go.Request[v1.GetWorkspaceDefaultImageRequest]) (*connect_go.Response[v1.GetWorkspaceDefaultImageResponse], error)
	// SendHeartBeat sends a heartbeat to activate the workspace
	SendHeartBeat(context.Context, *connect_go.Request[v1.SendHeartBeatRequest]) (*connect_go.Response[v1.SendHeartBeatResponse], error)
	// GetWorkspaceOwnerToken returns an owner token of workspace.
	GetWorkspaceOwnerToken(context.Context, *connect_go.Request[v1.GetWorkspaceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceOwnerTokenResponse], error)
	// GetWorkspaceEditorCredentials returns an credentials that is used in editor
	// to encrypt and decrypt secrets
	GetWorkspaceEditorCredentials(context.Context, *connect_go.Request[v1.GetWorkspaceEditorCredentialsRequest]) (*connect_go.Response[v1.GetWorkspaceEditorCredentialsResponse], error)
	// CreateWorkspaceSnapshot creates a snapshot of the workspace that can be
	// shared with others.
	CreateWorkspaceSnapshot(context.Context, *connect_go.Request[v1.CreateWorkspaceSnapshotRequest]) (*connect_go.Response[v1.CreateWorkspaceSnapshotResponse], error)
	// WaitWorkspaceSnapshot waits for the snapshot to be available or failed.
	WaitForWorkspaceSnapshot(context.Context, *connect_go.Request[v1.WaitForWorkspaceSnapshotRequest]) (*connect_go.Response[v1.WaitForWorkspaceSnapshotResponse], error)
	// UpdateWorkspacePort updates the port of workspace.
	UpdateWorkspacePort(context.Context, *connect_go.Request[v1.UpdateWorkspacePortRequest]) (*connect_go.Response[v1.UpdateWorkspacePortResponse], error)
}

// NewWorkspaceServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkspaceServiceHandler(svc WorkspaceServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.v1.WorkspaceService/GetWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/GetWorkspace",
		svc.GetWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/WatchWorkspaceStatus", connect_go.NewServerStreamHandler(
		"/gitpod.v1.WorkspaceService/WatchWorkspaceStatus",
		svc.WatchWorkspaceStatus,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/ListWorkspaces", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/ListWorkspaces",
		svc.ListWorkspaces,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/ListWorkspaceSessions", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/ListWorkspaceSessions",
		svc.ListWorkspaceSessions,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/CreateAndStartWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/CreateAndStartWorkspace",
		svc.CreateAndStartWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/StartWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/StartWorkspace",
		svc.StartWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/UpdateWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/UpdateWorkspace",
		svc.UpdateWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/StopWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/StopWorkspace",
		svc.StopWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/DeleteWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/DeleteWorkspace",
		svc.DeleteWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/ListWorkspaceClasses", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/ListWorkspaceClasses",
		svc.ListWorkspaceClasses,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/ParseContextURL", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/ParseContextURL",
		svc.ParseContextURL,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/GetWorkspaceDefaultImage", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/GetWorkspaceDefaultImage",
		svc.GetWorkspaceDefaultImage,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/SendHeartBeat", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/SendHeartBeat",
		svc.SendHeartBeat,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/GetWorkspaceOwnerToken", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/GetWorkspaceOwnerToken",
		svc.GetWorkspaceOwnerToken,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/GetWorkspaceEditorCredentials", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/GetWorkspaceEditorCredentials",
		svc.GetWorkspaceEditorCredentials,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/CreateWorkspaceSnapshot", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/CreateWorkspaceSnapshot",
		svc.CreateWorkspaceSnapshot,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/WaitForWorkspaceSnapshot", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/WaitForWorkspaceSnapshot",
		svc.WaitForWorkspaceSnapshot,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspaceService/UpdateWorkspacePort", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspaceService/UpdateWorkspacePort",
		svc.UpdateWorkspacePort,
		opts...,
	))
	return "/gitpod.v1.WorkspaceService/", mux
}

// UnimplementedWorkspaceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkspaceServiceHandler struct{}

func (UnimplementedWorkspaceServiceHandler) GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.GetWorkspace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) WatchWorkspaceStatus(context.Context, *connect_go.Request[v1.WatchWorkspaceStatusRequest], *connect_go.ServerStream[v1.WatchWorkspaceStatusResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.WatchWorkspaceStatus is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.ListWorkspaces is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) ListWorkspaceSessions(context.Context, *connect_go.Request[v1.ListWorkspaceSessionsRequest]) (*connect_go.Response[v1.ListWorkspaceSessionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.ListWorkspaceSessions is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.CreateAndStartWorkspace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) StartWorkspace(context.Context, *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.StartWorkspace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) UpdateWorkspace(context.Context, *connect_go.Request[v1.UpdateWorkspaceRequest]) (*connect_go.Response[v1.UpdateWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.UpdateWorkspace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.Response[v1.StopWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.StopWorkspace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) DeleteWorkspace(context.Context, *connect_go.Request[v1.DeleteWorkspaceRequest]) (*connect_go.Response[v1.DeleteWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.DeleteWorkspace is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) ListWorkspaceClasses(context.Context, *connect_go.Request[v1.ListWorkspaceClassesRequest]) (*connect_go.Response[v1.ListWorkspaceClassesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.ListWorkspaceClasses is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) ParseContextURL(context.Context, *connect_go.Request[v1.ParseContextURLRequest]) (*connect_go.Response[v1.ParseContextURLResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.ParseContextURL is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaceDefaultImage(context.Context, *connect_go.Request[v1.GetWorkspaceDefaultImageRequest]) (*connect_go.Response[v1.GetWorkspaceDefaultImageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.GetWorkspaceDefaultImage is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) SendHeartBeat(context.Context, *connect_go.Request[v1.SendHeartBeatRequest]) (*connect_go.Response[v1.SendHeartBeatResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.SendHeartBeat is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaceOwnerToken(context.Context, *connect_go.Request[v1.GetWorkspaceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceOwnerTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.GetWorkspaceOwnerToken is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaceEditorCredentials(context.Context, *connect_go.Request[v1.GetWorkspaceEditorCredentialsRequest]) (*connect_go.Response[v1.GetWorkspaceEditorCredentialsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.GetWorkspaceEditorCredentials is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) CreateWorkspaceSnapshot(context.Context, *connect_go.Request[v1.CreateWorkspaceSnapshotRequest]) (*connect_go.Response[v1.CreateWorkspaceSnapshotResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.CreateWorkspaceSnapshot is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) WaitForWorkspaceSnapshot(context.Context, *connect_go.Request[v1.WaitForWorkspaceSnapshotRequest]) (*connect_go.Response[v1.WaitForWorkspaceSnapshotResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.WaitForWorkspaceSnapshot is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) UpdateWorkspacePort(context.Context, *connect_go.Request[v1.UpdateWorkspacePortRequest]) (*connect_go.Response[v1.UpdateWorkspacePortResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspaceService.UpdateWorkspacePort is not implemented"))
}
