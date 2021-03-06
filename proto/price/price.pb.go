// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: price.proto

package price

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "github.com/noncepad/client/proto/base"
	exchange "github.com/noncepad/client/proto/exchange"
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

type Order_Status int32

const (
	Order_PENDING  Order_Status = 0
	Order_FILLING  Order_Status = 1
	Order_FILLED   Order_Status = 2
	Order_REJECTED Order_Status = 3
	Order_CANCELED Order_Status = 4
)

// Enum value maps for Order_Status.
var (
	Order_Status_name = map[int32]string{
		0: "PENDING",
		1: "FILLING",
		2: "FILLED",
		3: "REJECTED",
		4: "CANCELED",
	}
	Order_Status_value = map[string]int32{
		"PENDING":  0,
		"FILLING":  1,
		"FILLED":   2,
		"REJECTED": 3,
		"CANCELED": 4,
	}
)

func (x Order_Status) Enum() *Order_Status {
	p := new(Order_Status)
	*p = x
	return p
}

func (x Order_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_price_proto_enumTypes[0].Descriptor()
}

func (Order_Status) Type() protoreflect.EnumType {
	return &file_price_proto_enumTypes[0]
}

func (x Order_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_Status.Descriptor instead.
func (Order_Status) EnumDescriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{5, 0}
}

type MarketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *base.ShortCommodity `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Quote *base.ShortCommodity `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *MarketRequest) Reset() {
	*x = MarketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketRequest) ProtoMessage() {}

func (x *MarketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketRequest.ProtoReflect.Descriptor instead.
func (*MarketRequest) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{0}
}

func (x *MarketRequest) GetBase() *base.ShortCommodity {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *MarketRequest) GetQuote() *base.ShortCommodity {
	if x != nil {
		return x.Quote
	}
	return nil
}

type LastTrades struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trade []*Trade `protobuf:"bytes,1,rep,name=trade,proto3" json:"trade,omitempty"`
}

func (x *LastTrades) Reset() {
	*x = LastTrades{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastTrades) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastTrades) ProtoMessage() {}

func (x *LastTrades) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastTrades.ProtoReflect.Descriptor instead.
func (*LastTrades) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{1}
}

func (x *LastTrades) GetTrade() []*Trade {
	if x != nil {
		return x.Trade
	}
	return nil
}

type PriceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// Types that are assignable to Data:
	//	*PriceUpdate_RecentTrades
	//	*PriceUpdate_Book
	Data isPriceUpdate_Data `protobuf_oneof:"data"`
}

func (x *PriceUpdate) Reset() {
	*x = PriceUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceUpdate) ProtoMessage() {}

func (x *PriceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceUpdate.ProtoReflect.Descriptor instead.
func (*PriceUpdate) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{2}
}

func (x *PriceUpdate) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (m *PriceUpdate) GetData() isPriceUpdate_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *PriceUpdate) GetRecentTrades() *LastTrades {
	if x, ok := x.GetData().(*PriceUpdate_RecentTrades); ok {
		return x.RecentTrades
	}
	return nil
}

func (x *PriceUpdate) GetBook() *OrderBook {
	if x, ok := x.GetData().(*PriceUpdate_Book); ok {
		return x.Book
	}
	return nil
}

type isPriceUpdate_Data interface {
	isPriceUpdate_Data()
}

type PriceUpdate_RecentTrades struct {
	RecentTrades *LastTrades `protobuf:"bytes,2,opt,name=recent_trades,json=recentTrades,proto3,oneof"`
}

type PriceUpdate_Book struct {
	Book *OrderBook `protobuf:"bytes,3,opt,name=book,proto3,oneof"`
}

func (*PriceUpdate_RecentTrades) isPriceUpdate_Data() {}

func (*PriceUpdate_Book) isPriceUpdate_Data() {}

