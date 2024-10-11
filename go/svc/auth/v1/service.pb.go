// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: svc/auth/v1/service.proto

package authv1

import (
	v1 "github.com/humanlogio/api/go/types/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAuthURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional: if an org is specified
	OrganizationId string               `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	ReturnToUrl    string               `protobuf:"bytes,2,opt,name=return_to_url,json=returnToUrl,proto3" json:"return_to_url,omitempty"`
	Localhost      *LocalhostViaBrowser `protobuf:"bytes,3,opt,name=localhost,proto3" json:"localhost,omitempty"`
}

func (x *GetAuthURLRequest) Reset() {
	*x = GetAuthURLRequest{}
	mi := &file_svc_auth_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthURLRequest) ProtoMessage() {}

func (x *GetAuthURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_auth_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthURLRequest.ProtoReflect.Descriptor instead.
func (*GetAuthURLRequest) Descriptor() ([]byte, []int) {
	return file_svc_auth_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthURLRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *GetAuthURLRequest) GetReturnToUrl() string {
	if x != nil {
		return x.ReturnToUrl
	}
	return ""
}

func (x *GetAuthURLRequest) GetLocalhost() *LocalhostViaBrowser {
	if x != nil {
		return x.Localhost
	}
	return nil
}

type LocalhostViaBrowser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClaimAccountId  *int64      `protobuf:"varint,1,opt,name=claim_account_id,json=claimAccountId,proto3,oneof" json:"claim_account_id,omitempty"`
	ClaimMachineId  *int64      `protobuf:"varint,2,opt,name=claim_machine_id,json=claimMachineId,proto3,oneof" json:"claim_machine_id,omitempty"`
	Architecture    string      `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	OperatingSystem string      `protobuf:"bytes,4,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
	UsingVersion    *v1.Version `protobuf:"bytes,5,opt,name=using_version,json=usingVersion,proto3" json:"using_version,omitempty"`
}

func (x *LocalhostViaBrowser) Reset() {
	*x = LocalhostViaBrowser{}
	mi := &file_svc_auth_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalhostViaBrowser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalhostViaBrowser) ProtoMessage() {}

func (x *LocalhostViaBrowser) ProtoReflect() protoreflect.Message {
	mi := &file_svc_auth_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalhostViaBrowser.ProtoReflect.Descriptor instead.
func (*LocalhostViaBrowser) Descriptor() ([]byte, []int) {
	return file_svc_auth_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *LocalhostViaBrowser) GetClaimAccountId() int64 {
	if x != nil && x.ClaimAccountId != nil {
		return *x.ClaimAccountId
	}
	return 0
}

func (x *LocalhostViaBrowser) GetClaimMachineId() int64 {
	if x != nil && x.ClaimMachineId != nil {
		return *x.ClaimMachineId
	}
	return 0
}

func (x *LocalhostViaBrowser) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *LocalhostViaBrowser) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *LocalhostViaBrowser) GetUsingVersion() *v1.Version {
	if x != nil {
		return x.UsingVersion
	}
	return nil
}

type GetAuthURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthUrl string `protobuf:"bytes,1,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
}

func (x *GetAuthURLResponse) Reset() {
	*x = GetAuthURLResponse{}
	mi := &file_svc_auth_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthURLResponse) ProtoMessage() {}

func (x *GetAuthURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_auth_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthURLResponse.ProtoReflect.Descriptor instead.
func (*GetAuthURLResponse) Descriptor() ([]byte, []int) {
	return file_svc_auth_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAuthURLResponse) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

type BeginDeviceAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *BeginDeviceAuthRequest) Reset() {
	*x = BeginDeviceAuthRequest{}
	mi := &file_svc_auth_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeginDeviceAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginDeviceAuthRequest) ProtoMessage() {}

