// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: types/v1/meta.proto

package typesv1

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

type ReqMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId int64 `protobuf:"varint,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	MachineId     int64 `protobuf:"varint,2,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
}

func (x *ReqMeta) Reset() {
	*x = ReqMeta{}
	mi := &file_types_v1_meta_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMeta) ProtoMessage() {}

func (x *ReqMeta) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_meta_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMeta.ProtoReflect.Descriptor instead.
func (*ReqMeta) Descriptor() ([]byte, []int) {
	return file_types_v1_meta_proto_rawDescGZIP(), []int{0}
}

func (x *ReqMeta) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *ReqMeta) GetMachineId() int64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

type ResMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId int64 `protobuf:"varint,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	MachineId     int64 `protobuf:"varint,2,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
}

func (x *ResMeta) Reset() {
	*x = ResMeta{}
	mi := &file_types_v1_meta_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMeta) ProtoMessage() {}

func (x *ResMeta) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_meta_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMeta.ProtoReflect.Descriptor instead.
func (*ResMeta) Descriptor() ([]byte, []int) {
	return file_types_v1_meta_proto_rawDescGZIP(), []int{1}
}

func (x *ResMeta) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *ResMeta) GetMachineId() int64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

var File_types_v1_meta_proto protoreflect.FileDescriptor

var file_types_v1_meta_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22,
	0x4f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x22, 0x4f, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49,
	0x64, 0x42, 0x89, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x09, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_meta_proto_rawDescOnce sync.Once
	file_types_v1_meta_proto_rawDescData = file_types_v1_meta_proto_rawDesc
)

func file_types_v1_meta_proto_rawDescGZIP() []byte {
	file_types_v1_meta_proto_rawDescOnce.Do(func() {
		file_types_v1_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_meta_proto_rawDescData)
	})
	return file_types_v1_meta_proto_rawDescData
}

var file_types_v1_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_v1_meta_proto_goTypes = []any{
	(*ReqMeta)(nil), // 0: types.v1.ReqMeta
	(*ResMeta)(nil), // 1: types.v1.ResMeta
}
var file_types_v1_meta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_v1_meta_proto_init() }
func file_types_v1_meta_proto_init() {
	if File_types_v1_meta_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_v1_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_meta_proto_goTypes,
		DependencyIndexes: file_types_v1_meta_proto_depIdxs,
		MessageInfos:      file_types_v1_meta_proto_msgTypes,
	}.Build()
	File_types_v1_meta_proto = out.File
	file_types_v1_meta_proto_rawDesc = nil
	file_types_v1_meta_proto_goTypes = nil
	file_types_v1_meta_proto_depIdxs = nil
}
