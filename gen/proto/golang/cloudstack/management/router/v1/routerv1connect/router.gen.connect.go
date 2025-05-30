// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/management/router/v1/router.gen.proto

package routerv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/router/v1"
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
	// RouterServiceName is the fully-qualified name of the RouterService service.
	RouterServiceName = "cloudstack.management.router.v1.RouterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RouterServiceConfigureOvsElementProcedure is the fully-qualified name of the RouterService's
	// ConfigureOvsElement RPC.
	RouterServiceConfigureOvsElementProcedure = "/cloudstack.management.router.v1.RouterService/ConfigureOvsElement"
	// RouterServiceConfigureVirtualRouterElementProcedure is the fully-qualified name of the
	// RouterService's ConfigureVirtualRouterElement RPC.
	RouterServiceConfigureVirtualRouterElementProcedure = "/cloudstack.management.router.v1.RouterService/ConfigureVirtualRouterElement"
	// RouterServiceCreateVirtualRouterElementProcedure is the fully-qualified name of the
	// RouterService's CreateVirtualRouterElement RPC.
	RouterServiceCreateVirtualRouterElementProcedure = "/cloudstack.management.router.v1.RouterService/CreateVirtualRouterElement"
	// RouterServiceDestroyRouterProcedure is the fully-qualified name of the RouterService's
	// DestroyRouter RPC.
	RouterServiceDestroyRouterProcedure = "/cloudstack.management.router.v1.RouterService/DestroyRouter"
	// RouterServiceGetRouterHealthCheckResultsProcedure is the fully-qualified name of the
	// RouterService's GetRouterHealthCheckResults RPC.
	RouterServiceGetRouterHealthCheckResultsProcedure = "/cloudstack.management.router.v1.RouterService/GetRouterHealthCheckResults"
	// RouterServiceListOvsElementsProcedure is the fully-qualified name of the RouterService's
	// ListOvsElements RPC.
	RouterServiceListOvsElementsProcedure = "/cloudstack.management.router.v1.RouterService/ListOvsElements"
	// RouterServiceListRoutersProcedure is the fully-qualified name of the RouterService's ListRouters
	// RPC.
	RouterServiceListRoutersProcedure = "/cloudstack.management.router.v1.RouterService/ListRouters"
	// RouterServiceListVirtualRouterElementsProcedure is the fully-qualified name of the
	// RouterService's ListVirtualRouterElements RPC.
	RouterServiceListVirtualRouterElementsProcedure = "/cloudstack.management.router.v1.RouterService/ListVirtualRouterElements"
	// RouterServiceRebootRouterProcedure is the fully-qualified name of the RouterService's
	// RebootRouter RPC.
	RouterServiceRebootRouterProcedure = "/cloudstack.management.router.v1.RouterService/RebootRouter"
	// RouterServiceStartRouterProcedure is the fully-qualified name of the RouterService's StartRouter
	// RPC.
	RouterServiceStartRouterProcedure = "/cloudstack.management.router.v1.RouterService/StartRouter"
	// RouterServiceStopRouterProcedure is the fully-qualified name of the RouterService's StopRouter
	// RPC.
	RouterServiceStopRouterProcedure = "/cloudstack.management.router.v1.RouterService/StopRouter"
	// RouterServiceUpgradeRouterProcedure is the fully-qualified name of the RouterService's
	// UpgradeRouter RPC.
	RouterServiceUpgradeRouterProcedure = "/cloudstack.management.router.v1.RouterService/UpgradeRouter"
	// RouterServiceUpgradeRouterTemplateProcedure is the fully-qualified name of the RouterService's
	// UpgradeRouterTemplate RPC.
	RouterServiceUpgradeRouterTemplateProcedure = "/cloudstack.management.router.v1.RouterService/UpgradeRouterTemplate"
)