// ReadRequest - Order book read request.
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityId *base.SecurityID `protobuf:"bytes,1,opt,name=security_id,json=securityId,proto3" json:"security_id,omitempty"`
	// Exchange - Identifier of an exchange.
	Exchange exchange.Exchange `protobuf:"varint,2,opt,name=exchange,proto3,enum=exchange.Exchange" json:"exchange,omitempty"`
	// Start - Time of a first order to read.
	// If it is zero it will start reading from first record.
	Start int64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	// Finish - Time of a last order to read.
	// `Zero` will indicate we want to read until the end,
	// in case of historical data it will be last existing trade,
	// in case of real-time data it will stream updates forever.
	Finish int64 `protobuf:"varint,4,opt,name=finish,proto3" json:"finish,omitempty"`
	// MaxBatchSize - Maximum size of a single batch sent.
	MaxBatchSize int64 `protobuf:"varint,5,opt,name=max_batch_size,json=maxBatchSize,proto3" json:"max_batch_size,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetSecurityId() *base.SecurityID {
	if x != nil {
		return x.SecurityId
	}
	return nil
}

func (x *ReadRequest) GetExchange() exchange.Exchange {
	if x != nil {
		return x.Exchange
	}
	return exchange.Exchange(0)
}

func (x *ReadRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ReadRequest) GetFinish() int64 {
	if x != nil {
		return x.Finish
	}
	return 0
}

func (x *ReadRequest) GetMaxBatchSize() int64 {
	if x != nil {
		return x.MaxBatchSize
	}
	return 0
}

// Event - Order or Trade event.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id - Identifier of this event.
	// It can be for example `seq` number as in `Poloniex`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Event - Inner event, it can be order update or a trade.
	//
	// Types that are assignable to Event:
	//	*Event_Order
	//	*Event_Trade
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{4}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Event) GetOrder() *order.Order {
	if x, ok := x.GetEvent().(*Event_Order); ok {
		return x.Order
	}
	return nil
}

func (x *Event) GetTrade() *Trade {
	if x, ok := x.GetEvent().(*Event_Trade); ok {
		return x.Trade
	}
	return nil
}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Order struct {
	Order *order.Order `protobuf:"bytes,2,opt,name=order,proto3,oneof"`
}

type Event_Trade struct {
	Trade *Trade `protobuf:"bytes,3,opt,name=trade,proto3,oneof"`
}

func (*Event_Order) isEvent_Event() {}

func (*Event_Trade) isEvent_Event() {}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// true=bid, false=ask
	OrderId   string               `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	IsBid     bool                 `protobuf:"varint,2,opt,name=is_bid,json=isBid,proto3" json:"is_bid,omitempty"`
	Side      order.Side           `protobuf:"varint,3,opt,name=side,proto3,enum=order.Side" json:"side,omitempty"`
	Base      *base.ShortCommodity `protobuf:"bytes,4,opt,name=base,proto3" json:"base,omitempty"`
	Quote     *base.ShortCommodity `protobuf:"bytes,5,opt,name=quote,proto3" json:"quote,omitempty"`
	Size      int64                `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	SizeLots  int64                `protobuf:"varint,7,opt,name=size_lots,json=sizeLots,proto3" json:"size_lots,omitempty"`
	Price     int64                `protobuf:"varint,8,opt,name=price,proto3" json:"price,omitempty"`
	PriceLots int64                `protobuf:"varint,9,opt,name=price_lots,json=priceLots,proto3" json:"price_lots,omitempty"`
	// this is only filled for customer trade book queries
	Status Order_Status `protobuf:"varint,10,opt,name=status,proto3,enum=price.Order_Status" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[5]
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
	return file_price_proto_rawDescGZIP(), []int{5}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetIsBid() bool {
	if x != nil {
		return x.IsBid
	}
	return false
}

func (x *Order) GetSide() order.Side {
	if x != nil {
		return x.Side
	}
	return order.Side(0)
}

func (x *Order) GetBase() *base.ShortCommodity {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Order) GetQuote() *base.ShortCommodity {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *Order) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Order) GetSizeLots() int64 {
	if x != nil {
		return x.SizeLots
	}
	return 0
}

