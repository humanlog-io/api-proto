// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: svc/user/v1/service.proto

package userv1

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

type WhoamiRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WhoamiRequest) Reset() {
	*x = WhoamiRequest{}
	mi := &file_svc_user_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WhoamiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoamiRequest) ProtoMessage() {}

func (x *WhoamiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoamiRequest.ProtoReflect.Descriptor instead.
func (*WhoamiRequest) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{0}
}

type WhoamiResponse struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	User                *v1.User               `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	CurrentOrganization *v1.Organization       `protobuf:"bytes,2,opt,name=current_organization,json=currentOrganization,proto3" json:"current_organization,omitempty"`
	DefaultOrganization *v1.Organization       `protobuf:"bytes,3,opt,name=default_organization,json=defaultOrganization,proto3" json:"default_organization,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *WhoamiResponse) Reset() {
	*x = WhoamiResponse{}
	mi := &file_svc_user_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WhoamiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoamiResponse) ProtoMessage() {}

func (x *WhoamiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoamiResponse.ProtoReflect.Descriptor instead.
func (*WhoamiResponse) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *WhoamiResponse) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *WhoamiResponse) GetCurrentOrganization() *v1.Organization {
	if x != nil {
		return x.CurrentOrganization
	}
	return nil
}

func (x *WhoamiResponse) GetDefaultOrganization() *v1.Organization {
	if x != nil {
		return x.DefaultOrganization
	}
	return nil
}

type GetLogoutURLRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReturnTo      string                 `protobuf:"bytes,1,opt,name=return_to,json=returnTo,proto3" json:"return_to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLogoutURLRequest) Reset() {
	*x = GetLogoutURLRequest{}
	mi := &file_svc_user_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLogoutURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogoutURLRequest) ProtoMessage() {}

func (x *GetLogoutURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogoutURLRequest.ProtoReflect.Descriptor instead.
func (*GetLogoutURLRequest) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLogoutURLRequest) GetReturnTo() string {
	if x != nil {
		return x.ReturnTo
	}
	return ""
}

type GetLogoutURLResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LogoutUrl     string                 `protobuf:"bytes,1,opt,name=logout_url,json=logoutUrl,proto3" json:"logout_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLogoutURLResponse) Reset() {
	*x = GetLogoutURLResponse{}
	mi := &file_svc_user_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLogoutURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogoutURLResponse) ProtoMessage() {}

func (x *GetLogoutURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogoutURLResponse.ProtoReflect.Descriptor instead.
func (*GetLogoutURLResponse) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogoutURLResponse) GetLogoutUrl() string {
	if x != nil {
		return x.LogoutUrl
	}
	return ""
}

type CreateOrganizationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrganizationRequest) Reset() {
	*x = CreateOrganizationRequest{}
	mi := &file_svc_user_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrganizationRequest) ProtoMessage() {}

