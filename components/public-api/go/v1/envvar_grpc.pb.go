// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gitpod/v1/envvar.proto

package v1

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

// EnvironmentVariableServiceClient is the client API for EnvironmentVariableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentVariableServiceClient interface {
	// ListUserEnvironmentVariables returns all environment variables for the
	// authenticated user.
	ListUserEnvironmentVariables(ctx context.Context, in *ListUserEnvironmentVariablesRequest, opts ...grpc.CallOption) (*ListUserEnvironmentVariablesResponse, error)
	// UpdateUserEnvironmentVariable updates an environment variable for the
	// authenticated user.
	UpdateUserEnvironmentVariable(ctx context.Context, in *UpdateUserEnvironmentVariableRequest, opts ...grpc.CallOption) (*UpdateUserEnvironmentVariableResponse, error)
	// CreateUserEnvironmentVariable creates a new environment variable for the
	// authenticated user.
	CreateUserEnvironmentVariable(ctx context.Context, in *CreateUserEnvironmentVariableRequest, opts ...grpc.CallOption) (*CreateUserEnvironmentVariableResponse, error)
	// DeleteUserEnvironmentVariable deletes an environment variable for the
	// authenticated user.
	DeleteUserEnvironmentVariable(ctx context.Context, in *DeleteUserEnvironmentVariableRequest, opts ...grpc.CallOption) (*DeleteUserEnvironmentVariableResponse, error)
	// ListConfigurationEnvironmentVariables returns all environment variables in
	// a configuration.
	ListConfigurationEnvironmentVariables(ctx context.Context, in *ListConfigurationEnvironmentVariablesRequest, opts ...grpc.CallOption) (*ListConfigurationEnvironmentVariablesResponse, error)
	// UpdateConfigurationEnvironmentVariable updates an environment variable in
	// a configuration.
	UpdateConfigurationEnvironmentVariable(ctx context.Context, in *UpdateConfigurationEnvironmentVariableRequest, opts ...grpc.CallOption) (*UpdateConfigurationEnvironmentVariableResponse, error)
	// CreateConfigurationEnvironmentVariable creates a new environment variable
	// in a configuration.
	CreateConfigurationEnvironmentVariable(ctx context.Context, in *CreateConfigurationEnvironmentVariableRequest, opts ...grpc.CallOption) (*CreateConfigurationEnvironmentVariableResponse, error)
	// DeleteConfigurationEnvironmentVariable deletes an environment variable in
	// a configuration.
	DeleteConfigurationEnvironmentVariable(ctx context.Context, in *DeleteConfigurationEnvironmentVariableRequest, opts ...grpc.CallOption) (*DeleteConfigurationEnvironmentVariableResponse, error)
	ResolveWorkspaceEnvironmentVariables(ctx context.Context, in *ResolveWorkspaceEnvironmentVariablesRequest, opts ...grpc.CallOption) (*ResolveWorkspaceEnvironmentVariablesResponse, error)
}

type environmentVariableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentVariableServiceClient(cc grpc.ClientConnInterface) EnvironmentVariableServiceClient {
	return &environmentVariableServiceClient{cc}
}

