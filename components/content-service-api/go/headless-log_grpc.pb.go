// Copyright (c) 2023 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: headless-log.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HeadlessLogServiceClient is the client API for HeadlessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeadlessLogServiceClient interface {
	// LogDownloadURL provides a URL from where the content of a workspace can be downloaded from
	LogDownloadURL(ctx context.Context, in *LogDownloadURLRequest, opts ...grpc.CallOption) (*LogDownloadURLResponse, error)
	// ListLogs returns a list of taskIds for the specified workspace instance
	ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error)
}

type headlessLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeadlessLogServiceClient(cc grpc.ClientConnInterface) HeadlessLogServiceClient {
	return &headlessLogServiceClient{cc}
}

func (c *headlessLogServiceClient) LogDownloadURL(ctx context.Context, in *LogDownloadURLRequest, opts ...grpc.CallOption) (*LogDownloadURLResponse, error) {
	out := new(LogDownloadURLResponse)
	err := c.cc.Invoke(ctx, "/contentservice.HeadlessLogService/LogDownloadURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headlessLogServiceClient) ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error) {
	out := new(ListLogsResponse)
	err := c.cc.Invoke(ctx, "/contentservice.HeadlessLogService/ListLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeadlessLogServiceServer is the server API for HeadlessLogService service.
// All implementations must embed UnimplementedHeadlessLogServiceServer
// for forward compatibility
type HeadlessLogServiceServer interface {
	// LogDownloadURL provides a URL from where the content of a workspace can be downloaded from
	LogDownloadURL(context.Context, *LogDownloadURLRequest) (*LogDownloadURLResponse, error)
	// ListLogs returns a list of taskIds for the specified workspace instance
	ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error)
	mustEmbedUnimplementedHeadlessLogServiceServer()
}

// UnimplementedHeadlessLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeadlessLogServiceServer struct {
}

func (UnimplementedHeadlessLogServiceServer) LogDownloadURL(context.Context, *LogDownloadURLRequest) (*LogDownloadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogDownloadURL not implemented")
}
func (UnimplementedHeadlessLogServiceServer) ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogs not implemented")
}
func (UnimplementedHeadlessLogServiceServer) mustEmbedUnimplementedHeadlessLogServiceServer() {}

// UnsafeHeadlessLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeadlessLogServiceServer will
// result in compilation errors.
type UnsafeHeadlessLogServiceServer interface {
	mustEmbedUnimplementedHeadlessLogServiceServer()
}

func RegisterHeadlessLogServiceServer(s grpc.ServiceRegistrar, srv HeadlessLogServiceServer) {
	s.RegisterService(&HeadlessLogService_ServiceDesc, srv)
}

func _HeadlessLogService_LogDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogDownloadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadlessLogServiceServer).LogDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.HeadlessLogService/LogDownloadURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadlessLogServiceServer).LogDownloadURL(ctx, req.(*LogDownloadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeadlessLogService_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeadlessLogServiceServer).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.HeadlessLogService/ListLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeadlessLogServiceServer).ListLogs(ctx, req.(*ListLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeadlessLogService_ServiceDesc is the grpc.ServiceDesc for HeadlessLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeadlessLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contentservice.HeadlessLogService",
	HandlerType: (*HeadlessLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogDownloadURL",
			Handler:    _HeadlessLogService_LogDownloadURL_Handler,
		},
		{
			MethodName: "ListLogs",
			Handler:    _HeadlessLogService_ListLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "headless-log.proto",
}
