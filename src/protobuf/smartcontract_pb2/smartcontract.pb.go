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
// source: src/protobuf/smartcontract_pb2/smartcontract.proto

package smartcontract_pb2

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

type SmartContractValidationResponse_Status int32

const (
	SmartContractValidationResponse_STATUS_UNSET        SmartContractValidationResponse_Status = 0
	SmartContractValidationResponse_OK                  SmartContractValidationResponse_Status = 1
	SmartContractValidationResponse_INVALID_TRANSACTION SmartContractValidationResponse_Status = 2
	SmartContractValidationResponse_INTERNAL_ERROR      SmartContractValidationResponse_Status = 3
)

// Enum value maps for SmartContractValidationResponse_Status.
var (
	SmartContractValidationResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "INVALID_TRANSACTION",
		3: "INTERNAL_ERROR",
	}
	SmartContractValidationResponse_Status_value = map[string]int32{
		"STATUS_UNSET":        0,
		"OK":                  1,
		"INVALID_TRANSACTION": 2,
		"INTERNAL_ERROR":      3,
	}
)

func (x SmartContractValidationResponse_Status) Enum() *SmartContractValidationResponse_Status {
	p := new(SmartContractValidationResponse_Status)
	*p = x
	return p
}

func (x SmartContractValidationResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SmartContractValidationResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_enumTypes[0].Descriptor()
}

func (SmartContractValidationResponse_Status) Type() protoreflect.EnumType {
	return &file_src_protobuf_smartcontract_pb2_smartcontract_proto_enumTypes[0]
}

func (x SmartContractValidationResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SmartContractValidationResponse_Status.Descriptor instead.
func (SmartContractValidationResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescGZIP(), []int{1, 0}
}

type SmartContractRegisterResponse_Status int32

const (
	SmartContractRegisterResponse_STATUS_UNSET        SmartContractRegisterResponse_Status = 0
	SmartContractRegisterResponse_OK                  SmartContractRegisterResponse_Status = 1
	SmartContractRegisterResponse_INVALID_TRANSACTION SmartContractRegisterResponse_Status = 2
	SmartContractRegisterResponse_INTERNAL_ERROR      SmartContractRegisterResponse_Status = 3
)

// Enum value maps for SmartContractRegisterResponse_Status.
var (
	SmartContractRegisterResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "INVALID_TRANSACTION",
		3: "INTERNAL_ERROR",
	}
	SmartContractRegisterResponse_Status_value = map[string]int32{
		"STATUS_UNSET":        0,
		"OK":                  1,
		"INVALID_TRANSACTION": 2,
		"INTERNAL_ERROR":      3,
	}
)

func (x SmartContractRegisterResponse_Status) Enum() *SmartContractRegisterResponse_Status {
	p := new(SmartContractRegisterResponse_Status)
	*p = x
	return p
}

func (x SmartContractRegisterResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SmartContractRegisterResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_enumTypes[1].Descriptor()
}

func (SmartContractRegisterResponse_Status) Type() protoreflect.EnumType {
	return &file_src_protobuf_smartcontract_pb2_smartcontract_proto_enumTypes[1]
}

func (x SmartContractRegisterResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SmartContractRegisterResponse_Status.Descriptor instead.
func (SmartContractRegisterResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescGZIP(), []int{3, 0}
}

type SmartContractValidationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload              []byte `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	SmartContractAddress string `protobuf:"bytes,2,opt,name=SmartContractAddress,proto3" json:"SmartContractAddress,omitempty"`
}

func (x *SmartContractValidationRequest) Reset() {
	*x = SmartContractValidationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartContractValidationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartContractValidationRequest) ProtoMessage() {}

func (x *SmartContractValidationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartContractValidationRequest.ProtoReflect.Descriptor instead.
func (*SmartContractValidationRequest) Descriptor() ([]byte, []int) {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescGZIP(), []int{0}
}

func (x *SmartContractValidationRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SmartContractValidationRequest) GetSmartContractAddress() string {
	if x != nil {
		return x.SmartContractAddress
	}
	return ""
}

type SmartContractValidationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          SmartContractValidationResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=SmartContractValidationResponse_Status" json:"status,omitempty"`
	SignatureScheme string                                 `protobuf:"bytes,2,opt,name=SignatureScheme,proto3" json:"SignatureScheme,omitempty"`
	N               int32                                  `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	T               int32                                  `protobuf:"varint,4,opt,name=T,proto3" json:"T,omitempty"`
}

func (x *SmartContractValidationResponse) Reset() {
	*x = SmartContractValidationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartContractValidationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartContractValidationResponse) ProtoMessage() {}

func (x *SmartContractValidationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartContractValidationResponse.ProtoReflect.Descriptor instead.
func (*SmartContractValidationResponse) Descriptor() ([]byte, []int) {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescGZIP(), []int{1}
}

func (x *SmartContractValidationResponse) GetStatus() SmartContractValidationResponse_Status {
	if x != nil {
		return x.Status
	}
	return SmartContractValidationResponse_STATUS_UNSET
}

func (x *SmartContractValidationResponse) GetSignatureScheme() string {
	if x != nil {
		return x.SignatureScheme
	}
	return ""
}

func (x *SmartContractValidationResponse) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *SmartContractValidationResponse) GetT() int32 {
	if x != nil {
		return x.T
	}
	return 0
}

type SmartContractRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SmartContractAddress string `protobuf:"bytes,1,opt,name=SmartContractAddress,proto3" json:"SmartContractAddress,omitempty"`
}

func (x *SmartContractRegisterRequest) Reset() {
	*x = SmartContractRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartContractRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartContractRegisterRequest) ProtoMessage() {}

func (x *SmartContractRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartContractRegisterRequest.ProtoReflect.Descriptor instead.
func (*SmartContractRegisterRequest) Descriptor() ([]byte, []int) {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescGZIP(), []int{2}
}

func (x *SmartContractRegisterRequest) GetSmartContractAddress() string {
	if x != nil {
		return x.SmartContractAddress
	}
	return ""
}

type SmartContractRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SmartContractRegisterResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=SmartContractRegisterResponse_Status" json:"status,omitempty"`
}

