// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: zzzz-customer-trade.proto

package customer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "github.com/noncepad/client/proto/base"
	order "github.com/noncepad/client/proto/order"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderStatus_Status int32

const (
	OrderStatus_NEW              OrderStatus_Status = 0
	OrderStatus_PARTIALLY_FILLED OrderStatus_Status = 1
	OrderStatus_FILLED           OrderStatus_Status = 2
	OrderStatus_REJECTED         OrderStatus_Status = 3
	OrderStatus_CANCELED         OrderStatus_Status = 4
)

// Enum value maps for OrderStatus_Status.
var (
	OrderStatus_Status_name = map[int32]string{
		0: "NEW",
		1: "PARTIALLY_FILLED",
		2: "FILLED",
		3: "REJECTED",
		4: "CANCELED",
	}
	OrderStatus_Status_value = map[string]int32{
		"NEW":              0,
		"PARTIALLY_FILLED": 1,
		"FILLED":           2,
		"REJECTED":         3,
		"CANCELED":         4,
	}
)

func (x OrderStatus_Status) Enum() *OrderStatus_Status {
	p := new(OrderStatus_Status)
	*p = x
	return p
}

func (x OrderStatus_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_zzzz_customer_trade_proto_enumTypes[0].Descriptor()
}

func (OrderStatus_Status) Type() protoreflect.EnumType {
	return &file_zzzz_customer_trade_proto_enumTypes[0]
}

