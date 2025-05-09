// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/management/guest/v1/guest.gen.proto

package guestv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/guest/v1"
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
	// GuestServiceName is the fully-qualified name of the GuestService service.
	GuestServiceName = "cloudstack.management.guest.v1.GuestService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GuestServiceAddGuestOsProcedure is the fully-qualified name of the GuestService's AddGuestOs RPC.
	GuestServiceAddGuestOsProcedure = "/cloudstack.management.guest.v1.GuestService/AddGuestOs"
	// GuestServiceAddGuestOsMappingProcedure is the fully-qualified name of the GuestService's
	// AddGuestOsMapping RPC.
	GuestServiceAddGuestOsMappingProcedure = "/cloudstack.management.guest.v1.GuestService/AddGuestOsMapping"
	// GuestServiceGetHypervisorGuestOsNamesProcedure is the fully-qualified name of the GuestService's
	// GetHypervisorGuestOsNames RPC.
	GuestServiceGetHypervisorGuestOsNamesProcedure = "/cloudstack.management.guest.v1.GuestService/GetHypervisorGuestOsNames"
	// GuestServiceListGuestOsProcedure is the fully-qualified name of the GuestService's ListGuestOs
	// RPC.
	GuestServiceListGuestOsProcedure = "/cloudstack.management.guest.v1.GuestService/ListGuestOs"
	// GuestServiceListGuestOsCategoriesProcedure is the fully-qualified name of the GuestService's
	// ListGuestOsCategories RPC.
	GuestServiceListGuestOsCategoriesProcedure = "/cloudstack.management.guest.v1.GuestService/ListGuestOsCategories"
	// GuestServiceListGuestOsMappingProcedure is the fully-qualified name of the GuestService's
	// ListGuestOsMapping RPC.
	GuestServiceListGuestOsMappingProcedure = "/cloudstack.management.guest.v1.GuestService/ListGuestOsMapping"
	// GuestServiceRemoveGuestOsProcedure is the fully-qualified name of the GuestService's
	// RemoveGuestOs RPC.
	GuestServiceRemoveGuestOsProcedure = "/cloudstack.management.guest.v1.GuestService/RemoveGuestOs"
	// GuestServiceRemoveGuestOsMappingProcedure is the fully-qualified name of the GuestService's
	// RemoveGuestOsMapping RPC.
	GuestServiceRemoveGuestOsMappingProcedure = "/cloudstack.management.guest.v1.GuestService/RemoveGuestOsMapping"
	// GuestServiceUpdateGuestOsProcedure is the fully-qualified name of the GuestService's
	// UpdateGuestOs RPC.
	GuestServiceUpdateGuestOsProcedure = "/cloudstack.management.guest.v1.GuestService/UpdateGuestOs"
	// GuestServiceUpdateGuestOsMappingProcedure is the fully-qualified name of the GuestService's
	// UpdateGuestOsMapping RPC.
	GuestServiceUpdateGuestOsMappingProcedure = "/cloudstack.management.guest.v1.GuestService/UpdateGuestOsMapping"
)

