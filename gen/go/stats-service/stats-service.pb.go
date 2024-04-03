// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: stats-service/stats-service.proto

package statsv1

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

type OrderEnum int32

const (
	OrderEnum_ASC  OrderEnum = 0
	OrderEnum_DESC OrderEnum = 1
)

// Enum value maps for OrderEnum.
var (
	OrderEnum_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	OrderEnum_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x OrderEnum) Enum() *OrderEnum {
	p := new(OrderEnum)
	*p = x
	return p
}

func (x OrderEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_stats_service_stats_service_proto_enumTypes[0].Descriptor()
}

func (OrderEnum) Type() protoreflect.EnumType {
	return &file_stats_service_stats_service_proto_enumTypes[0]
}

func (x OrderEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderEnum.Descriptor instead.
func (OrderEnum) EnumDescriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{0}
}

type CreateAIServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       uint32 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateAIServiceRequest) Reset() {
	*x = CreateAIServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_service_stats_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAIServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAIServiceRequest) ProtoMessage() {}

func (x *CreateAIServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stats_service_stats_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAIServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateAIServiceRequest) Descriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAIServiceRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateAIServiceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateAIServiceRequest) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateAIServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceUid string `protobuf:"bytes,1,opt,name=service_uid,json=serviceUid,proto3" json:"service_uid,omitempty"`
}

func (x *CreateAIServiceResponse) Reset() {
	*x = CreateAIServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_service_stats_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAIServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAIServiceResponse) ProtoMessage() {}

func (x *CreateAIServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_service_stats_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAIServiceResponse.ProtoReflect.Descriptor instead.
func (*CreateAIServiceResponse) Descriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAIServiceResponse) GetServiceUid() string {
	if x != nil {
		return x.ServiceUid
	}
	return ""
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUid    string `protobuf:"bytes,1,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
	ServiceUid string `protobuf:"bytes,2,opt,name=service_uid,json=serviceUid,proto3" json:"service_uid,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_service_stats_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stats_service_stats_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{2}
}

func (x *CallRequest) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

func (x *CallRequest) GetServiceUid() string {
	if x != nil {
		return x.ServiceUid
	}
	return ""
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_service_stats_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_service_stats_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{3}
}

func (x *CallResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCallsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUid    string    `protobuf:"bytes,1,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
	ServiceUid string    `protobuf:"bytes,2,opt,name=service_uid,json=serviceUid,proto3" json:"service_uid,omitempty"`
	PageNumber uint32    `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   uint32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Order      OrderEnum `protobuf:"varint,5,opt,name=order,proto3,enum=stats.OrderEnum" json:"order,omitempty"`
}

func (x *GetCallsRequest) Reset() {
	*x = GetCallsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_service_stats_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCallsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCallsRequest) ProtoMessage() {}

func (x *GetCallsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stats_service_stats_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCallsRequest.ProtoReflect.Descriptor instead.
func (*GetCallsRequest) Descriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCallsRequest) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

func (x *GetCallsRequest) GetServiceUid() string {
	if x != nil {
		return x.ServiceUid
	}
	return ""
}

func (x *GetCallsRequest) GetPageNumber() uint32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetCallsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetCallsRequest) GetOrder() OrderEnum {
	if x != nil {
		return x.Order
	}
	return OrderEnum_ASC
}

type GetCallsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*Calls `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *GetCallsResponse) Reset() {
	*x = GetCallsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_service_stats_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCallsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCallsResponse) ProtoMessage() {}

func (x *GetCallsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stats_service_stats_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCallsResponse.ProtoReflect.Descriptor instead.
func (*GetCallsResponse) Descriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetCallsResponse) GetCalls() []*Calls {
	if x != nil {
		return x.Calls
	}
	return nil
}

type Calls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUid    string `protobuf:"bytes,1,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
	ServiceUid string `protobuf:"bytes,2,opt,name=service_uid,json=serviceUid,proto3" json:"service_uid,omitempty"`
	Count      uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Amount     uint32 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Calls) Reset() {
	*x = Calls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_service_stats_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calls) ProtoMessage() {}

