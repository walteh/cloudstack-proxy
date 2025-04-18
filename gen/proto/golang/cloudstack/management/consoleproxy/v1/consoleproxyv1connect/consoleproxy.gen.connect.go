// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/management/consoleproxy/v1/consoleproxy.gen.proto

package consoleproxyv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/consoleproxy/v1"
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
	// ConsoleproxyServiceName is the fully-qualified name of the ConsoleproxyService service.
	ConsoleproxyServiceName = "cloudstack.management.consoleproxy.v1.ConsoleproxyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ConsoleproxyServiceCreateConsoleEndpointProcedure is the fully-qualified name of the
	// ConsoleproxyService's CreateConsoleEndpoint RPC.
	ConsoleproxyServiceCreateConsoleEndpointProcedure = "/cloudstack.management.consoleproxy.v1.ConsoleproxyService/CreateConsoleEndpoint"
)

// ConsoleproxyServiceClient is a client for the
// cloudstack.management.consoleproxy.v1.ConsoleproxyService service.
type ConsoleproxyServiceClient interface {
	// CreateConsoleEndpoint Create a console endpoint to connect to a VM console
	CreateConsoleEndpoint(context.Context, *connect.Request[v1.CreateConsoleEndpointRequest]) (*connect.Response[v1.CreateConsoleEndpointResponse], error)
}

// NewConsoleproxyServiceClient constructs a client for the
// cloudstack.management.consoleproxy.v1.ConsoleproxyService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewConsoleproxyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ConsoleproxyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	consoleproxyServiceMethods := v1.File_cloudstack_management_consoleproxy_v1_consoleproxy_gen_proto.Services().ByName("ConsoleproxyService").Methods()
	return &consoleproxyServiceClient{
		createConsoleEndpoint: connect.NewClient[v1.CreateConsoleEndpointRequest, v1.CreateConsoleEndpointResponse](
			httpClient,
			baseURL+ConsoleproxyServiceCreateConsoleEndpointProcedure,
			connect.WithSchema(consoleproxyServiceMethods.ByName("CreateConsoleEndpoint")),
			connect.WithClientOptions(opts...),
		),
	}
}

// consoleproxyServiceClient implements ConsoleproxyServiceClient.
type consoleproxyServiceClient struct {
	createConsoleEndpoint *connect.Client[v1.CreateConsoleEndpointRequest, v1.CreateConsoleEndpointResponse]
}

// CreateConsoleEndpoint calls
// cloudstack.management.consoleproxy.v1.ConsoleproxyService.CreateConsoleEndpoint.
func (c *consoleproxyServiceClient) CreateConsoleEndpoint(ctx context.Context, req *connect.Request[v1.CreateConsoleEndpointRequest]) (*connect.Response[v1.CreateConsoleEndpointResponse], error) {
	return c.createConsoleEndpoint.CallUnary(ctx, req)
}

// ConsoleproxyServiceHandler is an implementation of the
// cloudstack.management.consoleproxy.v1.ConsoleproxyService service.
type ConsoleproxyServiceHandler interface {
	// CreateConsoleEndpoint Create a console endpoint to connect to a VM console
	CreateConsoleEndpoint(context.Context, *connect.Request[v1.CreateConsoleEndpointRequest]) (*connect.Response[v1.CreateConsoleEndpointResponse], error)
}

// NewConsoleproxyServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewConsoleproxyServiceHandler(svc ConsoleproxyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	consoleproxyServiceMethods := v1.File_cloudstack_management_consoleproxy_v1_consoleproxy_gen_proto.Services().ByName("ConsoleproxyService").Methods()
	consoleproxyServiceCreateConsoleEndpointHandler := connect.NewUnaryHandler(
		ConsoleproxyServiceCreateConsoleEndpointProcedure,
		svc.CreateConsoleEndpoint,
		connect.WithSchema(consoleproxyServiceMethods.ByName("CreateConsoleEndpoint")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.management.consoleproxy.v1.ConsoleproxyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ConsoleproxyServiceCreateConsoleEndpointProcedure:
			consoleproxyServiceCreateConsoleEndpointHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedConsoleproxyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedConsoleproxyServiceHandler struct{}

func (UnimplementedConsoleproxyServiceHandler) CreateConsoleEndpoint(context.Context, *connect.Request[v1.CreateConsoleEndpointRequest]) (*connect.Response[v1.CreateConsoleEndpointResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.consoleproxy.v1.ConsoleproxyService.CreateConsoleEndpoint is not implemented"))
}
