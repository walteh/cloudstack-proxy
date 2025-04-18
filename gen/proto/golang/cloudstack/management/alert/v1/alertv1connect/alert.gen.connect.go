// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/management/alert/v1/alert.gen.proto

package alertv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/alert/v1"
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
	// AlertServiceName is the fully-qualified name of the AlertService service.
	AlertServiceName = "cloudstack.management.alert.v1.AlertService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AlertServiceGenerateAlertProcedure is the fully-qualified name of the AlertService's
	// GenerateAlert RPC.
	AlertServiceGenerateAlertProcedure = "/cloudstack.management.alert.v1.AlertService/GenerateAlert"
)

// AlertServiceClient is a client for the cloudstack.management.alert.v1.AlertService service.
type AlertServiceClient interface {
	// GenerateAlert Generates an alert
	GenerateAlert(context.Context, *connect.Request[v1.GenerateAlertRequest]) (*connect.Response[v1.GenerateAlertResponse], error)
}

// NewAlertServiceClient constructs a client for the cloudstack.management.alert.v1.AlertService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAlertServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AlertServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	alertServiceMethods := v1.File_cloudstack_management_alert_v1_alert_gen_proto.Services().ByName("AlertService").Methods()
	return &alertServiceClient{
		generateAlert: connect.NewClient[v1.GenerateAlertRequest, v1.GenerateAlertResponse](
			httpClient,
			baseURL+AlertServiceGenerateAlertProcedure,
			connect.WithSchema(alertServiceMethods.ByName("GenerateAlert")),
			connect.WithClientOptions(opts...),
		),
	}
}

// alertServiceClient implements AlertServiceClient.
type alertServiceClient struct {
	generateAlert *connect.Client[v1.GenerateAlertRequest, v1.GenerateAlertResponse]
}

// GenerateAlert calls cloudstack.management.alert.v1.AlertService.GenerateAlert.
func (c *alertServiceClient) GenerateAlert(ctx context.Context, req *connect.Request[v1.GenerateAlertRequest]) (*connect.Response[v1.GenerateAlertResponse], error) {
	return c.generateAlert.CallUnary(ctx, req)
}

// AlertServiceHandler is an implementation of the cloudstack.management.alert.v1.AlertService
// service.
type AlertServiceHandler interface {
	// GenerateAlert Generates an alert
	GenerateAlert(context.Context, *connect.Request[v1.GenerateAlertRequest]) (*connect.Response[v1.GenerateAlertResponse], error)
}

// NewAlertServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAlertServiceHandler(svc AlertServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	alertServiceMethods := v1.File_cloudstack_management_alert_v1_alert_gen_proto.Services().ByName("AlertService").Methods()
	alertServiceGenerateAlertHandler := connect.NewUnaryHandler(
		AlertServiceGenerateAlertProcedure,
		svc.GenerateAlert,
		connect.WithSchema(alertServiceMethods.ByName("GenerateAlert")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.management.alert.v1.AlertService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AlertServiceGenerateAlertProcedure:
			alertServiceGenerateAlertHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAlertServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAlertServiceHandler struct{}

func (UnimplementedAlertServiceHandler) GenerateAlert(context.Context, *connect.Request[v1.GenerateAlertRequest]) (*connect.Response[v1.GenerateAlertResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.alert.v1.AlertService.GenerateAlert is not implemented"))
}
