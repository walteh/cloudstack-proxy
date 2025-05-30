// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cloudstack/management/bucket/v1/bucket.gen.proto

package bucketv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/walteh/cloudstack-proxy/gen/proto/golang/cloudstack/management/bucket/v1"
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
	// BucketServiceName is the fully-qualified name of the BucketService service.
	BucketServiceName = "cloudstack.management.bucket.v1.BucketService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BucketServiceCreateBucketProcedure is the fully-qualified name of the BucketService's
	// CreateBucket RPC.
	BucketServiceCreateBucketProcedure = "/cloudstack.management.bucket.v1.BucketService/CreateBucket"
	// BucketServiceDeleteBucketProcedure is the fully-qualified name of the BucketService's
	// DeleteBucket RPC.
	BucketServiceDeleteBucketProcedure = "/cloudstack.management.bucket.v1.BucketService/DeleteBucket"
	// BucketServiceListBucketsProcedure is the fully-qualified name of the BucketService's ListBuckets
	// RPC.
	BucketServiceListBucketsProcedure = "/cloudstack.management.bucket.v1.BucketService/ListBuckets"
	// BucketServiceUpdateBucketProcedure is the fully-qualified name of the BucketService's
	// UpdateBucket RPC.
	BucketServiceUpdateBucketProcedure = "/cloudstack.management.bucket.v1.BucketService/UpdateBucket"
)

// BucketServiceClient is a client for the cloudstack.management.bucket.v1.BucketService service.
type BucketServiceClient interface {
	// CreateBucket Creates a bucket in the specified object storage pool.
	CreateBucket(context.Context, *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error)
	// DeleteBucket Deletes an empty Bucket.
	DeleteBucket(context.Context, *connect.Request[v1.DeleteBucketRequest]) (*connect.Response[v1.DeleteBucketResponse], error)
	// ListBuckets Lists all Buckets.
	ListBuckets(context.Context, *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error)
	// UpdateBucket Updates Bucket properties
	UpdateBucket(context.Context, *connect.Request[v1.UpdateBucketRequest]) (*connect.Response[v1.UpdateBucketResponse], error)
}

// NewBucketServiceClient constructs a client for the cloudstack.management.bucket.v1.BucketService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBucketServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BucketServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	bucketServiceMethods := v1.File_cloudstack_management_bucket_v1_bucket_gen_proto.Services().ByName("BucketService").Methods()
	return &bucketServiceClient{
		createBucket: connect.NewClient[v1.CreateBucketRequest, v1.CreateBucketResponse](
			httpClient,
			baseURL+BucketServiceCreateBucketProcedure,
			connect.WithSchema(bucketServiceMethods.ByName("CreateBucket")),
			connect.WithClientOptions(opts...),
		),
		deleteBucket: connect.NewClient[v1.DeleteBucketRequest, v1.DeleteBucketResponse](
			httpClient,
			baseURL+BucketServiceDeleteBucketProcedure,
			connect.WithSchema(bucketServiceMethods.ByName("DeleteBucket")),
			connect.WithClientOptions(opts...),
		),
		listBuckets: connect.NewClient[v1.ListBucketsRequest, v1.ListBucketsResponse](
			httpClient,
			baseURL+BucketServiceListBucketsProcedure,
			connect.WithSchema(bucketServiceMethods.ByName("ListBuckets")),
			connect.WithClientOptions(opts...),
		),
		updateBucket: connect.NewClient[v1.UpdateBucketRequest, v1.UpdateBucketResponse](
			httpClient,
			baseURL+BucketServiceUpdateBucketProcedure,
			connect.WithSchema(bucketServiceMethods.ByName("UpdateBucket")),
			connect.WithClientOptions(opts...),
		),
	}
}

// bucketServiceClient implements BucketServiceClient.
type bucketServiceClient struct {
	createBucket *connect.Client[v1.CreateBucketRequest, v1.CreateBucketResponse]
	deleteBucket *connect.Client[v1.DeleteBucketRequest, v1.DeleteBucketResponse]
	listBuckets  *connect.Client[v1.ListBucketsRequest, v1.ListBucketsResponse]
	updateBucket *connect.Client[v1.UpdateBucketRequest, v1.UpdateBucketResponse]
}

// CreateBucket calls cloudstack.management.bucket.v1.BucketService.CreateBucket.
func (c *bucketServiceClient) CreateBucket(ctx context.Context, req *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error) {
	return c.createBucket.CallUnary(ctx, req)
}

