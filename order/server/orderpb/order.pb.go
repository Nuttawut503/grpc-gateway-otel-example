// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: protos/order.proto

package orderpb

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

type CreateCustomerOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string  `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Price      float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateCustomerOrderRequest) Reset() {
	*x = CreateCustomerOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerOrderRequest) ProtoMessage() {}

func (x *CreateCustomerOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerOrderRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCustomerOrderRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateCustomerOrderRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateCustomerOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *CreateCustomerOrderResponse) Reset() {
	*x = CreateCustomerOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerOrderResponse) ProtoMessage() {}

func (x *CreateCustomerOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerOrderResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCustomerOrderResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

var File_protos_order_proto protoreflect.FileDescriptor

var file_protos_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x18, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x32, 0x78, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x6f, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22,
	0x06, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x10, 0x5a, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_order_proto_rawDescOnce sync.Once
	file_protos_order_proto_rawDescData = file_protos_order_proto_rawDesc
)

func file_protos_order_proto_rawDescGZIP() []byte {
	file_protos_order_proto_rawDescOnce.Do(func() {
		file_protos_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_order_proto_rawDescData)
	})
	return file_protos_order_proto_rawDescData
}

var file_protos_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_order_proto_goTypes = []interface{}{
	(*CreateCustomerOrderRequest)(nil),  // 0: order.CreateCustomerOrderRequest
	(*CreateCustomerOrderResponse)(nil), // 1: order.CreateCustomerOrderResponse
}
var file_protos_order_proto_depIdxs = []int32{
	0, // 0: order.Order.CreateCustomerOrder:input_type -> order.CreateCustomerOrderRequest
	1, // 1: order.Order.CreateCustomerOrder:output_type -> order.CreateCustomerOrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_order_proto_init() }
func file_protos_order_proto_init() {
	if File_protos_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerOrderRequest); i {
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
		file_protos_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerOrderResponse); i {
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
			RawDescriptor: file_protos_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_order_proto_goTypes,
		DependencyIndexes: file_protos_order_proto_depIdxs,
		MessageInfos:      file_protos_order_proto_msgTypes,
	}.Build()
	File_protos_order_proto = out.File
	file_protos_order_proto_rawDesc = nil
	file_protos_order_proto_goTypes = nil
	file_protos_order_proto_depIdxs = nil
}