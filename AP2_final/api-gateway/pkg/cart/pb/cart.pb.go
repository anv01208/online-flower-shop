// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: pkg/cart/pb/cart.proto

package pb

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

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity  int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UserId    int64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddToCartRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddToCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddToCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddToCartResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddToCartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId int64 `protobuf:"varint,1,opt,name=cartItemId,proto3" json:"cartItemId,omitempty"`
	ProductId  int64 `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity   int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{2}
}

func (x *CartItem) GetCartItemId() int64 {
	if x != nil {
		return x.CartItemId
	}
	return 0
}

func (x *CartItem) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ViewCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ViewCartRequest) Reset() {
	*x = ViewCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartRequest) ProtoMessage() {}

func (x *ViewCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartRequest.ProtoReflect.Descriptor instead.
func (*ViewCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{3}
}

func (x *ViewCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ViewCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error    string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	CartItem []*CartItem `protobuf:"bytes,3,rep,name=cartItem,proto3" json:"cartItem,omitempty"`
}

func (x *ViewCartResponse) Reset() {
	*x = ViewCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResponse) ProtoMessage() {}

func (x *ViewCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartResponse.ProtoReflect.Descriptor instead.
func (*ViewCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{4}
}

func (x *ViewCartResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ViewCartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ViewCartResponse) GetCartItem() []*CartItem {
	if x != nil {
		return x.CartItem
	}
	return nil
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItemId int64 `protobuf:"varint,1,opt,name=cartItemId,proto3" json:"cartItemId,omitempty"`
	UserId     int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveFromCartRequest) GetCartItemId() int64 {
	if x != nil {
		return x.CartItemId
	}
	return 0
}

func (x *RemoveFromCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RemoveFromCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemoveFromCartResponse) Reset() {
	*x = RemoveFromCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFromCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartResponse) ProtoMessage() {}

func (x *RemoveFromCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartResponse.ProtoReflect.Descriptor instead.
func (*RemoveFromCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveFromCartResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RemoveFromCartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_cart_pb_cart_proto protoreflect.FileDescriptor

var file_pkg_cart_pb_cart_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x64,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x64, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x29, 0x0a,
	0x0f, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x10, 0x56, 0x69, 0x65, 0x77,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x4f, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xdf, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x3e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_cart_pb_cart_proto_rawDescOnce sync.Once
	file_pkg_cart_pb_cart_proto_rawDescData = file_pkg_cart_pb_cart_proto_rawDesc
)

func file_pkg_cart_pb_cart_proto_rawDescGZIP() []byte {
	file_pkg_cart_pb_cart_proto_rawDescOnce.Do(func() {
		file_pkg_cart_pb_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_cart_pb_cart_proto_rawDescData)
	})
	return file_pkg_cart_pb_cart_proto_rawDescData
}

var file_pkg_cart_pb_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_cart_pb_cart_proto_goTypes = []interface{}{
	(*AddToCartRequest)(nil),       // 0: cart.AddToCartRequest
	(*AddToCartResponse)(nil),      // 1: cart.AddToCartResponse
	(*CartItem)(nil),               // 2: cart.CartItem
	(*ViewCartRequest)(nil),        // 3: cart.ViewCartRequest
	(*ViewCartResponse)(nil),       // 4: cart.ViewCartResponse
	(*RemoveFromCartRequest)(nil),  // 5: cart.RemoveFromCartRequest
	(*RemoveFromCartResponse)(nil), // 6: cart.RemoveFromCartResponse
}
var file_pkg_cart_pb_cart_proto_depIdxs = []int32{
	2, // 0: cart.ViewCartResponse.cartItem:type_name -> cart.CartItem
	3, // 1: cart.CartService.GetCartItems:input_type -> cart.ViewCartRequest
	0, // 2: cart.CartService.AddToCart:input_type -> cart.AddToCartRequest
	5, // 3: cart.CartService.RemoveFromCart:input_type -> cart.RemoveFromCartRequest
	4, // 4: cart.CartService.GetCartItems:output_type -> cart.ViewCartResponse
	1, // 5: cart.CartService.AddToCart:output_type -> cart.AddToCartResponse
	6, // 6: cart.CartService.RemoveFromCart:output_type -> cart.RemoveFromCartResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_cart_pb_cart_proto_init() }
func file_pkg_cart_pb_cart_proto_init() {
	if File_pkg_cart_pb_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_cart_pb_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartRequest); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponse); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartRequest); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartResponse); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFromCartRequest); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFromCartResponse); i {
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
			RawDescriptor: file_pkg_cart_pb_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_cart_pb_cart_proto_goTypes,
		DependencyIndexes: file_pkg_cart_pb_cart_proto_depIdxs,
		MessageInfos:      file_pkg_cart_pb_cart_proto_msgTypes,
	}.Build()
	File_pkg_cart_pb_cart_proto = out.File
	file_pkg_cart_pb_cart_proto_rawDesc = nil
	file_pkg_cart_pb_cart_proto_goTypes = nil
	file_pkg_cart_pb_cart_proto_depIdxs = nil
}
