// Copyright 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: src/protobuf/authorization_pb2/authorization.proto

package authorization_pb2

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

type RoleType int32

const (
	RoleType_ROLE_TYPE_UNSET RoleType = 0
	// A shorthand request for asking for all allowed roles.
	RoleType_ALL RoleType = 1
	// Role defining validator to validator communication
	RoleType_NETWORK RoleType = 2
)

// Enum value maps for RoleType.
var (
	RoleType_name = map[int32]string{
		0: "ROLE_TYPE_UNSET",
		1: "ALL",
		2: "NETWORK",
	}
	RoleType_value = map[string]int32{
		"ROLE_TYPE_UNSET": 0,
		"ALL":             1,
		"NETWORK":         2,
	}
)

func (x RoleType) Enum() *RoleType {
	p := new(RoleType)
	*p = x
	return p
}

func (x RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_protobuf_authorization_pb2_authorization_proto_enumTypes[0].Descriptor()
}

func (RoleType) Type() protoreflect.EnumType {
	return &file_src_protobuf_authorization_pb2_authorization_proto_enumTypes[0]
}

func (x RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleType.Descriptor instead.
func (RoleType) EnumDescriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{0}
}

// Whether the connection can participate in authorization
type ConnectionResponse_Status int32

const (
	ConnectionResponse_STATUS_UNSET ConnectionResponse_Status = 0
	ConnectionResponse_OK           ConnectionResponse_Status = 1
	ConnectionResponse_ERROR        ConnectionResponse_Status = 2
)

// Enum value maps for ConnectionResponse_Status.
var (
	ConnectionResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "ERROR",
	}
	ConnectionResponse_Status_value = map[string]int32{
		"STATUS_UNSET": 0,
		"OK":           1,
		"ERROR":        2,
	}
)

func (x ConnectionResponse_Status) Enum() *ConnectionResponse_Status {
	p := new(ConnectionResponse_Status)
	*p = x
	return p
}

func (x ConnectionResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_src_protobuf_authorization_pb2_authorization_proto_enumTypes[1].Descriptor()
}

func (ConnectionResponse_Status) Type() protoreflect.EnumType {
	return &file_src_protobuf_authorization_pb2_authorization_proto_enumTypes[1]
}

func (x ConnectionResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionResponse_Status.Descriptor instead.
func (ConnectionResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{1, 0}
}

//Authorization Type required for the authorization procedure
type ConnectionResponse_AuthorizationType int32

const (
	ConnectionResponse_AUTHORIZATION_TYPE_UNSET ConnectionResponse_AuthorizationType = 0
	ConnectionResponse_TRUST                    ConnectionResponse_AuthorizationType = 1
	ConnectionResponse_CHALLENGE                ConnectionResponse_AuthorizationType = 2
)

// Enum value maps for ConnectionResponse_AuthorizationType.
var (
	ConnectionResponse_AuthorizationType_name = map[int32]string{
		0: "AUTHORIZATION_TYPE_UNSET",
		1: "TRUST",
		2: "CHALLENGE",
	}
	ConnectionResponse_AuthorizationType_value = map[string]int32{
		"AUTHORIZATION_TYPE_UNSET": 0,
		"TRUST":                    1,
		"CHALLENGE":                2,
	}
)

func (x ConnectionResponse_AuthorizationType) Enum() *ConnectionResponse_AuthorizationType {
	p := new(ConnectionResponse_AuthorizationType)
	*p = x
	return p
}

func (x ConnectionResponse_AuthorizationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionResponse_AuthorizationType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_protobuf_authorization_pb2_authorization_proto_enumTypes[2].Descriptor()
}

func (ConnectionResponse_AuthorizationType) Type() protoreflect.EnumType {
	return &file_src_protobuf_authorization_pb2_authorization_proto_enumTypes[2]
}

func (x ConnectionResponse_AuthorizationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionResponse_AuthorizationType.Descriptor instead.
func (ConnectionResponse_AuthorizationType) EnumDescriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{1, 1}
}

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the first message that must be sent to start off authorization.
	// The endpoint of the connection.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type ConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles  []*ConnectionResponse_RoleEntry `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	Status ConnectionResponse_Status       `protobuf:"varint,2,opt,name=status,proto3,enum=ConnectionResponse_Status" json:"status,omitempty"`
}

func (x *ConnectionResponse) Reset() {
	*x = ConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse) ProtoMessage() {}

func (x *ConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse.ProtoReflect.Descriptor instead.
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionResponse) GetRoles() []*ConnectionResponse_RoleEntry {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *ConnectionResponse) GetStatus() ConnectionResponse_Status {
	if x != nil {
		return x.Status
	}
	return ConnectionResponse_STATUS_UNSET
}

type AuthorizationTrustRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A set of requested RoleTypes
	Roles     []RoleType `protobuf:"varint,1,rep,packed,name=roles,proto3,enum=RoleType" json:"roles,omitempty"`
	PublicKey string     `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *AuthorizationTrustRequest) Reset() {
	*x = AuthorizationTrustRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationTrustRequest) ProtoMessage() {}

func (x *AuthorizationTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationTrustRequest.ProtoReflect.Descriptor instead.
func (*AuthorizationTrustRequest) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizationTrustRequest) GetRoles() []RoleType {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AuthorizationTrustRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type AuthorizationTrustResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual set the requester has access to
	Roles []RoleType `protobuf:"varint,1,rep,packed,name=roles,proto3,enum=RoleType" json:"roles,omitempty"`
}

func (x *AuthorizationTrustResponse) Reset() {
	*x = AuthorizationTrustResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationTrustResponse) ProtoMessage() {}

func (x *AuthorizationTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationTrustResponse.ProtoReflect.Descriptor instead.
func (*AuthorizationTrustResponse) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizationTrustResponse) GetRoles() []RoleType {
	if x != nil {
		return x.Roles
	}
	return nil
}

type AuthorizationViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Role the requester did not have access to
	Violation RoleType `protobuf:"varint,1,opt,name=violation,proto3,enum=RoleType" json:"violation,omitempty"`
}

func (x *AuthorizationViolation) Reset() {
	*x = AuthorizationViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationViolation) ProtoMessage() {}

func (x *AuthorizationViolation) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationViolation.ProtoReflect.Descriptor instead.
func (*AuthorizationViolation) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorizationViolation) GetViolation() RoleType {
	if x != nil {
		return x.Violation
	}
	return RoleType_ROLE_TYPE_UNSET
}

type AuthorizationChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorizationChallengeRequest) Reset() {
	*x = AuthorizationChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationChallengeRequest) ProtoMessage() {}

func (x *AuthorizationChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationChallengeRequest.ProtoReflect.Descriptor instead.
func (*AuthorizationChallengeRequest) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{5}
}

type AuthorizationChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Random payload that the connecting node must sign
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *AuthorizationChallengeResponse) Reset() {
	*x = AuthorizationChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationChallengeResponse) ProtoMessage() {}

func (x *AuthorizationChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationChallengeResponse.ProtoReflect.Descriptor instead.
func (*AuthorizationChallengeResponse) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{6}
}

func (x *AuthorizationChallengeResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type AuthorizationChallengeSubmit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public key of node
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// signature derived from signing the challenge payload
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// A set of requested Roles
	Roles []RoleType `protobuf:"varint,4,rep,packed,name=roles,proto3,enum=RoleType" json:"roles,omitempty"`
}

func (x *AuthorizationChallengeSubmit) Reset() {
	*x = AuthorizationChallengeSubmit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationChallengeSubmit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationChallengeSubmit) ProtoMessage() {}

func (x *AuthorizationChallengeSubmit) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationChallengeSubmit.ProtoReflect.Descriptor instead.
func (*AuthorizationChallengeSubmit) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{7}
}

func (x *AuthorizationChallengeSubmit) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *AuthorizationChallengeSubmit) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AuthorizationChallengeSubmit) GetRoles() []RoleType {
	if x != nil {
		return x.Roles
	}
	return nil
}

type AuthorizationChallengeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The approved roles for that connection
	Roles []RoleType `protobuf:"varint,1,rep,packed,name=roles,proto3,enum=RoleType" json:"roles,omitempty"`
}

func (x *AuthorizationChallengeResult) Reset() {
	*x = AuthorizationChallengeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationChallengeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationChallengeResult) ProtoMessage() {}

func (x *AuthorizationChallengeResult) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationChallengeResult.ProtoReflect.Descriptor instead.
func (*AuthorizationChallengeResult) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{8}
}

func (x *AuthorizationChallengeResult) GetRoles() []RoleType {
	if x != nil {
		return x.Roles
	}
	return nil
}

type ConnectionResponse_RoleEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The role type for this role entry
	Role RoleType `protobuf:"varint,1,opt,name=role,proto3,enum=RoleType" json:"role,omitempty"`
	// The Authorization Type required for the above role
	AuthType ConnectionResponse_AuthorizationType `protobuf:"varint,2,opt,name=auth_type,json=authType,proto3,enum=ConnectionResponse_AuthorizationType" json:"auth_type,omitempty"`
}

func (x *ConnectionResponse_RoleEntry) Reset() {
	*x = ConnectionResponse_RoleEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse_RoleEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse_RoleEntry) ProtoMessage() {}

func (x *ConnectionResponse_RoleEntry) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse_RoleEntry.ProtoReflect.Descriptor instead.
func (*ConnectionResponse_RoleEntry) Descriptor() ([]byte, []int) {
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ConnectionResponse_RoleEntry) GetRole() RoleType {
	if x != nil {
		return x.Role
	}
	return RoleType_ROLE_TYPE_UNSET
}

func (x *ConnectionResponse_RoleEntry) GetAuthType() ConnectionResponse_AuthorizationType {
	if x != nil {
		return x.AuthType
	}
	return ConnectionResponse_AUTHORIZATION_TYPE_UNSET
}

var File_src_protobuf_authorization_pb2_authorization_proto protoreflect.FileDescriptor

var file_src_protobuf_authorization_pb2_authorization_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x32,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xe9, 0x02, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x6e, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x42, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x22, 0x4b, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x55, 0x54,
	0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x55, 0x53, 0x54,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10,
	0x02, 0x22, 0x5b, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x3d,
	0x0a, 0x1a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x41, 0x0a,
	0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x69,
	0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x09, 0x76, 0x69, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x1f, 0x0a, 0x1d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3a, 0x0a, 0x1e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x7c, 0x0a,
	0x1c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x1c, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2a, 0x35, 0x0a, 0x08,
	0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x02, 0x42, 0x2c, 0x0a, 0x15, 0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x11,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_protobuf_authorization_pb2_authorization_proto_rawDescOnce sync.Once
	file_src_protobuf_authorization_pb2_authorization_proto_rawDescData = file_src_protobuf_authorization_pb2_authorization_proto_rawDesc
)

func file_src_protobuf_authorization_pb2_authorization_proto_rawDescGZIP() []byte {
	file_src_protobuf_authorization_pb2_authorization_proto_rawDescOnce.Do(func() {
		file_src_protobuf_authorization_pb2_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_protobuf_authorization_pb2_authorization_proto_rawDescData)
	})
	return file_src_protobuf_authorization_pb2_authorization_proto_rawDescData
}

var file_src_protobuf_authorization_pb2_authorization_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_src_protobuf_authorization_pb2_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_src_protobuf_authorization_pb2_authorization_proto_goTypes = []interface{}{
	(RoleType)(0),                             // 0: RoleType
	(ConnectionResponse_Status)(0),            // 1: ConnectionResponse.Status
	(ConnectionResponse_AuthorizationType)(0), // 2: ConnectionResponse.AuthorizationType
	(*ConnectionRequest)(nil),                 // 3: ConnectionRequest
	(*ConnectionResponse)(nil),                // 4: ConnectionResponse
	(*AuthorizationTrustRequest)(nil),         // 5: AuthorizationTrustRequest
	(*AuthorizationTrustResponse)(nil),        // 6: AuthorizationTrustResponse
	(*AuthorizationViolation)(nil),            // 7: AuthorizationViolation
	(*AuthorizationChallengeRequest)(nil),     // 8: AuthorizationChallengeRequest
	(*AuthorizationChallengeResponse)(nil),    // 9: AuthorizationChallengeResponse
	(*AuthorizationChallengeSubmit)(nil),      // 10: AuthorizationChallengeSubmit
	(*AuthorizationChallengeResult)(nil),      // 11: AuthorizationChallengeResult
	(*ConnectionResponse_RoleEntry)(nil),      // 12: ConnectionResponse.RoleEntry
}
var file_src_protobuf_authorization_pb2_authorization_proto_depIdxs = []int32{
	12, // 0: ConnectionResponse.roles:type_name -> ConnectionResponse.RoleEntry
	1,  // 1: ConnectionResponse.status:type_name -> ConnectionResponse.Status
	0,  // 2: AuthorizationTrustRequest.roles:type_name -> RoleType
	0,  // 3: AuthorizationTrustResponse.roles:type_name -> RoleType
	0,  // 4: AuthorizationViolation.violation:type_name -> RoleType
	0,  // 5: AuthorizationChallengeSubmit.roles:type_name -> RoleType
	0,  // 6: AuthorizationChallengeResult.roles:type_name -> RoleType
	0,  // 7: ConnectionResponse.RoleEntry.role:type_name -> RoleType
	2,  // 8: ConnectionResponse.RoleEntry.auth_type:type_name -> ConnectionResponse.AuthorizationType
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_src_protobuf_authorization_pb2_authorization_proto_init() }
func file_src_protobuf_authorization_pb2_authorization_proto_init() {
	if File_src_protobuf_authorization_pb2_authorization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationTrustRequest); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationTrustResponse); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationViolation); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationChallengeRequest); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationChallengeResponse); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationChallengeSubmit); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationChallengeResult); i {
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
		file_src_protobuf_authorization_pb2_authorization_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse_RoleEntry); i {
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
			RawDescriptor: file_src_protobuf_authorization_pb2_authorization_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_protobuf_authorization_pb2_authorization_proto_goTypes,
		DependencyIndexes: file_src_protobuf_authorization_pb2_authorization_proto_depIdxs,
		EnumInfos:         file_src_protobuf_authorization_pb2_authorization_proto_enumTypes,
		MessageInfos:      file_src_protobuf_authorization_pb2_authorization_proto_msgTypes,
	}.Build()
	File_src_protobuf_authorization_pb2_authorization_proto = out.File
	file_src_protobuf_authorization_pb2_authorization_proto_rawDesc = nil
	file_src_protobuf_authorization_pb2_authorization_proto_goTypes = nil
	file_src_protobuf_authorization_pb2_authorization_proto_depIdxs = nil
}
