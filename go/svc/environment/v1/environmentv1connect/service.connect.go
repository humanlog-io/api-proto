// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/environment/v1/service.proto

package environmentv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/humanlogio/api/go/svc/environment/v1"
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
	// EnvironmentServiceName is the fully-qualified name of the EnvironmentService service.
	EnvironmentServiceName = "svc.environment.v1.EnvironmentService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EnvironmentServiceListMachineProcedure is the fully-qualified name of the EnvironmentService's
	// ListMachine RPC.
	EnvironmentServiceListMachineProcedure = "/svc.environment.v1.EnvironmentService/ListMachine"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	environmentServiceServiceDescriptor           = v1.File_svc_environment_v1_service_proto.Services().ByName("EnvironmentService")
	environmentServiceListMachineMethodDescriptor = environmentServiceServiceDescriptor.Methods().ByName("ListMachine")
)

// EnvironmentServiceClient is a client for the svc.environment.v1.EnvironmentService service.
type EnvironmentServiceClient interface {
	ListMachine(context.Context, *connect.Request[v1.ListMachineRequest]) (*connect.Response[v1.ListMachineResponse], error)
}

// NewEnvironmentServiceClient constructs a client for the svc.environment.v1.EnvironmentService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEnvironmentServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) EnvironmentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &environmentServiceClient{
		listMachine: connect.NewClient[v1.ListMachineRequest, v1.ListMachineResponse](
			httpClient,
			baseURL+EnvironmentServiceListMachineProcedure,
			connect.WithSchema(environmentServiceListMachineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// environmentServiceClient implements EnvironmentServiceClient.
type environmentServiceClient struct {
	listMachine *connect.Client[v1.ListMachineRequest, v1.ListMachineResponse]
}

// ListMachine calls svc.environment.v1.EnvironmentService.ListMachine.
func (c *environmentServiceClient) ListMachine(ctx context.Context, req *connect.Request[v1.ListMachineRequest]) (*connect.Response[v1.ListMachineResponse], error) {
	return c.listMachine.CallUnary(ctx, req)
}

// EnvironmentServiceHandler is an implementation of the svc.environment.v1.EnvironmentService
// service.
type EnvironmentServiceHandler interface {
	ListMachine(context.Context, *connect.Request[v1.ListMachineRequest]) (*connect.Response[v1.ListMachineResponse], error)
}

// NewEnvironmentServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEnvironmentServiceHandler(svc EnvironmentServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	environmentServiceListMachineHandler := connect.NewUnaryHandler(
		EnvironmentServiceListMachineProcedure,
		svc.ListMachine,
		connect.WithSchema(environmentServiceListMachineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/svc.environment.v1.EnvironmentService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EnvironmentServiceListMachineProcedure:
			environmentServiceListMachineHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEnvironmentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEnvironmentServiceHandler struct{}

func (UnimplementedEnvironmentServiceHandler) ListMachine(context.Context, *connect.Request[v1.ListMachineRequest]) (*connect.Response[v1.ListMachineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.environment.v1.EnvironmentService.ListMachine is not implemented"))
}