func (x *Order) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetPriceLots() int64 {
	if x != nil {
		return x.PriceLots
	}
	return 0
}

func (x *Order) GetStatus() Order_Status {
	if x != nil {
		return x.Status
	}
	return Order_PENDING
}

type PriceLot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price int64 `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	Size  int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *PriceLot) Reset() {
	*x = PriceLot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceLot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceLot) ProtoMessage() {}

func (x *PriceLot) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceLot.ProtoReflect.Descriptor instead.
func (*PriceLot) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{6}
}

func (x *PriceLot) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PriceLot) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type OrderBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *base.ShortCommodity `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Quote *base.ShortCommodity `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
	Bid   []*PriceLot          `protobuf:"bytes,3,rep,name=bid,proto3" json:"bid,omitempty"`
	Ask   []*PriceLot          `protobuf:"bytes,4,rep,name=ask,proto3" json:"ask,omitempty"`
}

func (x *OrderBook) Reset() {
	*x = OrderBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBook) ProtoMessage() {}

func (x *OrderBook) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBook.ProtoReflect.Descriptor instead.
func (*OrderBook) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{7}
}

func (x *OrderBook) GetBase() *base.ShortCommodity {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrderBook) GetQuote() *base.ShortCommodity {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *OrderBook) GetBid() []*PriceLot {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *OrderBook) GetAsk() []*PriceLot {
	if x != nil {
		return x.Ask
	}
	return nil
}

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base          *base.ShortCommodity `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Quote         *base.ShortCommodity `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
	Taker         string               `protobuf:"bytes,3,opt,name=taker,proto3" json:"taker,omitempty"`
	Maker         string               `protobuf:"bytes,4,opt,name=maker,proto3" json:"maker,omitempty"`
	OrderId       string               `protobuf:"bytes,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ClientOrderId string               `protobuf:"bytes,6,opt,name=client_order_id,json=clientOrderId,proto3" json:"client_order_id,omitempty"`
	Side          order.Side           `protobuf:"varint,7,opt,name=side,proto3,enum=order.Side" json:"side,omitempty"`
	Timestamp     int64                `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	FeeCost       int64                `protobuf:"varint,9,opt,name=fee_cost,json=feeCost,proto3" json:"fee_cost,omitempty"`
	Price         int64                `protobuf:"varint,10,opt,name=price,proto3" json:"price,omitempty"`
	Size          int64                `protobuf:"varint,11,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_price_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_price_proto_rawDescGZIP(), []int{8}
}

func (x *Trade) GetBase() *base.ShortCommodity {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Trade) GetQuote() *base.ShortCommodity {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *Trade) GetTaker() string {
	if x != nil {
		return x.Taker
	}
	return ""
}

func (x *Trade) GetMaker() string {
	if x != nil {
		return x.Maker
	}
	return ""
}

func (x *Trade) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Trade) GetClientOrderId() string {
	if x != nil {
		return x.ClientOrderId
	}
	return ""
}

func (x *Trade) GetSide() order.Side {
	if x != nil {
		return x.Side
	}
	return order.Side(0)
}

func (x *Trade) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Trade) GetFeeCost() int64 {
	if x != nil {
		return x.FeeCost
	}
	return 0
}

func (x *Trade) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Trade) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_price_proto protoreflect.FileDescriptor

var file_price_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a,
	0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69,
	0x74, 0x79, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x22, 0x30, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x6f, 0x6f, 0x6b, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x0a, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d,
	0x61, 0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6c, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x48, 0x00, 0x52, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x9f, 0x03, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x69, 0x73, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x42, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x64, 0x65,
	0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x6f, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x08, 0x73, 0x69, 0x7a,
	0x65, 0x4c, 0x6f, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x21, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x6f,
	0x74, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x4a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x49, 0x4c, 0x4c, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x3c, 0x0a, 0x08, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x09, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x21, 0x0a,
	0x03, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x74, 0x52, 0x03, 0x62, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x03, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x74, 0x52, 0x03,
	0x61, 0x73, 0x6b, 0x22, 0xe0, 0x02, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x28, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74,
	0x79, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6b,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73,
	0x69, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x66, 0x65, 0x65,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x61, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_price_proto_rawDescOnce sync.Once
	file_price_proto_rawDescData = file_price_proto_rawDesc
)

