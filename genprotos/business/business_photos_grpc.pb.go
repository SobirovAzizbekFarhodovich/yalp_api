// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: business_photos.proto

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
	Business_Photos_CreateBusinessPhotos_FullMethodName     = "/business_photos.Business_Photos/CreateBusinessPhotos"
	Business_Photos_UpdateBusinessPhotos_FullMethodName     = "/business_photos.Business_Photos/UpdateBusinessPhotos"
	Business_Photos_DeleteBusinessPhotos_FullMethodName     = "/business_photos.Business_Photos/DeleteBusinessPhotos"
	Business_Photos_GetByBusinessId_FullMethodName          = "/business_photos.Business_Photos/GetByBusinessId"
	Business_Photos_GetBusinessPhotosByOwner_FullMethodName = "/business_photos.Business_Photos/GetBusinessPhotosByOwner"
)

// Business_PhotosClient is the client API for Business_Photos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Business_PhotosClient interface {
	CreateBusinessPhotos(ctx context.Context, in *CreateBusinessPhotosRequest, opts ...grpc.CallOption) (*CreateBusinessPhotosResponse, error)
	UpdateBusinessPhotos(ctx context.Context, in *UpdateBusinessPhotosRequest, opts ...grpc.CallOption) (*UpdateBusinessPhotosResponse, error)
	DeleteBusinessPhotos(ctx context.Context, in *DeleteBusinessPhotosRequest, opts ...grpc.CallOption) (*DeleteBusinessPhotosResponse, error)
	GetByBusinessId(ctx context.Context, in *GetBusinessIdRequest, opts ...grpc.CallOption) (*GetBusinessIdResponse, error)
	GetBusinessPhotosByOwner(ctx context.Context, in *GetBusinessPhotosByOwnerRequest, opts ...grpc.CallOption) (*GetBusinessPhotosByOwnerResponse, error)
}

type business_PhotosClient struct {
	cc grpc.ClientConnInterface
}

func NewBusiness_PhotosClient(cc grpc.ClientConnInterface) Business_PhotosClient {
	return &business_PhotosClient{cc}
}

