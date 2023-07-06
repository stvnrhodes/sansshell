// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: certfetcher.proto

package certfetcher

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

type GetCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The target to query, as accepted by tls.Dial
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *GetCertRequest) Reset() {
	*x = GetCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certfetcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertRequest) ProtoMessage() {}

func (x *GetCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_certfetcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertRequest.ProtoReflect.Descriptor instead.
func (*GetCertRequest) Descriptor() ([]byte, []int) {
	return file_certfetcher_proto_rawDescGZIP(), []int{0}
}

func (x *GetCertRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type GetCertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Complete ASN.1 DER content (certificate, signature algorithm and signature) of
	// all certificates in the certificate chain presented by the target. The first
	// element is the leaf certificate.
	RawCertificateChain [][]byte `protobuf:"bytes,1,rep,name=raw_certificate_chain,json=rawCertificateChain,proto3" json:"raw_certificate_chain,omitempty"`
}

func (x *GetCertReply) Reset() {
	*x = GetCertReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certfetcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCertReply) ProtoMessage() {}

func (x *GetCertReply) ProtoReflect() protoreflect.Message {
	mi := &file_certfetcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCertReply.ProtoReflect.Descriptor instead.
func (*GetCertReply) Descriptor() ([]byte, []int) {
	return file_certfetcher_proto_rawDescGZIP(), []int{1}
}

func (x *GetCertReply) GetRawCertificateChain() [][]byte {
	if x != nil {
		return x.RawCertificateChain
	}
	return nil
}

var File_certfetcher_proto protoreflect.FileDescriptor

var file_certfetcher_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x42, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x61,
	0x77, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x13, 0x72, 0x61, 0x77, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x32, 0x4e,
	0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x6f,
	0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x61, 0x6e, 0x73,
	0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x65, 0x72, 0x74,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certfetcher_proto_rawDescOnce sync.Once
	file_certfetcher_proto_rawDescData = file_certfetcher_proto_rawDesc
)

func file_certfetcher_proto_rawDescGZIP() []byte {
	file_certfetcher_proto_rawDescOnce.Do(func() {
		file_certfetcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_certfetcher_proto_rawDescData)
	})
	return file_certfetcher_proto_rawDescData
}

var file_certfetcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_certfetcher_proto_goTypes = []interface{}{
	(*GetCertRequest)(nil), // 0: CertFetcher.GetCertRequest
	(*GetCertReply)(nil),   // 1: CertFetcher.GetCertReply
}
var file_certfetcher_proto_depIdxs = []int32{
	0, // 0: CertFetcher.CertFetcher.Get:input_type -> CertFetcher.GetCertRequest
	1, // 1: CertFetcher.CertFetcher.Get:output_type -> CertFetcher.GetCertReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_certfetcher_proto_init() }
func file_certfetcher_proto_init() {
	if File_certfetcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certfetcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertRequest); i {
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
		file_certfetcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCertReply); i {
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
			RawDescriptor: file_certfetcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_certfetcher_proto_goTypes,
		DependencyIndexes: file_certfetcher_proto_depIdxs,
		MessageInfos:      file_certfetcher_proto_msgTypes,
	}.Build()
	File_certfetcher_proto = out.File
	file_certfetcher_proto_rawDesc = nil
	file_certfetcher_proto_goTypes = nil
	file_certfetcher_proto_depIdxs = nil
}
