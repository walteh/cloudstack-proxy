// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cloudstack/management/guest/v1/guest.gen.proto

package guestv1

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
	GuestService_AddGuestOs_FullMethodName                = "/cloudstack.management.guest.v1.GuestService/AddGuestOs"
	GuestService_AddGuestOsMapping_FullMethodName         = "/cloudstack.management.guest.v1.GuestService/AddGuestOsMapping"
	GuestService_GetHypervisorGuestOsNames_FullMethodName = "/cloudstack.management.guest.v1.GuestService/GetHypervisorGuestOsNames"
	GuestService_ListGuestOs_FullMethodName               = "/cloudstack.management.guest.v1.GuestService/ListGuestOs"
	GuestService_ListGuestOsCategories_FullMethodName     = "/cloudstack.management.guest.v1.GuestService/ListGuestOsCategories"
	GuestService_ListGuestOsMapping_FullMethodName        = "/cloudstack.management.guest.v1.GuestService/ListGuestOsMapping"
	GuestService_RemoveGuestOs_FullMethodName             = "/cloudstack.management.guest.v1.GuestService/RemoveGuestOs"
	GuestService_RemoveGuestOsMapping_FullMethodName      = "/cloudstack.management.guest.v1.GuestService/RemoveGuestOsMapping"
	GuestService_UpdateGuestOs_FullMethodName             = "/cloudstack.management.guest.v1.GuestService/UpdateGuestOs"
	GuestService_UpdateGuestOsMapping_FullMethodName      = "/cloudstack.management.guest.v1.GuestService/UpdateGuestOsMapping"
)

// GuestServiceClient is the client API for GuestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// GuestService provides operations for managing Guests
type GuestServiceClient interface {
	// AddGuestOs Add a new guest OS type
	AddGuestOs(ctx context.Context, in *AddGuestOsRequest, opts ...grpc.CallOption) (*AddGuestOsResponse, error)
	// AddGuestOsMapping Adds a guest OS name to hypervisor OS name mapping
	AddGuestOsMapping(ctx context.Context, in *AddGuestOsMappingRequest, opts ...grpc.CallOption) (*AddGuestOsMappingResponse, error)
	// GetHypervisorGuestOsNames Gets the guest OS names in the hypervisor
	GetHypervisorGuestOsNames(ctx context.Context, in *GetHypervisorGuestOsNamesRequest, opts ...grpc.CallOption) (*GetHypervisorGuestOsNamesResponse, error)
	// ListGuestOs Lists all supported OS types for this cloud.
	ListGuestOs(ctx context.Context, in *ListGuestOsRequest, opts ...grpc.CallOption) (*ListGuestOsResponse, error)
	// ListGuestOsCategories Lists all supported OS categories for this cloud.
	ListGuestOsCategories(ctx context.Context, in *ListGuestOsCategoriesRequest, opts ...grpc.CallOption) (*ListGuestOsCategoriesResponse, error)
	// ListGuestOsMapping Lists all available OS mappings for given hypervisor
	ListGuestOsMapping(ctx context.Context, in *ListGuestOsMappingRequest, opts ...grpc.CallOption) (*ListGuestOsMappingResponse, error)
	// RemoveGuestOs Removes a Guest OS from listing.
	RemoveGuestOs(ctx context.Context, in *RemoveGuestOsRequest, opts ...grpc.CallOption) (*RemoveGuestOsResponse, error)
	// RemoveGuestOsMapping Removes a Guest OS Mapping.
	RemoveGuestOsMapping(ctx context.Context, in *RemoveGuestOsMappingRequest, opts ...grpc.CallOption) (*RemoveGuestOsMappingResponse, error)
	// UpdateGuestOs Updates the information about Guest OS
	UpdateGuestOs(ctx context.Context, in *UpdateGuestOsRequest, opts ...grpc.CallOption) (*UpdateGuestOsResponse, error)
	// UpdateGuestOsMapping Updates the information about Guest OS to Hypervisor specific name mapping
	UpdateGuestOsMapping(ctx context.Context, in *UpdateGuestOsMappingRequest, opts ...grpc.CallOption) (*UpdateGuestOsMappingResponse, error)
}

type guestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuestServiceClient(cc grpc.ClientConnInterface) GuestServiceClient {
	return &guestServiceClient{cc}
}

