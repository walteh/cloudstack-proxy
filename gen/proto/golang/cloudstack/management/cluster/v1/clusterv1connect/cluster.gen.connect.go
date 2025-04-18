// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/management/cluster/v1/cluster.gen.proto

package clusterv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/cluster/v1"
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
	// ClusterServiceName is the fully-qualified name of the ClusterService service.
	ClusterServiceName = "cloudstack.management.cluster.v1.ClusterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClusterServiceAddClusterProcedure is the fully-qualified name of the ClusterService's AddCluster
	// RPC.
	ClusterServiceAddClusterProcedure = "/cloudstack.management.cluster.v1.ClusterService/AddCluster"
	// ClusterServiceDeleteClusterProcedure is the fully-qualified name of the ClusterService's
	// DeleteCluster RPC.
	ClusterServiceDeleteClusterProcedure = "/cloudstack.management.cluster.v1.ClusterService/DeleteCluster"
	// ClusterServiceExecuteClusterDrsPlanProcedure is the fully-qualified name of the ClusterService's
	// ExecuteClusterDrsPlan RPC.
	ClusterServiceExecuteClusterDrsPlanProcedure = "/cloudstack.management.cluster.v1.ClusterService/ExecuteClusterDrsPlan"
	// ClusterServiceGenerateClusterDrsPlanProcedure is the fully-qualified name of the ClusterService's
	// GenerateClusterDrsPlan RPC.
	ClusterServiceGenerateClusterDrsPlanProcedure = "/cloudstack.management.cluster.v1.ClusterService/GenerateClusterDrsPlan"
	// ClusterServiceListClusterDrsPlanProcedure is the fully-qualified name of the ClusterService's
	// ListClusterDrsPlan RPC.
	ClusterServiceListClusterDrsPlanProcedure = "/cloudstack.management.cluster.v1.ClusterService/ListClusterDrsPlan"
	// ClusterServiceListClustersProcedure is the fully-qualified name of the ClusterService's
	// ListClusters RPC.
	ClusterServiceListClustersProcedure = "/cloudstack.management.cluster.v1.ClusterService/ListClusters"
	// ClusterServiceUpdateClusterProcedure is the fully-qualified name of the ClusterService's
	// UpdateCluster RPC.
	ClusterServiceUpdateClusterProcedure = "/cloudstack.management.cluster.v1.ClusterService/UpdateCluster"
)

// ClusterServiceClient is a client for the cloudstack.management.cluster.v1.ClusterService service.
type ClusterServiceClient interface {
	// AddCluster Adds a new cluster
	AddCluster(context.Context, *connect.Request[v1.AddClusterRequest]) (*connect.Response[v1.AddClusterResponse], error)
	// DeleteCluster Deletes a cluster.
	DeleteCluster(context.Context, *connect.Request[v1.DeleteClusterRequest]) (*connect.Response[v1.DeleteClusterResponse], error)
	// ExecuteClusterDrsPlan Execute DRS for a cluster. If there is another plan in progress for the same cluster, this command will fail.
	ExecuteClusterDrsPlan(context.Context, *connect.Request[v1.ExecuteClusterDrsPlanRequest]) (*connect.Response[v1.ExecuteClusterDrsPlanResponse], error)
	// GenerateClusterDrsPlan Generate DRS plan for a cluster
	GenerateClusterDrsPlan(context.Context, *connect.Request[v1.GenerateClusterDrsPlanRequest]) (*connect.Response[v1.GenerateClusterDrsPlanResponse], error)
	// ListClusterDrsPlan List DRS plans for a clusters
	ListClusterDrsPlan(context.Context, *connect.Request[v1.ListClusterDrsPlanRequest]) (*connect.Response[v1.ListClusterDrsPlanResponse], error)
	// ListClusters Lists clusters.
	ListClusters(context.Context, *connect.Request[v1.ListClustersRequest]) (*connect.Response[v1.ListClustersResponse], error)
	// UpdateCluster Updates an existing cluster
	UpdateCluster(context.Context, *connect.Request[v1.UpdateClusterRequest]) (*connect.Response[v1.UpdateClusterResponse], error)
}

