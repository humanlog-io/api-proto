// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/token/v1/service.proto

package tokenv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/humanlogio/api/go/svc/token/v1"
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
	// TokenServiceName is the fully-qualified name of the TokenService service.
	TokenServiceName = "svc.token.v1.TokenService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TokenServiceGenerateEnvironmentTokenProcedure is the fully-qualified name of the TokenService's
	// GenerateEnvironmentToken RPC.
	TokenServiceGenerateEnvironmentTokenProcedure = "/svc.token.v1.TokenService/GenerateEnvironmentToken"
	// TokenServiceRevokeEnvironmentTokenProcedure is the fully-qualified name of the TokenService's
	// RevokeEnvironmentToken RPC.
	TokenServiceRevokeEnvironmentTokenProcedure = "/svc.token.v1.TokenService/RevokeEnvironmentToken"
	// TokenServiceGetEnvironmentTokenProcedure is the fully-qualified name of the TokenService's
	// GetEnvironmentToken RPC.
	TokenServiceGetEnvironmentTokenProcedure = "/svc.token.v1.TokenService/GetEnvironmentToken"
	// TokenServiceListEnvironmentTokenProcedure is the fully-qualified name of the TokenService's
	// ListEnvironmentToken RPC.
	TokenServiceListEnvironmentTokenProcedure = "/svc.token.v1.TokenService/ListEnvironmentToken"
)

// TokenServiceClient is a client for the svc.token.v1.TokenService service.
type TokenServiceClient interface {
	GenerateEnvironmentToken(context.Context, *connect.Request[v1.GenerateEnvironmentTokenRequest]) (*connect.Response[v1.GenerateEnvironmentTokenResponse], error)
	RevokeEnvironmentToken(context.Context, *connect.Request[v1.RevokeEnvironmentTokenRequest]) (*connect.Response[v1.RevokeEnvironmentTokenResponse], error)
	GetEnvironmentToken(context.Context, *connect.Request[v1.GetEnvironmentTokenRequest]) (*connect.Response[v1.GetEnvironmentTokenResponse], error)
	ListEnvironmentToken(context.Context, *connect.Request[v1.ListEnvironmentTokenRequest]) (*connect.Response[v1.ListEnvironmentTokenResponse], error)
}

// NewTokenServiceClient constructs a client for the svc.token.v1.TokenService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTokenServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TokenServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	tokenServiceMethods := v1.File_svc_token_v1_service_proto.Services().ByName("TokenService").Methods()
	return &tokenServiceClient{
		generateEnvironmentToken: connect.NewClient[v1.GenerateEnvironmentTokenRequest, v1.GenerateEnvironmentTokenResponse](
			httpClient,
			baseURL+TokenServiceGenerateEnvironmentTokenProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("GenerateEnvironmentToken")),
			connect.WithClientOptions(opts...),
		),
		revokeEnvironmentToken: connect.NewClient[v1.RevokeEnvironmentTokenRequest, v1.RevokeEnvironmentTokenResponse](
			httpClient,
			baseURL+TokenServiceRevokeEnvironmentTokenProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("RevokeEnvironmentToken")),
			connect.WithClientOptions(opts...),
		),
		getEnvironmentToken: connect.NewClient[v1.GetEnvironmentTokenRequest, v1.GetEnvironmentTokenResponse](
			httpClient,
			baseURL+TokenServiceGetEnvironmentTokenProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("GetEnvironmentToken")),
			connect.WithClientOptions(opts...),
		),
		listEnvironmentToken: connect.NewClient[v1.ListEnvironmentTokenRequest, v1.ListEnvironmentTokenResponse](
			httpClient,
			baseURL+TokenServiceListEnvironmentTokenProcedure,
			connect.WithSchema(tokenServiceMethods.ByName("ListEnvironmentToken")),
			connect.WithClientOptions(opts...),
		),
	}
}

// tokenServiceClient implements TokenServiceClient.
type tokenServiceClient struct {
	generateEnvironmentToken *connect.Client[v1.GenerateEnvironmentTokenRequest, v1.GenerateEnvironmentTokenResponse]
	revokeEnvironmentToken   *connect.Client[v1.RevokeEnvironmentTokenRequest, v1.RevokeEnvironmentTokenResponse]
	getEnvironmentToken      *connect.Client[v1.GetEnvironmentTokenRequest, v1.GetEnvironmentTokenResponse]
	listEnvironmentToken     *connect.Client[v1.ListEnvironmentTokenRequest, v1.ListEnvironmentTokenResponse]
}

// GenerateEnvironmentToken calls svc.token.v1.TokenService.GenerateEnvironmentToken.
func (c *tokenServiceClient) GenerateEnvironmentToken(ctx context.Context, req *connect.Request[v1.GenerateEnvironmentTokenRequest]) (*connect.Response[v1.GenerateEnvironmentTokenResponse], error) {
	return c.generateEnvironmentToken.CallUnary(ctx, req)
}