func (x *Calls) ProtoReflect() protoreflect.Message {
	mi := &file_stats_service_stats_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calls.ProtoReflect.Descriptor instead.
func (*Calls) Descriptor() ([]byte, []int) {
	return file_stats_service_stats_service_proto_rawDescGZIP(), []int{6}
}

func (x *Calls) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

func (x *Calls) GetServiceUid() string {
	if x != nil {
		return x.ServiceUid
	}
	return ""
}

func (x *Calls) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Calls) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_stats_service_stats_service_proto protoreflect.FileDescriptor

var file_stats_service_stats_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x66, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x3a, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x49, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x69, 0x64, 0x22, 0x49,
	0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c,
	0x73, 0x22, 0x71, 0x0a, 0x05, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x55, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x1e, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x01, 0x32, 0xc9, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12,
	0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1a, 0x5a, 0x18, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stats_service_stats_service_proto_rawDescOnce sync.Once
	file_stats_service_stats_service_proto_rawDescData = file_stats_service_stats_service_proto_rawDesc
)

func file_stats_service_stats_service_proto_rawDescGZIP() []byte {
	file_stats_service_stats_service_proto_rawDescOnce.Do(func() {
		file_stats_service_stats_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stats_service_stats_service_proto_rawDescData)
	})
	return file_stats_service_stats_service_proto_rawDescData
}

var file_stats_service_stats_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stats_service_stats_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_stats_service_stats_service_proto_goTypes = []interface{}{
	(OrderEnum)(0),                  // 0: stats.OrderEnum
	(*CreateAIServiceRequest)(nil),  // 1: stats.CreateAIServiceRequest
	(*CreateAIServiceResponse)(nil), // 2: stats.CreateAIServiceResponse
	(*CallRequest)(nil),             // 3: stats.CallRequest
	(*CallResponse)(nil),            // 4: stats.CallResponse
	(*GetCallsRequest)(nil),         // 5: stats.GetCallsRequest
	(*GetCallsResponse)(nil),        // 6: stats.GetCallsResponse
	(*Calls)(nil),                   // 7: stats.Calls
}
var file_stats_service_stats_service_proto_depIdxs = []int32{
	0, // 0: stats.GetCallsRequest.order:type_name -> stats.OrderEnum
	7, // 1: stats.GetCallsResponse.calls:type_name -> stats.Calls
	1, // 2: stats.StatisticService.Create:input_type -> stats.CreateAIServiceRequest
	3, // 3: stats.StatisticService.Call:input_type -> stats.CallRequest
	5, // 4: stats.StatisticService.GetCalls:input_type -> stats.GetCallsRequest
	2, // 5: stats.StatisticService.Create:output_type -> stats.CreateAIServiceResponse
	4, // 6: stats.StatisticService.Call:output_type -> stats.CallResponse
	6, // 7: stats.StatisticService.GetCalls:output_type -> stats.GetCallsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stats_service_stats_service_proto_init() }
func file_stats_service_stats_service_proto_init() {
	if File_stats_service_stats_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stats_service_stats_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAIServiceRequest); i {
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
		file_stats_service_stats_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAIServiceResponse); i {
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
		file_stats_service_stats_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_stats_service_stats_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_stats_service_stats_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCallsRequest); i {
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
		file_stats_service_stats_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCallsResponse); i {
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
		file_stats_service_stats_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calls); i {
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
			RawDescriptor: file_stats_service_stats_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stats_service_stats_service_proto_goTypes,
		DependencyIndexes: file_stats_service_stats_service_proto_depIdxs,
		EnumInfos:         file_stats_service_stats_service_proto_enumTypes,
		MessageInfos:      file_stats_service_stats_service_proto_msgTypes,
	}.Build()
	File_stats_service_stats_service_proto = out.File
	file_stats_service_stats_service_proto_rawDesc = nil
	file_stats_service_stats_service_proto_goTypes = nil
	file_stats_service_stats_service_proto_depIdxs = nil
}