func (x *SmartContractRegisterResponse) Reset() {
	*x = SmartContractRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartContractRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartContractRegisterResponse) ProtoMessage() {}

func (x *SmartContractRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartContractRegisterResponse.ProtoReflect.Descriptor instead.
func (*SmartContractRegisterResponse) Descriptor() ([]byte, []int) {
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescGZIP(), []int{3}
}

func (x *SmartContractRegisterResponse) GetStatus() SmartContractRegisterResponse_Status {
	if x != nil {
		return x.Status
	}
	return SmartContractRegisterResponse_STATUS_UNSET
}

var File_src_protobuf_smartcontract_pb2_smartcontract_proto protoreflect.FileDescriptor

var file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x62, 0x32,
	0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x1e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x32, 0x0a, 0x14, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0xf9, 0x01, 0x0a, 0x1f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x4e, 0x12, 0x0c, 0x0a, 0x01, 0x54, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x54, 0x22,
	0x4f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03,
	0x22, 0x52, 0x0a, 0x1c, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x14, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x1d, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x2c, 0x0a, 0x15, 0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f,
	0x74, 0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50,
	0x01, 0x5a, 0x11, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x70, 0x62, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescOnce sync.Once
	file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescData = file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDesc
)

func file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescGZIP() []byte {
	file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescOnce.Do(func() {
		file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescData)
	})
	return file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDescData
}

var file_src_protobuf_smartcontract_pb2_smartcontract_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_protobuf_smartcontract_pb2_smartcontract_proto_goTypes = []interface{}{
	(SmartContractValidationResponse_Status)(0), // 0: SmartContractValidationResponse.Status
	(SmartContractRegisterResponse_Status)(0),   // 1: SmartContractRegisterResponse.Status
	(*SmartContractValidationRequest)(nil),      // 2: SmartContractValidationRequest
	(*SmartContractValidationResponse)(nil),     // 3: SmartContractValidationResponse
	(*SmartContractRegisterRequest)(nil),        // 4: SmartContractRegisterRequest
	(*SmartContractRegisterResponse)(nil),       // 5: SmartContractRegisterResponse
}
var file_src_protobuf_smartcontract_pb2_smartcontract_proto_depIdxs = []int32{
	0, // 0: SmartContractValidationResponse.status:type_name -> SmartContractValidationResponse.Status
	1, // 1: SmartContractRegisterResponse.status:type_name -> SmartContractRegisterResponse.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_src_protobuf_smartcontract_pb2_smartcontract_proto_init() }
func file_src_protobuf_smartcontract_pb2_smartcontract_proto_init() {
	if File_src_protobuf_smartcontract_pb2_smartcontract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartContractValidationRequest); i {
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
		file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartContractValidationResponse); i {
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
		file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartContractRegisterRequest); i {
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
		file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartContractRegisterResponse); i {
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
			RawDescriptor: file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_protobuf_smartcontract_pb2_smartcontract_proto_goTypes,
		DependencyIndexes: file_src_protobuf_smartcontract_pb2_smartcontract_proto_depIdxs,
		EnumInfos:         file_src_protobuf_smartcontract_pb2_smartcontract_proto_enumTypes,
		MessageInfos:      file_src_protobuf_smartcontract_pb2_smartcontract_proto_msgTypes,
	}.Build()
	File_src_protobuf_smartcontract_pb2_smartcontract_proto = out.File
	file_src_protobuf_smartcontract_pb2_smartcontract_proto_rawDesc = nil
	file_src_protobuf_smartcontract_pb2_smartcontract_proto_goTypes = nil
	file_src_protobuf_smartcontract_pb2_smartcontract_proto_depIdxs = nil
}
