// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: bookmarked_businesses.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Bookmarked_Businesses_CreateBookmarkedBusiness_FullMethodName  = "/bookmarked_businesses.Bookmarked_Businesses/CreateBookmarkedBusiness"
	Bookmarked_Businesses_DeleteBookmarkedBusiness_FullMethodName  = "/bookmarked_businesses.Bookmarked_Businesses/DeleteBookmarkedBusiness"
	Bookmarked_Businesses_GetBookmarkedBusinessById_FullMethodName = "/bookmarked_businesses.Bookmarked_Businesses/GetBookmarkedBusinessById"
	Bookmarked_Businesses_GetAllBookmarkedBusiness_FullMethodName  = "/bookmarked_businesses.Bookmarked_Businesses/GetAllBookmarkedBusiness"
)

// Bookmarked_BusinessesClient is the client API for Bookmarked_Businesses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Bookmarked_BusinessesClient interface {
	CreateBookmarkedBusiness(ctx context.Context, in *CreateBookmarkedBusinessRequest, opts ...grpc.CallOption) (*CreateBookmarkedBusinessResponse, error)
	DeleteBookmarkedBusiness(ctx context.Context, in *DeleteBookmarkedBusinessRequest, opts ...grpc.CallOption) (*DeleteBookmarkedBusinessResponse, error)
	GetBookmarkedBusinessById(ctx context.Context, in *GetBookmarkedBusinessByIdRequest, opts ...grpc.CallOption) (*GetBookmarkedBusinessByIdResponse, error)
	GetAllBookmarkedBusiness(ctx context.Context, in *GetAllBookmarkedBusinessRequest, opts ...grpc.CallOption) (*GetAllBookmarkedBusinessResponse, error)
}

type bookmarked_BusinessesClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarked_BusinessesClient(cc grpc.ClientConnInterface) Bookmarked_BusinessesClient {
	return &bookmarked_BusinessesClient{cc}
}

