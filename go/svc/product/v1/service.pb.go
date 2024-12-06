// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: svc/product/v1/service.proto

package productv1

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

type ListProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor   *v1.Cursor `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit    int32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Category string     `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// filters
	Scope v1.Product_Scope `protobuf:"varint,400,opt,name=scope,proto3,enum=types.v1.Product_Scope" json:"scope,omitempty"`
}

func (x *ListProductRequest) Reset() {
	*x = ListProductRequest{}
	mi := &file_svc_product_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductRequest) ProtoMessage() {}

func (x *ListProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_product_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductRequest.ProtoReflect.Descriptor instead.
func (*ListProductRequest) Descriptor() ([]byte, []int) {
	return file_svc_product_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListProductRequest) GetCursor() *v1.Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ListProductRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ListProductRequest) GetScope() v1.Product_Scope {
	if x != nil {
		return x.Scope
	}
	return v1.Product_Scope(0)
}

type ListProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next           *v1.Cursor                      `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items          []*ListProductResponse_ListItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	DefaultProduct *v1.Product                     `protobuf:"bytes,3,opt,name=default_product,json=defaultProduct,proto3" json:"default_product,omitempty"`
}

func (x *ListProductResponse) Reset() {
	*x = ListProductResponse{}
	mi := &file_svc_product_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductResponse) ProtoMessage() {}

func (x *ListProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_product_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductResponse.ProtoReflect.Descriptor instead.
func (*ListProductResponse) Descriptor() ([]byte, []int) {
	return file_svc_product_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListProductResponse) GetNext() *v1.Cursor {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *ListProductResponse) GetItems() []*ListProductResponse_ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListProductResponse) GetDefaultProduct() *v1.Product {
	if x != nil {
		return x.DefaultProduct
	}
	return nil
}

type ListProductResponse_ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *v1.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Prices  []*v1.Price `protobuf:"bytes,2,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *ListProductResponse_ListItem) Reset() {
	*x = ListProductResponse_ListItem{}
	mi := &file_svc_product_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductResponse_ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductResponse_ListItem) ProtoMessage() {}

func (x *ListProductResponse_ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_svc_product_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductResponse_ListItem.ProtoReflect.Descriptor instead.
func (*ListProductResponse_ListItem) Descriptor() ([]byte, []int) {
	return file_svc_product_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListProductResponse_ListItem) GetProduct() *v1.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ListProductResponse_ListItem) GetPrices() []*v1.Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

var File_svc_product_v1_service_proto protoreflect.FileDescriptor

var file_svc_product_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x15,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x90, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3a, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x1a, 0x60, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x2b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x27, 0x0a,
	0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x32, 0x68, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xb3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x50, 0x58, 0xaa, 0x02, 0x0e, 0x53, 0x76, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x53, 0x76, 0x63, 0x5c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x53, 0x76, 0x63, 0x5c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x76, 0x63, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_product_v1_service_proto_rawDescOnce sync.Once
	file_svc_product_v1_service_proto_rawDescData = file_svc_product_v1_service_proto_rawDesc
)

func file_svc_product_v1_service_proto_rawDescGZIP() []byte {
	file_svc_product_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_product_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_product_v1_service_proto_rawDescData)
	})
	return file_svc_product_v1_service_proto_rawDescData
}

var file_svc_product_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_svc_product_v1_service_proto_goTypes = []any{
	(*ListProductRequest)(nil),           // 0: svc.product.v1.ListProductRequest
	(*ListProductResponse)(nil),          // 1: svc.product.v1.ListProductResponse
	(*ListProductResponse_ListItem)(nil), // 2: svc.product.v1.ListProductResponse.ListItem
	(*v1.Cursor)(nil),                    // 3: types.v1.Cursor
	(v1.Product_Scope)(0),                // 4: types.v1.Product.Scope
	(*v1.Product)(nil),                   // 5: types.v1.Product
	(*v1.Price)(nil),                     // 6: types.v1.Price
}
var file_svc_product_v1_service_proto_depIdxs = []int32{
	3, // 0: svc.product.v1.ListProductRequest.cursor:type_name -> types.v1.Cursor
	4, // 1: svc.product.v1.ListProductRequest.scope:type_name -> types.v1.Product.Scope
	3, // 2: svc.product.v1.ListProductResponse.next:type_name -> types.v1.Cursor
	2, // 3: svc.product.v1.ListProductResponse.items:type_name -> svc.product.v1.ListProductResponse.ListItem
	5, // 4: svc.product.v1.ListProductResponse.default_product:type_name -> types.v1.Product
	5, // 5: svc.product.v1.ListProductResponse.ListItem.product:type_name -> types.v1.Product
	6, // 6: svc.product.v1.ListProductResponse.ListItem.prices:type_name -> types.v1.Price
	0, // 7: svc.product.v1.ProductService.ListProduct:input_type -> svc.product.v1.ListProductRequest
	1, // 8: svc.product.v1.ProductService.ListProduct:output_type -> svc.product.v1.ListProductResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_svc_product_v1_service_proto_init() }
func file_svc_product_v1_service_proto_init() {
	if File_svc_product_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_product_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_product_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_product_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_product_v1_service_proto_msgTypes,
	}.Build()
	File_svc_product_v1_service_proto = out.File
	file_svc_product_v1_service_proto_rawDesc = nil
	file_svc_product_v1_service_proto_goTypes = nil
	file_svc_product_v1_service_proto_depIdxs = nil
}