func (x OrderStatus_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus_Status.Descriptor instead.
func (OrderStatus_Status) EnumDescriptor() ([]byte, []int) {
	return file_zzzz_customer_trade_proto_rawDescGZIP(), []int{3, 0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityId *base.SecurityID `protobuf:"bytes,1,opt,name=security_id,json=securityId,proto3" json:"security_id,omitempty"`
	BuyOrSell  bool             `protobuf:"varint,2,opt,name=buy_or_sell,json=buyOrSell,proto3" json:"buy_or_sell,omitempty"`
	// Types that are assignable to OrderType:
	//	*Order_Market
	//	*Order_Limit
	//	*Order_Stop
	//	*Order_StopLimit
	OrderType isOrder_OrderType `protobuf_oneof:"order_type"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zzzz_customer_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_zzzz_customer_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_zzzz_customer_trade_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetSecurityId() *base.SecurityID {
	if x != nil {
		return x.SecurityId
	}
	return nil
}

func (x *Order) GetBuyOrSell() bool {
	if x != nil {
		return x.BuyOrSell
	}
	return false
}

func (m *Order) GetOrderType() isOrder_OrderType {
	if m != nil {
		return m.OrderType
	}
	return nil
}

func (x *Order) GetMarket() *order.Market {
	if x, ok := x.GetOrderType().(*Order_Market); ok {
		return x.Market
	}
	return nil
}

func (x *Order) GetLimit() *order.Limit {
	if x, ok := x.GetOrderType().(*Order_Limit); ok {
		return x.Limit
	}
	return nil
}

func (x *Order) GetStop() *order.Stop {
	if x, ok := x.GetOrderType().(*Order_Stop); ok {
		return x.Stop
	}
	return nil
}

func (x *Order) GetStopLimit() *order.StopLimit {
	if x, ok := x.GetOrderType().(*Order_StopLimit); ok {
		return x.StopLimit
	}
	return nil
}

type isOrder_OrderType interface {
	isOrder_OrderType()
}

type Order_Market struct {
	Market *order.Market `protobuf:"bytes,3,opt,name=market,proto3,oneof"`
}

type Order_Limit struct {
	Limit *order.Limit `protobuf:"bytes,4,opt,name=limit,proto3,oneof"`
}

type Order_Stop struct {
	Stop *order.Stop `protobuf:"bytes,5,opt,name=stop,proto3,oneof"`
}

type Order_StopLimit struct {
	StopLimit *order.StopLimit `protobuf:"bytes,6,opt,name=stop_limit,json=stopLimit,proto3,oneof"`
}

func (*Order_Market) isOrder_OrderType() {}

func (*Order_Limit) isOrder_OrderType() {}

func (*Order_Stop) isOrder_OrderType() {}

func (*Order_StopLimit) isOrder_OrderType() {}

type CancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CancelRequest) Reset() {
	*x = CancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zzzz_customer_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRequest) ProtoMessage() {}

func (x *CancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zzzz_customer_trade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRequest.ProtoReflect.Descriptor instead.
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return file_zzzz_customer_trade_proto_rawDescGZIP(), []int{1}
}

func (x *CancelRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Stream  bool   `protobuf:"varint,2,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zzzz_customer_trade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zzzz_customer_trade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_zzzz_customer_trade_proto_rawDescGZIP(), []int{2}
}

func (x *StatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *StatusRequest) GetStream() bool {
	if x != nil {
		return x.Stream
	}
	return false
}

type OrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string             `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status  OrderStatus_Status `protobuf:"varint,2,opt,name=status,proto3,enum=customer.OrderStatus_Status" json:"status,omitempty"`
}

func (x *OrderStatus) Reset() {
	*x = OrderStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zzzz_customer_trade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatus) ProtoMessage() {}

func (x *OrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_zzzz_customer_trade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatus.ProtoReflect.Descriptor instead.
func (*OrderStatus) Descriptor() ([]byte, []int) {
	return file_zzzz_customer_trade_proto_rawDescGZIP(), []int{3}
}

func (x *OrderStatus) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderStatus) GetStatus() OrderStatus_Status {
	if x != nil {
		return x.Status
	}
	return OrderStatus_NEW
}

type CancelEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*CancelEvent_Report
	//	*CancelEvent_Reject
	Event isCancelEvent_Event `protobuf_oneof:"event"`
}

func (x *CancelEvent) Reset() {
	*x = CancelEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zzzz_customer_trade_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelEvent) ProtoMessage() {}

func (x *CancelEvent) ProtoReflect() protoreflect.Message {
	mi := &file_zzzz_customer_trade_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelEvent.ProtoReflect.Descriptor instead.
func (*CancelEvent) Descriptor() ([]byte, []int) {
	return file_zzzz_customer_trade_proto_rawDescGZIP(), []int{4}
}

func (m *CancelEvent) GetEvent() isCancelEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *CancelEvent) GetReport() *order.ExecutionReport {
	if x, ok := x.GetEvent().(*CancelEvent_Report); ok {
		return x.Report
	}
	return nil
}

func (x *CancelEvent) GetReject() *order.CancelReject {
	if x, ok := x.GetEvent().(*CancelEvent_Reject); ok {
		return x.Reject
	}
	return nil
}

type isCancelEvent_Event interface {
	isCancelEvent_Event()
}

type CancelEvent_Report struct {
	Report *order.ExecutionReport `protobuf:"bytes,1,opt,name=report,proto3,oneof"`
}

type CancelEvent_Reject struct {
	Reject *order.CancelReject `protobuf:"bytes,2,opt,name=reject,proto3,oneof"`
}

func (*CancelEvent_Report) isCancelEvent_Event() {}

func (*CancelEvent_Reject) isCancelEvent_Event() {}

var File_zzzz_customer_trade_proto protoreflect.FileDescriptor

var file_zzzz_customer_trade_proto_rawDesc = []byte{
	0x0a, 0x19, 0x7a, 0x7a, 0x7a, 0x7a, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d,
	0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52,
	0x0a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x62,
	0x75, 0x79, 0x5f, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x27, 0x0a, 0x06, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x74,
	0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x31, 0x0a,
	0x0a, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x42, 0x0c, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2a,
	0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0xaf,
	0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x4f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x4c, 0x59, 0x5f,
	0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x49, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x22, 0x77, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x30, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0xc3, 0x01, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x3f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x17, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x21, 0x5a, 0x1f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x61, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zzzz_customer_trade_proto_rawDescOnce sync.Once
	file_zzzz_customer_trade_proto_rawDescData = file_zzzz_customer_trade_proto_rawDesc
)

func file_zzzz_customer_trade_proto_rawDescGZIP() []byte {
	file_zzzz_customer_trade_proto_rawDescOnce.Do(func() {
		file_zzzz_customer_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_zzzz_customer_trade_proto_rawDescData)
	})
	return file_zzzz_customer_trade_proto_rawDescData
}

var file_zzzz_customer_trade_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zzzz_customer_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_zzzz_customer_trade_proto_goTypes = []interface{}{
	(OrderStatus_Status)(0),       // 0: customer.OrderStatus.Status
	(*Order)(nil),                 // 1: customer.Order
	(*CancelRequest)(nil),         // 2: customer.CancelRequest
	(*StatusRequest)(nil),         // 3: customer.StatusRequest
	(*OrderStatus)(nil),           // 4: customer.OrderStatus
	(*CancelEvent)(nil),           // 5: customer.CancelEvent
	(*base.SecurityID)(nil),       // 6: base.SecurityID
	(*order.Market)(nil),          // 7: order.Market
	(*order.Limit)(nil),           // 8: order.Limit
	(*order.Stop)(nil),            // 9: order.Stop
	(*order.StopLimit)(nil),       // 10: order.StopLimit
	(*order.ExecutionReport)(nil), // 11: order.ExecutionReport
	(*order.CancelReject)(nil),    // 12: order.CancelReject
}
var file_zzzz_customer_trade_proto_depIdxs = []int32{
	6,  // 0: customer.Order.security_id:type_name -> base.SecurityID
	7,  // 1: customer.Order.market:type_name -> order.Market
	8,  // 2: customer.Order.limit:type_name -> order.Limit
	9,  // 3: customer.Order.stop:type_name -> order.Stop
	10, // 4: customer.Order.stop_limit:type_name -> order.StopLimit
	0,  // 5: customer.OrderStatus.status:type_name -> customer.OrderStatus.Status
	11, // 6: customer.CancelEvent.report:type_name -> order.ExecutionReport
	12, // 7: customer.CancelEvent.reject:type_name -> order.CancelReject
	1,  // 8: customer.OrderManager.Submit:input_type -> customer.Order
	3,  // 9: customer.OrderManager.GetStatus:input_type -> customer.StatusRequest
	2,  // 10: customer.OrderManager.Cancel:input_type -> customer.CancelRequest
	4,  // 11: customer.OrderManager.Submit:output_type -> customer.OrderStatus
	4,  // 12: customer.OrderManager.GetStatus:output_type -> customer.OrderStatus
	5,  // 13: customer.OrderManager.Cancel:output_type -> customer.CancelEvent
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_zzzz_customer_trade_proto_init() }
func file_zzzz_customer_trade_proto_init() {
	if File_zzzz_customer_trade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zzzz_customer_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_zzzz_customer_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRequest); i {
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
		file_zzzz_customer_trade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_zzzz_customer_trade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatus); i {
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
		file_zzzz_customer_trade_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelEvent); i {
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
	file_zzzz_customer_trade_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Order_Market)(nil),
		(*Order_Limit)(nil),
		(*Order_Stop)(nil),
		(*Order_StopLimit)(nil),
	}
	file_zzzz_customer_trade_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*CancelEvent_Report)(nil),
		(*CancelEvent_Reject)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zzzz_customer_trade_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zzzz_customer_trade_proto_goTypes,
		DependencyIndexes: file_zzzz_customer_trade_proto_depIdxs,
		EnumInfos:         file_zzzz_customer_trade_proto_enumTypes,
		MessageInfos:      file_zzzz_customer_trade_proto_msgTypes,
	}.Build()
	File_zzzz_customer_trade_proto = out.File
	file_zzzz_customer_trade_proto_rawDesc = nil
	file_zzzz_customer_trade_proto_goTypes = nil
	file_zzzz_customer_trade_proto_depIdxs = nil
}
