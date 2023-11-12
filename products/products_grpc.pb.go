// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: products.proto

package products

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

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsClient interface {
	// Demonstrates a request response call
	GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error)
	// Demonstrates a server streamer call
	GetProducts(ctx context.Context, in *GetProductsReq, opts ...grpc.CallOption) (Products_GetProductsClient, error)
	// Demonstrates a client streaming call
	CreateProducts(ctx context.Context, opts ...grpc.CallOption) (Products_CreateProductsClient, error)
	// Demonstrates a bidirectional streaming call
	CreateProductsInStream(ctx context.Context, opts ...grpc.CallOption) (Products_CreateProductsInStreamClient, error)
}

type productsClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsClient(cc grpc.ClientConnInterface) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error) {
	out := new(GetProductResp)
	err := c.cc.Invoke(ctx, "/rpc.Products/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProducts(ctx context.Context, in *GetProductsReq, opts ...grpc.CallOption) (Products_GetProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Products_ServiceDesc.Streams[0], "/rpc.Products/GetProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productsGetProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Products_GetProductsClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type productsGetProductsClient struct {
	grpc.ClientStream
}

func (x *productsGetProductsClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productsClient) CreateProducts(ctx context.Context, opts ...grpc.CallOption) (Products_CreateProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Products_ServiceDesc.Streams[1], "/rpc.Products/CreateProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productsCreateProductsClient{stream}
	return x, nil
}

type Products_CreateProductsClient interface {
	Send(*Product) error
	CloseAndRecv() (*CreateProductsResp, error)
	grpc.ClientStream
}

type productsCreateProductsClient struct {
	grpc.ClientStream
}

func (x *productsCreateProductsClient) Send(m *Product) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productsCreateProductsClient) CloseAndRecv() (*CreateProductsResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateProductsResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productsClient) CreateProductsInStream(ctx context.Context, opts ...grpc.CallOption) (Products_CreateProductsInStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Products_ServiceDesc.Streams[2], "/rpc.Products/CreateProductsInStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &productsCreateProductsInStreamClient{stream}
	return x, nil
}

type Products_CreateProductsInStreamClient interface {
	Send(*Product) error
	Recv() (*Product, error)
	grpc.ClientStream
}

type productsCreateProductsInStreamClient struct {
	grpc.ClientStream
}

func (x *productsCreateProductsInStreamClient) Send(m *Product) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productsCreateProductsInStreamClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductsServer is the server API for Products service.
// All implementations must embed UnimplementedProductsServer
// for forward compatibility
type ProductsServer interface {
	// Demonstrates a request response call
	GetProduct(context.Context, *GetProductReq) (*GetProductResp, error)
	// Demonstrates a server streamer call
	GetProducts(*GetProductsReq, Products_GetProductsServer) error
	// Demonstrates a client streaming call
	CreateProducts(Products_CreateProductsServer) error
	// Demonstrates a bidirectional streaming call
	CreateProductsInStream(Products_CreateProductsInStreamServer) error
	mustEmbedUnimplementedProductsServer()
}

// UnimplementedProductsServer must be embedded to have forward compatible implementations.
type UnimplementedProductsServer struct {
}

func (UnimplementedProductsServer) GetProduct(context.Context, *GetProductReq) (*GetProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductsServer) GetProducts(*GetProductsReq, Products_GetProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductsServer) CreateProducts(Products_CreateProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProducts not implemented")
}
func (UnimplementedProductsServer) CreateProductsInStream(Products_CreateProductsInStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProductsInStream not implemented")
}
func (UnimplementedProductsServer) mustEmbedUnimplementedProductsServer() {}

// UnsafeProductsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServer will
// result in compilation errors.
type UnsafeProductsServer interface {
	mustEmbedUnimplementedProductsServer()
}

func RegisterProductsServer(s grpc.ServiceRegistrar, srv ProductsServer) {
	s.RegisterService(&Products_ServiceDesc, srv)
}

func _Products_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Products/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProduct(ctx, req.(*GetProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetProductsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductsServer).GetProducts(m, &productsGetProductsServer{stream})
}

type Products_GetProductsServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type productsGetProductsServer struct {
	grpc.ServerStream
}

func (x *productsGetProductsServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _Products_CreateProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductsServer).CreateProducts(&productsCreateProductsServer{stream})
}

type Products_CreateProductsServer interface {
	SendAndClose(*CreateProductsResp) error
	Recv() (*Product, error)
	grpc.ServerStream
}

type productsCreateProductsServer struct {
	grpc.ServerStream
}

func (x *productsCreateProductsServer) SendAndClose(m *CreateProductsResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productsCreateProductsServer) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Products_CreateProductsInStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductsServer).CreateProductsInStream(&productsCreateProductsInStreamServer{stream})
}

type Products_CreateProductsInStreamServer interface {
	Send(*Product) error
	Recv() (*Product, error)
	grpc.ServerStream
}

type productsCreateProductsInStreamServer struct {
	grpc.ServerStream
}

func (x *productsCreateProductsInStreamServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productsCreateProductsInStreamServer) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Products_ServiceDesc is the grpc.ServiceDesc for Products service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Products_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _Products_GetProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProducts",
			Handler:       _Products_GetProducts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateProducts",
			Handler:       _Products_CreateProducts_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateProductsInStream",
			Handler:       _Products_CreateProductsInStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "products.proto",
}