func file_price_proto_rawDescGZIP() []byte {
	file_price_proto_rawDescOnce.Do(func() {
		file_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_price_proto_rawDescData)
	})
	return file_price_proto_rawDescData
}

var file_price_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_price_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_price_proto_goTypes = []interface{}{
	(Order_Status)(0),           // 0: price.Order.Status
	(*MarketRequest)(nil),       // 1: price.MarketRequest
	(*LastTrades)(nil),          // 2: price.LastTrades
	(*PriceUpdate)(nil),         // 3: price.PriceUpdate
	(*ReadRequest)(nil),         // 4: price.ReadRequest
	(*Event)(nil),               // 5: price.Event
	(*Order)(nil),               // 6: price.Order
	(*PriceLot)(nil),            // 7: price.PriceLot
	(*OrderBook)(nil),           // 8: price.OrderBook
	(*Trade)(nil),               // 9: price.Trade
	(*base.ShortCommodity)(nil), // 10: base.ShortCommodity
	(*base.SecurityID)(nil),     // 11: base.SecurityID
	(exchange.Exchange)(0),      // 12: exchange.Exchange
	(*order.Order)(nil),         // 13: order.Order
	(order.Side)(0),             // 14: order.Side
}
var file_price_proto_depIdxs = []int32{
	10, // 0: price.MarketRequest.base:type_name -> base.ShortCommodity
	10, // 1: price.MarketRequest.quote:type_name -> base.ShortCommodity
	9,  // 2: price.LastTrades.trade:type_name -> price.Trade
	2,  // 3: price.PriceUpdate.recent_trades:type_name -> price.LastTrades
	8,  // 4: price.PriceUpdate.book:type_name -> price.OrderBook
	11, // 5: price.ReadRequest.security_id:type_name -> base.SecurityID
	12, // 6: price.ReadRequest.exchange:type_name -> exchange.Exchange
	13, // 7: price.Event.order:type_name -> order.Order
	9,  // 8: price.Event.trade:type_name -> price.Trade
	14, // 9: price.Order.side:type_name -> order.Side
	10, // 10: price.Order.base:type_name -> base.ShortCommodity
	10, // 11: price.Order.quote:type_name -> base.ShortCommodity
	0,  // 12: price.Order.status:type_name -> price.Order.Status
	10, // 13: price.OrderBook.base:type_name -> base.ShortCommodity
	10, // 14: price.OrderBook.quote:type_name -> base.ShortCommodity
	7,  // 15: price.OrderBook.bid:type_name -> price.PriceLot
	7,  // 16: price.OrderBook.ask:type_name -> price.PriceLot
	10, // 17: price.Trade.base:type_name -> base.ShortCommodity
	10, // 18: price.Trade.quote:type_name -> base.ShortCommodity
	14, // 19: price.Trade.side:type_name -> order.Side
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_price_proto_init() }
func file_price_proto_init() {
	if File_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketRequest); i {
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
		file_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastTrades); i {
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
		file_price_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceUpdate); i {
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
		file_price_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_price_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_price_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_price_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceLot); i {
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
		file_price_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBook); i {
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
		file_price_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trade); i {
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
	file_price_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PriceUpdate_RecentTrades)(nil),
		(*PriceUpdate_Book)(nil),
	}
	file_price_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Event_Order)(nil),
		(*Event_Trade)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_price_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_price_proto_goTypes,
		DependencyIndexes: file_price_proto_depIdxs,
		EnumInfos:         file_price_proto_enumTypes,
		MessageInfos:      file_price_proto_msgTypes,
	}.Build()
	File_price_proto = out.File
	file_price_proto_rawDesc = nil
	file_price_proto_goTypes = nil
	file_price_proto_depIdxs = nil
}