// NewClusterServiceClient constructs a client for the
// cloudstack.management.cluster.v1.ClusterService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClusterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ClusterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	clusterServiceMethods := v1.File_cloudstack_management_cluster_v1_cluster_gen_proto.Services().ByName("ClusterService").Methods()
	return &clusterServiceClient{
		addCluster: connect.NewClient[v1.AddClusterRequest, v1.AddClusterResponse](
			httpClient,
			baseURL+ClusterServiceAddClusterProcedure,
			connect.WithSchema(clusterServiceMethods.ByName("AddCluster")),
			connect.WithClientOptions(opts...),
		),
		deleteCluster: connect.NewClient[v1.DeleteClusterRequest, v1.DeleteClusterResponse](
			httpClient,
			baseURL+ClusterServiceDeleteClusterProcedure,
			connect.WithSchema(clusterServiceMethods.ByName("DeleteCluster")),
			connect.WithClientOptions(opts...),
		),
		executeClusterDrsPlan: connect.NewClient[v1.ExecuteClusterDrsPlanRequest, v1.ExecuteClusterDrsPlanResponse](
			httpClient,
			baseURL+ClusterServiceExecuteClusterDrsPlanProcedure,
			connect.WithSchema(clusterServiceMethods.ByName("ExecuteClusterDrsPlan")),
			connect.WithClientOptions(opts...),
		),
		generateClusterDrsPlan: connect.NewClient[v1.GenerateClusterDrsPlanRequest, v1.GenerateClusterDrsPlanResponse](
			httpClient,
			baseURL+ClusterServiceGenerateClusterDrsPlanProcedure,
			connect.WithSchema(clusterServiceMethods.ByName("GenerateClusterDrsPlan")),
			connect.WithClientOptions(opts...),
		),
		listClusterDrsPlan: connect.NewClient[v1.ListClusterDrsPlanRequest, v1.ListClusterDrsPlanResponse](
			httpClient,
			baseURL+ClusterServiceListClusterDrsPlanProcedure,
			connect.WithSchema(clusterServiceMethods.ByName("ListClusterDrsPlan")),
			connect.WithClientOptions(opts...),
		),
		listClusters: connect.NewClient[v1.ListClustersRequest, v1.ListClustersResponse](
			httpClient,
			baseURL+ClusterServiceListClustersProcedure,
			connect.WithSchema(clusterServiceMethods.ByName("ListClusters")),
			connect.WithClientOptions(opts...),
		),
		updateCluster: connect.NewClient[v1.UpdateClusterRequest, v1.UpdateClusterResponse](
			httpClient,
			baseURL+ClusterServiceUpdateClusterProcedure,
			connect.WithSchema(clusterServiceMethods.ByName("UpdateCluster")),
			connect.WithClientOptions(opts...),
		),
	}
}

// clusterServiceClient implements ClusterServiceClient.
type clusterServiceClient struct {
	addCluster             *connect.Client[v1.AddClusterRequest, v1.AddClusterResponse]
	deleteCluster          *connect.Client[v1.DeleteClusterRequest, v1.DeleteClusterResponse]
	executeClusterDrsPlan  *connect.Client[v1.ExecuteClusterDrsPlanRequest, v1.ExecuteClusterDrsPlanResponse]
	generateClusterDrsPlan *connect.Client[v1.GenerateClusterDrsPlanRequest, v1.GenerateClusterDrsPlanResponse]
	listClusterDrsPlan     *connect.Client[v1.ListClusterDrsPlanRequest, v1.ListClusterDrsPlanResponse]
	listClusters           *connect.Client[v1.ListClustersRequest, v1.ListClustersResponse]
	updateCluster          *connect.Client[v1.UpdateClusterRequest, v1.UpdateClusterResponse]
}

// AddCluster calls cloudstack.management.cluster.v1.ClusterService.AddCluster.
func (c *clusterServiceClient) AddCluster(ctx context.Context, req *connect.Request[v1.AddClusterRequest]) (*connect.Response[v1.AddClusterResponse], error) {
	return c.addCluster.CallUnary(ctx, req)
}

// DeleteCluster calls cloudstack.management.cluster.v1.ClusterService.DeleteCluster.
func (c *clusterServiceClient) DeleteCluster(ctx context.Context, req *connect.Request[v1.DeleteClusterRequest]) (*connect.Response[v1.DeleteClusterResponse], error) {
	return c.deleteCluster.CallUnary(ctx, req)
}

// ExecuteClusterDrsPlan calls
// cloudstack.management.cluster.v1.ClusterService.ExecuteClusterDrsPlan.
func (c *clusterServiceClient) ExecuteClusterDrsPlan(ctx context.Context, req *connect.Request[v1.ExecuteClusterDrsPlanRequest]) (*connect.Response[v1.ExecuteClusterDrsPlanResponse], error) {
	return c.executeClusterDrsPlan.CallUnary(ctx, req)
}

