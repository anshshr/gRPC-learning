// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: demo.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnshResposne struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Data             string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Age              int32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Height           int32                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Marks            int32                  `protobuf:"varint,4,opt,name=marks,proto3" json:"marks,omitempty"`
	IsResultDeclared bool                   `protobuf:"varint,5,opt,name=isResultDeclared,proto3" json:"isResultDeclared,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AnshResposne) Reset() {
	*x = AnshResposne{}
	mi := &file_demo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnshResposne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnshResposne) ProtoMessage() {}

func (x *AnshResposne) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnshResposne.ProtoReflect.Descriptor instead.
func (*AnshResposne) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

func (x *AnshResposne) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *AnshResposne) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *AnshResposne) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *AnshResposne) GetMarks() int32 {
	if x != nil {
		return x.Marks
	}
	return 0
}

func (x *AnshResposne) GetIsResultDeclared() bool {
	if x != nil {
		return x.IsResultDeclared
	}
	return false
}

type RonakReply struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Data             string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Age              int32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Height           int32                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Marks            int32                  `protobuf:"varint,4,opt,name=marks,proto3" json:"marks,omitempty"`
	IsResultDeclared bool                   `protobuf:"varint,5,opt,name=isResultDeclared,proto3" json:"isResultDeclared,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *RonakReply) Reset() {
	*x = RonakReply{}
	mi := &file_demo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RonakReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RonakReply) ProtoMessage() {}

func (x *RonakReply) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RonakReply.ProtoReflect.Descriptor instead.
func (*RonakReply) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *RonakReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *RonakReply) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *RonakReply) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *RonakReply) GetMarks() int32 {
	if x != nil {
		return x.Marks
	}
	return 0
}

func (x *RonakReply) GetIsResultDeclared() bool {
	if x != nil {
		return x.IsResultDeclared
	}
	return false
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65,
	0x6d, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x41, 0x6e, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x73, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x10, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x63, 0x6c, 0x61,
	0x72, 0x65, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x6e, 0x61, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72,
	0x65, 0x64, 0x32, 0x3b, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x41,
	0x6e, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73, 0x6e, 0x65, 0x1a, 0x10, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x52, 0x6f, 0x6e, 0x61, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData []byte
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_demo_proto_rawDesc), len(file_demo_proto_rawDesc)))
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demo_proto_goTypes = []any{
	(*AnshResposne)(nil), // 0: demo.AnshResposne
	(*RonakReply)(nil),   // 1: demo.RonakReply
}
var file_demo_proto_depIdxs = []int32{
	0, // 0: demo.chat.StartChat:input_type -> demo.AnshResposne
	1, // 1: demo.chat.StartChat:output_type -> demo.RonakReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_demo_proto_rawDesc), len(file_demo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}
