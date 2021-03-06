// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: genlink.proto

package api

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

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_genlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_genlink_proto_rawDescGZIP(), []int{0}
}

func (x *URL) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type ShortURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL string `protobuf:"bytes,1,opt,name=ShortURL,proto3" json:"ShortURL,omitempty"`
}

func (x *ShortURL) Reset() {
	*x = ShortURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortURL) ProtoMessage() {}

func (x *ShortURL) ProtoReflect() protoreflect.Message {
	mi := &file_genlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortURL.ProtoReflect.Descriptor instead.
func (*ShortURL) Descriptor() ([]byte, []int) {
	return file_genlink_proto_rawDescGZIP(), []int{1}
}

func (x *ShortURL) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

var File_genlink_proto protoreflect.FileDescriptor

var file_genlink_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22,
	0x26, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x32, 0x58, 0x0a, 0x07, 0x67, 0x65, 0x6e, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x27, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x52, 0x4c, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x52, 0x4c, 0x22,
	0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_genlink_proto_rawDescOnce sync.Once
	file_genlink_proto_rawDescData = file_genlink_proto_rawDesc
)

func file_genlink_proto_rawDescGZIP() []byte {
	file_genlink_proto_rawDescOnce.Do(func() {
		file_genlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_genlink_proto_rawDescData)
	})
	return file_genlink_proto_rawDescData
}

var file_genlink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_genlink_proto_goTypes = []interface{}{
	(*URL)(nil),      // 0: proto.URL
	(*ShortURL)(nil), // 1: proto.ShortURL
}
var file_genlink_proto_depIdxs = []int32{
	0, // 0: proto.genLink.Create:input_type -> proto.URL
	1, // 1: proto.genLink.Get:input_type -> proto.ShortURL
	1, // 2: proto.genLink.Create:output_type -> proto.ShortURL
	0, // 3: proto.genLink.Get:output_type -> proto.URL
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_genlink_proto_init() }
func file_genlink_proto_init() {
	if File_genlink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_genlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
		file_genlink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortURL); i {
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
			RawDescriptor: file_genlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_genlink_proto_goTypes,
		DependencyIndexes: file_genlink_proto_depIdxs,
		MessageInfos:      file_genlink_proto_msgTypes,
	}.Build()
	File_genlink_proto = out.File
	file_genlink_proto_rawDesc = nil
	file_genlink_proto_goTypes = nil
	file_genlink_proto_depIdxs = nil
}