// RouterServiceClient is a client for the cloudstack.management.router.v1.RouterService service.
type RouterServiceClient interface {
	// ConfigureOvsElement Configures an ovs element.
	ConfigureOvsElement(context.Context, *connect.Request[v1.ConfigureOvsElementRequest]) (*connect.Response[v1.ConfigureOvsElementResponse], error)
	// ConfigureVirtualRouterElement Configures a virtual router element.
	ConfigureVirtualRouterElement(context.Context, *connect.Request[v1.ConfigureVirtualRouterElementRequest]) (*connect.Response[v1.ConfigureVirtualRouterElementResponse], error)
	// CreateVirtualRouterElement Create a virtual router element.
	CreateVirtualRouterElement(context.Context, *connect.Request[v1.CreateVirtualRouterElementRequest]) (*connect.Response[v1.CreateVirtualRouterElementResponse], error)
	// DestroyRouter Destroys a router.
	DestroyRouter(context.Context, *connect.Request[v1.DestroyRouterRequest]) (*connect.Response[v1.DestroyRouterResponse], error)
	// GetRouterHealthCheckResults Starts a router.
	GetRouterHealthCheckResults(context.Context, *connect.Request[v1.GetRouterHealthCheckResultsRequest]) (*connect.Response[v1.GetRouterHealthCheckResultsResponse], error)
	// ListOvsElements Lists all available ovs elements.
	ListOvsElements(context.Context, *connect.Request[v1.ListOvsElementsRequest]) (*connect.Response[v1.ListOvsElementsResponse], error)
	// ListRouters List routers.
	ListRouters(context.Context, *connect.Request[v1.ListRoutersRequest]) (*connect.Response[v1.ListRoutersResponse], error)
	// ListVirtualRouterElements Lists all available virtual router elements.
	ListVirtualRouterElements(context.Context, *connect.Request[v1.ListVirtualRouterElementsRequest]) (*connect.Response[v1.ListVirtualRouterElementsResponse], error)
	// RebootRouter Starts a router.
	RebootRouter(context.Context, *connect.Request[v1.RebootRouterRequest]) (*connect.Response[v1.RebootRouterResponse], error)
	// StartRouter Starts a router.
	StartRouter(context.Context, *connect.Request[v1.StartRouterRequest]) (*connect.Response[v1.StartRouterResponse], error)
	// StopRouter Stops a router.
	StopRouter(context.Context, *connect.Request[v1.StopRouterRequest]) (*connect.Response[v1.StopRouterResponse], error)
	// UpgradeRouter Upgrades domain router to a new service offering
	UpgradeRouter(context.Context, *connect.Request[v1.UpgradeRouterRequest]) (*connect.Response[v1.UpgradeRouterResponse], error)
	// UpgradeRouterTemplate Upgrades router to use newer template
	UpgradeRouterTemplate(context.Context, *connect.Request[v1.UpgradeRouterTemplateRequest]) (*connect.Response[v1.UpgradeRouterTemplateResponse], error)
}

