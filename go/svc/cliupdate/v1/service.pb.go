// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: svc/cliupdate/v1/service.proto

package cliupdatev1

import (
	v1 "github.com/humanlogio/api/go/types/v1"
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

type GetNextUpdateRequest struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	ProjectName            string                 `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	CurrentVersion         *v1.Version            `protobuf:"bytes,2,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	MachineArchitecture    string                 `protobuf:"bytes,3,opt,name=machine_architecture,json=machineArchitecture,proto3" json:"machine_architecture,omitempty"`
	MachineOperatingSystem string                 `protobuf:"bytes,4,opt,name=machine_operating_system,json=machineOperatingSystem,proto3" json:"machine_operating_system,omitempty"`
	ReleaseChannelName     *string                `protobuf:"bytes,5,opt,name=release_channel_name,json=releaseChannelName,proto3,oneof" json:"release_channel_name,omitempty"`
	Meta                   *v1.ReqMeta            `protobuf:"bytes,1000,opt,name=meta,proto3" json:"meta,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetNextUpdateRequest) Reset() {
	*x = GetNextUpdateRequest{}
	mi := &file_svc_cliupdate_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNextUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextUpdateRequest) ProtoMessage() {}

func (x *GetNextUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_cliupdate_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextUpdateRequest.ProtoReflect.Descriptor instead.
func (*GetNextUpdateRequest) Descriptor() ([]byte, []int) {
	return file_svc_cliupdate_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetNextUpdateRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *GetNextUpdateRequest) GetCurrentVersion() *v1.Version {
	if x != nil {
		return x.CurrentVersion
	}
	return nil
}

func (x *GetNextUpdateRequest) GetMachineArchitecture() string {
	if x != nil {
		return x.MachineArchitecture
	}
	return ""
}

func (x *GetNextUpdateRequest) GetMachineOperatingSystem() string {
	if x != nil {
		return x.MachineOperatingSystem
	}
	return ""
}

func (x *GetNextUpdateRequest) GetReleaseChannelName() string {
	if x != nil && x.ReleaseChannelName != nil {
		return *x.ReleaseChannelName
	}
	return ""
}

func (x *GetNextUpdateRequest) GetMeta() *v1.ReqMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GetNextUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NextVersion   *v1.Version            `protobuf:"bytes,1,opt,name=next_version,json=nextVersion,proto3" json:"next_version,omitempty"`
	NextArtifact  *v1.VersionArtifact    `protobuf:"bytes,2,opt,name=next_artifact,json=nextArtifact,proto3" json:"next_artifact,omitempty"`
	Meta          *v1.ResMeta            `protobuf:"bytes,1000,opt,name=meta,proto3" json:"meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNextUpdateResponse) Reset() {
	*x = GetNextUpdateResponse{}
	mi := &file_svc_cliupdate_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNextUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNextUpdateResponse) ProtoMessage() {}

func (x *GetNextUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_cliupdate_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNextUpdateResponse.ProtoReflect.Descriptor instead.
func (*GetNextUpdateResponse) Descriptor() ([]byte, []int) {
	return file_svc_cliupdate_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetNextUpdateResponse) GetNextVersion() *v1.Version {
	if x != nil {
		return x.NextVersion
	}
	return nil
}

func (x *GetNextUpdateResponse) GetNextArtifact() *v1.VersionArtifact {
	if x != nil {
		return x.NextArtifact
	}
	return nil
}

func (x *GetNextUpdateResponse) GetMeta() *v1.ResMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_svc_cliupdate_v1_service_proto protoreflect.FileDescriptor

var file_svc_cliupdate_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x76, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x76, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xda, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x14, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x35, 0x0a, 0x14, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x12, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb5, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x0c, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x26, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x32, 0x73, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x63, 0x6c,
	0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc1, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x75, 0x6d, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x43, 0x58, 0xaa, 0x02, 0x10, 0x53, 0x76, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x53, 0x76, 0x63, 0x5c, 0x43, 0x6c,
	0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x53, 0x76, 0x63,
	0x5c, 0x43, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x76, 0x63, 0x3a,
	0x3a, 0x43, 0x6c, 0x69, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_cliupdate_v1_service_proto_rawDescOnce sync.Once
	file_svc_cliupdate_v1_service_proto_rawDescData = file_svc_cliupdate_v1_service_proto_rawDesc
)

func file_svc_cliupdate_v1_service_proto_rawDescGZIP() []byte {
	file_svc_cliupdate_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_cliupdate_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_cliupdate_v1_service_proto_rawDescData)
	})
	return file_svc_cliupdate_v1_service_proto_rawDescData
}

var file_svc_cliupdate_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svc_cliupdate_v1_service_proto_goTypes = []any{
	(*GetNextUpdateRequest)(nil),  // 0: svc.cliupdate.v1.GetNextUpdateRequest
	(*GetNextUpdateResponse)(nil), // 1: svc.cliupdate.v1.GetNextUpdateResponse
	(*v1.Version)(nil),            // 2: types.v1.Version
	(*v1.ReqMeta)(nil),            // 3: types.v1.ReqMeta
	(*v1.VersionArtifact)(nil),    // 4: types.v1.VersionArtifact
	(*v1.ResMeta)(nil),            // 5: types.v1.ResMeta
}
var file_svc_cliupdate_v1_service_proto_depIdxs = []int32{
	2, // 0: svc.cliupdate.v1.GetNextUpdateRequest.current_version:type_name -> types.v1.Version
	3, // 1: svc.cliupdate.v1.GetNextUpdateRequest.meta:type_name -> types.v1.ReqMeta
	2, // 2: svc.cliupdate.v1.GetNextUpdateResponse.next_version:type_name -> types.v1.Version
	4, // 3: svc.cliupdate.v1.GetNextUpdateResponse.next_artifact:type_name -> types.v1.VersionArtifact
	5, // 4: svc.cliupdate.v1.GetNextUpdateResponse.meta:type_name -> types.v1.ResMeta
	0, // 5: svc.cliupdate.v1.UpdateService.GetNextUpdate:input_type -> svc.cliupdate.v1.GetNextUpdateRequest
	1, // 6: svc.cliupdate.v1.UpdateService.GetNextUpdate:output_type -> svc.cliupdate.v1.GetNextUpdateResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_svc_cliupdate_v1_service_proto_init() }
func file_svc_cliupdate_v1_service_proto_init() {
	if File_svc_cliupdate_v1_service_proto != nil {
		return
	}
	file_svc_cliupdate_v1_service_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_cliupdate_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_cliupdate_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_cliupdate_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_cliupdate_v1_service_proto_msgTypes,
	}.Build()
	File_svc_cliupdate_v1_service_proto = out.File
	file_svc_cliupdate_v1_service_proto_rawDesc = nil
	file_svc_cliupdate_v1_service_proto_goTypes = nil
	file_svc_cliupdate_v1_service_proto_depIdxs = nil
}
