// Copyright (c) 2019 Snowflake Inc. All rights reserved.
//
//Licensed under the Apache License, Version 2.0 (the
//"License"); you may not use this file except in compliance
//with the License.  You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing,
//software distributed under the License is distributed on an
//"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//KIND, either express or implied.  See the License for the
//specific language governing permissions and limitations
//under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.27.1
// source: testservice.proto

package testdata

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

type MyNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fine           string `protobuf:"bytes,1,opt,name=fine,proto3" json:"fine,omitempty"`
	Sensitive      string `protobuf:"bytes,2,opt,name=sensitive,proto3" json:"sensitive,omitempty"`
	SensitiveBytes []byte `protobuf:"bytes,3,opt,name=sensitive_bytes,json=sensitiveBytes,proto3" json:"sensitive_bytes,omitempty"`
	// Types that are assignable to OneofField:
	//
	//	*MyNested_OneofFine
	//	*MyNested_OneofSensitive
	OneofField             isMyNested_OneofField `protobuf_oneof:"oneof_field"`
	SensitiveRepeatedBytes [][]byte              `protobuf:"bytes,6,rep,name=sensitive_repeated_bytes,json=sensitiveRepeatedBytes,proto3" json:"sensitive_repeated_bytes,omitempty"`
	SensitiveInt           int64                 `protobuf:"varint,7,opt,name=sensitive_int,json=sensitiveInt,proto3" json:"sensitive_int,omitempty"`
}

func (x *MyNested) Reset() {
	*x = MyNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyNested) ProtoMessage() {}

func (x *MyNested) ProtoReflect() protoreflect.Message {
	mi := &file_testservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyNested.ProtoReflect.Descriptor instead.
func (*MyNested) Descriptor() ([]byte, []int) {
	return file_testservice_proto_rawDescGZIP(), []int{0}
}

func (x *MyNested) GetFine() string {
	if x != nil {
		return x.Fine
	}
	return ""
}

func (x *MyNested) GetSensitive() string {
	if x != nil {
		return x.Sensitive
	}
	return ""
}

func (x *MyNested) GetSensitiveBytes() []byte {
	if x != nil {
		return x.SensitiveBytes
	}
	return nil
}

func (m *MyNested) GetOneofField() isMyNested_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (x *MyNested) GetOneofFine() string {
	if x, ok := x.GetOneofField().(*MyNested_OneofFine); ok {
		return x.OneofFine
	}
	return ""
}

func (x *MyNested) GetOneofSensitive() string {
	if x, ok := x.GetOneofField().(*MyNested_OneofSensitive); ok {
		return x.OneofSensitive
	}
	return ""
}

func (x *MyNested) GetSensitiveRepeatedBytes() [][]byte {
	if x != nil {
		return x.SensitiveRepeatedBytes
	}
	return nil
}

func (x *MyNested) GetSensitiveInt() int64 {
	if x != nil {
		return x.SensitiveInt
	}
	return 0
}

type isMyNested_OneofField interface {
	isMyNested_OneofField()
}

type MyNested_OneofFine struct {
	OneofFine string `protobuf:"bytes,4,opt,name=oneof_fine,json=oneofFine,proto3,oneof"`
}

type MyNested_OneofSensitive struct {
	OneofSensitive string `protobuf:"bytes,5,opt,name=oneof_sensitive,json=oneofSensitive,proto3,oneof"`
}

func (*MyNested_OneofFine) isMyNested_OneofField() {}

func (*MyNested_OneofSensitive) isMyNested_OneofField() {}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input      string               `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	ListScalar []string             `protobuf:"bytes,2,rep,name=list_scalar,json=listScalar,proto3" json:"list_scalar,omitempty"`
	ListMsg    []*MyNested          `protobuf:"bytes,3,rep,name=list_msg,json=listMsg,proto3" json:"list_msg,omitempty"`
	MapScalar  map[string]string    `protobuf:"bytes,4,rep,name=map_scalar,json=mapScalar,proto3" json:"map_scalar,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapMsg     map[string]*MyNested `protobuf:"bytes,5,rep,name=map_msg,json=mapMsg,proto3" json:"map_msg,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_testservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_testservice_proto_rawDescGZIP(), []int{1}
}

func (x *TestRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TestRequest) GetListScalar() []string {
	if x != nil {
		return x.ListScalar
	}
	return nil
}

func (x *TestRequest) GetListMsg() []*MyNested {
	if x != nil {
		return x.ListMsg
	}
	return nil
}

