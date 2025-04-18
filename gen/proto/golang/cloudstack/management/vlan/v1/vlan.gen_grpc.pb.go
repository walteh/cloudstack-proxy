// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cloudstack/management/vlan/v1/vlan.gen.proto

package vlanv1

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
	VlanService_CreateVlanIpRange_FullMethodName     = "/cloudstack.management.vlan.v1.VlanService/CreateVlanIpRange"
	VlanService_DedicatePublicIpRange_FullMethodName = "/cloudstack.management.vlan.v1.VlanService/DedicatePublicIpRange"
	VlanService_DeleteVlanIpRange_FullMethodName     = "/cloudstack.management.vlan.v1.VlanService/DeleteVlanIpRange"
	VlanService_ListVlanIpRanges_FullMethodName      = "/cloudstack.management.vlan.v1.VlanService/ListVlanIpRanges"
	VlanService_ReleasePublicIpRange_FullMethodName  = "/cloudstack.management.vlan.v1.VlanService/ReleasePublicIpRange"
	VlanService_UpdateVlanIpRange_FullMethodName     = "/cloudstack.management.vlan.v1.VlanService/UpdateVlanIpRange"
)

// VlanServiceClient is the client API for VlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// VlanService provides operations for managing Vlans
type VlanServiceClient interface {
	// CreateVlanIpRange Creates a VLAN IP range.
	CreateVlanIpRange(ctx context.Context, in *CreateVlanIpRangeRequest, opts ...grpc.CallOption) (*CreateVlanIpRangeResponse, error)
	// DedicatePublicIpRange Dedicates a Public IP range to an account
	DedicatePublicIpRange(ctx context.Context, in *DedicatePublicIpRangeRequest, opts ...grpc.CallOption) (*DedicatePublicIpRangeResponse, error)
	// DeleteVlanIpRange Deletes a VLAN IP range.
	DeleteVlanIpRange(ctx context.Context, in *DeleteVlanIpRangeRequest, opts ...grpc.CallOption) (*DeleteVlanIpRangeResponse, error)
	// ListVlanIpRanges Lists all VLAN IP ranges.
	ListVlanIpRanges(ctx context.Context, in *ListVlanIpRangesRequest, opts ...grpc.CallOption) (*ListVlanIpRangesResponse, error)
	// ReleasePublicIpRange Releases a Public IP range back to the system pool
	ReleasePublicIpRange(ctx context.Context, in *ReleasePublicIpRangeRequest, opts ...grpc.CallOption) (*ReleasePublicIpRangeResponse, error)
	// UpdateVlanIpRange Updates a VLAN IP range.
	UpdateVlanIpRange(ctx context.Context, in *UpdateVlanIpRangeRequest, opts ...grpc.CallOption) (*UpdateVlanIpRangeResponse, error)
}

type vlanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVlanServiceClient(cc grpc.ClientConnInterface) VlanServiceClient {
	return &vlanServiceClient{cc}
}

