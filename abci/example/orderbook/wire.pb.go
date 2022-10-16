// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: wire.proto

package orderbook

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

type OrderAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity  float64 `protobuf:"fixed64,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	AskPrice  float64 `protobuf:"fixed64,2,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	Owner     uint64  `protobuf:"varint,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Signature []byte  `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *OrderAsk) Reset() {
	*x = OrderAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wire_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAsk) ProtoMessage() {}

func (x *OrderAsk) ProtoReflect() protoreflect.Message {
	mi := &file_wire_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAsk.ProtoReflect.Descriptor instead.
func (*OrderAsk) Descriptor() ([]byte, []int) {
	return file_wire_proto_rawDescGZIP(), []int{0}
}

func (x *OrderAsk) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderAsk) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *OrderAsk) GetOwner() uint64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *OrderAsk) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type OrderBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxQuantity float64 `protobuf:"fixed64,1,opt,name=max_quantity,json=maxQuantity,proto3" json:"max_quantity,omitempty"`
	MaxPrice    float64 `protobuf:"fixed64,2,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Owner       uint64  `protobuf:"varint,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Signature   []byte  `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *OrderBid) Reset() {
	*x = OrderBid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wire_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBid) ProtoMessage() {}

func (x *OrderBid) ProtoReflect() protoreflect.Message {
	mi := &file_wire_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBid.ProtoReflect.Descriptor instead.
func (*OrderBid) Descriptor() ([]byte, []int) {
	return file_wire_proto_rawDescGZIP(), []int{1}
}

func (x *OrderBid) GetMaxQuantity() float64 {
	if x != nil {
		return x.MaxQuantity
	}
	return 0
}

func (x *OrderBid) GetMaxPrice() float64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *OrderBid) GetOwner() uint64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *OrderBid) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the denomination that the buyer receives i.e. EUR
	BuyersDenomination string `protobuf:"bytes,1,opt,name=buyers_denomination,json=buyersDenomination,proto3" json:"buyers_denomination,omitempty"`
	// the denomination that the seller receives i.e. USD
	SellersDenomination string `protobuf:"bytes,2,opt,name=sellers_denomination,json=sellersDenomination,proto3" json:"sellers_denomination,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wire_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_wire_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_wire_proto_rawDescGZIP(), []int{2}
}

func (x *Pair) GetBuyersDenomination() string {
	if x != nil {
		return x.BuyersDenomination
	}
	return ""
}

func (x *Pair) GetSellersDenomination() string {
	if x != nil {
		return x.SellersDenomination
	}
	return ""
}

type Commodity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom    string  `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Quantity float64 `protobuf:"fixed64,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Commodity) Reset() {
	*x = Commodity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wire_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commodity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commodity) ProtoMessage() {}

func (x *Commodity) ProtoReflect() protoreflect.Message {
	mi := &file_wire_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commodity.ProtoReflect.Descriptor instead.
func (*Commodity) Descriptor() ([]byte, []int) {
	return file_wire_proto_rawDescGZIP(), []int{3}
}

func (x *Commodity) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Commodity) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	PublicKey [][]byte `protobuf:"bytes,2,rep,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// the set of commodities that the account has
	Commodities []*Commodity `protobuf:"bytes,3,rep,name=commodities,proto3" json:"commodities,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wire_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_wire_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_wire_proto_rawDescGZIP(), []int{4}
}

func (x *Account) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Account) GetPublicKey() [][]byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Account) GetCommodities() []*Commodity {
	if x != nil {
		return x.Commodities
	}
	return nil
}

// TradeSet is the transaction that eventually is committed in a block
// It is derived from a group of MsgBid and MsgAsk's
type TradeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair *Pair `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"` // i.e. EUR/USD
	// the set of matched trades for that peer
	MatchedOrders []*MatchedOrder `protobuf:"bytes,2,rep,name=matched_orders,json=matchedOrders,proto3" json:"matched_orders,omitempty"`
}

func (x *TradeSet) Reset() {
	*x = TradeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wire_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeSet) ProtoMessage() {}