func (c *environmentVariableServiceClient) ListUserEnvironmentVariables(ctx context.Context, in *ListUserEnvironmentVariablesRequest, opts ...grpc.CallOption) (*ListUserEnvironmentVariablesResponse, error) {
	out := new(ListUserEnvironmentVariablesResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/ListUserEnvironmentVariables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) UpdateUserEnvironmentVariable(ctx context.Context, in *UpdateUserEnvironmentVariableRequest, opts ...grpc.CallOption) (*UpdateUserEnvironmentVariableResponse, error) {
	out := new(UpdateUserEnvironmentVariableResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/UpdateUserEnvironmentVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) CreateUserEnvironmentVariable(ctx context.Context, in *CreateUserEnvironmentVariableRequest, opts ...grpc.CallOption) (*CreateUserEnvironmentVariableResponse, error) {
	out := new(CreateUserEnvironmentVariableResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/CreateUserEnvironmentVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) DeleteUserEnvironmentVariable(ctx context.Context, in *DeleteUserEnvironmentVariableRequest, opts ...grpc.CallOption) (*DeleteUserEnvironmentVariableResponse, error) {
	out := new(DeleteUserEnvironmentVariableResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/DeleteUserEnvironmentVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) ListConfigurationEnvironmentVariables(ctx context.Context, in *ListConfigurationEnvironmentVariablesRequest, opts ...grpc.CallOption) (*ListConfigurationEnvironmentVariablesResponse, error) {
	out := new(ListConfigurationEnvironmentVariablesResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/ListConfigurationEnvironmentVariables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) UpdateConfigurationEnvironmentVariable(ctx context.Context, in *UpdateConfigurationEnvironmentVariableRequest, opts ...grpc.CallOption) (*UpdateConfigurationEnvironmentVariableResponse, error) {
	out := new(UpdateConfigurationEnvironmentVariableResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/UpdateConfigurationEnvironmentVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) CreateConfigurationEnvironmentVariable(ctx context.Context, in *CreateConfigurationEnvironmentVariableRequest, opts ...grpc.CallOption) (*CreateConfigurationEnvironmentVariableResponse, error) {
	out := new(CreateConfigurationEnvironmentVariableResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/CreateConfigurationEnvironmentVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) DeleteConfigurationEnvironmentVariable(ctx context.Context, in *DeleteConfigurationEnvironmentVariableRequest, opts ...grpc.CallOption) (*DeleteConfigurationEnvironmentVariableResponse, error) {
	out := new(DeleteConfigurationEnvironmentVariableResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/DeleteConfigurationEnvironmentVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentVariableServiceClient) ResolveWorkspaceEnvironmentVariables(ctx context.Context, in *ResolveWorkspaceEnvironmentVariablesRequest, opts ...grpc.CallOption) (*ResolveWorkspaceEnvironmentVariablesResponse, error) {
	out := new(ResolveWorkspaceEnvironmentVariablesResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.EnvironmentVariableService/ResolveWorkspaceEnvironmentVariables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentVariableServiceServer is the server API for EnvironmentVariableService service.
// All implementations must embed UnimplementedEnvironmentVariableServiceServer
// for forward compatibility
type EnvironmentVariableServiceServer interface {
	// ListUserEnvironmentVariables returns all environment variables for the
	// authenticated user.
	ListUserEnvironmentVariables(context.Context, *ListUserEnvironmentVariablesRequest) (*ListUserEnvironmentVariablesResponse, error)
	// UpdateUserEnvironmentVariable updates an environment variable for the
	// authenticated user.
	UpdateUserEnvironmentVariable(context.Context, *UpdateUserEnvironmentVariableRequest) (*UpdateUserEnvironmentVariableResponse, error)
	// CreateUserEnvironmentVariable creates a new environment variable for the
	// authenticated user.
	CreateUserEnvironmentVariable(context.Context, *CreateUserEnvironmentVariableRequest) (*CreateUserEnvironmentVariableResponse, error)
	// DeleteUserEnvironmentVariable deletes an environment variable for the
	// authenticated user.
	DeleteUserEnvironmentVariable(context.Context, *DeleteUserEnvironmentVariableRequest) (*DeleteUserEnvironmentVariableResponse, error)
	// ListConfigurationEnvironmentVariables returns all environment variables in
	// a configuration.
	ListConfigurationEnvironmentVariables(context.Context, *ListConfigurationEnvironmentVariablesRequest) (*ListConfigurationEnvironmentVariablesResponse, error)
	// UpdateConfigurationEnvironmentVariable updates an environment variable in
	// a configuration.
	UpdateConfigurationEnvironmentVariable(context.Context, *UpdateConfigurationEnvironmentVariableRequest) (*UpdateConfigurationEnvironmentVariableResponse, error)
	// CreateConfigurationEnvironmentVariable creates a new environment variable
	// in a configuration.
	CreateConfigurationEnvironmentVariable(context.Context, *CreateConfigurationEnvironmentVariableRequest) (*CreateConfigurationEnvironmentVariableResponse, error)
	// DeleteConfigurationEnvironmentVariable deletes an environment variable in
	// a configuration.
	DeleteConfigurationEnvironmentVariable(context.Context, *DeleteConfigurationEnvironmentVariableRequest) (*DeleteConfigurationEnvironmentVariableResponse, error)
	ResolveWorkspaceEnvironmentVariables(context.Context, *ResolveWorkspaceEnvironmentVariablesRequest) (*ResolveWorkspaceEnvironmentVariablesResponse, error)
	mustEmbedUnimplementedEnvironmentVariableServiceServer()
}

// UnimplementedEnvironmentVariableServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEnvironmentVariableServiceServer struct {
}

func (UnimplementedEnvironmentVariableServiceServer) ListUserEnvironmentVariables(context.Context, *ListUserEnvironmentVariablesRequest) (*ListUserEnvironmentVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserEnvironmentVariables not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) UpdateUserEnvironmentVariable(context.Context, *UpdateUserEnvironmentVariableRequest) (*UpdateUserEnvironmentVariableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserEnvironmentVariable not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) CreateUserEnvironmentVariable(context.Context, *CreateUserEnvironmentVariableRequest) (*CreateUserEnvironmentVariableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserEnvironmentVariable not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) DeleteUserEnvironmentVariable(context.Context, *DeleteUserEnvironmentVariableRequest) (*DeleteUserEnvironmentVariableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserEnvironmentVariable not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) ListConfigurationEnvironmentVariables(context.Context, *ListConfigurationEnvironmentVariablesRequest) (*ListConfigurationEnvironmentVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigurationEnvironmentVariables not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) UpdateConfigurationEnvironmentVariable(context.Context, *UpdateConfigurationEnvironmentVariableRequest) (*UpdateConfigurationEnvironmentVariableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfigurationEnvironmentVariable not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) CreateConfigurationEnvironmentVariable(context.Context, *CreateConfigurationEnvironmentVariableRequest) (*CreateConfigurationEnvironmentVariableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfigurationEnvironmentVariable not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) DeleteConfigurationEnvironmentVariable(context.Context, *DeleteConfigurationEnvironmentVariableRequest) (*DeleteConfigurationEnvironmentVariableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfigurationEnvironmentVariable not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) ResolveWorkspaceEnvironmentVariables(context.Context, *ResolveWorkspaceEnvironmentVariablesRequest) (*ResolveWorkspaceEnvironmentVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveWorkspaceEnvironmentVariables not implemented")
}
func (UnimplementedEnvironmentVariableServiceServer) mustEmbedUnimplementedEnvironmentVariableServiceServer() {
}

// UnsafeEnvironmentVariableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentVariableServiceServer will
// result in compilation errors.
type UnsafeEnvironmentVariableServiceServer interface {
	mustEmbedUnimplementedEnvironmentVariableServiceServer()
}

func RegisterEnvironmentVariableServiceServer(s grpc.ServiceRegistrar, srv EnvironmentVariableServiceServer) {
	s.RegisterService(&EnvironmentVariableService_ServiceDesc, srv)
}

func _EnvironmentVariableService_ListUserEnvironmentVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserEnvironmentVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).ListUserEnvironmentVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/ListUserEnvironmentVariables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).ListUserEnvironmentVariables(ctx, req.(*ListUserEnvironmentVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_UpdateUserEnvironmentVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserEnvironmentVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).UpdateUserEnvironmentVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/UpdateUserEnvironmentVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).UpdateUserEnvironmentVariable(ctx, req.(*UpdateUserEnvironmentVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_CreateUserEnvironmentVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserEnvironmentVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).CreateUserEnvironmentVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/CreateUserEnvironmentVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).CreateUserEnvironmentVariable(ctx, req.(*CreateUserEnvironmentVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_DeleteUserEnvironmentVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserEnvironmentVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).DeleteUserEnvironmentVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/DeleteUserEnvironmentVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).DeleteUserEnvironmentVariable(ctx, req.(*DeleteUserEnvironmentVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_ListConfigurationEnvironmentVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigurationEnvironmentVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).ListConfigurationEnvironmentVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/ListConfigurationEnvironmentVariables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).ListConfigurationEnvironmentVariables(ctx, req.(*ListConfigurationEnvironmentVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_UpdateConfigurationEnvironmentVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigurationEnvironmentVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).UpdateConfigurationEnvironmentVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/UpdateConfigurationEnvironmentVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).UpdateConfigurationEnvironmentVariable(ctx, req.(*UpdateConfigurationEnvironmentVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_CreateConfigurationEnvironmentVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigurationEnvironmentVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).CreateConfigurationEnvironmentVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/CreateConfigurationEnvironmentVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).CreateConfigurationEnvironmentVariable(ctx, req.(*CreateConfigurationEnvironmentVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_DeleteConfigurationEnvironmentVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigurationEnvironmentVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).DeleteConfigurationEnvironmentVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/DeleteConfigurationEnvironmentVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).DeleteConfigurationEnvironmentVariable(ctx, req.(*DeleteConfigurationEnvironmentVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentVariableService_ResolveWorkspaceEnvironmentVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveWorkspaceEnvironmentVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentVariableServiceServer).ResolveWorkspaceEnvironmentVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.EnvironmentVariableService/ResolveWorkspaceEnvironmentVariables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentVariableServiceServer).ResolveWorkspaceEnvironmentVariables(ctx, req.(*ResolveWorkspaceEnvironmentVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentVariableService_ServiceDesc is the grpc.ServiceDesc for EnvironmentVariableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentVariableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.v1.EnvironmentVariableService",
	HandlerType: (*EnvironmentVariableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUserEnvironmentVariables",
			Handler:    _EnvironmentVariableService_ListUserEnvironmentVariables_Handler,
		},
		{
			MethodName: "UpdateUserEnvironmentVariable",
			Handler:    _EnvironmentVariableService_UpdateUserEnvironmentVariable_Handler,
		},
		{
			MethodName: "CreateUserEnvironmentVariable",
			Handler:    _EnvironmentVariableService_CreateUserEnvironmentVariable_Handler,
		},
		{
			MethodName: "DeleteUserEnvironmentVariable",
			Handler:    _EnvironmentVariableService_DeleteUserEnvironmentVariable_Handler,
		},
		{
			MethodName: "ListConfigurationEnvironmentVariables",
			Handler:    _EnvironmentVariableService_ListConfigurationEnvironmentVariables_Handler,
		},
		{
			MethodName: "UpdateConfigurationEnvironmentVariable",
			Handler:    _EnvironmentVariableService_UpdateConfigurationEnvironmentVariable_Handler,
		},
		{
			MethodName: "CreateConfigurationEnvironmentVariable",
			Handler:    _EnvironmentVariableService_CreateConfigurationEnvironmentVariable_Handler,
		},
		{
			MethodName: "DeleteConfigurationEnvironmentVariable",
			Handler:    _EnvironmentVariableService_DeleteConfigurationEnvironmentVariable_Handler,
		},
		{
			MethodName: "ResolveWorkspaceEnvironmentVariables",
			Handler:    _EnvironmentVariableService_ResolveWorkspaceEnvironmentVariables_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitpod/v1/envvar.proto",
}
