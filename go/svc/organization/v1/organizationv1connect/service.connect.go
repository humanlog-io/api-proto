// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/organization/v1/service.proto

package organizationv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/humanlogio/api/go/svc/organization/v1"
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
	// OrganizationServiceName is the fully-qualified name of the OrganizationService service.
	OrganizationServiceName = "svc.organization.v1.OrganizationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OrganizationServiceCreateEnvironmentProcedure is the fully-qualified name of the
	// OrganizationService's CreateEnvironment RPC.
	OrganizationServiceCreateEnvironmentProcedure = "/svc.organization.v1.OrganizationService/CreateEnvironment"
	// OrganizationServiceListEnvironmentProcedure is the fully-qualified name of the
	// OrganizationService's ListEnvironment RPC.
	OrganizationServiceListEnvironmentProcedure = "/svc.organization.v1.OrganizationService/ListEnvironment"
	// OrganizationServiceListUserProcedure is the fully-qualified name of the OrganizationService's
	// ListUser RPC.
	OrganizationServiceListUserProcedure = "/svc.organization.v1.OrganizationService/ListUser"
	// OrganizationServiceInviteUserProcedure is the fully-qualified name of the OrganizationService's
	// InviteUser RPC.
	OrganizationServiceInviteUserProcedure = "/svc.organization.v1.OrganizationService/InviteUser"
	// OrganizationServiceRevokeUserProcedure is the fully-qualified name of the OrganizationService's
	// RevokeUser RPC.
	OrganizationServiceRevokeUserProcedure = "/svc.organization.v1.OrganizationService/RevokeUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	organizationServiceServiceDescriptor                 = v1.File_svc_organization_v1_service_proto.Services().ByName("OrganizationService")
	organizationServiceCreateEnvironmentMethodDescriptor = organizationServiceServiceDescriptor.Methods().ByName("CreateEnvironment")
	organizationServiceListEnvironmentMethodDescriptor   = organizationServiceServiceDescriptor.Methods().ByName("ListEnvironment")
	organizationServiceListUserMethodDescriptor          = organizationServiceServiceDescriptor.Methods().ByName("ListUser")
	organizationServiceInviteUserMethodDescriptor        = organizationServiceServiceDescriptor.Methods().ByName("InviteUser")
	organizationServiceRevokeUserMethodDescriptor        = organizationServiceServiceDescriptor.Methods().ByName("RevokeUser")
)

// OrganizationServiceClient is a client for the svc.organization.v1.OrganizationService service.
type OrganizationServiceClient interface {
	CreateEnvironment(context.Context, *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error)
	ListEnvironment(context.Context, *connect.Request[v1.ListEnvironmentRequest]) (*connect.Response[v1.ListEnvironmentResponse], error)
	ListUser(context.Context, *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error)
	InviteUser(context.Context, *connect.Request[v1.InviteUserRequest]) (*connect.Response[v1.InviteUserResponse], error)
	RevokeUser(context.Context, *connect.Request[v1.RevokeUserRequest]) (*connect.Response[v1.RevokeUserResponse], error)
}