func (x *CreateOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrganizationRequest.ProtoReflect.Descriptor instead.
func (*CreateOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrganizationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateOrganizationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Organization  *v1.Organization       `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrganizationResponse) Reset() {
	*x = CreateOrganizationResponse{}
	mi := &file_svc_user_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrganizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrganizationResponse) ProtoMessage() {}

func (x *CreateOrganizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrganizationResponse.ProtoReflect.Descriptor instead.
func (*CreateOrganizationResponse) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateOrganizationResponse) GetOrganization() *v1.Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

type ListOrganizationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cursor        *v1.Cursor             `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrganizationRequest) Reset() {
	*x = ListOrganizationRequest{}
	mi := &file_svc_user_v1_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationRequest) ProtoMessage() {}

func (x *ListOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationRequest.ProtoReflect.Descriptor instead.
func (*ListOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrganizationRequest) GetCursor() *v1.Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ListOrganizationRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListOrganizationResponse struct {
	state               protoimpl.MessageState               `protogen:"open.v1"`
	Next                *v1.Cursor                           `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items               []*ListOrganizationResponse_ListItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	DefaultOrganization *v1.Organization                     `protobuf:"bytes,3,opt,name=default_organization,json=defaultOrganization,proto3" json:"default_organization,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ListOrganizationResponse) Reset() {
	*x = ListOrganizationResponse{}
	mi := &file_svc_user_v1_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrganizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationResponse) ProtoMessage() {}

func (x *ListOrganizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationResponse.ProtoReflect.Descriptor instead.
func (*ListOrganizationResponse) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrganizationResponse) GetNext() *v1.Cursor {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *ListOrganizationResponse) GetItems() []*ListOrganizationResponse_ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListOrganizationResponse) GetDefaultOrganization() *v1.Organization {
	if x != nil {
		return x.DefaultOrganization
	}
	return nil
}

type ListOrganizationResponse_ListItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Organization  *v1.Organization       `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrganizationResponse_ListItem) Reset() {
	*x = ListOrganizationResponse_ListItem{}
	mi := &file_svc_user_v1_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrganizationResponse_ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationResponse_ListItem) ProtoMessage() {}

func (x *ListOrganizationResponse_ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_svc_user_v1_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationResponse_ListItem.ProtoReflect.Descriptor instead.
func (*ListOrganizationResponse_ListItem) Descriptor() ([]byte, []int) {
	return file_svc_user_v1_service_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListOrganizationResponse_ListItem) GetOrganization() *v1.Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

var File_svc_user_v1_service_proto protoreflect.FileDescriptor

var file_svc_user_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x76, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x57, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x14, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x54, 0x6f, 0x22, 0x35, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x99, 0x02, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x49, 0x0a, 0x14, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x13, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x46, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x3a, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xf5, 0x02, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x06,
	0x57, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x12, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x52,
	0x4c, 0x12, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x61, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x9e, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x55, 0x58, 0xaa,
	0x02, 0x0b, 0x53, 0x76, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b,
	0x53, 0x76, 0x63, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x53, 0x76,
	0x63, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x76, 0x63, 0x3a, 0x3a, 0x55, 0x73, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_user_v1_service_proto_rawDescOnce sync.Once
	file_svc_user_v1_service_proto_rawDescData = file_svc_user_v1_service_proto_rawDesc
)

func file_svc_user_v1_service_proto_rawDescGZIP() []byte {
	file_svc_user_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_user_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_user_v1_service_proto_rawDescData)
	})
	return file_svc_user_v1_service_proto_rawDescData
}

var file_svc_user_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_svc_user_v1_service_proto_goTypes = []any{
	(*WhoamiRequest)(nil),                     // 0: svc.user.v1.WhoamiRequest
	(*WhoamiResponse)(nil),                    // 1: svc.user.v1.WhoamiResponse
	(*GetLogoutURLRequest)(nil),               // 2: svc.user.v1.GetLogoutURLRequest
	(*GetLogoutURLResponse)(nil),              // 3: svc.user.v1.GetLogoutURLResponse
	(*CreateOrganizationRequest)(nil),         // 4: svc.user.v1.CreateOrganizationRequest
	(*CreateOrganizationResponse)(nil),        // 5: svc.user.v1.CreateOrganizationResponse
	(*ListOrganizationRequest)(nil),           // 6: svc.user.v1.ListOrganizationRequest
	(*ListOrganizationResponse)(nil),          // 7: svc.user.v1.ListOrganizationResponse
	(*ListOrganizationResponse_ListItem)(nil), // 8: svc.user.v1.ListOrganizationResponse.ListItem
	(*v1.User)(nil),                           // 9: types.v1.User
	(*v1.Organization)(nil),                   // 10: types.v1.Organization
	(*v1.Cursor)(nil),                         // 11: types.v1.Cursor
}
var file_svc_user_v1_service_proto_depIdxs = []int32{
	9,  // 0: svc.user.v1.WhoamiResponse.user:type_name -> types.v1.User
	10, // 1: svc.user.v1.WhoamiResponse.current_organization:type_name -> types.v1.Organization
	10, // 2: svc.user.v1.WhoamiResponse.default_organization:type_name -> types.v1.Organization
	10, // 3: svc.user.v1.CreateOrganizationResponse.organization:type_name -> types.v1.Organization
	11, // 4: svc.user.v1.ListOrganizationRequest.cursor:type_name -> types.v1.Cursor
	11, // 5: svc.user.v1.ListOrganizationResponse.next:type_name -> types.v1.Cursor
	8,  // 6: svc.user.v1.ListOrganizationResponse.items:type_name -> svc.user.v1.ListOrganizationResponse.ListItem
	10, // 7: svc.user.v1.ListOrganizationResponse.default_organization:type_name -> types.v1.Organization
	10, // 8: svc.user.v1.ListOrganizationResponse.ListItem.organization:type_name -> types.v1.Organization
	0,  // 9: svc.user.v1.UserService.Whoami:input_type -> svc.user.v1.WhoamiRequest
	2,  // 10: svc.user.v1.UserService.GetLogoutURL:input_type -> svc.user.v1.GetLogoutURLRequest
	4,  // 11: svc.user.v1.UserService.CreateOrganization:input_type -> svc.user.v1.CreateOrganizationRequest
	6,  // 12: svc.user.v1.UserService.ListOrganization:input_type -> svc.user.v1.ListOrganizationRequest
	1,  // 13: svc.user.v1.UserService.Whoami:output_type -> svc.user.v1.WhoamiResponse
	3,  // 14: svc.user.v1.UserService.GetLogoutURL:output_type -> svc.user.v1.GetLogoutURLResponse
	5,  // 15: svc.user.v1.UserService.CreateOrganization:output_type -> svc.user.v1.CreateOrganizationResponse
	7,  // 16: svc.user.v1.UserService.ListOrganization:output_type -> svc.user.v1.ListOrganizationResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_svc_user_v1_service_proto_init() }
func file_svc_user_v1_service_proto_init() {
	if File_svc_user_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_user_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_user_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_user_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_user_v1_service_proto_msgTypes,
	}.Build()
	File_svc_user_v1_service_proto = out.File
	file_svc_user_v1_service_proto_rawDesc = nil
	file_svc_user_v1_service_proto_goTypes = nil
	file_svc_user_v1_service_proto_depIdxs = nil
}