// NewRouterServiceClient constructs a client for the cloudstack.management.router.v1.RouterService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRouterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RouterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	routerServiceMethods := v1.File_cloudstack_management_router_v1_router_gen_proto.Services().ByName("RouterService").Methods()
	return &routerServiceClient{
		configureOvsElement: connect.NewClient[v1.ConfigureOvsElementRequest, v1.ConfigureOvsElementResponse](
			httpClient,
			baseURL+RouterServiceConfigureOvsElementProcedure,
			connect.WithSchema(routerServiceMethods.ByName("ConfigureOvsElement")),
			connect.WithClientOptions(opts...),
		),
		configureVirtualRouterElement: connect.NewClient[v1.ConfigureVirtualRouterElementRequest, v1.ConfigureVirtualRouterElementResponse](
			httpClient,
			baseURL+RouterServiceConfigureVirtualRouterElementProcedure,
			connect.WithSchema(routerServiceMethods.ByName("ConfigureVirtualRouterElement")),
			connect.WithClientOptions(opts...),
		),
		createVirtualRouterElement: connect.NewClient[v1.CreateVirtualRouterElementRequest, v1.CreateVirtualRouterElementResponse](
			httpClient,
			baseURL+RouterServiceCreateVirtualRouterElementProcedure,
			connect.WithSchema(routerServiceMethods.ByName("CreateVirtualRouterElement")),
			connect.WithClientOptions(opts...),
		),
		destroyRouter: connect.NewClient[v1.DestroyRouterRequest, v1.DestroyRouterResponse](
			httpClient,
			baseURL+RouterServiceDestroyRouterProcedure,
			connect.WithSchema(routerServiceMethods.ByName("DestroyRouter")),
			connect.WithClientOptions(opts...),
		),
		getRouterHealthCheckResults: connect.NewClient[v1.GetRouterHealthCheckResultsRequest, v1.GetRouterHealthCheckResultsResponse](
			httpClient,
			baseURL+RouterServiceGetRouterHealthCheckResultsProcedure,
			connect.WithSchema(routerServiceMethods.ByName("GetRouterHealthCheckResults")),
			connect.WithClientOptions(opts...),
		),
		listOvsElements: connect.NewClient[v1.ListOvsElementsRequest, v1.ListOvsElementsResponse](
			httpClient,
			baseURL+RouterServiceListOvsElementsProcedure,
			connect.WithSchema(routerServiceMethods.ByName("ListOvsElements")),
			connect.WithClientOptions(opts...),
		),
		listRouters: connect.NewClient[v1.ListRoutersRequest, v1.ListRoutersResponse](
			httpClient,
			baseURL+RouterServiceListRoutersProcedure,
			connect.WithSchema(routerServiceMethods.ByName("ListRouters")),
			connect.WithClientOptions(opts...),
		),
		listVirtualRouterElements: connect.NewClient[v1.ListVirtualRouterElementsRequest, v1.ListVirtualRouterElementsResponse](
			httpClient,
			baseURL+RouterServiceListVirtualRouterElementsProcedure,
			connect.WithSchema(routerServiceMethods.ByName("ListVirtualRouterElements")),
			connect.WithClientOptions(opts...),
		),
		rebootRouter: connect.NewClient[v1.RebootRouterRequest, v1.RebootRouterResponse](
			httpClient,
			baseURL+RouterServiceRebootRouterProcedure,
			connect.WithSchema(routerServiceMethods.ByName("RebootRouter")),
			connect.WithClientOptions(opts...),
		),
		startRouter: connect.NewClient[v1.StartRouterRequest, v1.StartRouterResponse](
			httpClient,
			baseURL+RouterServiceStartRouterProcedure,
			connect.WithSchema(routerServiceMethods.ByName("StartRouter")),
			connect.WithClientOptions(opts...),
		),
		stopRouter: connect.NewClient[v1.StopRouterRequest, v1.StopRouterResponse](
			httpClient,
			baseURL+RouterServiceStopRouterProcedure,
			connect.WithSchema(routerServiceMethods.ByName("StopRouter")),
			connect.WithClientOptions(opts...),
		),
		upgradeRouter: connect.NewClient[v1.UpgradeRouterRequest, v1.UpgradeRouterResponse](
			httpClient,
			baseURL+RouterServiceUpgradeRouterProcedure,
			connect.WithSchema(routerServiceMethods.ByName("UpgradeRouter")),
			connect.WithClientOptions(opts...),
		),
		upgradeRouterTemplate: connect.NewClient[v1.UpgradeRouterTemplateRequest, v1.UpgradeRouterTemplateResponse](
			httpClient,
			baseURL+RouterServiceUpgradeRouterTemplateProcedure,
			connect.WithSchema(routerServiceMethods.ByName("UpgradeRouterTemplate")),
			connect.WithClientOptions(opts...),
		),
	}
}

// routerServiceClient implements RouterServiceClient.
type routerServiceClient struct {
	configureOvsElement           *connect.Client[v1.ConfigureOvsElementRequest, v1.ConfigureOvsElementResponse]
	configureVirtualRouterElement *connect.Client[v1.ConfigureVirtualRouterElementRequest, v1.ConfigureVirtualRouterElementResponse]
	createVirtualRouterElement    *connect.Client[v1.CreateVirtualRouterElementRequest, v1.CreateVirtualRouterElementResponse]
	destroyRouter                 *connect.Client[v1.DestroyRouterRequest, v1.DestroyRouterResponse]
	getRouterHealthCheckResults   *connect.Client[v1.GetRouterHealthCheckResultsRequest, v1.GetRouterHealthCheckResultsResponse]
	listOvsElements               *connect.Client[v1.ListOvsElementsRequest, v1.ListOvsElementsResponse]
	listRouters                   *connect.Client[v1.ListRoutersRequest, v1.ListRoutersResponse]
	listVirtualRouterElements     *connect.Client[v1.ListVirtualRouterElementsRequest, v1.ListVirtualRouterElementsResponse]
	rebootRouter                  *connect.Client[v1.RebootRouterRequest, v1.RebootRouterResponse]
	startRouter                   *connect.Client[v1.StartRouterRequest, v1.StartRouterResponse]
	stopRouter                    *connect.Client[v1.StopRouterRequest, v1.StopRouterResponse]
	upgradeRouter                 *connect.Client[v1.UpgradeRouterRequest, v1.UpgradeRouterResponse]
	upgradeRouterTemplate         *connect.Client[v1.UpgradeRouterTemplateRequest, v1.UpgradeRouterTemplateResponse]
}