func (x *TestRequest) GetMapScalar() map[string]string {
	if x != nil {
		return x.MapScalar
	}
	return nil
}

func (x *TestRequest) GetMapMsg() map[string]*MyNested {
	if x != nil {
		return x.MapMsg
	}
	return nil
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_testservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_testservice_proto_rawDescGZIP(), []int{2}
}

func (x *TestResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_testservice_proto protoreflect.FileDescriptor

var file_testservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb8, 0x02,
	0x0a, 0x08, 0x4d, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x21,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0x80, 0x01, 0x01, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0x80, 0x01, 0x01, 0x52,
	0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x1f, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x46, 0x69, 0x6e, 0x65,
	0x12, 0x2e, 0x0a, 0x0f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0x80, 0x01, 0x01, 0x48, 0x00,
	0x52, 0x0e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x3d, 0x0a, 0x18, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0c, 0x42, 0x03, 0x80, 0x01, 0x01, 0x52, 0x16, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x28, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0x80, 0x01, 0x01, 0x52, 0x0c, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x81, 0x03, 0x0a, 0x0b, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12,
	0x2d, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x79, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x43,
	0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x4d,
	0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x4d, 0x73, 0x67, 0x1a,
	0x3c, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a,
	0x0b, 0x4d, 0x61, 0x70, 0x4d, 0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x0c,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x32, 0xa0, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x55, 0x6e, 0x61, 0x72,
	0x79, 0x12, 0x15, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x10, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x12, 0x45, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x15, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x1b, 0x0a, 0x19, 0x54, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x4c, 0x61, 0x62,
	0x73, 0x2f, 0x73, 0x61, 0x6e, 0x73, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_testservice_proto_rawDescOnce sync.Once
	file_testservice_proto_rawDescData = file_testservice_proto_rawDesc
)

func file_testservice_proto_rawDescGZIP() []byte {
	file_testservice_proto_rawDescOnce.Do(func() {
		file_testservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_testservice_proto_rawDescData)
	})
	return file_testservice_proto_rawDescData
}

var file_testservice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_testservice_proto_goTypes = []interface{}{
	(*MyNested)(nil),     // 0: Testdata.MyNested
	(*TestRequest)(nil),  // 1: Testdata.TestRequest
	(*TestResponse)(nil), // 2: Testdata.TestResponse
	nil,                  // 3: Testdata.TestRequest.MapScalarEntry
	nil,                  // 4: Testdata.TestRequest.MapMsgEntry
}
var file_testservice_proto_depIdxs = []int32{
	0, // 0: Testdata.TestRequest.list_msg:type_name -> Testdata.MyNested
	3, // 1: Testdata.TestRequest.map_scalar:type_name -> Testdata.TestRequest.MapScalarEntry
	4, // 2: Testdata.TestRequest.map_msg:type_name -> Testdata.TestRequest.MapMsgEntry
	0, // 3: Testdata.TestRequest.MapMsgEntry.value:type_name -> Testdata.MyNested
	1, // 4: Testdata.TestService.TestUnary:input_type -> Testdata.TestRequest
	1, // 5: Testdata.TestService.TestServerStream:input_type -> Testdata.TestRequest
	1, // 6: Testdata.TestService.TestClientStream:input_type -> Testdata.TestRequest
	1, // 7: Testdata.TestService.TestBidiStream:input_type -> Testdata.TestRequest
	2, // 8: Testdata.TestService.TestUnary:output_type -> Testdata.TestResponse
	2, // 9: Testdata.TestService.TestServerStream:output_type -> Testdata.TestResponse
	2, // 10: Testdata.TestService.TestClientStream:output_type -> Testdata.TestResponse
	2, // 11: Testdata.TestService.TestBidiStream:output_type -> Testdata.TestResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_testservice_proto_init() }
func file_testservice_proto_init() {
	if File_testservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyNested); i {
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
		file_testservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_testservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
	file_testservice_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MyNested_OneofFine)(nil),
		(*MyNested_OneofSensitive)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_testservice_proto_goTypes,
		DependencyIndexes: file_testservice_proto_depIdxs,
		MessageInfos:      file_testservice_proto_msgTypes,
	}.Build()
	File_testservice_proto = out.File
	file_testservice_proto_rawDesc = nil
	file_testservice_proto_goTypes = nil
	file_testservice_proto_depIdxs = nil
}
