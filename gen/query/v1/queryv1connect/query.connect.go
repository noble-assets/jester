// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: query/v1/query.proto

package queryv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/noble-assets/jester/gen/query/v1"
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
	// QueryServiceName is the fully-qualified name of the QueryService service.
	QueryServiceName = "query.v1.QueryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QueryServiceGetVaasProcedure is the fully-qualified name of the QueryService's GetVaas RPC.
	QueryServiceGetVaasProcedure = "/query.v1.QueryService/GetVaas"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	queryServiceServiceDescriptor       = v1.File_query_v1_query_proto.Services().ByName("QueryService")
	queryServiceGetVaasMethodDescriptor = queryServiceServiceDescriptor.Methods().ByName("GetVaas")
)

// QueryServiceClient is a client for the query.v1.QueryService service.
type QueryServiceClient interface {
	GetVaas(context.Context, *connect.Request[v1.GetVaasRequest]) (*connect.Response[v1.GetVaasResponse], error)
}

// NewQueryServiceClient constructs a client for the query.v1.QueryService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQueryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) QueryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &queryServiceClient{
		getVaas: connect.NewClient[v1.GetVaasRequest, v1.GetVaasResponse](
			httpClient,
			baseURL+QueryServiceGetVaasProcedure,
			connect.WithSchema(queryServiceGetVaasMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// queryServiceClient implements QueryServiceClient.
type queryServiceClient struct {
	getVaas *connect.Client[v1.GetVaasRequest, v1.GetVaasResponse]
}

// GetVaas calls query.v1.QueryService.GetVaas.
func (c *queryServiceClient) GetVaas(ctx context.Context, req *connect.Request[v1.GetVaasRequest]) (*connect.Response[v1.GetVaasResponse], error) {
	return c.getVaas.CallUnary(ctx, req)
}

// QueryServiceHandler is an implementation of the query.v1.QueryService service.
type QueryServiceHandler interface {
	GetVaas(context.Context, *connect.Request[v1.GetVaasRequest]) (*connect.Response[v1.GetVaasResponse], error)
}

// NewQueryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueryServiceHandler(svc QueryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queryServiceGetVaasHandler := connect.NewUnaryHandler(
		QueryServiceGetVaasProcedure,
		svc.GetVaas,
		connect.WithSchema(queryServiceGetVaasMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/query.v1.QueryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueryServiceGetVaasProcedure:
			queryServiceGetVaasHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQueryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQueryServiceHandler struct{}

func (UnimplementedQueryServiceHandler) GetVaas(context.Context, *connect.Request[v1.GetVaasRequest]) (*connect.Response[v1.GetVaasResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("query.v1.QueryService.GetVaas is not implemented"))
}