// ConfigureOvsElement calls cloudstack.management.router.v1.RouterService.ConfigureOvsElement.
func (c *routerServiceClient) ConfigureOvsElement(ctx context.Context, req *connect.Request[v1.ConfigureOvsElementRequest]) (*connect.Response[v1.ConfigureOvsElementResponse], error) {
	return c.configureOvsElement.CallUnary(ctx, req)
}

// ConfigureVirtualRouterElement calls
// cloudstack.management.router.v1.RouterService.ConfigureVirtualRouterElement.
func (c *routerServiceClient) ConfigureVirtualRouterElement(ctx context.Context, req *connect.Request[v1.ConfigureVirtualRouterElementRequest]) (*connect.Response[v1.ConfigureVirtualRouterElementResponse], error) {
	return c.configureVirtualRouterElement.CallUnary(ctx, req)
}

// CreateVirtualRouterElement calls
// cloudstack.management.router.v1.RouterService.CreateVirtualRouterElement.
func (c *routerServiceClient) CreateVirtualRouterElement(ctx context.Context, req *connect.Request[v1.CreateVirtualRouterElementRequest]) (*connect.Response[v1.CreateVirtualRouterElementResponse], error) {
	return c.createVirtualRouterElement.CallUnary(ctx, req)
}

// DestroyRouter calls cloudstack.management.router.v1.RouterService.DestroyRouter.
func (c *routerServiceClient) DestroyRouter(ctx context.Context, req *connect.Request[v1.DestroyRouterRequest]) (*connect.Response[v1.DestroyRouterResponse], error) {
	return c.destroyRouter.CallUnary(ctx, req)
}

// GetRouterHealthCheckResults calls
// cloudstack.management.router.v1.RouterService.GetRouterHealthCheckResults.
func (c *routerServiceClient) GetRouterHealthCheckResults(ctx context.Context, req *connect.Request[v1.GetRouterHealthCheckResultsRequest]) (*connect.Response[v1.GetRouterHealthCheckResultsResponse], error) {
	return c.getRouterHealthCheckResults.CallUnary(ctx, req)
}

// ListOvsElements calls cloudstack.management.router.v1.RouterService.ListOvsElements.
func (c *routerServiceClient) ListOvsElements(ctx context.Context, req *connect.Request[v1.ListOvsElementsRequest]) (*connect.Response[v1.ListOvsElementsResponse], error) {
	return c.listOvsElements.CallUnary(ctx, req)
}

// ListRouters calls cloudstack.management.router.v1.RouterService.ListRouters.
func (c *routerServiceClient) ListRouters(ctx context.Context, req *connect.Request[v1.ListRoutersRequest]) (*connect.Response[v1.ListRoutersResponse], error) {
	return c.listRouters.CallUnary(ctx, req)
}

// ListVirtualRouterElements calls
// cloudstack.management.router.v1.RouterService.ListVirtualRouterElements.
func (c *routerServiceClient) ListVirtualRouterElements(ctx context.Context, req *connect.Request[v1.ListVirtualRouterElementsRequest]) (*connect.Response[v1.ListVirtualRouterElementsResponse], error) {
	return c.listVirtualRouterElements.CallUnary(ctx, req)
}

// RebootRouter calls cloudstack.management.router.v1.RouterService.RebootRouter.
func (c *routerServiceClient) RebootRouter(ctx context.Context, req *connect.Request[v1.RebootRouterRequest]) (*connect.Response[v1.RebootRouterResponse], error) {
	return c.rebootRouter.CallUnary(ctx, req)
}

// StartRouter calls cloudstack.management.router.v1.RouterService.StartRouter.
func (c *routerServiceClient) StartRouter(ctx context.Context, req *connect.Request[v1.StartRouterRequest]) (*connect.Response[v1.StartRouterResponse], error) {
	return c.startRouter.CallUnary(ctx, req)
}

