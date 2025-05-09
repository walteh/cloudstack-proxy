// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cloudstack/management/bucket/v1/bucket.gen.proto

package bucketv1

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
	BucketService_CreateBucket_FullMethodName = "/cloudstack.management.bucket.v1.BucketService/CreateBucket"
	BucketService_DeleteBucket_FullMethodName = "/cloudstack.management.bucket.v1.BucketService/DeleteBucket"
	BucketService_ListBuckets_FullMethodName  = "/cloudstack.management.bucket.v1.BucketService/ListBuckets"
	BucketService_UpdateBucket_FullMethodName = "/cloudstack.management.bucket.v1.BucketService/UpdateBucket"
)

// BucketServiceClient is the client API for BucketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// BucketService provides operations for managing Buckets
type BucketServiceClient interface {
	// CreateBucket Creates a bucket in the specified object storage pool.
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error)
	// DeleteBucket Deletes an empty Bucket.
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error)
	// ListBuckets Lists all Buckets.
	ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error)
	// UpdateBucket Updates Bucket properties
	UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*UpdateBucketResponse, error)
}

type bucketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBucketServiceClient(cc grpc.ClientConnInterface) BucketServiceClient {
	return &bucketServiceClient{cc}
}

func (c *bucketServiceClient) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBucketResponse)
	err := c.cc.Invoke(ctx, BucketService_CreateBucket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBucketResponse)
	err := c.cc.Invoke(ctx, BucketService_DeleteBucket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, BucketService_ListBuckets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*UpdateBucketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBucketResponse)
	err := c.cc.Invoke(ctx, BucketService_UpdateBucket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BucketServiceServer is the server API for BucketService service.
// All implementations must embed UnimplementedBucketServiceServer
// for forward compatibility.
//
// BucketService provides operations for managing Buckets
type BucketServiceServer interface {
	// CreateBucket Creates a bucket in the specified object storage pool.
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error)
	// DeleteBucket Deletes an empty Bucket.
	DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error)
	// ListBuckets Lists all Buckets.
	ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error)
	// UpdateBucket Updates Bucket properties
	UpdateBucket(context.Context, *UpdateBucketRequest) (*UpdateBucketResponse, error)
	mustEmbedUnimplementedBucketServiceServer()
}

// UnimplementedBucketServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBucketServiceServer struct{}

func (UnimplementedBucketServiceServer) CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (UnimplementedBucketServiceServer) DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (UnimplementedBucketServiceServer) ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (UnimplementedBucketServiceServer) UpdateBucket(context.Context, *UpdateBucketRequest) (*UpdateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBucket not implemented")
}
func (UnimplementedBucketServiceServer) mustEmbedUnimplementedBucketServiceServer() {}
func (UnimplementedBucketServiceServer) testEmbeddedByValue()                       {}

// UnsafeBucketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BucketServiceServer will
// result in compilation errors.
type UnsafeBucketServiceServer interface {
	mustEmbedUnimplementedBucketServiceServer()
}

func RegisterBucketServiceServer(s grpc.ServiceRegistrar, srv BucketServiceServer) {
	// If the following call pancis, it indicates UnimplementedBucketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BucketService_ServiceDesc, srv)
}

func _BucketService_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_CreateBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_DeleteBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).DeleteBucket(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_ListBuckets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).ListBuckets(ctx, req.(*ListBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_UpdateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).UpdateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_UpdateBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).UpdateBucket(ctx, req.(*UpdateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BucketService_ServiceDesc is the grpc.ServiceDesc for BucketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BucketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudstack.management.bucket.v1.BucketService",
	HandlerType: (*BucketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _BucketService_CreateBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _BucketService_DeleteBucket_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _BucketService_ListBuckets_Handler,
		},
		{
			MethodName: "UpdateBucket",
			Handler:    _BucketService_UpdateBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudstack/management/bucket/v1/bucket.gen.proto",
}