func (x *TradeSet) ProtoReflect() protoreflect.Message {
	mi := &file_wire_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeSet.ProtoReflect.Descriptor instead.
func (*TradeSet) Descriptor() ([]byte, []int) {
	return file_wire_proto_rawDescGZIP(), []int{5}
}

func (x *TradeSet) GetPair() *Pair {
	if x != nil {
		return x.Pair
	}
	return nil
}

func (x *TradeSet) GetMatchedOrders() []*MatchedOrder {
	if x != nil {
		return x.MatchedOrders
	}
	return nil
}

type MatchedOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderAsk *OrderAsk `protobuf:"bytes,1,opt,name=order_ask,json=orderAsk,proto3" json:"order_ask,omitempty"`
	OrderBid *OrderBid `protobuf:"bytes,2,opt,name=order_bid,json=orderBid,proto3" json:"order_bid,omitempty"`
}

func (x *MatchedOrder) Reset() {
	*x = MatchedOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wire_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchedOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchedOrder) ProtoMessage() {}

func (x *MatchedOrder) ProtoReflect() protoreflect.Message {
	mi := &file_wire_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchedOrder.ProtoReflect.Descriptor instead.
func (*MatchedOrder) Descriptor() ([]byte, []int) {
	return file_wire_proto_rawDescGZIP(), []int{6}
}

func (x *MatchedOrder) GetOrderAsk() *OrderAsk {
	if x != nil {
		return x.OrderAsk
	}
	return nil
}

func (x *MatchedOrder) GetOrderBid() *OrderBid {
	if x != nil {
		return x.OrderBid
	}
	return nil
}

var File_wire_proto protoreflect.FileDescriptor

var file_wire_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x77, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x77, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x7e, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x61, 0x78, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x6a, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x2f, 0x0a, 0x13, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x73, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x62, 0x75, 0x79, 0x65, 0x72, 0x73, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x73, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x09,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x76, 0x0a, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x23, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04,
	0x70, 0x61, 0x69, 0x72, 0x12, 0x3e, 0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x22, 0x72, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x73,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x73, 0x6b, 0x52, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x41, 0x73, 0x6b, 0x12, 0x30, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x69, 0x64, 0x52, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x69, 0x64, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x62, 0x63,
	0x69, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wire_proto_rawDescOnce sync.Once
	file_wire_proto_rawDescData = file_wire_proto_rawDesc
)

func file_wire_proto_rawDescGZIP() []byte {
	file_wire_proto_rawDescOnce.Do(func() {
		file_wire_proto_rawDescData = protoimpl.X.CompressGZIP(file_wire_proto_rawDescData)
	})
	return file_wire_proto_rawDescData
}

var file_wire_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wire_proto_goTypes = []interface{}{
	(*OrderAsk)(nil),     // 0: orderbook.OrderAsk
	(*OrderBid)(nil),     // 1: orderbook.OrderBid
	(*Pair)(nil),         // 2: orderbook.Pair
	(*Commodity)(nil),    // 3: orderbook.Commodity
	(*Account)(nil),      // 4: orderbook.Account
	(*TradeSet)(nil),     // 5: orderbook.TradeSet
	(*MatchedOrder)(nil), // 6: orderbook.MatchedOrder
}
var file_wire_proto_depIdxs = []int32{
	3, // 0: orderbook.Account.commodities:type_name -> orderbook.Commodity
	2, // 1: orderbook.TradeSet.pair:type_name -> orderbook.Pair
	6, // 2: orderbook.TradeSet.matched_orders:type_name -> orderbook.MatchedOrder
	0, // 3: orderbook.MatchedOrder.order_ask:type_name -> orderbook.OrderAsk
	1, // 4: orderbook.MatchedOrder.order_bid:type_name -> orderbook.OrderBid
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_wire_proto_init() }
func file_wire_proto_init() {
	if File_wire_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wire_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAsk); i {
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
		file_wire_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBid); i {
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
		file_wire_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_wire_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commodity); i {
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
		file_wire_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_wire_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeSet); i {
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
		file_wire_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchedOrder); i {
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
			RawDescriptor: file_wire_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wire_proto_goTypes,
		DependencyIndexes: file_wire_proto_depIdxs,
		MessageInfos:      file_wire_proto_msgTypes,
	}.Build()
	File_wire_proto = out.File
	file_wire_proto_rawDesc = nil
	file_wire_proto_goTypes = nil
	file_wire_proto_depIdxs = nil
}
