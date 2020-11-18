// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: echoer.proto

package echoer

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Enum int32

const (
	Enum_One  Enum = 0
	Enum_Two  Enum = 1
	Enum_Tree Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "One",
		1: "Two",
		2: "Tree",
	}
	Enum_value = map[string]int32{
		"One":  0,
		"Two":  1,
		"Tree": 2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_echoer_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_echoer_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_echoer_proto_rawDescGZIP(), []int{0}
}

type StringMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value   string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Version Enum   `protobuf:"varint,2,opt,name=version,json=ver,proto3,enum=echoer.Enum" json:"version,omitempty"`
}

func (x *StringMessage) Reset() {
	*x = StringMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echoer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMessage) ProtoMessage() {}

func (x *StringMessage) ProtoReflect() protoreflect.Message {
	mi := &file_echoer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMessage.ProtoReflect.Descriptor instead.
func (*StringMessage) Descriptor() ([]byte, []int) {
	return file_echoer_proto_rawDescGZIP(), []int{0}
}

func (x *StringMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *StringMessage) GetVersion() Enum {
	if x != nil {
		return x.Version
	}
	return Enum_One
}

var File_echoer_proto protoreflect.FileDescriptor

var file_echoer_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x65,
	0x63, 0x68, 0x6f, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x03, 0x76, 0x65, 0x72, 0x2a,
	0x22, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x72, 0x65,
	0x65, 0x10, 0x02, 0x32, 0x5b, 0x0a, 0x06, 0x45, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x12, 0x51, 0x0a,
	0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x15, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x65,
	0x63, 0x68, 0x6f, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a,
	0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x61, 0x72, 0x61, 0x73, 0x6d, 0x61, 0x74, 0x73, 0x79, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x65, 0x63, 0x68, 0x6f, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_echoer_proto_rawDescOnce sync.Once
	file_echoer_proto_rawDescData = file_echoer_proto_rawDesc
)

func file_echoer_proto_rawDescGZIP() []byte {
	file_echoer_proto_rawDescOnce.Do(func() {
		file_echoer_proto_rawDescData = protoimpl.X.CompressGZIP(file_echoer_proto_rawDescData)
	})
	return file_echoer_proto_rawDescData
}

var file_echoer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_echoer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_echoer_proto_goTypes = []interface{}{
	(Enum)(0),             // 0: echoer.Enum
	(*StringMessage)(nil), // 1: echoer.StringMessage
}
var file_echoer_proto_depIdxs = []int32{
	0, // 0: echoer.StringMessage.version:type_name -> echoer.Enum
	1, // 1: echoer.Echoer.Echo:input_type -> echoer.StringMessage
	1, // 2: echoer.Echoer.Echo:output_type -> echoer.StringMessage
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_echoer_proto_init() }
func file_echoer_proto_init() {
	if File_echoer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echoer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMessage); i {
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
			RawDescriptor: file_echoer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echoer_proto_goTypes,
		DependencyIndexes: file_echoer_proto_depIdxs,
		EnumInfos:         file_echoer_proto_enumTypes,
		MessageInfos:      file_echoer_proto_msgTypes,
	}.Build()
	File_echoer_proto = out.File
	file_echoer_proto_rawDesc = nil
	file_echoer_proto_goTypes = nil
	file_echoer_proto_depIdxs = nil
}
