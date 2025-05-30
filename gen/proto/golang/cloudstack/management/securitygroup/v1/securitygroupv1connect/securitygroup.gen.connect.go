// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/management/securitygroup/v1/securitygroup.gen.proto

package securitygroupv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/securitygroup/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SecuritygroupServiceName is the fully-qualified name of the SecuritygroupService service.
	SecuritygroupServiceName = "cloudstack.management.securitygroup.v1.SecuritygroupService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SecuritygroupServiceAuthorizeSecurityGroupEgressProcedure is the fully-qualified name of the
	// SecuritygroupService's AuthorizeSecurityGroupEgress RPC.
	SecuritygroupServiceAuthorizeSecurityGroupEgressProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/AuthorizeSecurityGroupEgress"
	// SecuritygroupServiceAuthorizeSecurityGroupIngressProcedure is the fully-qualified name of the
	// SecuritygroupService's AuthorizeSecurityGroupIngress RPC.
	SecuritygroupServiceAuthorizeSecurityGroupIngressProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/AuthorizeSecurityGroupIngress"
	// SecuritygroupServiceCreateSecurityGroupProcedure is the fully-qualified name of the
	// SecuritygroupService's CreateSecurityGroup RPC.
	SecuritygroupServiceCreateSecurityGroupProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/CreateSecurityGroup"
	// SecuritygroupServiceDeleteSecurityGroupProcedure is the fully-qualified name of the
	// SecuritygroupService's DeleteSecurityGroup RPC.
	SecuritygroupServiceDeleteSecurityGroupProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/DeleteSecurityGroup"
	// SecuritygroupServiceListSecurityGroupsProcedure is the fully-qualified name of the
	// SecuritygroupService's ListSecurityGroups RPC.
	SecuritygroupServiceListSecurityGroupsProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/ListSecurityGroups"
	// SecuritygroupServiceRevokeSecurityGroupEgressProcedure is the fully-qualified name of the
	// SecuritygroupService's RevokeSecurityGroupEgress RPC.
	SecuritygroupServiceRevokeSecurityGroupEgressProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/RevokeSecurityGroupEgress"
	// SecuritygroupServiceRevokeSecurityGroupIngressProcedure is the fully-qualified name of the
	// SecuritygroupService's RevokeSecurityGroupIngress RPC.
	SecuritygroupServiceRevokeSecurityGroupIngressProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/RevokeSecurityGroupIngress"
	// SecuritygroupServiceUpdateSecurityGroupProcedure is the fully-qualified name of the
	// SecuritygroupService's UpdateSecurityGroup RPC.
	SecuritygroupServiceUpdateSecurityGroupProcedure = "/cloudstack.management.securitygroup.v1.SecuritygroupService/UpdateSecurityGroup"
)

// SecuritygroupServiceClient is a client for the
// cloudstack.management.securitygroup.v1.SecuritygroupService service.
type SecuritygroupServiceClient interface {
	// AuthorizeSecurityGroupEgress Authorizes a particular egress rule for this security group
	AuthorizeSecurityGroupEgress(context.Context, *connect.Request[v1.AuthorizeSecurityGroupEgressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupEgressResponse], error)
	// AuthorizeSecurityGroupIngress Authorizes a particular ingress rule for this security group
	AuthorizeSecurityGroupIngress(context.Context, *connect.Request[v1.AuthorizeSecurityGroupIngressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupIngressResponse], error)
	// CreateSecurityGroup Creates a security group
	CreateSecurityGroup(context.Context, *connect.Request[v1.CreateSecurityGroupRequest]) (*connect.Response[v1.CreateSecurityGroupResponse], error)
	// DeleteSecurityGroup Deletes security group
	DeleteSecurityGroup(context.Context, *connect.Request[v1.DeleteSecurityGroupRequest]) (*connect.Response[v1.DeleteSecurityGroupResponse], error)
	// ListSecurityGroups Lists security groups
	ListSecurityGroups(context.Context, *connect.Request[v1.ListSecurityGroupsRequest]) (*connect.Response[v1.ListSecurityGroupsResponse], error)
	// RevokeSecurityGroupEgress Deletes a particular egress rule from this security group
	RevokeSecurityGroupEgress(context.Context, *connect.Request[v1.RevokeSecurityGroupEgressRequest]) (*connect.Response[v1.RevokeSecurityGroupEgressResponse], error)
	// RevokeSecurityGroupIngress Deletes a particular ingress rule from this security group
	RevokeSecurityGroupIngress(context.Context, *connect.Request[v1.RevokeSecurityGroupIngressRequest]) (*connect.Response[v1.RevokeSecurityGroupIngressResponse], error)
	// UpdateSecurityGroup Updates a security group
	UpdateSecurityGroup(context.Context, *connect.Request[v1.UpdateSecurityGroupRequest]) (*connect.Response[v1.UpdateSecurityGroupResponse], error)
}

