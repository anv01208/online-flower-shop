// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: pkg/cart/pb/cart.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CartService_GetCartItems_FullMethodName   = "/cart.CartService/GetCartItems"
	CartService_AddToCart_FullMethodName      = "/cart.CartService/AddToCart"
	CartService_RemoveFromCart_FullMethodName = "/cart.CartService/RemoveFromCart"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	GetCartItems(ctx context.Context, in *ViewCartRequest, opts ...grpc.CallOption) (CartService_GetCartItemsClient, error)
	AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error)
	RemoveFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*RemoveFromCartResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) GetCartItems(ctx context.Context, in *ViewCartRequest, opts ...grpc.CallOption) (CartService_GetCartItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CartService_ServiceDesc.Streams[0], CartService_GetCartItems_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &cartServiceGetCartItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CartService_GetCartItemsClient interface {
	Recv() (*ViewCartResponse, error)
	grpc.ClientStream
}

type cartServiceGetCartItemsClient struct {
	grpc.ClientStream
}

func (x *cartServiceGetCartItemsClient) Recv() (*ViewCartResponse, error) {
	m := new(ViewCartResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cartServiceClient) AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error) {
	out := new(AddToCartResponse)
	err := c.cc.Invoke(ctx, CartService_AddToCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*RemoveFromCartResponse, error) {
	out := new(RemoveFromCartResponse)
	err := c.cc.Invoke(ctx, CartService_RemoveFromCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	GetCartItems(*ViewCartRequest, CartService_GetCartItemsServer) error
	AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error)
	RemoveFromCart(context.Context, *RemoveFromCartRequest) (*RemoveFromCartResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) GetCartItems(*ViewCartRequest, CartService_GetCartItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCartItems not implemented")
}
func (UnimplementedCartServiceServer) AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveFromCart(context.Context, *RemoveFromCartRequest) (*RemoveFromCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCart not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_GetCartItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ViewCartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CartServiceServer).GetCartItems(m, &cartServiceGetCartItemsServer{stream})
}

type CartService_GetCartItemsServer interface {
	Send(*ViewCartResponse) error
	grpc.ServerStream
}

type cartServiceGetCartItemsServer struct {
	grpc.ServerStream
}

func (x *cartServiceGetCartItemsServer) Send(m *ViewCartResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CartService_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_AddToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddToCart(ctx, req.(*AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_RemoveFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveFromCart(ctx, req.(*RemoveFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToCart",
			Handler:    _CartService_AddToCart_Handler,
		},
		{
			MethodName: "RemoveFromCart",
			Handler:    _CartService_RemoveFromCart_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCartItems",
			Handler:       _CartService_GetCartItems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/cart/pb/cart.proto",
}