// NewOrganizationServiceClient constructs a client for the svc.organization.v1.OrganizationService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrganizationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OrganizationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &organizationServiceClient{
		createEnvironment: connect.NewClient[v1.CreateEnvironmentRequest, v1.CreateEnvironmentResponse](
			httpClient,
			baseURL+OrganizationServiceCreateEnvironmentProcedure,
			connect.WithSchema(organizationServiceCreateEnvironmentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listEnvironment: connect.NewClient[v1.ListEnvironmentRequest, v1.ListEnvironmentResponse](
			httpClient,
			baseURL+OrganizationServiceListEnvironmentProcedure,
			connect.WithSchema(organizationServiceListEnvironmentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listUser: connect.NewClient[v1.ListUserRequest, v1.ListUserResponse](
			httpClient,
			baseURL+OrganizationServiceListUserProcedure,
			connect.WithSchema(organizationServiceListUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		inviteUser: connect.NewClient[v1.InviteUserRequest, v1.InviteUserResponse](
			httpClient,
			baseURL+OrganizationServiceInviteUserProcedure,
			connect.WithSchema(organizationServiceInviteUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		revokeUser: connect.NewClient[v1.RevokeUserRequest, v1.RevokeUserResponse](
			httpClient,
			baseURL+OrganizationServiceRevokeUserProcedure,
			connect.WithSchema(organizationServiceRevokeUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// organizationServiceClient implements OrganizationServiceClient.
type organizationServiceClient struct {
	createEnvironment *connect.Client[v1.CreateEnvironmentRequest, v1.CreateEnvironmentResponse]
	listEnvironment   *connect.Client[v1.ListEnvironmentRequest, v1.ListEnvironmentResponse]
	listUser          *connect.Client[v1.ListUserRequest, v1.ListUserResponse]
	inviteUser        *connect.Client[v1.InviteUserRequest, v1.InviteUserResponse]
	revokeUser        *connect.Client[v1.RevokeUserRequest, v1.RevokeUserResponse]
}

// CreateEnvironment calls svc.organization.v1.OrganizationService.CreateEnvironment.
func (c *organizationServiceClient) CreateEnvironment(ctx context.Context, req *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error) {
	return c.createEnvironment.CallUnary(ctx, req)
}

// ListEnvironment calls svc.organization.v1.OrganizationService.ListEnvironment.
func (c *organizationServiceClient) ListEnvironment(ctx context.Context, req *connect.Request[v1.ListEnvironmentRequest]) (*connect.Response[v1.ListEnvironmentResponse], error) {
	return c.listEnvironment.CallUnary(ctx, req)
}

// ListUser calls svc.organization.v1.OrganizationService.ListUser.
func (c *organizationServiceClient) ListUser(ctx context.Context, req *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error) {
	return c.listUser.CallUnary(ctx, req)
}

// InviteUser calls svc.organization.v1.OrganizationService.InviteUser.
func (c *organizationServiceClient) InviteUser(ctx context.Context, req *connect.Request[v1.InviteUserRequest]) (*connect.Response[v1.InviteUserResponse], error) {
	return c.inviteUser.CallUnary(ctx, req)
}

// RevokeUser calls svc.organization.v1.OrganizationService.RevokeUser.
func (c *organizationServiceClient) RevokeUser(ctx context.Context, req *connect.Request[v1.RevokeUserRequest]) (*connect.Response[v1.RevokeUserResponse], error) {
	return c.revokeUser.CallUnary(ctx, req)
}

// OrganizationServiceHandler is an implementation of the svc.organization.v1.OrganizationService
// service.
type OrganizationServiceHandler interface {
	CreateEnvironment(context.Context, *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error)
	ListEnvironment(context.Context, *connect.Request[v1.ListEnvironmentRequest]) (*connect.Response[v1.ListEnvironmentResponse], error)
	ListUser(context.Context, *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error)
	InviteUser(context.Context, *connect.Request[v1.InviteUserRequest]) (*connect.Response[v1.InviteUserResponse], error)
	RevokeUser(context.Context, *connect.Request[v1.RevokeUserRequest]) (*connect.Response[v1.RevokeUserResponse], error)
}

// NewOrganizationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrganizationServiceHandler(svc OrganizationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	organizationServiceCreateEnvironmentHandler := connect.NewUnaryHandler(
		OrganizationServiceCreateEnvironmentProcedure,
		svc.CreateEnvironment,
		connect.WithSchema(organizationServiceCreateEnvironmentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationServiceListEnvironmentHandler := connect.NewUnaryHandler(
		OrganizationServiceListEnvironmentProcedure,
		svc.ListEnvironment,
		connect.WithSchema(organizationServiceListEnvironmentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationServiceListUserHandler := connect.NewUnaryHandler(
		OrganizationServiceListUserProcedure,
		svc.ListUser,
		connect.WithSchema(organizationServiceListUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationServiceInviteUserHandler := connect.NewUnaryHandler(
		OrganizationServiceInviteUserProcedure,
		svc.InviteUser,
		connect.WithSchema(organizationServiceInviteUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationServiceRevokeUserHandler := connect.NewUnaryHandler(
		OrganizationServiceRevokeUserProcedure,
		svc.RevokeUser,
		connect.WithSchema(organizationServiceRevokeUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/svc.organization.v1.OrganizationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OrganizationServiceCreateEnvironmentProcedure:
			organizationServiceCreateEnvironmentHandler.ServeHTTP(w, r)
		case OrganizationServiceListEnvironmentProcedure:
			organizationServiceListEnvironmentHandler.ServeHTTP(w, r)
		case OrganizationServiceListUserProcedure:
			organizationServiceListUserHandler.ServeHTTP(w, r)
		case OrganizationServiceInviteUserProcedure:
			organizationServiceInviteUserHandler.ServeHTTP(w, r)
		case OrganizationServiceRevokeUserProcedure:
			organizationServiceRevokeUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOrganizationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrganizationServiceHandler struct{}

func (UnimplementedOrganizationServiceHandler) CreateEnvironment(context.Context, *connect.Request[v1.CreateEnvironmentRequest]) (*connect.Response[v1.CreateEnvironmentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.organization.v1.OrganizationService.CreateEnvironment is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) ListEnvironment(context.Context, *connect.Request[v1.ListEnvironmentRequest]) (*connect.Response[v1.ListEnvironmentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.organization.v1.OrganizationService.ListEnvironment is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) ListUser(context.Context, *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.organization.v1.OrganizationService.ListUser is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) InviteUser(context.Context, *connect.Request[v1.InviteUserRequest]) (*connect.Response[v1.InviteUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.organization.v1.OrganizationService.InviteUser is not implemented"))
}

func (UnimplementedOrganizationServiceHandler) RevokeUser(context.Context, *connect.Request[v1.RevokeUserRequest]) (*connect.Response[v1.RevokeUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.organization.v1.OrganizationService.RevokeUser is not implemented"))
}