// GuestServiceClient is a client for the cloudstack.management.guest.v1.GuestService service.
type GuestServiceClient interface {
	// AddGuestOs Add a new guest OS type
	AddGuestOs(context.Context, *connect.Request[v1.AddGuestOsRequest]) (*connect.Response[v1.AddGuestOsResponse], error)
	// AddGuestOsMapping Adds a guest OS name to hypervisor OS name mapping
	AddGuestOsMapping(context.Context, *connect.Request[v1.AddGuestOsMappingRequest]) (*connect.Response[v1.AddGuestOsMappingResponse], error)
	// GetHypervisorGuestOsNames Gets the guest OS names in the hypervisor
	GetHypervisorGuestOsNames(context.Context, *connect.Request[v1.GetHypervisorGuestOsNamesRequest]) (*connect.Response[v1.GetHypervisorGuestOsNamesResponse], error)
	// ListGuestOs Lists all supported OS types for this cloud.
	ListGuestOs(context.Context, *connect.Request[v1.ListGuestOsRequest]) (*connect.Response[v1.ListGuestOsResponse], error)
	// ListGuestOsCategories Lists all supported OS categories for this cloud.
	ListGuestOsCategories(context.Context, *connect.Request[v1.ListGuestOsCategoriesRequest]) (*connect.Response[v1.ListGuestOsCategoriesResponse], error)
	// ListGuestOsMapping Lists all available OS mappings for given hypervisor
	ListGuestOsMapping(context.Context, *connect.Request[v1.ListGuestOsMappingRequest]) (*connect.Response[v1.ListGuestOsMappingResponse], error)
	// RemoveGuestOs Removes a Guest OS from listing.
	RemoveGuestOs(context.Context, *connect.Request[v1.RemoveGuestOsRequest]) (*connect.Response[v1.RemoveGuestOsResponse], error)
	// RemoveGuestOsMapping Removes a Guest OS Mapping.
	RemoveGuestOsMapping(context.Context, *connect.Request[v1.RemoveGuestOsMappingRequest]) (*connect.Response[v1.RemoveGuestOsMappingResponse], error)
	// UpdateGuestOs Updates the information about Guest OS
	UpdateGuestOs(context.Context, *connect.Request[v1.UpdateGuestOsRequest]) (*connect.Response[v1.UpdateGuestOsResponse], error)
	// UpdateGuestOsMapping Updates the information about Guest OS to Hypervisor specific name mapping
	UpdateGuestOsMapping(context.Context, *connect.Request[v1.UpdateGuestOsMappingRequest]) (*connect.Response[v1.UpdateGuestOsMappingResponse], error)
}

// NewGuestServiceClient constructs a client for the cloudstack.management.guest.v1.GuestService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGuestServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GuestServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	guestServiceMethods := v1.File_cloudstack_management_guest_v1_guest_gen_proto.Services().ByName("GuestService").Methods()
	return &guestServiceClient{
		addGuestOs: connect.NewClient[v1.AddGuestOsRequest, v1.AddGuestOsResponse](
			httpClient,
			baseURL+GuestServiceAddGuestOsProcedure,
			connect.WithSchema(guestServiceMethods.ByName("AddGuestOs")),
			connect.WithClientOptions(opts...),
		),
		addGuestOsMapping: connect.NewClient[v1.AddGuestOsMappingRequest, v1.AddGuestOsMappingResponse](
			httpClient,
			baseURL+GuestServiceAddGuestOsMappingProcedure,
			connect.WithSchema(guestServiceMethods.ByName("AddGuestOsMapping")),
			connect.WithClientOptions(opts...),
		),
		getHypervisorGuestOsNames: connect.NewClient[v1.GetHypervisorGuestOsNamesRequest, v1.GetHypervisorGuestOsNamesResponse](
			httpClient,
			baseURL+GuestServiceGetHypervisorGuestOsNamesProcedure,
			connect.WithSchema(guestServiceMethods.ByName("GetHypervisorGuestOsNames")),
			connect.WithClientOptions(opts...),
		),
		listGuestOs: connect.NewClient[v1.ListGuestOsRequest, v1.ListGuestOsResponse](
			httpClient,
			baseURL+GuestServiceListGuestOsProcedure,
			connect.WithSchema(guestServiceMethods.ByName("ListGuestOs")),
			connect.WithClientOptions(opts...),
		),
		listGuestOsCategories: connect.NewClient[v1.ListGuestOsCategoriesRequest, v1.ListGuestOsCategoriesResponse](
			httpClient,
			baseURL+GuestServiceListGuestOsCategoriesProcedure,
			connect.WithSchema(guestServiceMethods.ByName("ListGuestOsCategories")),
			connect.WithClientOptions(opts...),
		),
		listGuestOsMapping: connect.NewClient[v1.ListGuestOsMappingRequest, v1.ListGuestOsMappingResponse](
			httpClient,
			baseURL+GuestServiceListGuestOsMappingProcedure,
			connect.WithSchema(guestServiceMethods.ByName("ListGuestOsMapping")),
			connect.WithClientOptions(opts...),
		),
		removeGuestOs: connect.NewClient[v1.RemoveGuestOsRequest, v1.RemoveGuestOsResponse](
			httpClient,
			baseURL+GuestServiceRemoveGuestOsProcedure,
			connect.WithSchema(guestServiceMethods.ByName("RemoveGuestOs")),
			connect.WithClientOptions(opts...),
		),
		removeGuestOsMapping: connect.NewClient[v1.RemoveGuestOsMappingRequest, v1.RemoveGuestOsMappingResponse](
			httpClient,
			baseURL+GuestServiceRemoveGuestOsMappingProcedure,
			connect.WithSchema(guestServiceMethods.ByName("RemoveGuestOsMapping")),
			connect.WithClientOptions(opts...),
		),
		updateGuestOs: connect.NewClient[v1.UpdateGuestOsRequest, v1.UpdateGuestOsResponse](
			httpClient,
			baseURL+GuestServiceUpdateGuestOsProcedure,
			connect.WithSchema(guestServiceMethods.ByName("UpdateGuestOs")),
			connect.WithClientOptions(opts...),
		),
		updateGuestOsMapping: connect.NewClient[v1.UpdateGuestOsMappingRequest, v1.UpdateGuestOsMappingResponse](
			httpClient,
			baseURL+GuestServiceUpdateGuestOsMappingProcedure,
			connect.WithSchema(guestServiceMethods.ByName("UpdateGuestOsMapping")),
			connect.WithClientOptions(opts...),
		),
	}
}