// RevokeEnvironmentToken calls svc.token.v1.TokenService.RevokeEnvironmentToken.
func (c *tokenServiceClient) RevokeEnvironmentToken(ctx context.Context, req *connect.Request[v1.RevokeEnvironmentTokenRequest]) (*connect.Response[v1.RevokeEnvironmentTokenResponse], error) {
	return c.revokeEnvironmentToken.CallUnary(ctx, req)
}

// GetEnvironmentToken calls svc.token.v1.TokenService.GetEnvironmentToken.
func (c *tokenServiceClient) GetEnvironmentToken(ctx context.Context, req *connect.Request[v1.GetEnvironmentTokenRequest]) (*connect.Response[v1.GetEnvironmentTokenResponse], error) {
	return c.getEnvironmentToken.CallUnary(ctx, req)
}

// ListEnvironmentToken calls svc.token.v1.TokenService.ListEnvironmentToken.
func (c *tokenServiceClient) ListEnvironmentToken(ctx context.Context, req *connect.Request[v1.ListEnvironmentTokenRequest]) (*connect.Response[v1.ListEnvironmentTokenResponse], error) {
	return c.listEnvironmentToken.CallUnary(ctx, req)
}

// TokenServiceHandler is an implementation of the svc.token.v1.TokenService service.
type TokenServiceHandler interface {
	GenerateEnvironmentToken(context.Context, *connect.Request[v1.GenerateEnvironmentTokenRequest]) (*connect.Response[v1.GenerateEnvironmentTokenResponse], error)
	RevokeEnvironmentToken(context.Context, *connect.Request[v1.RevokeEnvironmentTokenRequest]) (*connect.Response[v1.RevokeEnvironmentTokenResponse], error)
	GetEnvironmentToken(context.Context, *connect.Request[v1.GetEnvironmentTokenRequest]) (*connect.Response[v1.GetEnvironmentTokenResponse], error)
	ListEnvironmentToken(context.Context, *connect.Request[v1.ListEnvironmentTokenRequest]) (*connect.Response[v1.ListEnvironmentTokenResponse], error)
}

// NewTokenServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTokenServiceHandler(svc TokenServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tokenServiceMethods := v1.File_svc_token_v1_service_proto.Services().ByName("TokenService").Methods()
	tokenServiceGenerateEnvironmentTokenHandler := connect.NewUnaryHandler(
		TokenServiceGenerateEnvironmentTokenProcedure,
		svc.GenerateEnvironmentToken,
		connect.WithSchema(tokenServiceMethods.ByName("GenerateEnvironmentToken")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceRevokeEnvironmentTokenHandler := connect.NewUnaryHandler(
		TokenServiceRevokeEnvironmentTokenProcedure,
		svc.RevokeEnvironmentToken,
		connect.WithSchema(tokenServiceMethods.ByName("RevokeEnvironmentToken")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceGetEnvironmentTokenHandler := connect.NewUnaryHandler(
		TokenServiceGetEnvironmentTokenProcedure,
		svc.GetEnvironmentToken,
		connect.WithSchema(tokenServiceMethods.ByName("GetEnvironmentToken")),
		connect.WithHandlerOptions(opts...),
	)
	tokenServiceListEnvironmentTokenHandler := connect.NewUnaryHandler(
		TokenServiceListEnvironmentTokenProcedure,
		svc.ListEnvironmentToken,
		connect.WithSchema(tokenServiceMethods.ByName("ListEnvironmentToken")),
		connect.WithHandlerOptions(opts...),
	)
	return "/svc.token.v1.TokenService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TokenServiceGenerateEnvironmentTokenProcedure:
			tokenServiceGenerateEnvironmentTokenHandler.ServeHTTP(w, r)
		case TokenServiceRevokeEnvironmentTokenProcedure:
			tokenServiceRevokeEnvironmentTokenHandler.ServeHTTP(w, r)
		case TokenServiceGetEnvironmentTokenProcedure:
			tokenServiceGetEnvironmentTokenHandler.ServeHTTP(w, r)
		case TokenServiceListEnvironmentTokenProcedure:
			tokenServiceListEnvironmentTokenHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTokenServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTokenServiceHandler struct{}

func (UnimplementedTokenServiceHandler) GenerateEnvironmentToken(context.Context, *connect.Request[v1.GenerateEnvironmentTokenRequest]) (*connect.Response[v1.GenerateEnvironmentTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.token.v1.TokenService.GenerateEnvironmentToken is not implemented"))
}

func (UnimplementedTokenServiceHandler) RevokeEnvironmentToken(context.Context, *connect.Request[v1.RevokeEnvironmentTokenRequest]) (*connect.Response[v1.RevokeEnvironmentTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.token.v1.TokenService.RevokeEnvironmentToken is not implemented"))
}

func (UnimplementedTokenServiceHandler) GetEnvironmentToken(context.Context, *connect.Request[v1.GetEnvironmentTokenRequest]) (*connect.Response[v1.GetEnvironmentTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.token.v1.TokenService.GetEnvironmentToken is not implemented"))
}

func (UnimplementedTokenServiceHandler) ListEnvironmentToken(context.Context, *connect.Request[v1.ListEnvironmentTokenRequest]) (*connect.Response[v1.ListEnvironmentTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.token.v1.TokenService.ListEnvironmentToken is not implemented"))
}