func (x *BeginDeviceAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_auth_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginDeviceAuthRequest.ProtoReflect.Descriptor instead.
func (*BeginDeviceAuthRequest) Descriptor() ([]byte, []int) {
	return file_svc_auth_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *BeginDeviceAuthRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

type BeginDeviceAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// device_code must be used to retrieve an authenticated user token with `CompleteDeviceAuth`
	DeviceCode   string                 `protobuf:"bytes,2,opt,name=device_code,json=deviceCode,proto3" json:"device_code,omitempty"`
	UserCode     string                 `protobuf:"bytes,3,opt,name=user_code,json=userCode,proto3" json:"user_code,omitempty"`
	ExpiresAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	PollInterval *durationpb.Duration   `protobuf:"bytes,5,opt,name=poll_interval,json=pollInterval,proto3" json:"poll_interval,omitempty"`
}

func (x *BeginDeviceAuthResponse) Reset() {
	*x = BeginDeviceAuthResponse{}
	mi := &file_svc_auth_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeginDeviceAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginDeviceAuthResponse) ProtoMessage() {}

func (x *BeginDeviceAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_auth_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginDeviceAuthResponse.ProtoReflect.Descriptor instead.
func (*BeginDeviceAuthResponse) Descriptor() ([]byte, []int) {
	return file_svc_auth_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *BeginDeviceAuthResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BeginDeviceAuthResponse) GetDeviceCode() string {
	if x != nil {
		return x.DeviceCode
	}
	return ""
}

func (x *BeginDeviceAuthResponse) GetUserCode() string {
	if x != nil {
		return x.UserCode
	}
	return ""
}

func (x *BeginDeviceAuthResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *BeginDeviceAuthResponse) GetPollInterval() *durationpb.Duration {
	if x != nil {
		return x.PollInterval
	}
	return nil
}

type CompleteDeviceAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// device_code is returned when calling `BeginDeviceAuth` and allows an auth flow
	// to retrieve the user token that is generated when a user is authenticated
	DeviceCode      string `protobuf:"bytes,1,opt,name=device_code,json=deviceCode,proto3" json:"device_code,omitempty"`
	UserCode        string `protobuf:"bytes,2,opt,name=user_code,json=userCode,proto3" json:"user_code,omitempty"`
	ClaimAccountId  *int64 `protobuf:"varint,3,opt,name=claim_account_id,json=claimAccountId,proto3,oneof" json:"claim_account_id,omitempty"`
	ClaimMachineId  *int64 `protobuf:"varint,4,opt,name=claim_machine_id,json=claimMachineId,proto3,oneof" json:"claim_machine_id,omitempty"`
	Architecture    string `protobuf:"bytes,5,opt,name=architecture,proto3" json:"architecture,omitempty"`
	OperatingSystem string `protobuf:"bytes,6,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
}

func (x *CompleteDeviceAuthRequest) Reset() {
	*x = CompleteDeviceAuthRequest{}
	mi := &file_svc_auth_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteDeviceAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteDeviceAuthRequest) ProtoMessage() {}

func (x *CompleteDeviceAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_auth_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteDeviceAuthRequest.ProtoReflect.Descriptor instead.
func (*CompleteDeviceAuthRequest) Descriptor() ([]byte, []int) {
	return file_svc_auth_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *CompleteDeviceAuthRequest) GetDeviceCode() string {
	if x != nil {
		return x.DeviceCode
	}
	return ""
}

func (x *CompleteDeviceAuthRequest) GetUserCode() string {
	if x != nil {
		return x.UserCode
	}
	return ""
}

func (x *CompleteDeviceAuthRequest) GetClaimAccountId() int64 {
	if x != nil && x.ClaimAccountId != nil {
		return *x.ClaimAccountId
	}
	return 0
}

func (x *CompleteDeviceAuthRequest) GetClaimMachineId() int64 {
	if x != nil && x.ClaimMachineId != nil {
		return *x.ClaimMachineId
	}
	return 0
}

func (x *CompleteDeviceAuthRequest) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *CompleteDeviceAuthRequest) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

type CompleteDeviceAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     *v1.UserToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AccountId int64         `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	MachineId int64         `protobuf:"varint,3,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
}

func (x *CompleteDeviceAuthResponse) Reset() {
	*x = CompleteDeviceAuthResponse{}
	mi := &file_svc_auth_v1_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteDeviceAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteDeviceAuthResponse) ProtoMessage() {}

func (x *CompleteDeviceAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_auth_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteDeviceAuthResponse.ProtoReflect.Descriptor instead.
func (*CompleteDeviceAuthResponse) Descriptor() ([]byte, []int) {
	return file_svc_auth_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *CompleteDeviceAuthResponse) GetToken() *v1.UserToken {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *CompleteDeviceAuthResponse) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CompleteDeviceAuthResponse) GetMachineId() int64 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

var File_svc_auth_v1_service_proto protoreflect.FileDescriptor

var file_svc_auth_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x3e, 0x0a, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x61, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x22, 0xa4, 0x02, 0x0a,
	0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x61, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x0e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52,
	0x0e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x36, 0x0a, 0x0d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x75, 0x73, 0x69,
	0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x13,
	0x0a, 0x11, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x55, 0x72, 0x6c, 0x22, 0x3c, 0x0a, 0x16, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xe4, 0x01, 0x0a, 0x17, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x70, 0x6f, 0x6c,
	0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x6f, 0x6c,
	0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0xb0, 0x02, 0x0a, 0x19, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x0e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x85, 0x01, 0x0a,
	0x1a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x49, 0x64, 0x32, 0xa7, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55,
	0x52, 0x4c, 0x12, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0f, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x23, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x26, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9e,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x76, 0x63, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x53, 0x76, 0x63, 0x5c, 0x41, 0x75,
	0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x53, 0x76, 0x63, 0x5c, 0x41, 0x75, 0x74, 0x68,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0d, 0x53, 0x76, 0x63, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_auth_v1_service_proto_rawDescOnce sync.Once
	file_svc_auth_v1_service_proto_rawDescData = file_svc_auth_v1_service_proto_rawDesc
)

func file_svc_auth_v1_service_proto_rawDescGZIP() []byte {
	file_svc_auth_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_auth_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_auth_v1_service_proto_rawDescData)
	})
	return file_svc_auth_v1_service_proto_rawDescData
}

var file_svc_auth_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_svc_auth_v1_service_proto_goTypes = []any{
	(*GetAuthURLRequest)(nil),          // 0: svc.auth.v1.GetAuthURLRequest
	(*LocalhostViaBrowser)(nil),        // 1: svc.auth.v1.LocalhostViaBrowser
	(*GetAuthURLResponse)(nil),         // 2: svc.auth.v1.GetAuthURLResponse
	(*BeginDeviceAuthRequest)(nil),     // 3: svc.auth.v1.BeginDeviceAuthRequest
	(*BeginDeviceAuthResponse)(nil),    // 4: svc.auth.v1.BeginDeviceAuthResponse
	(*CompleteDeviceAuthRequest)(nil),  // 5: svc.auth.v1.CompleteDeviceAuthRequest
	(*CompleteDeviceAuthResponse)(nil), // 6: svc.auth.v1.CompleteDeviceAuthResponse
	(*v1.Version)(nil),                 // 7: types.v1.Version
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),        // 9: google.protobuf.Duration
	(*v1.UserToken)(nil),               // 10: types.v1.UserToken
}
var file_svc_auth_v1_service_proto_depIdxs = []int32{
	1,  // 0: svc.auth.v1.GetAuthURLRequest.localhost:type_name -> svc.auth.v1.LocalhostViaBrowser
	7,  // 1: svc.auth.v1.LocalhostViaBrowser.using_version:type_name -> types.v1.Version
	8,  // 2: svc.auth.v1.BeginDeviceAuthResponse.expires_at:type_name -> google.protobuf.Timestamp
	9,  // 3: svc.auth.v1.BeginDeviceAuthResponse.poll_interval:type_name -> google.protobuf.Duration
	10, // 4: svc.auth.v1.CompleteDeviceAuthResponse.token:type_name -> types.v1.UserToken
	0,  // 5: svc.auth.v1.AuthService.GetAuthURL:input_type -> svc.auth.v1.GetAuthURLRequest
	3,  // 6: svc.auth.v1.AuthService.BeginDeviceAuth:input_type -> svc.auth.v1.BeginDeviceAuthRequest
	5,  // 7: svc.auth.v1.AuthService.CompleteDeviceAuth:input_type -> svc.auth.v1.CompleteDeviceAuthRequest
	2,  // 8: svc.auth.v1.AuthService.GetAuthURL:output_type -> svc.auth.v1.GetAuthURLResponse
	4,  // 9: svc.auth.v1.AuthService.BeginDeviceAuth:output_type -> svc.auth.v1.BeginDeviceAuthResponse
	6,  // 10: svc.auth.v1.AuthService.CompleteDeviceAuth:output_type -> svc.auth.v1.CompleteDeviceAuthResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_svc_auth_v1_service_proto_init() }
func file_svc_auth_v1_service_proto_init() {
	if File_svc_auth_v1_service_proto != nil {
		return
	}
	file_svc_auth_v1_service_proto_msgTypes[1].OneofWrappers = []any{}
	file_svc_auth_v1_service_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_auth_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_auth_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_auth_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_auth_v1_service_proto_msgTypes,
	}.Build()
	File_svc_auth_v1_service_proto = out.File
	file_svc_auth_v1_service_proto_rawDesc = nil
	file_svc_auth_v1_service_proto_goTypes = nil
	file_svc_auth_v1_service_proto_depIdxs = nil
}