func (c *guestServiceClient) AddGuestOs(ctx context.Context, in *AddGuestOsRequest, opts ...grpc.CallOption) (*AddGuestOsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddGuestOsResponse)
	err := c.cc.Invoke(ctx, GuestService_AddGuestOs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) AddGuestOsMapping(ctx context.Context, in *AddGuestOsMappingRequest, opts ...grpc.CallOption) (*AddGuestOsMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddGuestOsMappingResponse)
	err := c.cc.Invoke(ctx, GuestService_AddGuestOsMapping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GetHypervisorGuestOsNames(ctx context.Context, in *GetHypervisorGuestOsNamesRequest, opts ...grpc.CallOption) (*GetHypervisorGuestOsNamesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHypervisorGuestOsNamesResponse)
	err := c.cc.Invoke(ctx, GuestService_GetHypervisorGuestOsNames_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) ListGuestOs(ctx context.Context, in *ListGuestOsRequest, opts ...grpc.CallOption) (*ListGuestOsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGuestOsResponse)
	err := c.cc.Invoke(ctx, GuestService_ListGuestOs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) ListGuestOsCategories(ctx context.Context, in *ListGuestOsCategoriesRequest, opts ...grpc.CallOption) (*ListGuestOsCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGuestOsCategoriesResponse)
	err := c.cc.Invoke(ctx, GuestService_ListGuestOsCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) ListGuestOsMapping(ctx context.Context, in *ListGuestOsMappingRequest, opts ...grpc.CallOption) (*ListGuestOsMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGuestOsMappingResponse)
	err := c.cc.Invoke(ctx, GuestService_ListGuestOsMapping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) RemoveGuestOs(ctx context.Context, in *RemoveGuestOsRequest, opts ...grpc.CallOption) (*RemoveGuestOsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveGuestOsResponse)
	err := c.cc.Invoke(ctx, GuestService_RemoveGuestOs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) RemoveGuestOsMapping(ctx context.Context, in *RemoveGuestOsMappingRequest, opts ...grpc.CallOption) (*RemoveGuestOsMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveGuestOsMappingResponse)
	err := c.cc.Invoke(ctx, GuestService_RemoveGuestOsMapping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) UpdateGuestOs(ctx context.Context, in *UpdateGuestOsRequest, opts ...grpc.CallOption) (*UpdateGuestOsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGuestOsResponse)
	err := c.cc.Invoke(ctx, GuestService_UpdateGuestOs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) UpdateGuestOsMapping(ctx context.Context, in *UpdateGuestOsMappingRequest, opts ...grpc.CallOption) (*UpdateGuestOsMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGuestOsMappingResponse)
	err := c.cc.Invoke(ctx, GuestService_UpdateGuestOsMapping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuestServiceServer is the server API for GuestService service.
// All implementations must embed UnimplementedGuestServiceServer
// for forward compatibility.
//
// GuestService provides operations for managing Guests
type GuestServiceServer interface {
	// AddGuestOs Add a new guest OS type
	AddGuestOs(context.Context, *AddGuestOsRequest) (*AddGuestOsResponse, error)
	// AddGuestOsMapping Adds a guest OS name to hypervisor OS name mapping
	AddGuestOsMapping(context.Context, *AddGuestOsMappingRequest) (*AddGuestOsMappingResponse, error)
	// GetHypervisorGuestOsNames Gets the guest OS names in the hypervisor
	GetHypervisorGuestOsNames(context.Context, *GetHypervisorGuestOsNamesRequest) (*GetHypervisorGuestOsNamesResponse, error)
	// ListGuestOs Lists all supported OS types for this cloud.
	ListGuestOs(context.Context, *ListGuestOsRequest) (*ListGuestOsResponse, error)
	// ListGuestOsCategories Lists all supported OS categories for this cloud.
	ListGuestOsCategories(context.Context, *ListGuestOsCategoriesRequest) (*ListGuestOsCategoriesResponse, error)
	// ListGuestOsMapping Lists all available OS mappings for given hypervisor
	ListGuestOsMapping(context.Context, *ListGuestOsMappingRequest) (*ListGuestOsMappingResponse, error)
	// RemoveGuestOs Removes a Guest OS from listing.
	RemoveGuestOs(context.Context, *RemoveGuestOsRequest) (*RemoveGuestOsResponse, error)
	// RemoveGuestOsMapping Removes a Guest OS Mapping.
	RemoveGuestOsMapping(context.Context, *RemoveGuestOsMappingRequest) (*RemoveGuestOsMappingResponse, error)
	// UpdateGuestOs Updates the information about Guest OS
	UpdateGuestOs(context.Context, *UpdateGuestOsRequest) (*UpdateGuestOsResponse, error)
	// UpdateGuestOsMapping Updates the information about Guest OS to Hypervisor specific name mapping
	UpdateGuestOsMapping(context.Context, *UpdateGuestOsMappingRequest) (*UpdateGuestOsMappingResponse, error)
	mustEmbedUnimplementedGuestServiceServer()
}

// UnimplementedGuestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGuestServiceServer struct{}

func (UnimplementedGuestServiceServer) AddGuestOs(context.Context, *AddGuestOsRequest) (*AddGuestOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGuestOs not implemented")
}
func (UnimplementedGuestServiceServer) AddGuestOsMapping(context.Context, *AddGuestOsMappingRequest) (*AddGuestOsMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGuestOsMapping not implemented")
}
func (UnimplementedGuestServiceServer) GetHypervisorGuestOsNames(context.Context, *GetHypervisorGuestOsNamesRequest) (*GetHypervisorGuestOsNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHypervisorGuestOsNames not implemented")
}
func (UnimplementedGuestServiceServer) ListGuestOs(context.Context, *ListGuestOsRequest) (*ListGuestOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuestOs not implemented")
}
func (UnimplementedGuestServiceServer) ListGuestOsCategories(context.Context, *ListGuestOsCategoriesRequest) (*ListGuestOsCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuestOsCategories not implemented")
}
func (UnimplementedGuestServiceServer) ListGuestOsMapping(context.Context, *ListGuestOsMappingRequest) (*ListGuestOsMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGuestOsMapping not implemented")
}
func (UnimplementedGuestServiceServer) RemoveGuestOs(context.Context, *RemoveGuestOsRequest) (*RemoveGuestOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGuestOs not implemented")
}
func (UnimplementedGuestServiceServer) RemoveGuestOsMapping(context.Context, *RemoveGuestOsMappingRequest) (*RemoveGuestOsMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGuestOsMapping not implemented")
}
func (UnimplementedGuestServiceServer) UpdateGuestOs(context.Context, *UpdateGuestOsRequest) (*UpdateGuestOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuestOs not implemented")
}
func (UnimplementedGuestServiceServer) UpdateGuestOsMapping(context.Context, *UpdateGuestOsMappingRequest) (*UpdateGuestOsMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuestOsMapping not implemented")
}
func (UnimplementedGuestServiceServer) mustEmbedUnimplementedGuestServiceServer() {}
func (UnimplementedGuestServiceServer) testEmbeddedByValue()                      {}

// UnsafeGuestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuestServiceServer will
// result in compilation errors.
type UnsafeGuestServiceServer interface {
	mustEmbedUnimplementedGuestServiceServer()
}

func RegisterGuestServiceServer(s grpc.ServiceRegistrar, srv GuestServiceServer) {
	// If the following call pancis, it indicates UnimplementedGuestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GuestService_ServiceDesc, srv)
}

func _GuestService_AddGuestOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGuestOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).AddGuestOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_AddGuestOs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).AddGuestOs(ctx, req.(*AddGuestOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_AddGuestOsMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGuestOsMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).AddGuestOsMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_AddGuestOsMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).AddGuestOsMapping(ctx, req.(*AddGuestOsMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GetHypervisorGuestOsNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHypervisorGuestOsNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GetHypervisorGuestOsNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GetHypervisorGuestOsNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GetHypervisorGuestOsNames(ctx, req.(*GetHypervisorGuestOsNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_ListGuestOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGuestOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).ListGuestOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_ListGuestOs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).ListGuestOs(ctx, req.(*ListGuestOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_ListGuestOsCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGuestOsCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).ListGuestOsCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_ListGuestOsCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).ListGuestOsCategories(ctx, req.(*ListGuestOsCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_ListGuestOsMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGuestOsMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).ListGuestOsMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_ListGuestOsMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).ListGuestOsMapping(ctx, req.(*ListGuestOsMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_RemoveGuestOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGuestOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).RemoveGuestOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_RemoveGuestOs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).RemoveGuestOs(ctx, req.(*RemoveGuestOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_RemoveGuestOsMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGuestOsMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).RemoveGuestOsMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_RemoveGuestOsMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).RemoveGuestOsMapping(ctx, req.(*RemoveGuestOsMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_UpdateGuestOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGuestOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).UpdateGuestOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_UpdateGuestOs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).UpdateGuestOs(ctx, req.(*UpdateGuestOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_UpdateGuestOsMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGuestOsMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).UpdateGuestOsMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_UpdateGuestOsMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).UpdateGuestOsMapping(ctx, req.(*UpdateGuestOsMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GuestService_ServiceDesc is the grpc.ServiceDesc for GuestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudstack.management.guest.v1.GuestService",
	HandlerType: (*GuestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGuestOs",
			Handler:    _GuestService_AddGuestOs_Handler,
		},
		{
			MethodName: "AddGuestOsMapping",
			Handler:    _GuestService_AddGuestOsMapping_Handler,
		},
		{
			MethodName: "GetHypervisorGuestOsNames",
			Handler:    _GuestService_GetHypervisorGuestOsNames_Handler,
		},
		{
			MethodName: "ListGuestOs",
			Handler:    _GuestService_ListGuestOs_Handler,
		},
		{
			MethodName: "ListGuestOsCategories",
			Handler:    _GuestService_ListGuestOsCategories_Handler,
		},
		{
			MethodName: "ListGuestOsMapping",
			Handler:    _GuestService_ListGuestOsMapping_Handler,
		},
		{
			MethodName: "RemoveGuestOs",
			Handler:    _GuestService_RemoveGuestOs_Handler,
		},
		{
			MethodName: "RemoveGuestOsMapping",
			Handler:    _GuestService_RemoveGuestOsMapping_Handler,
		},
		{
			MethodName: "UpdateGuestOs",
			Handler:    _GuestService_UpdateGuestOs_Handler,
		},
		{
			MethodName: "UpdateGuestOsMapping",
			Handler:    _GuestService_UpdateGuestOsMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudstack/management/guest/v1/guest.gen.proto",
}