// GenerateClusterDrsPlan calls
// cloudstack.management.cluster.v1.ClusterService.GenerateClusterDrsPlan.
func (c *clusterServiceClient) GenerateClusterDrsPlan(ctx context.Context, req *connect.Request[v1.GenerateClusterDrsPlanRequest]) (*connect.Response[v1.GenerateClusterDrsPlanResponse], error) {
	return c.generateClusterDrsPlan.CallUnary(ctx, req)
}

// ListClusterDrsPlan calls cloudstack.management.cluster.v1.ClusterService.ListClusterDrsPlan.
func (c *clusterServiceClient) ListClusterDrsPlan(ctx context.Context, req *connect.Request[v1.ListClusterDrsPlanRequest]) (*connect.Response[v1.ListClusterDrsPlanResponse], error) {
	return c.listClusterDrsPlan.CallUnary(ctx, req)
}

// ListClusters calls cloudstack.management.cluster.v1.ClusterService.ListClusters.
func (c *clusterServiceClient) ListClusters(ctx context.Context, req *connect.Request[v1.ListClustersRequest]) (*connect.Response[v1.ListClustersResponse], error) {
	return c.listClusters.CallUnary(ctx, req)
}

// UpdateCluster calls cloudstack.management.cluster.v1.ClusterService.UpdateCluster.
func (c *clusterServiceClient) UpdateCluster(ctx context.Context, req *connect.Request[v1.UpdateClusterRequest]) (*connect.Response[v1.UpdateClusterResponse], error) {
	return c.updateCluster.CallUnary(ctx, req)
}

// ClusterServiceHandler is an implementation of the cloudstack.management.cluster.v1.ClusterService
// service.
type ClusterServiceHandler interface {
	// AddCluster Adds a new cluster
	AddCluster(context.Context, *connect.Request[v1.AddClusterRequest]) (*connect.Response[v1.AddClusterResponse], error)
	// DeleteCluster Deletes a cluster.
	DeleteCluster(context.Context, *connect.Request[v1.DeleteClusterRequest]) (*connect.Response[v1.DeleteClusterResponse], error)
	// ExecuteClusterDrsPlan Execute DRS for a cluster. If there is another plan in progress for the same cluster, this command will fail.
	ExecuteClusterDrsPlan(context.Context, *connect.Request[v1.ExecuteClusterDrsPlanRequest]) (*connect.Response[v1.ExecuteClusterDrsPlanResponse], error)
	// GenerateClusterDrsPlan Generate DRS plan for a cluster
	GenerateClusterDrsPlan(context.Context, *connect.Request[v1.GenerateClusterDrsPlanRequest]) (*connect.Response[v1.GenerateClusterDrsPlanResponse], error)
	// ListClusterDrsPlan List DRS plans for a clusters
	ListClusterDrsPlan(context.Context, *connect.Request[v1.ListClusterDrsPlanRequest]) (*connect.Response[v1.ListClusterDrsPlanResponse], error)
	// ListClusters Lists clusters.
	ListClusters(context.Context, *connect.Request[v1.ListClustersRequest]) (*connect.Response[v1.ListClustersResponse], error)
	// UpdateCluster Updates an existing cluster
	UpdateCluster(context.Context, *connect.Request[v1.UpdateClusterRequest]) (*connect.Response[v1.UpdateClusterResponse], error)
}

// NewClusterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClusterServiceHandler(svc ClusterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	clusterServiceMethods := v1.File_cloudstack_management_cluster_v1_cluster_gen_proto.Services().ByName("ClusterService").Methods()
	clusterServiceAddClusterHandler := connect.NewUnaryHandler(
		ClusterServiceAddClusterProcedure,
		svc.AddCluster,
		connect.WithSchema(clusterServiceMethods.ByName("AddCluster")),
		connect.WithHandlerOptions(opts...),
	)
	clusterServiceDeleteClusterHandler := connect.NewUnaryHandler(
		ClusterServiceDeleteClusterProcedure,
		svc.DeleteCluster,
		connect.WithSchema(clusterServiceMethods.ByName("DeleteCluster")),
		connect.WithHandlerOptions(opts...),
	)
	clusterServiceExecuteClusterDrsPlanHandler := connect.NewUnaryHandler(
		ClusterServiceExecuteClusterDrsPlanProcedure,
		svc.ExecuteClusterDrsPlan,
		connect.WithSchema(clusterServiceMethods.ByName("ExecuteClusterDrsPlan")),
		connect.WithHandlerOptions(opts...),
	)
	clusterServiceGenerateClusterDrsPlanHandler := connect.NewUnaryHandler(
		ClusterServiceGenerateClusterDrsPlanProcedure,
		svc.GenerateClusterDrsPlan,
		connect.WithSchema(clusterServiceMethods.ByName("GenerateClusterDrsPlan")),
		connect.WithHandlerOptions(opts...),
	)
	clusterServiceListClusterDrsPlanHandler := connect.NewUnaryHandler(
		ClusterServiceListClusterDrsPlanProcedure,
		svc.ListClusterDrsPlan,
		connect.WithSchema(clusterServiceMethods.ByName("ListClusterDrsPlan")),
		connect.WithHandlerOptions(opts...),
	)
	clusterServiceListClustersHandler := connect.NewUnaryHandler(
		ClusterServiceListClustersProcedure,
		svc.ListClusters,
		connect.WithSchema(clusterServiceMethods.ByName("ListClusters")),
		connect.WithHandlerOptions(opts...),
	)
	clusterServiceUpdateClusterHandler := connect.NewUnaryHandler(
		ClusterServiceUpdateClusterProcedure,
		svc.UpdateCluster,
		connect.WithSchema(clusterServiceMethods.ByName("UpdateCluster")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.management.cluster.v1.ClusterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ClusterServiceAddClusterProcedure:
			clusterServiceAddClusterHandler.ServeHTTP(w, r)
		case ClusterServiceDeleteClusterProcedure:
			clusterServiceDeleteClusterHandler.ServeHTTP(w, r)
		case ClusterServiceExecuteClusterDrsPlanProcedure:
			clusterServiceExecuteClusterDrsPlanHandler.ServeHTTP(w, r)
		case ClusterServiceGenerateClusterDrsPlanProcedure:
			clusterServiceGenerateClusterDrsPlanHandler.ServeHTTP(w, r)
		case ClusterServiceListClusterDrsPlanProcedure:
			clusterServiceListClusterDrsPlanHandler.ServeHTTP(w, r)
		case ClusterServiceListClustersProcedure:
			clusterServiceListClustersHandler.ServeHTTP(w, r)
		case ClusterServiceUpdateClusterProcedure:
			clusterServiceUpdateClusterHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedClusterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClusterServiceHandler struct{}

func (UnimplementedClusterServiceHandler) AddCluster(context.Context, *connect.Request[v1.AddClusterRequest]) (*connect.Response[v1.AddClusterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.cluster.v1.ClusterService.AddCluster is not implemented"))
}

func (UnimplementedClusterServiceHandler) DeleteCluster(context.Context, *connect.Request[v1.DeleteClusterRequest]) (*connect.Response[v1.DeleteClusterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.cluster.v1.ClusterService.DeleteCluster is not implemented"))
}

func (UnimplementedClusterServiceHandler) ExecuteClusterDrsPlan(context.Context, *connect.Request[v1.ExecuteClusterDrsPlanRequest]) (*connect.Response[v1.ExecuteClusterDrsPlanResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.cluster.v1.ClusterService.ExecuteClusterDrsPlan is not implemented"))
}

func (UnimplementedClusterServiceHandler) GenerateClusterDrsPlan(context.Context, *connect.Request[v1.GenerateClusterDrsPlanRequest]) (*connect.Response[v1.GenerateClusterDrsPlanResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.cluster.v1.ClusterService.GenerateClusterDrsPlan is not implemented"))
}

func (UnimplementedClusterServiceHandler) ListClusterDrsPlan(context.Context, *connect.Request[v1.ListClusterDrsPlanRequest]) (*connect.Response[v1.ListClusterDrsPlanResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.cluster.v1.ClusterService.ListClusterDrsPlan is not implemented"))
}

func (UnimplementedClusterServiceHandler) ListClusters(context.Context, *connect.Request[v1.ListClustersRequest]) (*connect.Response[v1.ListClustersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.cluster.v1.ClusterService.ListClusters is not implemented"))
}

func (UnimplementedClusterServiceHandler) UpdateCluster(context.Context, *connect.Request[v1.UpdateClusterRequest]) (*connect.Response[v1.UpdateClusterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.cluster.v1.ClusterService.UpdateCluster is not implemented"))
}
