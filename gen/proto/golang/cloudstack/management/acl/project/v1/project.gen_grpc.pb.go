// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cloudstack/management/acl/project/v1/project.gen.proto

package projectv1

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
	ProjectService_CreateProjectRole_FullMethodName           = "/cloudstack.management.acl.project.v1.ProjectService/CreateProjectRole"
	ProjectService_CreateProjectRolePermission_FullMethodName = "/cloudstack.management.acl.project.v1.ProjectService/CreateProjectRolePermission"
	ProjectService_DeleteProjectRole_FullMethodName           = "/cloudstack.management.acl.project.v1.ProjectService/DeleteProjectRole"
	ProjectService_DeleteProjectRolePermission_FullMethodName = "/cloudstack.management.acl.project.v1.ProjectService/DeleteProjectRolePermission"
	ProjectService_ListProjectRolePermissions_FullMethodName  = "/cloudstack.management.acl.project.v1.ProjectService/ListProjectRolePermissions"
	ProjectService_ListProjectRoles_FullMethodName            = "/cloudstack.management.acl.project.v1.ProjectService/ListProjectRoles"
	ProjectService_UpdateProjectRole_FullMethodName           = "/cloudstack.management.acl.project.v1.ProjectService/UpdateProjectRole"
	ProjectService_UpdateProjectRolePermission_FullMethodName = "/cloudstack.management.acl.project.v1.ProjectService/UpdateProjectRolePermission"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ProjectService provides operations for managing Acl.Projects