// guestServiceClient implements GuestServiceClient.
type guestServiceClient struct {
	addGuestOs                *connect.Client[v1.AddGuestOsRequest, v1.AddGuestOsResponse]
	addGuestOsMapping         *connect.Client[v1.AddGuestOsMappingRequest, v1.AddGuestOsMappingResponse]
	getHypervisorGuestOsNames *connect.Client[v1.GetHypervisorGuestOsNamesRequest, v1.GetHypervisorGuestOsNamesResponse]
	listGuestOs               *connect.Client[v1.ListGuestOsRequest, v1.ListGuestOsResponse]
	listGuestOsCategories     *connect.Client[v1.ListGuestOsCategoriesRequest, v1.ListGuestOsCategoriesResponse]
	listGuestOsMapping        *connect.Client[v1.ListGuestOsMappingRequest, v1.ListGuestOsMappingResponse]
	removeGuestOs             *connect.Client[v1.RemoveGuestOsRequest, v1.RemoveGuestOsResponse]
	removeGuestOsMapping      *connect.Client[v1.RemoveGuestOsMappingRequest, v1.RemoveGuestOsMappingResponse]
	updateGuestOs             *connect.Client[v1.UpdateGuestOsRequest, v1.UpdateGuestOsResponse]
	updateGuestOsMapping      *connect.Client[v1.UpdateGuestOsMappingRequest, v1.UpdateGuestOsMappingResponse]
}

// AddGuestOs calls cloudstack.management.guest.v1.GuestService.AddGuestOs.
func (c *guestServiceClient) AddGuestOs(ctx context.Context, req *connect.Request[v1.AddGuestOsRequest]) (*connect.Response[v1.AddGuestOsResponse], error) {
	return c.addGuestOs.CallUnary(ctx, req)
}

// AddGuestOsMapping calls cloudstack.management.guest.v1.GuestService.AddGuestOsMapping.
func (c *guestServiceClient) AddGuestOsMapping(ctx context.Context, req *connect.Request[v1.AddGuestOsMappingRequest]) (*connect.Response[v1.AddGuestOsMappingResponse], error) {
	return c.addGuestOsMapping.CallUnary(ctx, req)
}

// GetHypervisorGuestOsNames calls
// cloudstack.management.guest.v1.GuestService.GetHypervisorGuestOsNames.
func (c *guestServiceClient) GetHypervisorGuestOsNames(ctx context.Context, req *connect.Request[v1.GetHypervisorGuestOsNamesRequest]) (*connect.Response[v1.GetHypervisorGuestOsNamesResponse], error) {
	return c.getHypervisorGuestOsNames.CallUnary(ctx, req)
}

