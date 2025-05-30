// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cloudstack/management/resource/icon/v1/icon.gen.proto

package iconv1

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
	IconService_DeleteResourceIcon_FullMethodName = "/cloudstack.management.resource.icon.v1.IconService/DeleteResourceIcon"
	IconService_ListResourceIcon_FullMethodName   = "/cloudstack.management.resource.icon.v1.IconService/ListResourceIcon"
	IconService_UploadResourceIcon_FullMethodName = "/cloudstack.management.resource.icon.v1.IconService/UploadResourceIcon"
)

// IconServiceClient is the client API for IconService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// IconService provides operations for managing Resource.Icons
type IconServiceClient interface {
	// DeleteResourceIcon deletes the resource icon from the specified resource(s)
	DeleteResourceIcon(ctx context.Context, in *DeleteResourceIconRequest, opts ...grpc.CallOption) (*DeleteResourceIconResponse, error)
	// ListResourceIcon Lists the resource icon for the specified resource(s)
	ListResourceIcon(ctx context.Context, in *ListResourceIconRequest, opts ...grpc.CallOption) (*ListResourceIconResponse, error)
	// UploadResourceIcon Uploads an icon for the specified resource(s)
	UploadResourceIcon(ctx context.Context, in *UploadResourceIconRequest, opts ...grpc.CallOption) (*UploadResourceIconResponse, error)
}

type iconServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIconServiceClient(cc grpc.ClientConnInterface) IconServiceClient {
	return &iconServiceClient{cc}
}

func (c *iconServiceClient) DeleteResourceIcon(ctx context.Context, in *DeleteResourceIconRequest, opts ...grpc.CallOption) (*DeleteResourceIconResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResourceIconResponse)
	err := c.cc.Invoke(ctx, IconService_DeleteResourceIcon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iconServiceClient) ListResourceIcon(ctx context.Context, in *ListResourceIconRequest, opts ...grpc.CallOption) (*ListResourceIconResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResourceIconResponse)
	err := c.cc.Invoke(ctx, IconService_ListResourceIcon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iconServiceClient) UploadResourceIcon(ctx context.Context, in *UploadResourceIconRequest, opts ...grpc.CallOption) (*UploadResourceIconResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResourceIconResponse)
	err := c.cc.Invoke(ctx, IconService_UploadResourceIcon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IconServiceServer is the server API for IconService service.
// All implementations must embed UnimplementedIconServiceServer
// for forward compatibility.
//
// IconService provides operations for managing Resource.Icons
type IconServiceServer interface {
	// DeleteResourceIcon deletes the resource icon from the specified resource(s)
	DeleteResourceIcon(context.Context, *DeleteResourceIconRequest) (*DeleteResourceIconResponse, error)
	// ListResourceIcon Lists the resource icon for the specified resource(s)
	ListResourceIcon(context.Context, *ListResourceIconRequest) (*ListResourceIconResponse, error)
	// UploadResourceIcon Uploads an icon for the specified resource(s)
	UploadResourceIcon(context.Context, *UploadResourceIconRequest) (*UploadResourceIconResponse, error)
	mustEmbedUnimplementedIconServiceServer()
}

// UnimplementedIconServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIconServiceServer struct{}

func (UnimplementedIconServiceServer) DeleteResourceIcon(context.Context, *DeleteResourceIconRequest) (*DeleteResourceIconResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResourceIcon not implemented")
}
func (UnimplementedIconServiceServer) ListResourceIcon(context.Context, *ListResourceIconRequest) (*ListResourceIconResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResourceIcon not implemented")
}
func (UnimplementedIconServiceServer) UploadResourceIcon(context.Context, *UploadResourceIconRequest) (*UploadResourceIconResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadResourceIcon not implemented")
}
func (UnimplementedIconServiceServer) mustEmbedUnimplementedIconServiceServer() {}
func (UnimplementedIconServiceServer) testEmbeddedByValue()                     {}

// UnsafeIconServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IconServiceServer will
// result in compilation errors.
type UnsafeIconServiceServer interface {
	mustEmbedUnimplementedIconServiceServer()
}

func RegisterIconServiceServer(s grpc.ServiceRegistrar, srv IconServiceServer) {
	// If the following call pancis, it indicates UnimplementedIconServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IconService_ServiceDesc, srv)
}

func _IconService_DeleteResourceIcon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceIconRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IconServiceServer).DeleteResourceIcon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IconService_DeleteResourceIcon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IconServiceServer).DeleteResourceIcon(ctx, req.(*DeleteResourceIconRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IconService_ListResourceIcon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourceIconRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IconServiceServer).ListResourceIcon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IconService_ListResourceIcon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IconServiceServer).ListResourceIcon(ctx, req.(*ListResourceIconRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IconService_UploadResourceIcon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadResourceIconRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IconServiceServer).UploadResourceIcon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IconService_UploadResourceIcon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IconServiceServer).UploadResourceIcon(ctx, req.(*UploadResourceIconRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IconService_ServiceDesc is the grpc.ServiceDesc for IconService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IconService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudstack.management.resource.icon.v1.IconService",
	HandlerType: (*IconServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteResourceIcon",
			Handler:    _IconService_DeleteResourceIcon_Handler,
		},
		{
			MethodName: "ListResourceIcon",
			Handler:    _IconService_ListResourceIcon_Handler,
		},
		{
			MethodName: "UploadResourceIcon",
			Handler:    _IconService_UploadResourceIcon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudstack/management/resource/icon/v1/icon.gen.proto",
}