type ProjectServiceClient interface {
	// CreateProjectRole Creates a Project role
	CreateProjectRole(ctx context.Context, in *CreateProjectRoleRequest, opts ...grpc.CallOption) (*CreateProjectRoleResponse, error)
	// CreateProjectRolePermission Adds API permissions to a project role
	CreateProjectRolePermission(ctx context.Context, in *CreateProjectRolePermissionRequest, opts ...grpc.CallOption) (*CreateProjectRolePermissionResponse, error)
	// DeleteProjectRole Delete Project roles in CloudStack
	DeleteProjectRole(ctx context.Context, in *DeleteProjectRoleRequest, opts ...grpc.CallOption) (*DeleteProjectRoleResponse, error)
	// DeleteProjectRolePermission Deletes a project role permission in the project
	DeleteProjectRolePermission(ctx context.Context, in *DeleteProjectRolePermissionRequest, opts ...grpc.CallOption) (*DeleteProjectRolePermissionResponse, error)
	// ListProjectRolePermissions Lists a project's project role permissions
	ListProjectRolePermissions(ctx context.Context, in *ListProjectRolePermissionsRequest, opts ...grpc.CallOption) (*ListProjectRolePermissionsResponse, error)
	// ListProjectRoles Lists Project roles in CloudStack
	ListProjectRoles(ctx context.Context, in *ListProjectRolesRequest, opts ...grpc.CallOption) (*ListProjectRolesResponse, error)
	// UpdateProjectRole Creates a Project role
	UpdateProjectRole(ctx context.Context, in *UpdateProjectRoleRequest, opts ...grpc.CallOption) (*UpdateProjectRoleResponse, error)
	// UpdateProjectRolePermission Updates a project role permission and/or order
	UpdateProjectRolePermission(ctx context.Context, in *UpdateProjectRolePermissionRequest, opts ...grpc.CallOption) (*UpdateProjectRolePermissionResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateProjectRole(ctx context.Context, in *CreateProjectRoleRequest, opts ...grpc.CallOption) (*CreateProjectRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProjectRoleResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProjectRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) CreateProjectRolePermission(ctx context.Context, in *CreateProjectRolePermissionRequest, opts ...grpc.CallOption) (*CreateProjectRolePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProjectRolePermissionResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProjectRolePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProjectRole(ctx context.Context, in *DeleteProjectRoleRequest, opts ...grpc.CallOption) (*DeleteProjectRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProjectRoleResponse)
	err := c.cc.Invoke(ctx, ProjectService_DeleteProjectRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProjectRolePermission(ctx context.Context, in *DeleteProjectRolePermissionRequest, opts ...grpc.CallOption) (*DeleteProjectRolePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProjectRolePermissionResponse)
	err := c.cc.Invoke(ctx, ProjectService_DeleteProjectRolePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ListProjectRolePermissions(ctx context.Context, in *ListProjectRolePermissionsRequest, opts ...grpc.CallOption) (*ListProjectRolePermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProjectRolePermissionsResponse)
	err := c.cc.Invoke(ctx, ProjectService_ListProjectRolePermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ListProjectRoles(ctx context.Context, in *ListProjectRolesRequest, opts ...grpc.CallOption) (*ListProjectRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProjectRolesResponse)
	err := c.cc.Invoke(ctx, ProjectService_ListProjectRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProjectRole(ctx context.Context, in *UpdateProjectRoleRequest, opts ...grpc.CallOption) (*UpdateProjectRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProjectRoleResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProjectRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProjectRolePermission(ctx context.Context, in *UpdateProjectRolePermissionRequest, opts ...grpc.CallOption) (*UpdateProjectRolePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProjectRolePermissionResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProjectRolePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility.
//
// ProjectService provides operations for managing Acl.Projects
type ProjectServiceServer interface {
	// CreateProjectRole Creates a Project role
	CreateProjectRole(context.Context, *CreateProjectRoleRequest) (*CreateProjectRoleResponse, error)
	// CreateProjectRolePermission Adds API permissions to a project role
	CreateProjectRolePermission(context.Context, *CreateProjectRolePermissionRequest) (*CreateProjectRolePermissionResponse, error)
	// DeleteProjectRole Delete Project roles in CloudStack
	DeleteProjectRole(context.Context, *DeleteProjectRoleRequest) (*DeleteProjectRoleResponse, error)
	// DeleteProjectRolePermission Deletes a project role permission in the project
	DeleteProjectRolePermission(context.Context, *DeleteProjectRolePermissionRequest) (*DeleteProjectRolePermissionResponse, error)
	// ListProjectRolePermissions Lists a project's project role permissions
	ListProjectRolePermissions(context.Context, *ListProjectRolePermissionsRequest) (*ListProjectRolePermissionsResponse, error)
	// ListProjectRoles Lists Project roles in CloudStack
	ListProjectRoles(context.Context, *ListProjectRolesRequest) (*ListProjectRolesResponse, error)
	// UpdateProjectRole Creates a Project role
	UpdateProjectRole(context.Context, *UpdateProjectRoleRequest) (*UpdateProjectRoleResponse, error)
	// UpdateProjectRolePermission Updates a project role permission and/or order
	UpdateProjectRolePermission(context.Context, *UpdateProjectRolePermissionRequest) (*UpdateProjectRolePermissionResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProjectServiceServer struct{}

func (UnimplementedProjectServiceServer) CreateProjectRole(context.Context, *CreateProjectRoleRequest) (*CreateProjectRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProjectRole not implemented")
}
func (UnimplementedProjectServiceServer) CreateProjectRolePermission(context.Context, *CreateProjectRolePermissionRequest) (*CreateProjectRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProjectRolePermission not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProjectRole(context.Context, *DeleteProjectRoleRequest) (*DeleteProjectRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectRole not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProjectRolePermission(context.Context, *DeleteProjectRolePermissionRequest) (*DeleteProjectRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectRolePermission not implemented")
}
func (UnimplementedProjectServiceServer) ListProjectRolePermissions(context.Context, *ListProjectRolePermissionsRequest) (*ListProjectRolePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjectRolePermissions not implemented")
}
func (UnimplementedProjectServiceServer) ListProjectRoles(context.Context, *ListProjectRolesRequest) (*ListProjectRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjectRoles not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProjectRole(context.Context, *UpdateProjectRoleRequest) (*UpdateProjectRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectRole not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProjectRolePermission(context.Context, *UpdateProjectRolePermissionRequest) (*UpdateProjectRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectRolePermission not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}
func (UnimplementedProjectServiceServer) testEmbeddedByValue()                        {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	// If the following call pancis, it indicates UnimplementedProjectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateProjectRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProjectRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProjectRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProjectRole(ctx, req.(*CreateProjectRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_CreateProjectRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProjectRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProjectRolePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProjectRolePermission(ctx, req.(*CreateProjectRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProjectRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProjectRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_DeleteProjectRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProjectRole(ctx, req.(*DeleteProjectRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProjectRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProjectRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_DeleteProjectRolePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProjectRolePermission(ctx, req.(*DeleteProjectRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ListProjectRolePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectRolePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ListProjectRolePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_ListProjectRolePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ListProjectRolePermissions(ctx, req.(*ListProjectRolePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ListProjectRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ListProjectRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_ListProjectRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ListProjectRoles(ctx, req.(*ListProjectRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProjectRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProjectRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProjectRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProjectRole(ctx, req.(*UpdateProjectRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProjectRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProjectRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProjectRolePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProjectRolePermission(ctx, req.(*UpdateProjectRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudstack.management.acl.project.v1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProjectRole",
			Handler:    _ProjectService_CreateProjectRole_Handler,
		},
		{
			MethodName: "CreateProjectRolePermission",
			Handler:    _ProjectService_CreateProjectRolePermission_Handler,
		},
		{
			MethodName: "DeleteProjectRole",
			Handler:    _ProjectService_DeleteProjectRole_Handler,
		},
		{
			MethodName: "DeleteProjectRolePermission",
			Handler:    _ProjectService_DeleteProjectRolePermission_Handler,
		},
		{
			MethodName: "ListProjectRolePermissions",
			Handler:    _ProjectService_ListProjectRolePermissions_Handler,
		},
		{
			MethodName: "ListProjectRoles",
			Handler:    _ProjectService_ListProjectRoles_Handler,
		},
		{
			MethodName: "UpdateProjectRole",
			Handler:    _ProjectService_UpdateProjectRole_Handler,
		},
		{
			MethodName: "UpdateProjectRolePermission",
			Handler:    _ProjectService_UpdateProjectRolePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudstack/management/acl/project/v1/project.gen.proto",
}