// NewSecuritygroupServiceClient constructs a client for the
// cloudstack.management.securitygroup.v1.SecuritygroupService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSecuritygroupServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SecuritygroupServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	securitygroupServiceMethods := v1.File_cloudstack_management_securitygroup_v1_securitygroup_gen_proto.Services().ByName("SecuritygroupService").Methods()
	return &securitygroupServiceClient{
		authorizeSecurityGroupEgress: connect.NewClient[v1.AuthorizeSecurityGroupEgressRequest, v1.AuthorizeSecurityGroupEgressResponse](
			httpClient,
			baseURL+SecuritygroupServiceAuthorizeSecurityGroupEgressProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("AuthorizeSecurityGroupEgress")),
			connect.WithClientOptions(opts...),
		),
		authorizeSecurityGroupIngress: connect.NewClient[v1.AuthorizeSecurityGroupIngressRequest, v1.AuthorizeSecurityGroupIngressResponse](
			httpClient,
			baseURL+SecuritygroupServiceAuthorizeSecurityGroupIngressProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("AuthorizeSecurityGroupIngress")),
			connect.WithClientOptions(opts...),
		),
		createSecurityGroup: connect.NewClient[v1.CreateSecurityGroupRequest, v1.CreateSecurityGroupResponse](
			httpClient,
			baseURL+SecuritygroupServiceCreateSecurityGroupProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("CreateSecurityGroup")),
			connect.WithClientOptions(opts...),
		),
		deleteSecurityGroup: connect.NewClient[v1.DeleteSecurityGroupRequest, v1.DeleteSecurityGroupResponse](
			httpClient,
			baseURL+SecuritygroupServiceDeleteSecurityGroupProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("DeleteSecurityGroup")),
			connect.WithClientOptions(opts...),
		),
		listSecurityGroups: connect.NewClient[v1.ListSecurityGroupsRequest, v1.ListSecurityGroupsResponse](
			httpClient,
			baseURL+SecuritygroupServiceListSecurityGroupsProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("ListSecurityGroups")),
			connect.WithClientOptions(opts...),
		),
		revokeSecurityGroupEgress: connect.NewClient[v1.RevokeSecurityGroupEgressRequest, v1.RevokeSecurityGroupEgressResponse](
			httpClient,
			baseURL+SecuritygroupServiceRevokeSecurityGroupEgressProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("RevokeSecurityGroupEgress")),
			connect.WithClientOptions(opts...),
		),
		revokeSecurityGroupIngress: connect.NewClient[v1.RevokeSecurityGroupIngressRequest, v1.RevokeSecurityGroupIngressResponse](
			httpClient,
			baseURL+SecuritygroupServiceRevokeSecurityGroupIngressProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("RevokeSecurityGroupIngress")),
			connect.WithClientOptions(opts...),
		),
		updateSecurityGroup: connect.NewClient[v1.UpdateSecurityGroupRequest, v1.UpdateSecurityGroupResponse](
			httpClient,
			baseURL+SecuritygroupServiceUpdateSecurityGroupProcedure,
			connect.WithSchema(securitygroupServiceMethods.ByName("UpdateSecurityGroup")),
			connect.WithClientOptions(opts...),
		),
	}
}