// DeleteBucket calls cloudstack.management.bucket.v1.BucketService.DeleteBucket.
func (c *bucketServiceClient) DeleteBucket(ctx context.Context, req *connect.Request[v1.DeleteBucketRequest]) (*connect.Response[v1.DeleteBucketResponse], error) {
	return c.deleteBucket.CallUnary(ctx, req)
}

// ListBuckets calls cloudstack.management.bucket.v1.BucketService.ListBuckets.
func (c *bucketServiceClient) ListBuckets(ctx context.Context, req *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error) {
	return c.listBuckets.CallUnary(ctx, req)
}

// UpdateBucket calls cloudstack.management.bucket.v1.BucketService.UpdateBucket.
func (c *bucketServiceClient) UpdateBucket(ctx context.Context, req *connect.Request[v1.UpdateBucketRequest]) (*connect.Response[v1.UpdateBucketResponse], error) {
	return c.updateBucket.CallUnary(ctx, req)
}

// BucketServiceHandler is an implementation of the cloudstack.management.bucket.v1.BucketService
// service.
type BucketServiceHandler interface {
	// CreateBucket Creates a bucket in the specified object storage pool.
	CreateBucket(context.Context, *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error)
	// DeleteBucket Deletes an empty Bucket.
	DeleteBucket(context.Context, *connect.Request[v1.DeleteBucketRequest]) (*connect.Response[v1.DeleteBucketResponse], error)
	// ListBuckets Lists all Buckets.
	ListBuckets(context.Context, *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error)
	// UpdateBucket Updates Bucket properties
	UpdateBucket(context.Context, *connect.Request[v1.UpdateBucketRequest]) (*connect.Response[v1.UpdateBucketResponse], error)
}

// NewBucketServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBucketServiceHandler(svc BucketServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	bucketServiceMethods := v1.File_cloudstack_management_bucket_v1_bucket_gen_proto.Services().ByName("BucketService").Methods()
	bucketServiceCreateBucketHandler := connect.NewUnaryHandler(
		BucketServiceCreateBucketProcedure,
		svc.CreateBucket,
		connect.WithSchema(bucketServiceMethods.ByName("CreateBucket")),
		connect.WithHandlerOptions(opts...),
	)
	bucketServiceDeleteBucketHandler := connect.NewUnaryHandler(
		BucketServiceDeleteBucketProcedure,
		svc.DeleteBucket,
		connect.WithSchema(bucketServiceMethods.ByName("DeleteBucket")),
		connect.WithHandlerOptions(opts...),
	)
	bucketServiceListBucketsHandler := connect.NewUnaryHandler(
		BucketServiceListBucketsProcedure,
		svc.ListBuckets,
		connect.WithSchema(bucketServiceMethods.ByName("ListBuckets")),
		connect.WithHandlerOptions(opts...),
	)
	bucketServiceUpdateBucketHandler := connect.NewUnaryHandler(
		BucketServiceUpdateBucketProcedure,
		svc.UpdateBucket,
		connect.WithSchema(bucketServiceMethods.ByName("UpdateBucket")),
		connect.WithHandlerOptions(opts...),
	)
	return "/cloudstack.management.bucket.v1.BucketService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BucketServiceCreateBucketProcedure:
			bucketServiceCreateBucketHandler.ServeHTTP(w, r)
		case BucketServiceDeleteBucketProcedure:
			bucketServiceDeleteBucketHandler.ServeHTTP(w, r)
		case BucketServiceListBucketsProcedure:
			bucketServiceListBucketsHandler.ServeHTTP(w, r)
		case BucketServiceUpdateBucketProcedure:
			bucketServiceUpdateBucketHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBucketServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBucketServiceHandler struct{}

func (UnimplementedBucketServiceHandler) CreateBucket(context.Context, *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.bucket.v1.BucketService.CreateBucket is not implemented"))
}

func (UnimplementedBucketServiceHandler) DeleteBucket(context.Context, *connect.Request[v1.DeleteBucketRequest]) (*connect.Response[v1.DeleteBucketResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.bucket.v1.BucketService.DeleteBucket is not implemented"))
}

func (UnimplementedBucketServiceHandler) ListBuckets(context.Context, *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.bucket.v1.BucketService.ListBuckets is not implemented"))
}

func (UnimplementedBucketServiceHandler) UpdateBucket(context.Context, *connect.Request[v1.UpdateBucketRequest]) (*connect.Response[v1.UpdateBucketResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cloudstack.management.bucket.v1.BucketService.UpdateBucket is not implemented"))
}
