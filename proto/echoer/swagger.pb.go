// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: echoer/swagger.proto

package echoer

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_echoer_swagger_proto protoreflect.FileDescriptor

var file_echoer_swagger_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xf7,
	0x02, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61,
	0x72, 0x61, 0x73, 0x6d, 0x61, 0x74, 0x73, 0x79, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x92, 0x41, 0xba, 0x02, 0x12, 0xf6,
	0x01, 0x0a, 0x13, 0x41, 0x20, 0x42, 0x69, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x45, 0x76, 0x65, 0x72,
	0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x58, 0x0a, 0x14, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x10,
	0x6e, 0x6f, 0x6e, 0x65, 0x40, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2a, 0x5e, 0x0a, 0x14, 0x42, 0x53, 0x44, 0x20, 0x33, 0x2d, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65,
	0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2e, 0x74, 0x78, 0x74,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x3a, 0x20, 0x0a, 0x15, 0x78, 0x2d, 0x73, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x07,
	0x1a, 0x05, 0x79, 0x61, 0x64, 0x64, 0x61, 0x5a, 0x32, 0x0a, 0x30, 0x0a, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x08, 0x02, 0x12, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x09,
	0x58, 0x2d, 0x41, 0x50, 0x49, 0x2d, 0x4b, 0x65, 0x79, 0x20, 0x02, 0x62, 0x0b, 0x0a, 0x09, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_echoer_swagger_proto_goTypes = []interface{}{}
var file_echoer_swagger_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echoer_swagger_proto_init() }
func file_echoer_swagger_proto_init() {
	if File_echoer_swagger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_echoer_swagger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_echoer_swagger_proto_goTypes,
		DependencyIndexes: file_echoer_swagger_proto_depIdxs,
	}.Build()
	File_echoer_swagger_proto = out.File
	file_echoer_swagger_proto_rawDesc = nil
	file_echoer_swagger_proto_goTypes = nil
	file_echoer_swagger_proto_depIdxs = nil
}