// ListGuestOs calls cloudstack.management.guest.v1.GuestService.ListGuestOs.
func (c *guestServiceClient) ListGuestOs(ctx context.Context, req *connect.Request[v1.ListGuestOsRequest]) (*connect.Response[v1.ListGuestOsResponse], error) {
	return c.listGuestOs.CallUnary(ctx, req)
}

// ListGuestOsCategories calls cloudstack.management.guest.v1.GuestService.ListGuestOsCategories.
func (c *guestServiceClient) ListGuestOsCategories(ctx context.Context, req *connect.Request[v1.ListGuestOsCategoriesRequest]) (*connect.Response[v1.ListGuestOsCategoriesResponse], error) {
	return c.listGuestOsCategories.CallUnary(ctx, req)
}

// ListGuestOsMapping calls cloudstack.management.guest.v1.GuestService.ListGuestOsMapping.
func (c *guestServiceClient) ListGuestOsMapping(ctx context.Context, req *connect.Request[v1.ListGuestOsMappingRequest]) (*connect.Response[v1.ListGuestOsMappingResponse], error) {
	return c.listGuestOsMapping.CallUnary(ctx, req)
}

// RemoveGuestOs calls cloudstack.management.guest.v1.GuestService.RemoveGuestOs.
func (c *guestServiceClient) RemoveGuestOs(ctx context.Context, req *connect.Request[v1.RemoveGuestOsRequest]) (*connect.Response[v1.RemoveGuestOsResponse], error) {
	return c.removeGuestOs.CallUnary(ctx, req)
}

// RemoveGuestOsMapping calls cloudstack.management.guest.v1.GuestService.RemoveGuestOsMapping.
func (c *guestServiceClient) RemoveGuestOsMapping(ctx context.Context, req *connect.Request[v1.RemoveGuestOsMappingRequest]) (*connect.Response[v1.RemoveGuestOsMappingResponse], error) {
	return c.removeGuestOsMapping.CallUnary(ctx, req)
}

// UpdateGuestOs calls cloudstack.management.guest.v1.GuestService.UpdateGuestOs.
func (c *guestServiceClient) UpdateGuestOs(ctx context.Context, req *connect.Request[v1.UpdateGuestOsRequest]) (*connect.Response[v1.UpdateGuestOsResponse], error) {
	return c.updateGuestOs.CallUnary(ctx, req)
}

// UpdateGuestOsMapping calls cloudstack.management.guest.v1.GuestService.UpdateGuestOsMapping.
func (c *guestServiceClient) UpdateGuestOsMapping(ctx context.Context, req *connect.Request[v1.UpdateGuestOsMappingRequest]) (*connect.Response[v1.UpdateGuestOsMappingResponse], error) {
	return c.updateGuestOsMapping.CallUnary(ctx, req)
}

// GuestServiceHandler is an implementation of the cloudstack.management.guest.v1.GuestService
// service.
type GuestServiceHandler interface {
	// AddGuestOs Add a new guest OS type
	AddGuestOs(context.Context, *connect.Request[v1.AddGuestOsRequest]) (*connect.Response[v1.AddGuestOsResponse], error)
	// AddGuestOsMapping Adds a guest OS name to hypervisor OS name mapping
	AddGuestOsMapping(context.Context, *connect.Request[v1.AddGuestOsMappingRequest]) (*connect.Response[v1.AddGuestOsMappingResponse], error)
	// GetHypervisorGuestOsNames Gets the guest OS names in the hypervisor
	GetHypervisorGuestOsNames(context.Context, *connect.Request[v1.GetHypervisorGuestOsNamesRequest]) (*connect.Response[v1.GetHypervisorGuestOsNamesResponse], error)
	// ListGuestOs Lists all supported OS types for this cloud.
	ListGuestOs(context.Context, *connect.Request[v1.ListGuestOsRequest]) (*connect.Response[v1.ListGuestOsResponse], error)
	// ListGuestOsCategories Lists all supported OS categories for this cloud.
	ListGuestOsCategories(context.Context, *connect.Request[v1.ListGuestOsCategoriesRequest]) (*connect.Response[v1.ListGuestOsCategoriesResponse], error)
	// ListGuestOsMapping Lists all available OS mappings for given hypervisor
	ListGuestOsMapping(context.Context, *connect.Request[v1.ListGuestOsMappingRequest]) (*connect.Response[v1.ListGuestOsMappingResponse], error)
	// RemoveGuestOs Removes a Guest OS from listing.
	RemoveGuestOs(context.Context, *connect.Request[v1.RemoveGuestOsRequest]) (*connect.Response[v1.RemoveGuestOsResponse], error)
	// RemoveGuestOsMapping Removes a Guest OS Mapping.
	RemoveGuestOsMapping(context.Context, *connect.Request[v1.RemoveGuestOsMappingRequest]) (*connect.Response[v1.RemoveGuestOsMappingResponse], error)
	// UpdateGuestOs Updates the information about Guest OS
	UpdateGuestOs(context.Context, *connect.Request[v1.UpdateGuestOsRequest]) (*connect.Response[v1.UpdateGuestOsResponse], error)
	// UpdateGuestOsMapping Updates the information about Guest OS to Hypervisor specific name mapping
	UpdateGuestOsMapping(context.Context, *connect.Request[v1.UpdateGuestOsMappingRequest]) (*connect.Response[v1.UpdateGuestOsMappingResponse], error)
}

// NewGuestServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGuestServiceHandler(svc GuestServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	guestServiceMethods := v1.File_cloudstack_management_guest_v1_guest_gen_proto.Services().ByName("GuestService").Methods()
	guestServiceAddGuestOsHandler := connect.NewUnaryHandler(
		GuestServiceAddGuestOsProcedure,
		svc.AddGuestOs,
		connect.WithSchema(guestServiceMethods.ByName("AddGuestOs")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceAddGuestOsMappingHandler := connect.NewUnaryHandler(
		GuestServiceAddGuestOsMappingProcedure,
		svc.AddGuestOsMapping,
		connect.WithSchema(guestServiceMethods.ByName("AddGuestOsMapping")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceGetHypervisorGuestOsNamesHandler := connect.NewUnaryHandler(
		GuestServiceGetHypervisorGuestOsNamesProcedure,
		svc.GetHypervisorGuestOsNames,
		connect.WithSchema(guestServiceMethods.ByName("GetHypervisorGuestOsNames")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceListGuestOsHandler := connect.NewUnaryHandler(
		GuestServiceListGuestOsProcedure,
		svc.ListGuestOs,
		connect.WithSchema(guestServiceMethods.ByName("ListGuestOs")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceListGuestOsCategoriesHandler := connect.NewUnaryHandler(
		GuestServiceListGuestOsCategoriesProcedure,
		svc.ListGuestOsCategories,
		connect.WithSchema(guestServiceMethods.ByName("ListGuestOsCategories")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceListGuestOsMappingHandler := connect.NewUnaryHandler(
		GuestServiceListGuestOsMappingProcedure,
		svc.ListGuestOsMapping,
		connect.WithSchema(guestServiceMethods.ByName("ListGuestOsMapping")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceRemoveGuestOsHandler := connect.NewUnaryHandler(
		GuestServiceRemoveGuestOsProcedure,
		svc.RemoveGuestOs,
		connect.WithSchema(guestServiceMethods.ByName("RemoveGuestOs")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceRemoveGuestOsMappingHandler := connect.NewUnaryHandler(
		GuestServiceRemoveGuestOsMappingProcedure,
		svc.RemoveGuestOsMapping,
		connect.WithSchema(guestServiceMethods.ByName("RemoveGuestOsMapping")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceUpdateGuestOsHandler := connect.NewUnaryHandler(
		GuestServiceUpdateGuestOsProcedure,
		svc.UpdateGuestOs,
		connect.WithSchema(guestServiceMethods.ByName("UpdateGuestOs")),
		connect.WithHandlerOptions(opts...),
	)
	guestServiceUpdateGuestOsMappingHandler := connect.NewUnaryHandler(
		GuestServiceUpdateGuestOsMappingProcedure,
		svc.UpdateGuestOsMapping,
		connect.WithSchema(guestServiceMethods.ByName("UpdateGuestOsMapping")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.management.guest.v1.GuestService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GuestServiceAddGuestOsProcedure:
			guestServiceAddGuestOsHandler.ServeHTTP(w, r)
		case GuestServiceAddGuestOsMappingProcedure:
			guestServiceAddGuestOsMappingHandler.ServeHTTP(w, r)
		case GuestServiceGetHypervisorGuestOsNamesProcedure:
			guestServiceGetHypervisorGuestOsNamesHandler.ServeHTTP(w, r)
		case GuestServiceListGuestOsProcedure:
			guestServiceListGuestOsHandler.ServeHTTP(w, r)
		case GuestServiceListGuestOsCategoriesProcedure:
			guestServiceListGuestOsCategoriesHandler.ServeHTTP(w, r)
		case GuestServiceListGuestOsMappingProcedure:
			guestServiceListGuestOsMappingHandler.ServeHTTP(w, r)
		case GuestServiceRemoveGuestOsProcedure:
			guestServiceRemoveGuestOsHandler.ServeHTTP(w, r)
		case GuestServiceRemoveGuestOsMappingProcedure:
			guestServiceRemoveGuestOsMappingHandler.ServeHTTP(w, r)
		case GuestServiceUpdateGuestOsProcedure:
			guestServiceUpdateGuestOsHandler.ServeHTTP(w, r)
		case GuestServiceUpdateGuestOsMappingProcedure:
			guestServiceUpdateGuestOsMappingHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGuestServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGuestServiceHandler struct{}

func (UnimplementedGuestServiceHandler) AddGuestOs(context.Context, *connect.Request[v1.AddGuestOsRequest]) (*connect.Response[v1.AddGuestOsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.AddGuestOs is not implemented"))
}

func (UnimplementedGuestServiceHandler) AddGuestOsMapping(context.Context, *connect.Request[v1.AddGuestOsMappingRequest]) (*connect.Response[v1.AddGuestOsMappingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.AddGuestOsMapping is not implemented"))
}

func (UnimplementedGuestServiceHandler) GetHypervisorGuestOsNames(context.Context, *connect.Request[v1.GetHypervisorGuestOsNamesRequest]) (*connect.Response[v1.GetHypervisorGuestOsNamesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.GetHypervisorGuestOsNames is not implemented"))
}

func (UnimplementedGuestServiceHandler) ListGuestOs(context.Context, *connect.Request[v1.ListGuestOsRequest]) (*connect.Response[v1.ListGuestOsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.ListGuestOs is not implemented"))
}

func (UnimplementedGuestServiceHandler) ListGuestOsCategories(context.Context, *connect.Request[v1.ListGuestOsCategoriesRequest]) (*connect.Response[v1.ListGuestOsCategoriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.ListGuestOsCategories is not implemented"))
}

func (UnimplementedGuestServiceHandler) ListGuestOsMapping(context.Context, *connect.Request[v1.ListGuestOsMappingRequest]) (*connect.Response[v1.ListGuestOsMappingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.ListGuestOsMapping is not implemented"))
}

func (UnimplementedGuestServiceHandler) RemoveGuestOs(context.Context, *connect.Request[v1.RemoveGuestOsRequest]) (*connect.Response[v1.RemoveGuestOsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.RemoveGuestOs is not implemented"))
}

func (UnimplementedGuestServiceHandler) RemoveGuestOsMapping(context.Context, *connect.Request[v1.RemoveGuestOsMappingRequest]) (*connect.Response[v1.RemoveGuestOsMappingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.RemoveGuestOsMapping is not implemented"))
}

func (UnimplementedGuestServiceHandler) UpdateGuestOs(context.Context, *connect.Request[v1.UpdateGuestOsRequest]) (*connect.Response[v1.UpdateGuestOsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.UpdateGuestOs is not implemented"))
}

func (UnimplementedGuestServiceHandler) UpdateGuestOsMapping(context.Context, *connect.Request[v1.UpdateGuestOsMappingRequest]) (*connect.Response[v1.UpdateGuestOsMappingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.guest.v1.GuestService.UpdateGuestOsMapping is not implemented"))
}
