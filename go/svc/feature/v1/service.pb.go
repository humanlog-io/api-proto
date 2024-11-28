// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: svc/feature/v1/service.proto

package featurev1

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

type HasFeatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature string `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *HasFeatureRequest) Reset() {
	*x = HasFeatureRequest{}
	mi := &file_svc_feature_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HasFeatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasFeatureRequest) ProtoMessage() {}

func (x *HasFeatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_feature_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasFeatureRequest.ProtoReflect.Descriptor instead.
func (*HasFeatureRequest) Descriptor() ([]byte, []int) {
	return file_svc_feature_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *HasFeatureRequest) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

type HasFeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available bool `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *HasFeatureResponse) Reset() {
	*x = HasFeatureResponse{}
	mi := &file_svc_feature_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HasFeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasFeatureResponse) ProtoMessage() {}

func (x *HasFeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_feature_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasFeatureResponse.ProtoReflect.Descriptor instead.
func (*HasFeatureResponse) Descriptor() ([]byte, []int) {
	return file_svc_feature_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *HasFeatureResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

var File_svc_feature_v1_service_proto protoreflect.FileDescriptor

var file_svc_feature_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x76, 0x63, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x73, 0x76, 0x63, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2d,
	0x0a, 0x11, 0x48, 0x61, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x32, 0x0a,
	0x12, 0x48, 0x61, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x32, 0x65, 0x0a, 0x0e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x48, 0x61, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x76,
	0x63, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x46, 0x58, 0xaa, 0x02, 0x0e, 0x53,
	0x76, 0x63, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e,
	0x53, 0x76, 0x63, 0x5c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1a, 0x53, 0x76, 0x63, 0x5c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x76,
	0x63, 0x3a, 0x3a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_feature_v1_service_proto_rawDescOnce sync.Once
	file_svc_feature_v1_service_proto_rawDescData = file_svc_feature_v1_service_proto_rawDesc
)

func file_svc_feature_v1_service_proto_rawDescGZIP() []byte {
	file_svc_feature_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_feature_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_feature_v1_service_proto_rawDescData)
	})
	return file_svc_feature_v1_service_proto_rawDescData
}

var file_svc_feature_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svc_feature_v1_service_proto_goTypes = []any{
	(*HasFeatureRequest)(nil),  // 0: svc.feature.v1.HasFeatureRequest
	(*HasFeatureResponse)(nil), // 1: svc.feature.v1.HasFeatureResponse
}
var file_svc_feature_v1_service_proto_depIdxs = []int32{
	0, // 0: svc.feature.v1.FeatureService.HasFeature:input_type -> svc.feature.v1.HasFeatureRequest
	1, // 1: svc.feature.v1.FeatureService.HasFeature:output_type -> svc.feature.v1.HasFeatureResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_feature_v1_service_proto_init() }
func file_svc_feature_v1_service_proto_init() {
	if File_svc_feature_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_feature_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_feature_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_feature_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_feature_v1_service_proto_msgTypes,
	}.Build()
	File_svc_feature_v1_service_proto = out.File
	file_svc_feature_v1_service_proto_rawDesc = nil
	file_svc_feature_v1_service_proto_goTypes = nil
	file_svc_feature_v1_service_proto_depIdxs = nil
}