// securitygroupServiceClient implements SecuritygroupServiceClient.
type securitygroupServiceClient struct {
	authorizeSecurityGroupEgress  *connect.Client[v1.AuthorizeSecurityGroupEgressRequest, v1.AuthorizeSecurityGroupEgressResponse]
	authorizeSecurityGroupIngress *connect.Client[v1.AuthorizeSecurityGroupIngressRequest, v1.AuthorizeSecurityGroupIngressResponse]
	createSecurityGroup           *connect.Client[v1.CreateSecurityGroupRequest, v1.CreateSecurityGroupResponse]
	deleteSecurityGroup           *connect.Client[v1.DeleteSecurityGroupRequest, v1.DeleteSecurityGroupResponse]
	listSecurityGroups            *connect.Client[v1.ListSecurityGroupsRequest, v1.ListSecurityGroupsResponse]
	revokeSecurityGroupEgress     *connect.Client[v1.RevokeSecurityGroupEgressRequest, v1.RevokeSecurityGroupEgressResponse]
	revokeSecurityGroupIngress    *connect.Client[v1.RevokeSecurityGroupIngressRequest, v1.RevokeSecurityGroupIngressResponse]
	updateSecurityGroup           *connect.Client[v1.UpdateSecurityGroupRequest, v1.UpdateSecurityGroupResponse]
}

// AuthorizeSecurityGroupEgress calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.AuthorizeSecurityGroupEgress.
func (c *securitygroupServiceClient) AuthorizeSecurityGroupEgress(ctx context.Context, req *connect.Request[v1.AuthorizeSecurityGroupEgressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupEgressResponse], error) {
	return c.authorizeSecurityGroupEgress.CallUnary(ctx, req)
}

// AuthorizeSecurityGroupIngress calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.AuthorizeSecurityGroupIngress.
func (c *securitygroupServiceClient) AuthorizeSecurityGroupIngress(ctx context.Context, req *connect.Request[v1.AuthorizeSecurityGroupIngressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupIngressResponse], error) {
	return c.authorizeSecurityGroupIngress.CallUnary(ctx, req)
}

// CreateSecurityGroup calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.CreateSecurityGroup.
func (c *securitygroupServiceClient) CreateSecurityGroup(ctx context.Context, req *connect.Request[v1.CreateSecurityGroupRequest]) (*connect.Response[v1.CreateSecurityGroupResponse], error) {
	return c.createSecurityGroup.CallUnary(ctx, req)
}

// DeleteSecurityGroup calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.DeleteSecurityGroup.
func (c *securitygroupServiceClient) DeleteSecurityGroup(ctx context.Context, req *connect.Request[v1.DeleteSecurityGroupRequest]) (*connect.Response[v1.DeleteSecurityGroupResponse], error) {
	return c.deleteSecurityGroup.CallUnary(ctx, req)
}

// ListSecurityGroups calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.ListSecurityGroups.
func (c *securitygroupServiceClient) ListSecurityGroups(ctx context.Context, req *connect.Request[v1.ListSecurityGroupsRequest]) (*connect.Response[v1.ListSecurityGroupsResponse], error) {
	return c.listSecurityGroups.CallUnary(ctx, req)
}

// RevokeSecurityGroupEgress calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.RevokeSecurityGroupEgress.
func (c *securitygroupServiceClient) RevokeSecurityGroupEgress(ctx context.Context, req *connect.Request[v1.RevokeSecurityGroupEgressRequest]) (*connect.Response[v1.RevokeSecurityGroupEgressResponse], error) {
	return c.revokeSecurityGroupEgress.CallUnary(ctx, req)
}

// RevokeSecurityGroupIngress calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.RevokeSecurityGroupIngress.
func (c *securitygroupServiceClient) RevokeSecurityGroupIngress(ctx context.Context, req *connect.Request[v1.RevokeSecurityGroupIngressRequest]) (*connect.Response[v1.RevokeSecurityGroupIngressResponse], error) {
	return c.revokeSecurityGroupIngress.CallUnary(ctx, req)
}

// UpdateSecurityGroup calls
// cloudstack.management.securitygroup.v1.SecuritygroupService.UpdateSecurityGroup.
func (c *securitygroupServiceClient) UpdateSecurityGroup(ctx context.Context, req *connect.Request[v1.UpdateSecurityGroupRequest]) (*connect.Response[v1.UpdateSecurityGroupResponse], error) {
	return c.updateSecurityGroup.CallUnary(ctx, req)
}

