// Copyright (c) 2023 Snowflake Inc. All rights reserved.
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
// 	protoc        v5.26.1
// source: tlsinfo.proto

package tlsinfo

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

type TLSCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddress      string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"` // Server's address, including port (e.g. "example.com:443")
	ServerName         string `protobuf:"bytes,2,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	InsecureSkipVerify bool   `protobuf:"varint,3,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify,omitempty"`
}

func (x *TLSCertificateRequest) Reset() {
	*x = TLSCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlsinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSCertificateRequest) ProtoMessage() {}

func (x *TLSCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tlsinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSCertificateRequest.ProtoReflect.Descriptor instead.
func (*TLSCertificateRequest) Descriptor() ([]byte, []int) {
	return file_tlsinfo_proto_rawDescGZIP(), []int{0}
}

func (x *TLSCertificateRequest) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *TLSCertificateRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *TLSCertificateRequest) GetInsecureSkipVerify() bool {
	if x != nil {
		return x.InsecureSkipVerify
	}
	return false
}

type TLSCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer      string   `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject     string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	NotBefore   int64    `protobuf:"varint,3,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	NotAfter    int64    `protobuf:"varint,4,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	DnsNames    []string `protobuf:"bytes,5,rep,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty"`
	IpAddresses []string `protobuf:"bytes,6,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
}

func (x *TLSCertificate) Reset() {
	*x = TLSCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlsinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSCertificate) ProtoMessage() {}

func (x *TLSCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_tlsinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSCertificate.ProtoReflect.Descriptor instead.
func (*TLSCertificate) Descriptor() ([]byte, []int) {
	return file_tlsinfo_proto_rawDescGZIP(), []int{1}
}

func (x *TLSCertificate) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *TLSCertificate) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *TLSCertificate) GetNotBefore() int64 {
	if x != nil {
		return x.NotBefore
	}
	return 0
}

func (x *TLSCertificate) GetNotAfter() int64 {
	if x != nil {
		return x.NotAfter
	}
	return 0
}

func (x *TLSCertificate) GetDnsNames() []string {
	if x != nil {
		return x.DnsNames
	}
	return nil
}

func (x *TLSCertificate) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

type TLSCertificateChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificates []*TLSCertificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
}

func (x *TLSCertificateChain) Reset() {
	*x = TLSCertificateChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tlsinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSCertificateChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSCertificateChain) ProtoMessage() {}

func (x *TLSCertificateChain) ProtoReflect() protoreflect.Message {
	mi := &file_tlsinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSCertificateChain.ProtoReflect.Descriptor instead.
func (*TLSCertificateChain) Descriptor() ([]byte, []int) {
	return file_tlsinfo_proto_rawDescGZIP(), []int{2}
}

func (x *TLSCertificateChain) GetCertificates() []*TLSCertificate {
	if x != nil {
		return x.Certificates
	}
	return nil
}

var File_tlsinfo_proto protoreflect.FileDescriptor

var file_tlsinfo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x6c, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x54, 0x4c, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x15, 0x54, 0x4c, 0x53,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0xbe, 0x01, 0x0a,
	0x0e, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6e, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x70,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x52, 0x0a,
	0x13, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x12, 0x3b, 0x0a, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x54, 0x4c, 0x53,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x32, 0x5c, 0x0a, 0x07, 0x54, 0x4c, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x54, 0x4c, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x4c, 0x53, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x54, 0x4c, 0x53, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x4c, 0x53, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e,
	0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x61, 0x6e,
	0x73, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2f, 0x74, 0x6c, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tlsinfo_proto_rawDescOnce sync.Once
	file_tlsinfo_proto_rawDescData = file_tlsinfo_proto_rawDesc
)

func file_tlsinfo_proto_rawDescGZIP() []byte {
	file_tlsinfo_proto_rawDescOnce.Do(func() {
		file_tlsinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_tlsinfo_proto_rawDescData)
	})
	return file_tlsinfo_proto_rawDescData
}

var file_tlsinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tlsinfo_proto_goTypes = []interface{}{
	(*TLSCertificateRequest)(nil), // 0: TLSInfo.TLSCertificateRequest
	(*TLSCertificate)(nil),        // 1: TLSInfo.TLSCertificate
	(*TLSCertificateChain)(nil),   // 2: TLSInfo.TLSCertificateChain
}
var file_tlsinfo_proto_depIdxs = []int32{
	1, // 0: TLSInfo.TLSCertificateChain.certificates:type_name -> TLSInfo.TLSCertificate
	0, // 1: TLSInfo.TLSInfo.GetTLSCertificate:input_type -> TLSInfo.TLSCertificateRequest
	2, // 2: TLSInfo.TLSInfo.GetTLSCertificate:output_type -> TLSInfo.TLSCertificateChain
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tlsinfo_proto_init() }
func file_tlsinfo_proto_init() {
	if File_tlsinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tlsinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSCertificateRequest); i {
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
		file_tlsinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSCertificate); i {
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
		file_tlsinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSCertificateChain); i {
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
			RawDescriptor: file_tlsinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tlsinfo_proto_goTypes,
		DependencyIndexes: file_tlsinfo_proto_depIdxs,
		MessageInfos:      file_tlsinfo_proto_msgTypes,
	}.Build()
	File_tlsinfo_proto = out.File
	file_tlsinfo_proto_rawDesc = nil
	file_tlsinfo_proto_goTypes = nil
	file_tlsinfo_proto_depIdxs = nil
}