func (c *bookmarked_BusinessesClient) CreateBookmarkedBusiness(ctx context.Context, in *CreateBookmarkedBusinessRequest, opts ...grpc.CallOption) (*CreateBookmarkedBusinessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBookmarkedBusinessResponse)
	err := c.cc.Invoke(ctx, Bookmarked_Businesses_CreateBookmarkedBusiness_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarked_BusinessesClient) DeleteBookmarkedBusiness(ctx context.Context, in *DeleteBookmarkedBusinessRequest, opts ...grpc.CallOption) (*DeleteBookmarkedBusinessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBookmarkedBusinessResponse)
	err := c.cc.Invoke(ctx, Bookmarked_Businesses_DeleteBookmarkedBusiness_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarked_BusinessesClient) GetBookmarkedBusinessById(ctx context.Context, in *GetBookmarkedBusinessByIdRequest, opts ...grpc.CallOption) (*GetBookmarkedBusinessByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookmarkedBusinessByIdResponse)
	err := c.cc.Invoke(ctx, Bookmarked_Businesses_GetBookmarkedBusinessById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookmarked_BusinessesClient) GetAllBookmarkedBusiness(ctx context.Context, in *GetAllBookmarkedBusinessRequest, opts ...grpc.CallOption) (*GetAllBookmarkedBusinessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllBookmarkedBusinessResponse)
	err := c.cc.Invoke(ctx, Bookmarked_Businesses_GetAllBookmarkedBusiness_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Bookmarked_BusinessesServer is the server API for Bookmarked_Businesses service.
// All implementations must embed UnimplementedBookmarked_BusinessesServer
// for forward compatibility.
type Bookmarked_BusinessesServer interface {
	CreateBookmarkedBusiness(context.Context, *CreateBookmarkedBusinessRequest) (*CreateBookmarkedBusinessResponse, error)
	DeleteBookmarkedBusiness(context.Context, *DeleteBookmarkedBusinessRequest) (*DeleteBookmarkedBusinessResponse, error)
	GetBookmarkedBusinessById(context.Context, *GetBookmarkedBusinessByIdRequest) (*GetBookmarkedBusinessByIdResponse, error)
	GetAllBookmarkedBusiness(context.Context, *GetAllBookmarkedBusinessRequest) (*GetAllBookmarkedBusinessResponse, error)
	mustEmbedUnimplementedBookmarked_BusinessesServer()
}

// UnimplementedBookmarked_BusinessesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookmarked_BusinessesServer struct{}

func (UnimplementedBookmarked_BusinessesServer) CreateBookmarkedBusiness(context.Context, *CreateBookmarkedBusinessRequest) (*CreateBookmarkedBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookmarkedBusiness not implemented")
}
func (UnimplementedBookmarked_BusinessesServer) DeleteBookmarkedBusiness(context.Context, *DeleteBookmarkedBusinessRequest) (*DeleteBookmarkedBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookmarkedBusiness not implemented")
}
func (UnimplementedBookmarked_BusinessesServer) GetBookmarkedBusinessById(context.Context, *GetBookmarkedBusinessByIdRequest) (*GetBookmarkedBusinessByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookmarkedBusinessById not implemented")
}
func (UnimplementedBookmarked_BusinessesServer) GetAllBookmarkedBusiness(context.Context, *GetAllBookmarkedBusinessRequest) (*GetAllBookmarkedBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBookmarkedBusiness not implemented")
}
func (UnimplementedBookmarked_BusinessesServer) mustEmbedUnimplementedBookmarked_BusinessesServer() {}
func (UnimplementedBookmarked_BusinessesServer) testEmbeddedByValue()                               {}

// UnsafeBookmarked_BusinessesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Bookmarked_BusinessesServer will
// result in compilation errors.
type UnsafeBookmarked_BusinessesServer interface {
	mustEmbedUnimplementedBookmarked_BusinessesServer()
}

func RegisterBookmarked_BusinessesServer(s grpc.ServiceRegistrar, srv Bookmarked_BusinessesServer) {
	// If the following call pancis, it indicates UnimplementedBookmarked_BusinessesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Bookmarked_Businesses_ServiceDesc, srv)
}

func _Bookmarked_Businesses_CreateBookmarkedBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookmarkedBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Bookmarked_BusinessesServer).CreateBookmarkedBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookmarked_Businesses_CreateBookmarkedBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Bookmarked_BusinessesServer).CreateBookmarkedBusiness(ctx, req.(*CreateBookmarkedBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookmarked_Businesses_DeleteBookmarkedBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookmarkedBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Bookmarked_BusinessesServer).DeleteBookmarkedBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookmarked_Businesses_DeleteBookmarkedBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Bookmarked_BusinessesServer).DeleteBookmarkedBusiness(ctx, req.(*DeleteBookmarkedBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookmarked_Businesses_GetBookmarkedBusinessById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookmarkedBusinessByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Bookmarked_BusinessesServer).GetBookmarkedBusinessById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookmarked_Businesses_GetBookmarkedBusinessById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Bookmarked_BusinessesServer).GetBookmarkedBusinessById(ctx, req.(*GetBookmarkedBusinessByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookmarked_Businesses_GetAllBookmarkedBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBookmarkedBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Bookmarked_BusinessesServer).GetAllBookmarkedBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookmarked_Businesses_GetAllBookmarkedBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Bookmarked_BusinessesServer).GetAllBookmarkedBusiness(ctx, req.(*GetAllBookmarkedBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bookmarked_Businesses_ServiceDesc is the grpc.ServiceDesc for Bookmarked_Businesses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookmarked_Businesses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookmarked_businesses.Bookmarked_Businesses",
	HandlerType: (*Bookmarked_BusinessesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBookmarkedBusiness",
			Handler:    _Bookmarked_Businesses_CreateBookmarkedBusiness_Handler,
		},
		{
			MethodName: "DeleteBookmarkedBusiness",
			Handler:    _Bookmarked_Businesses_DeleteBookmarkedBusiness_Handler,
		},
		{
			MethodName: "GetBookmarkedBusinessById",
			Handler:    _Bookmarked_Businesses_GetBookmarkedBusinessById_Handler,
		},
		{
			MethodName: "GetAllBookmarkedBusiness",
			Handler:    _Bookmarked_Businesses_GetAllBookmarkedBusiness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookmarked_businesses.proto",
}
