// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/common/core/v1/health.proto

package corev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/common/core/v1"
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
	// HealthServiceName is the fully-qualified name of the HealthService service.
	HealthServiceName = "cloudstack.core.v1.HealthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthServiceCheckProcedure is the fully-qualified name of the HealthService's Check RPC.
	HealthServiceCheckProcedure = "/cloudstack.core.v1.HealthService/Check"
)

// HealthServiceClient is a client for the cloudstack.core.v1.HealthService service.
type HealthServiceClient interface {
	// Check returns the current health status of the service
	Check(context.Context, *connect.Request[v1.HealthCheckRequest]) (*connect.Response[v1.HealthCheckResponse], error)
}

// NewHealthServiceClient constructs a client for the cloudstack.core.v1.HealthService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) HealthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	healthServiceMethods := v1.File_cloudstack_common_core_v1_health_proto.Services().ByName("HealthService").Methods()
	return &healthServiceClient{
		check: connect.NewClient[v1.HealthCheckRequest, v1.HealthCheckResponse](
			httpClient,
			baseURL+HealthServiceCheckProcedure,
			connect.WithSchema(healthServiceMethods.ByName("Check")),
			connect.WithClientOptions(opts...),
		),
	}
}

// healthServiceClient implements HealthServiceClient.
type healthServiceClient struct {
	check *connect.Client[v1.HealthCheckRequest, v1.HealthCheckResponse]
}

// Check calls cloudstack.core.v1.HealthService.Check.
func (c *healthServiceClient) Check(ctx context.Context, req *connect.Request[v1.HealthCheckRequest]) (*connect.Response[v1.HealthCheckResponse], error) {
	return c.check.CallUnary(ctx, req)
}

// HealthServiceHandler is an implementation of the cloudstack.core.v1.HealthService service.
type HealthServiceHandler interface {
	// Check returns the current health status of the service
	Check(context.Context, *connect.Request[v1.HealthCheckRequest]) (*connect.Response[v1.HealthCheckResponse], error)
}

// NewHealthServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthServiceHandler(svc HealthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	healthServiceMethods := v1.File_cloudstack_common_core_v1_health_proto.Services().ByName("HealthService").Methods()
	healthServiceCheckHandler := connect.NewUnaryHandler(
		HealthServiceCheckProcedure,
		svc.Check,
		connect.WithSchema(healthServiceMethods.ByName("Check")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.core.v1.HealthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthServiceCheckProcedure:
			healthServiceCheckHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHealthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthServiceHandler struct{}

func (UnimplementedHealthServiceHandler) Check(context.Context, *connect.Request[v1.HealthCheckRequest]) (*connect.Response[v1.HealthCheckResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.core.v1.HealthService.Check is not implemented"))
}