func (c *vlanServiceClient) CreateVlanIpRange(ctx context.Context, in *CreateVlanIpRangeRequest, opts ...grpc.CallOption) (*CreateVlanIpRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVlanIpRangeResponse)
	err := c.cc.Invoke(ctx, VlanService_CreateVlanIpRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vlanServiceClient) DedicatePublicIpRange(ctx context.Context, in *DedicatePublicIpRangeRequest, opts ...grpc.CallOption) (*DedicatePublicIpRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DedicatePublicIpRangeResponse)
	err := c.cc.Invoke(ctx, VlanService_DedicatePublicIpRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vlanServiceClient) DeleteVlanIpRange(ctx context.Context, in *DeleteVlanIpRangeRequest, opts ...grpc.CallOption) (*DeleteVlanIpRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVlanIpRangeResponse)
	err := c.cc.Invoke(ctx, VlanService_DeleteVlanIpRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vlanServiceClient) ListVlanIpRanges(ctx context.Context, in *ListVlanIpRangesRequest, opts ...grpc.CallOption) (*ListVlanIpRangesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVlanIpRangesResponse)
	err := c.cc.Invoke(ctx, VlanService_ListVlanIpRanges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vlanServiceClient) ReleasePublicIpRange(ctx context.Context, in *ReleasePublicIpRangeRequest, opts ...grpc.CallOption) (*ReleasePublicIpRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReleasePublicIpRangeResponse)
	err := c.cc.Invoke(ctx, VlanService_ReleasePublicIpRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vlanServiceClient) UpdateVlanIpRange(ctx context.Context, in *UpdateVlanIpRangeRequest, opts ...grpc.CallOption) (*UpdateVlanIpRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVlanIpRangeResponse)
	err := c.cc.Invoke(ctx, VlanService_UpdateVlanIpRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VlanServiceServer is the server API for VlanService service.
// All implementations must embed UnimplementedVlanServiceServer
// for forward compatibility.
//
// VlanService provides operations for managing Vlans
type VlanServiceServer interface {
	// CreateVlanIpRange Creates a VLAN IP range.
	CreateVlanIpRange(context.Context, *CreateVlanIpRangeRequest) (*CreateVlanIpRangeResponse, error)
	// DedicatePublicIpRange Dedicates a Public IP range to an account
	DedicatePublicIpRange(context.Context, *DedicatePublicIpRangeRequest) (*DedicatePublicIpRangeResponse, error)
	// DeleteVlanIpRange Deletes a VLAN IP range.
	DeleteVlanIpRange(context.Context, *DeleteVlanIpRangeRequest) (*DeleteVlanIpRangeResponse, error)
	// ListVlanIpRanges Lists all VLAN IP ranges.
	ListVlanIpRanges(context.Context, *ListVlanIpRangesRequest) (*ListVlanIpRangesResponse, error)
	// ReleasePublicIpRange Releases a Public IP range back to the system pool
	ReleasePublicIpRange(context.Context, *ReleasePublicIpRangeRequest) (*ReleasePublicIpRangeResponse, error)
	// UpdateVlanIpRange Updates a VLAN IP range.
	UpdateVlanIpRange(context.Context, *UpdateVlanIpRangeRequest) (*UpdateVlanIpRangeResponse, error)
	mustEmbedUnimplementedVlanServiceServer()
}

// UnimplementedVlanServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVlanServiceServer struct{}

func (UnimplementedVlanServiceServer) CreateVlanIpRange(context.Context, *CreateVlanIpRangeRequest) (*CreateVlanIpRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVlanIpRange not implemented")
}
func (UnimplementedVlanServiceServer) DedicatePublicIpRange(context.Context, *DedicatePublicIpRangeRequest) (*DedicatePublicIpRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DedicatePublicIpRange not implemented")
}
func (UnimplementedVlanServiceServer) DeleteVlanIpRange(context.Context, *DeleteVlanIpRangeRequest) (*DeleteVlanIpRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVlanIpRange not implemented")
}
func (UnimplementedVlanServiceServer) ListVlanIpRanges(context.Context, *ListVlanIpRangesRequest) (*ListVlanIpRangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVlanIpRanges not implemented")
}
func (UnimplementedVlanServiceServer) ReleasePublicIpRange(context.Context, *ReleasePublicIpRangeRequest) (*ReleasePublicIpRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleasePublicIpRange not implemented")
}
func (UnimplementedVlanServiceServer) UpdateVlanIpRange(context.Context, *UpdateVlanIpRangeRequest) (*UpdateVlanIpRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVlanIpRange not implemented")
}
func (UnimplementedVlanServiceServer) mustEmbedUnimplementedVlanServiceServer() {}
func (UnimplementedVlanServiceServer) testEmbeddedByValue()                     {}

// UnsafeVlanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VlanServiceServer will
// result in compilation errors.
type UnsafeVlanServiceServer interface {
	mustEmbedUnimplementedVlanServiceServer()
}

func RegisterVlanServiceServer(s grpc.ServiceRegistrar, srv VlanServiceServer) {
	// If the following call pancis, it indicates UnimplementedVlanServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VlanService_ServiceDesc, srv)
}

func _VlanService_CreateVlanIpRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVlanIpRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VlanServiceServer).CreateVlanIpRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VlanService_CreateVlanIpRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VlanServiceServer).CreateVlanIpRange(ctx, req.(*CreateVlanIpRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VlanService_DedicatePublicIpRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DedicatePublicIpRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VlanServiceServer).DedicatePublicIpRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VlanService_DedicatePublicIpRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VlanServiceServer).DedicatePublicIpRange(ctx, req.(*DedicatePublicIpRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VlanService_DeleteVlanIpRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVlanIpRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VlanServiceServer).DeleteVlanIpRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VlanService_DeleteVlanIpRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VlanServiceServer).DeleteVlanIpRange(ctx, req.(*DeleteVlanIpRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VlanService_ListVlanIpRanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVlanIpRangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VlanServiceServer).ListVlanIpRanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VlanService_ListVlanIpRanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VlanServiceServer).ListVlanIpRanges(ctx, req.(*ListVlanIpRangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VlanService_ReleasePublicIpRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasePublicIpRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VlanServiceServer).ReleasePublicIpRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VlanService_ReleasePublicIpRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VlanServiceServer).ReleasePublicIpRange(ctx, req.(*ReleasePublicIpRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VlanService_UpdateVlanIpRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVlanIpRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VlanServiceServer).UpdateVlanIpRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VlanService_UpdateVlanIpRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VlanServiceServer).UpdateVlanIpRange(ctx, req.(*UpdateVlanIpRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VlanService_ServiceDesc is the grpc.ServiceDesc for VlanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VlanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudstack.management.vlan.v1.VlanService",
	HandlerType: (*VlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVlanIpRange",
			Handler:    _VlanService_CreateVlanIpRange_Handler,
		},
		{
			MethodName: "DedicatePublicIpRange",
			Handler:    _VlanService_DedicatePublicIpRange_Handler,
		},
		{
			MethodName: "DeleteVlanIpRange",
			Handler:    _VlanService_DeleteVlanIpRange_Handler,
		},
		{
			MethodName: "ListVlanIpRanges",
			Handler:    _VlanService_ListVlanIpRanges_Handler,
		},
		{
			MethodName: "ReleasePublicIpRange",
			Handler:    _VlanService_ReleasePublicIpRange_Handler,
		},
		{
			MethodName: "UpdateVlanIpRange",
			Handler:    _VlanService_UpdateVlanIpRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudstack/management/vlan/v1/vlan.gen.proto",
}