func (c *business_PhotosClient) CreateBusinessPhotos(ctx context.Context, in *CreateBusinessPhotosRequest, opts ...grpc.CallOption) (*CreateBusinessPhotosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBusinessPhotosResponse)
	err := c.cc.Invoke(ctx, Business_Photos_CreateBusinessPhotos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *business_PhotosClient) UpdateBusinessPhotos(ctx context.Context, in *UpdateBusinessPhotosRequest, opts ...grpc.CallOption) (*UpdateBusinessPhotosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBusinessPhotosResponse)
	err := c.cc.Invoke(ctx, Business_Photos_UpdateBusinessPhotos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *business_PhotosClient) DeleteBusinessPhotos(ctx context.Context, in *DeleteBusinessPhotosRequest, opts ...grpc.CallOption) (*DeleteBusinessPhotosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBusinessPhotosResponse)
	err := c.cc.Invoke(ctx, Business_Photos_DeleteBusinessPhotos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *business_PhotosClient) GetByBusinessId(ctx context.Context, in *GetBusinessIdRequest, opts ...grpc.CallOption) (*GetBusinessIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBusinessIdResponse)
	err := c.cc.Invoke(ctx, Business_Photos_GetByBusinessId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *business_PhotosClient) GetBusinessPhotosByOwner(ctx context.Context, in *GetBusinessPhotosByOwnerRequest, opts ...grpc.CallOption) (*GetBusinessPhotosByOwnerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBusinessPhotosByOwnerResponse)
	err := c.cc.Invoke(ctx, Business_Photos_GetBusinessPhotosByOwner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Business_PhotosServer is the server API for Business_Photos service.
// All implementations must embed UnimplementedBusiness_PhotosServer
// for forward compatibility.
type Business_PhotosServer interface {
	CreateBusinessPhotos(context.Context, *CreateBusinessPhotosRequest) (*CreateBusinessPhotosResponse, error)
	UpdateBusinessPhotos(context.Context, *UpdateBusinessPhotosRequest) (*UpdateBusinessPhotosResponse, error)
	DeleteBusinessPhotos(context.Context, *DeleteBusinessPhotosRequest) (*DeleteBusinessPhotosResponse, error)
	GetByBusinessId(context.Context, *GetBusinessIdRequest) (*GetBusinessIdResponse, error)
	GetBusinessPhotosByOwner(context.Context, *GetBusinessPhotosByOwnerRequest) (*GetBusinessPhotosByOwnerResponse, error)
	mustEmbedUnimplementedBusiness_PhotosServer()
}

// UnimplementedBusiness_PhotosServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBusiness_PhotosServer struct{}

func (UnimplementedBusiness_PhotosServer) CreateBusinessPhotos(context.Context, *CreateBusinessPhotosRequest) (*CreateBusinessPhotosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusinessPhotos not implemented")
}
func (UnimplementedBusiness_PhotosServer) UpdateBusinessPhotos(context.Context, *UpdateBusinessPhotosRequest) (*UpdateBusinessPhotosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBusinessPhotos not implemented")
}
func (UnimplementedBusiness_PhotosServer) DeleteBusinessPhotos(context.Context, *DeleteBusinessPhotosRequest) (*DeleteBusinessPhotosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBusinessPhotos not implemented")
}
func (UnimplementedBusiness_PhotosServer) GetByBusinessId(context.Context, *GetBusinessIdRequest) (*GetBusinessIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByBusinessId not implemented")
}
func (UnimplementedBusiness_PhotosServer) GetBusinessPhotosByOwner(context.Context, *GetBusinessPhotosByOwnerRequest) (*GetBusinessPhotosByOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessPhotosByOwner not implemented")
}
func (UnimplementedBusiness_PhotosServer) mustEmbedUnimplementedBusiness_PhotosServer() {}
func (UnimplementedBusiness_PhotosServer) testEmbeddedByValue()                         {}

// UnsafeBusiness_PhotosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Business_PhotosServer will
// result in compilation errors.
type UnsafeBusiness_PhotosServer interface {
	mustEmbedUnimplementedBusiness_PhotosServer()
}

func RegisterBusiness_PhotosServer(s grpc.ServiceRegistrar, srv Business_PhotosServer) {
	// If the following call pancis, it indicates UnimplementedBusiness_PhotosServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Business_Photos_ServiceDesc, srv)
}

func _Business_Photos_CreateBusinessPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusinessPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Business_PhotosServer).CreateBusinessPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_Photos_CreateBusinessPhotos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Business_PhotosServer).CreateBusinessPhotos(ctx, req.(*CreateBusinessPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Business_Photos_UpdateBusinessPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBusinessPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Business_PhotosServer).UpdateBusinessPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_Photos_UpdateBusinessPhotos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Business_PhotosServer).UpdateBusinessPhotos(ctx, req.(*UpdateBusinessPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Business_Photos_DeleteBusinessPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBusinessPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Business_PhotosServer).DeleteBusinessPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_Photos_DeleteBusinessPhotos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Business_PhotosServer).DeleteBusinessPhotos(ctx, req.(*DeleteBusinessPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Business_Photos_GetByBusinessId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Business_PhotosServer).GetByBusinessId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_Photos_GetByBusinessId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Business_PhotosServer).GetByBusinessId(ctx, req.(*GetBusinessIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Business_Photos_GetBusinessPhotosByOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessPhotosByOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Business_PhotosServer).GetBusinessPhotosByOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_Photos_GetBusinessPhotosByOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Business_PhotosServer).GetBusinessPhotosByOwner(ctx, req.(*GetBusinessPhotosByOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Business_Photos_ServiceDesc is the grpc.ServiceDesc for Business_Photos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Business_Photos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "business_photos.Business_Photos",
	HandlerType: (*Business_PhotosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBusinessPhotos",
			Handler:    _Business_Photos_CreateBusinessPhotos_Handler,
		},
		{
			MethodName: "UpdateBusinessPhotos",
			Handler:    _Business_Photos_UpdateBusinessPhotos_Handler,
		},
		{
			MethodName: "DeleteBusinessPhotos",
			Handler:    _Business_Photos_DeleteBusinessPhotos_Handler,
		},
		{
			MethodName: "GetByBusinessId",
			Handler:    _Business_Photos_GetByBusinessId_Handler,
		},
		{
			MethodName: "GetBusinessPhotosByOwner",
			Handler:    _Business_Photos_GetBusinessPhotosByOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "business_photos.proto",
}