// SecuritygroupServiceHandler is an implementation of the
// cloudstack.management.securitygroup.v1.SecuritygroupService service.
type SecuritygroupServiceHandler interface {
	// AuthorizeSecurityGroupEgress Authorizes a particular egress rule for this security group
	AuthorizeSecurityGroupEgress(context.Context, *connect.Request[v1.AuthorizeSecurityGroupEgressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupEgressResponse], error)
	// AuthorizeSecurityGroupIngress Authorizes a particular ingress rule for this security group
	AuthorizeSecurityGroupIngress(context.Context, *connect.Request[v1.AuthorizeSecurityGroupIngressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupIngressResponse], error)
	// CreateSecurityGroup Creates a security group
	CreateSecurityGroup(context.Context, *connect.Request[v1.CreateSecurityGroupRequest]) (*connect.Response[v1.CreateSecurityGroupResponse], error)
	// DeleteSecurityGroup Deletes security group
	DeleteSecurityGroup(context.Context, *connect.Request[v1.DeleteSecurityGroupRequest]) (*connect.Response[v1.DeleteSecurityGroupResponse], error)
	// ListSecurityGroups Lists security groups
	ListSecurityGroups(context.Context, *connect.Request[v1.ListSecurityGroupsRequest]) (*connect.Response[v1.ListSecurityGroupsResponse], error)
	// RevokeSecurityGroupEgress Deletes a particular egress rule from this security group
	RevokeSecurityGroupEgress(context.Context, *connect.Request[v1.RevokeSecurityGroupEgressRequest]) (*connect.Response[v1.RevokeSecurityGroupEgressResponse], error)
	// RevokeSecurityGroupIngress Deletes a particular ingress rule from this security group
	RevokeSecurityGroupIngress(context.Context, *connect.Request[v1.RevokeSecurityGroupIngressRequest]) (*connect.Response[v1.RevokeSecurityGroupIngressResponse], error)
	// UpdateSecurityGroup Updates a security group
	UpdateSecurityGroup(context.Context, *connect.Request[v1.UpdateSecurityGroupRequest]) (*connect.Response[v1.UpdateSecurityGroupResponse], error)
}

// NewSecuritygroupServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSecuritygroupServiceHandler(svc SecuritygroupServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	securitygroupServiceMethods := v1.File_cloudstack_management_securitygroup_v1_securitygroup_gen_proto.Services().ByName("SecuritygroupService").Methods()
	securitygroupServiceAuthorizeSecurityGroupEgressHandler := connect.NewUnaryHandler(
		SecuritygroupServiceAuthorizeSecurityGroupEgressProcedure,
		svc.AuthorizeSecurityGroupEgress,
		connect.WithSchema(securitygroupServiceMethods.ByName("AuthorizeSecurityGroupEgress")),
		connect.WithHandlerOptions(opts...),
	)
	securitygroupServiceAuthorizeSecurityGroupIngressHandler := connect.NewUnaryHandler(
		SecuritygroupServiceAuthorizeSecurityGroupIngressProcedure,
		svc.AuthorizeSecurityGroupIngress,
		connect.WithSchema(securitygroupServiceMethods.ByName("AuthorizeSecurityGroupIngress")),
		connect.WithHandlerOptions(opts...),
	)
	securitygroupServiceCreateSecurityGroupHandler := connect.NewUnaryHandler(
		SecuritygroupServiceCreateSecurityGroupProcedure,
		svc.CreateSecurityGroup,
		connect.WithSchema(securitygroupServiceMethods.ByName("CreateSecurityGroup")),
		connect.WithHandlerOptions(opts...),
	)
	securitygroupServiceDeleteSecurityGroupHandler := connect.NewUnaryHandler(
		SecuritygroupServiceDeleteSecurityGroupProcedure,
		svc.DeleteSecurityGroup,
		connect.WithSchema(securitygroupServiceMethods.ByName("DeleteSecurityGroup")),
		connect.WithHandlerOptions(opts...),
	)
	securitygroupServiceListSecurityGroupsHandler := connect.NewUnaryHandler(
		SecuritygroupServiceListSecurityGroupsProcedure,
		svc.ListSecurityGroups,
		connect.WithSchema(securitygroupServiceMethods.ByName("ListSecurityGroups")),
		connect.WithHandlerOptions(opts...),
	)
	securitygroupServiceRevokeSecurityGroupEgressHandler := connect.NewUnaryHandler(
		SecuritygroupServiceRevokeSecurityGroupEgressProcedure,
		svc.RevokeSecurityGroupEgress,
		connect.WithSchema(securitygroupServiceMethods.ByName("RevokeSecurityGroupEgress")),
		connect.WithHandlerOptions(opts...),
	)
	securitygroupServiceRevokeSecurityGroupIngressHandler := connect.NewUnaryHandler(
		SecuritygroupServiceRevokeSecurityGroupIngressProcedure,
		svc.RevokeSecurityGroupIngress,
		connect.WithSchema(securitygroupServiceMethods.ByName("RevokeSecurityGroupIngress")),
		connect.WithHandlerOptions(opts...),
	)
	securitygroupServiceUpdateSecurityGroupHandler := connect.NewUnaryHandler(
		SecuritygroupServiceUpdateSecurityGroupProcedure,
		svc.UpdateSecurityGroup,
		connect.WithSchema(securitygroupServiceMethods.ByName("UpdateSecurityGroup")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.management.securitygroup.v1.SecuritygroupService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SecuritygroupServiceAuthorizeSecurityGroupEgressProcedure:
			securitygroupServiceAuthorizeSecurityGroupEgressHandler.ServeHTTP(w, r)
		case SecuritygroupServiceAuthorizeSecurityGroupIngressProcedure:
			securitygroupServiceAuthorizeSecurityGroupIngressHandler.ServeHTTP(w, r)
		case SecuritygroupServiceCreateSecurityGroupProcedure:
			securitygroupServiceCreateSecurityGroupHandler.ServeHTTP(w, r)
		case SecuritygroupServiceDeleteSecurityGroupProcedure:
			securitygroupServiceDeleteSecurityGroupHandler.ServeHTTP(w, r)
		case SecuritygroupServiceListSecurityGroupsProcedure:
			securitygroupServiceListSecurityGroupsHandler.ServeHTTP(w, r)
		case SecuritygroupServiceRevokeSecurityGroupEgressProcedure:
			securitygroupServiceRevokeSecurityGroupEgressHandler.ServeHTTP(w, r)
		case SecuritygroupServiceRevokeSecurityGroupIngressProcedure:
			securitygroupServiceRevokeSecurityGroupIngressHandler.ServeHTTP(w, r)
		case SecuritygroupServiceUpdateSecurityGroupProcedure:
			securitygroupServiceUpdateSecurityGroupHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSecuritygroupServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSecuritygroupServiceHandler struct{}

func (UnimplementedSecuritygroupServiceHandler) AuthorizeSecurityGroupEgress(context.Context, *connect.Request[v1.AuthorizeSecurityGroupEgressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupEgressResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.AuthorizeSecurityGroupEgress is not implemented"))
}

func (UnimplementedSecuritygroupServiceHandler) AuthorizeSecurityGroupIngress(context.Context, *connect.Request[v1.AuthorizeSecurityGroupIngressRequest]) (*connect.Response[v1.AuthorizeSecurityGroupIngressResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.AuthorizeSecurityGroupIngress is not implemented"))
}

func (UnimplementedSecuritygroupServiceHandler) CreateSecurityGroup(context.Context, *connect.Request[v1.CreateSecurityGroupRequest]) (*connect.Response[v1.CreateSecurityGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.CreateSecurityGroup is not implemented"))
}

func (UnimplementedSecuritygroupServiceHandler) DeleteSecurityGroup(context.Context, *connect.Request[v1.DeleteSecurityGroupRequest]) (*connect.Response[v1.DeleteSecurityGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.DeleteSecurityGroup is not implemented"))
}

func (UnimplementedSecuritygroupServiceHandler) ListSecurityGroups(context.Context, *connect.Request[v1.ListSecurityGroupsRequest]) (*connect.Response[v1.ListSecurityGroupsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.ListSecurityGroups is not implemented"))
}

func (UnimplementedSecuritygroupServiceHandler) RevokeSecurityGroupEgress(context.Context, *connect.Request[v1.RevokeSecurityGroupEgressRequest]) (*connect.Response[v1.RevokeSecurityGroupEgressResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.RevokeSecurityGroupEgress is not implemented"))
}

func (UnimplementedSecuritygroupServiceHandler) RevokeSecurityGroupIngress(context.Context, *connect.Request[v1.RevokeSecurityGroupIngressRequest]) (*connect.Response[v1.RevokeSecurityGroupIngressResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.RevokeSecurityGroupIngress is not implemented"))
}

func (UnimplementedSecuritygroupServiceHandler) UpdateSecurityGroup(context.Context, *connect.Request[v1.UpdateSecurityGroupRequest]) (*connect.Response[v1.UpdateSecurityGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.securitygroup.v1.SecuritygroupService.UpdateSecurityGroup is not implemented"))
}