// StopRouter calls cloudstack.management.router.v1.RouterService.StopRouter.
func (c *routerServiceClient) StopRouter(ctx context.Context, req *connect.Request[v1.StopRouterRequest]) (*connect.Response[v1.StopRouterResponse], error) {
	return c.stopRouter.CallUnary(ctx, req)
}

// UpgradeRouter calls cloudstack.management.router.v1.RouterService.UpgradeRouter.
func (c *routerServiceClient) UpgradeRouter(ctx context.Context, req *connect.Request[v1.UpgradeRouterRequest]) (*connect.Response[v1.UpgradeRouterResponse], error) {
	return c.upgradeRouter.CallUnary(ctx, req)
}

// UpgradeRouterTemplate calls cloudstack.management.router.v1.RouterService.UpgradeRouterTemplate.
func (c *routerServiceClient) UpgradeRouterTemplate(ctx context.Context, req *connect.Request[v1.UpgradeRouterTemplateRequest]) (*connect.Response[v1.UpgradeRouterTemplateResponse], error) {
	return c.upgradeRouterTemplate.CallUnary(ctx, req)
}

// RouterServiceHandler is an implementation of the cloudstack.management.router.v1.RouterService
// service.
type RouterServiceHandler interface {
	// ConfigureOvsElement Configures an ovs element.
	ConfigureOvsElement(context.Context, *connect.Request[v1.ConfigureOvsElementRequest]) (*connect.Response[v1.ConfigureOvsElementResponse], error)
	// ConfigureVirtualRouterElement Configures a virtual router element.
	ConfigureVirtualRouterElement(context.Context, *connect.Request[v1.ConfigureVirtualRouterElementRequest]) (*connect.Response[v1.ConfigureVirtualRouterElementResponse], error)
	// CreateVirtualRouterElement Create a virtual router element.
	CreateVirtualRouterElement(context.Context, *connect.Request[v1.CreateVirtualRouterElementRequest]) (*connect.Response[v1.CreateVirtualRouterElementResponse], error)
	// DestroyRouter Destroys a router.
	DestroyRouter(context.Context, *connect.Request[v1.DestroyRouterRequest]) (*connect.Response[v1.DestroyRouterResponse], error)
	// GetRouterHealthCheckResults Starts a router.
	GetRouterHealthCheckResults(context.Context, *connect.Request[v1.GetRouterHealthCheckResultsRequest]) (*connect.Response[v1.GetRouterHealthCheckResultsResponse], error)
	// ListOvsElements Lists all available ovs elements.
	ListOvsElements(context.Context, *connect.Request[v1.ListOvsElementsRequest]) (*connect.Response[v1.ListOvsElementsResponse], error)
	// ListRouters List routers.
	ListRouters(context.Context, *connect.Request[v1.ListRoutersRequest]) (*connect.Response[v1.ListRoutersResponse], error)
	// ListVirtualRouterElements Lists all available virtual router elements.
	ListVirtualRouterElements(context.Context, *connect.Request[v1.ListVirtualRouterElementsRequest]) (*connect.Response[v1.ListVirtualRouterElementsResponse], error)
	// RebootRouter Starts a router.
	RebootRouter(context.Context, *connect.Request[v1.RebootRouterRequest]) (*connect.Response[v1.RebootRouterResponse], error)
	// StartRouter Starts a router.
	StartRouter(context.Context, *connect.Request[v1.StartRouterRequest]) (*connect.Response[v1.StartRouterResponse], error)
	// StopRouter Stops a router.
	StopRouter(context.Context, *connect.Request[v1.StopRouterRequest]) (*connect.Response[v1.StopRouterResponse], error)
	// UpgradeRouter Upgrades domain router to a new service offering
	UpgradeRouter(context.Context, *connect.Request[v1.UpgradeRouterRequest]) (*connect.Response[v1.UpgradeRouterResponse], error)
	// UpgradeRouterTemplate Upgrades router to use newer template
	UpgradeRouterTemplate(context.Context, *connect.Request[v1.UpgradeRouterTemplateRequest]) (*connect.Response[v1.UpgradeRouterTemplateResponse], error)
}

// NewRouterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRouterServiceHandler(svc RouterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	routerServiceMethods := v1.File_cloudstack_management_router_v1_router_gen_proto.Services().ByName("RouterService").Methods()
	routerServiceConfigureOvsElementHandler := connect.NewUnaryHandler(
		RouterServiceConfigureOvsElementProcedure,
		svc.ConfigureOvsElement,
		connect.WithSchema(routerServiceMethods.ByName("ConfigureOvsElement")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceConfigureVirtualRouterElementHandler := connect.NewUnaryHandler(
		RouterServiceConfigureVirtualRouterElementProcedure,
		svc.ConfigureVirtualRouterElement,
		connect.WithSchema(routerServiceMethods.ByName("ConfigureVirtualRouterElement")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceCreateVirtualRouterElementHandler := connect.NewUnaryHandler(
		RouterServiceCreateVirtualRouterElementProcedure,
		svc.CreateVirtualRouterElement,
		connect.WithSchema(routerServiceMethods.ByName("CreateVirtualRouterElement")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceDestroyRouterHandler := connect.NewUnaryHandler(
		RouterServiceDestroyRouterProcedure,
		svc.DestroyRouter,
		connect.WithSchema(routerServiceMethods.ByName("DestroyRouter")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceGetRouterHealthCheckResultsHandler := connect.NewUnaryHandler(
		RouterServiceGetRouterHealthCheckResultsProcedure,
		svc.GetRouterHealthCheckResults,
		connect.WithSchema(routerServiceMethods.ByName("GetRouterHealthCheckResults")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceListOvsElementsHandler := connect.NewUnaryHandler(
		RouterServiceListOvsElementsProcedure,
		svc.ListOvsElements,
		connect.WithSchema(routerServiceMethods.ByName("ListOvsElements")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceListRoutersHandler := connect.NewUnaryHandler(
		RouterServiceListRoutersProcedure,
		svc.ListRouters,
		connect.WithSchema(routerServiceMethods.ByName("ListRouters")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceListVirtualRouterElementsHandler := connect.NewUnaryHandler(
		RouterServiceListVirtualRouterElementsProcedure,
		svc.ListVirtualRouterElements,
		connect.WithSchema(routerServiceMethods.ByName("ListVirtualRouterElements")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceRebootRouterHandler := connect.NewUnaryHandler(
		RouterServiceRebootRouterProcedure,
		svc.RebootRouter,
		connect.WithSchema(routerServiceMethods.ByName("RebootRouter")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceStartRouterHandler := connect.NewUnaryHandler(
		RouterServiceStartRouterProcedure,
		svc.StartRouter,
		connect.WithSchema(routerServiceMethods.ByName("StartRouter")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceStopRouterHandler := connect.NewUnaryHandler(
		RouterServiceStopRouterProcedure,
		svc.StopRouter,
		connect.WithSchema(routerServiceMethods.ByName("StopRouter")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceUpgradeRouterHandler := connect.NewUnaryHandler(
		RouterServiceUpgradeRouterProcedure,
		svc.UpgradeRouter,
		connect.WithSchema(routerServiceMethods.ByName("UpgradeRouter")),
		connect.WithHandlerOptions(opts...),
	)
	routerServiceUpgradeRouterTemplateHandler := connect.NewUnaryHandler(
		RouterServiceUpgradeRouterTemplateProcedure,
		svc.UpgradeRouterTemplate,
		connect.WithSchema(routerServiceMethods.ByName("UpgradeRouterTemplate")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.management.router.v1.RouterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RouterServiceConfigureOvsElementProcedure:
			routerServiceConfigureOvsElementHandler.ServeHTTP(w, r)
		case RouterServiceConfigureVirtualRouterElementProcedure:
			routerServiceConfigureVirtualRouterElementHandler.ServeHTTP(w, r)
		case RouterServiceCreateVirtualRouterElementProcedure:
			routerServiceCreateVirtualRouterElementHandler.ServeHTTP(w, r)
		case RouterServiceDestroyRouterProcedure:
			routerServiceDestroyRouterHandler.ServeHTTP(w, r)
		case RouterServiceGetRouterHealthCheckResultsProcedure:
			routerServiceGetRouterHealthCheckResultsHandler.ServeHTTP(w, r)
		case RouterServiceListOvsElementsProcedure:
			routerServiceListOvsElementsHandler.ServeHTTP(w, r)
		case RouterServiceListRoutersProcedure:
			routerServiceListRoutersHandler.ServeHTTP(w, r)
		case RouterServiceListVirtualRouterElementsProcedure:
			routerServiceListVirtualRouterElementsHandler.ServeHTTP(w, r)
		case RouterServiceRebootRouterProcedure:
			routerServiceRebootRouterHandler.ServeHTTP(w, r)
		case RouterServiceStartRouterProcedure:
			routerServiceStartRouterHandler.ServeHTTP(w, r)
		case RouterServiceStopRouterProcedure:
			routerServiceStopRouterHandler.ServeHTTP(w, r)
		case RouterServiceUpgradeRouterProcedure:
			routerServiceUpgradeRouterHandler.ServeHTTP(w, r)
		case RouterServiceUpgradeRouterTemplateProcedure:
			routerServiceUpgradeRouterTemplateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRouterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRouterServiceHandler struct{}

func (UnimplementedRouterServiceHandler) ConfigureOvsElement(context.Context, *connect.Request[v1.ConfigureOvsElementRequest]) (*connect.Response[v1.ConfigureOvsElementResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.ConfigureOvsElement is not implemented"))
}

func (UnimplementedRouterServiceHandler) ConfigureVirtualRouterElement(context.Context, *connect.Request[v1.ConfigureVirtualRouterElementRequest]) (*connect.Response[v1.ConfigureVirtualRouterElementResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.ConfigureVirtualRouterElement is not implemented"))
}

func (UnimplementedRouterServiceHandler) CreateVirtualRouterElement(context.Context, *connect.Request[v1.CreateVirtualRouterElementRequest]) (*connect.Response[v1.CreateVirtualRouterElementResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.CreateVirtualRouterElement is not implemented"))
}

func (UnimplementedRouterServiceHandler) DestroyRouter(context.Context, *connect.Request[v1.DestroyRouterRequest]) (*connect.Response[v1.DestroyRouterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.DestroyRouter is not implemented"))
}

func (UnimplementedRouterServiceHandler) GetRouterHealthCheckResults(context.Context, *connect.Request[v1.GetRouterHealthCheckResultsRequest]) (*connect.Response[v1.GetRouterHealthCheckResultsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.GetRouterHealthCheckResults is not implemented"))
}

func (UnimplementedRouterServiceHandler) ListOvsElements(context.Context, *connect.Request[v1.ListOvsElementsRequest]) (*connect.Response[v1.ListOvsElementsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.ListOvsElements is not implemented"))
}

func (UnimplementedRouterServiceHandler) ListRouters(context.Context, *connect.Request[v1.ListRoutersRequest]) (*connect.Response[v1.ListRoutersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.ListRouters is not implemented"))
}

func (UnimplementedRouterServiceHandler) ListVirtualRouterElements(context.Context, *connect.Request[v1.ListVirtualRouterElementsRequest]) (*connect.Response[v1.ListVirtualRouterElementsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.ListVirtualRouterElements is not implemented"))
}

func (UnimplementedRouterServiceHandler) RebootRouter(context.Context, *connect.Request[v1.RebootRouterRequest]) (*connect.Response[v1.RebootRouterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.RebootRouter is not implemented"))
}

func (UnimplementedRouterServiceHandler) StartRouter(context.Context, *connect.Request[v1.StartRouterRequest]) (*connect.Response[v1.StartRouterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.StartRouter is not implemented"))
}

func (UnimplementedRouterServiceHandler) StopRouter(context.Context, *connect.Request[v1.StopRouterRequest]) (*connect.Response[v1.StopRouterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.StopRouter is not implemented"))
}

func (UnimplementedRouterServiceHandler) UpgradeRouter(context.Context, *connect.Request[v1.UpgradeRouterRequest]) (*connect.Response[v1.UpgradeRouterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.UpgradeRouter is not implemented"))
}

func (UnimplementedRouterServiceHandler) UpgradeRouterTemplate(context.Context, *connect.Request[v1.UpgradeRouterTemplateRequest]) (*connect.Response[v1.UpgradeRouterTemplateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.router.v1.RouterService.UpgradeRouterTemplate is not implemented"))
}
