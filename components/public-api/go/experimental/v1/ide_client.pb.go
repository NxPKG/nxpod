// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/experimental/v1/ide_client.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendHeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *SendHeartbeatRequest) Reset() {
	*x = SendHeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendHeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendHeartbeatRequest) ProtoMessage() {}

func (x *SendHeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendHeartbeatRequest.ProtoReflect.Descriptor instead.
func (*SendHeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_ide_client_proto_rawDescGZIP(), []int{0}
}

func (x *SendHeartbeatRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type SendHeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendHeartbeatResponse) Reset() {
	*x = SendHeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendHeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendHeartbeatResponse) ProtoMessage() {}

func (x *SendHeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendHeartbeatResponse.ProtoReflect.Descriptor instead.
func (*SendHeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_ide_client_proto_rawDescGZIP(), []int{1}
}

type SendDidCloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *SendDidCloseRequest) Reset() {
	*x = SendDidCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDidCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDidCloseRequest) ProtoMessage() {}

func (x *SendDidCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDidCloseRequest.ProtoReflect.Descriptor instead.
func (*SendDidCloseRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_ide_client_proto_rawDescGZIP(), []int{2}
}

func (x *SendDidCloseRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type SendDidCloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendDidCloseResponse) Reset() {
	*x = SendDidCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDidCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDidCloseResponse) ProtoMessage() {}

func (x *SendDidCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDidCloseResponse.ProtoReflect.Descriptor instead.
func (*SendDidCloseResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_ide_client_proto_rawDescGZIP(), []int{3}
}

type UpdateGitStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string     `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Status      *GitStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateGitStatusRequest) Reset() {
	*x = UpdateGitStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGitStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGitStatusRequest) ProtoMessage() {}

func (x *UpdateGitStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGitStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateGitStatusRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_ide_client_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGitStatusRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *UpdateGitStatusRequest) GetStatus() *GitStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type UpdateGitStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGitStatusResponse) Reset() {
	*x = UpdateGitStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGitStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGitStatusResponse) ProtoMessage() {}

func (x *UpdateGitStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_experimental_v1_ide_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGitStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateGitStatusResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_experimental_v1_ide_client_proto_rawDescGZIP(), []int{5}
}

var File_gitpod_experimental_v1_ide_client_proto protoreflect.FileDescriptor

var file_gitpod_experimental_v1_ide_client_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x27, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x14, 0x53, 0x65,
	0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38,
	0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x69, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x76, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xe5, 0x02, 0x0a, 0x10, 0x49, 0x44, 0x45, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x69, 0x74, 0x70,
	0x6f, 0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x69, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x44, 0x69, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6b, 0x0a, 0x23, 0x69,
	0x6f, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_experimental_v1_ide_client_proto_rawDescOnce sync.Once
	file_gitpod_experimental_v1_ide_client_proto_rawDescData = file_gitpod_experimental_v1_ide_client_proto_rawDesc
)

func file_gitpod_experimental_v1_ide_client_proto_rawDescGZIP() []byte {
	file_gitpod_experimental_v1_ide_client_proto_rawDescOnce.Do(func() {
		file_gitpod_experimental_v1_ide_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_experimental_v1_ide_client_proto_rawDescData)
	})
	return file_gitpod_experimental_v1_ide_client_proto_rawDescData
}

var file_gitpod_experimental_v1_ide_client_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gitpod_experimental_v1_ide_client_proto_goTypes = []interface{}{
	(*SendHeartbeatRequest)(nil),    // 0: gitpod.experimental.v1.SendHeartbeatRequest
	(*SendHeartbeatResponse)(nil),   // 1: gitpod.experimental.v1.SendHeartbeatResponse
	(*SendDidCloseRequest)(nil),     // 2: gitpod.experimental.v1.SendDidCloseRequest
	(*SendDidCloseResponse)(nil),    // 3: gitpod.experimental.v1.SendDidCloseResponse
	(*UpdateGitStatusRequest)(nil),  // 4: gitpod.experimental.v1.UpdateGitStatusRequest
	(*UpdateGitStatusResponse)(nil), // 5: gitpod.experimental.v1.UpdateGitStatusResponse
	(*GitStatus)(nil),               // 6: gitpod.experimental.v1.GitStatus
}
var file_gitpod_experimental_v1_ide_client_proto_depIdxs = []int32{
	6, // 0: gitpod.experimental.v1.UpdateGitStatusRequest.status:type_name -> gitpod.experimental.v1.GitStatus
	0, // 1: gitpod.experimental.v1.IDEClientService.SendHeartbeat:input_type -> gitpod.experimental.v1.SendHeartbeatRequest
	2, // 2: gitpod.experimental.v1.IDEClientService.SendDidClose:input_type -> gitpod.experimental.v1.SendDidCloseRequest
	4, // 3: gitpod.experimental.v1.IDEClientService.UpdateGitStatus:input_type -> gitpod.experimental.v1.UpdateGitStatusRequest
	1, // 4: gitpod.experimental.v1.IDEClientService.SendHeartbeat:output_type -> gitpod.experimental.v1.SendHeartbeatResponse
	3, // 5: gitpod.experimental.v1.IDEClientService.SendDidClose:output_type -> gitpod.experimental.v1.SendDidCloseResponse
	5, // 6: gitpod.experimental.v1.IDEClientService.UpdateGitStatus:output_type -> gitpod.experimental.v1.UpdateGitStatusResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gitpod_experimental_v1_ide_client_proto_init() }
func file_gitpod_experimental_v1_ide_client_proto_init() {
	if File_gitpod_experimental_v1_ide_client_proto != nil {
		return
	}
	file_gitpod_experimental_v1_workspaces_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gitpod_experimental_v1_ide_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendHeartbeatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_experimental_v1_ide_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendHeartbeatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_experimental_v1_ide_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDidCloseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_experimental_v1_ide_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDidCloseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_experimental_v1_ide_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGitStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitpod_experimental_v1_ide_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGitStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gitpod_experimental_v1_ide_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitpod_experimental_v1_ide_client_proto_goTypes,
		DependencyIndexes: file_gitpod_experimental_v1_ide_client_proto_depIdxs,
		MessageInfos:      file_gitpod_experimental_v1_ide_client_proto_msgTypes,
	}.Build()
	File_gitpod_experimental_v1_ide_client_proto = out.File
	file_gitpod_experimental_v1_ide_client_proto_rawDesc = nil
	file_gitpod_experimental_v1_ide_client_proto_goTypes = nil
	file_gitpod_experimental_v1_ide_client_proto_depIdxs = nil
}
