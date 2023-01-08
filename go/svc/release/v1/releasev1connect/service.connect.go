// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/release/v1/service.proto

package releasev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/humanlogio/api/go/svc/release/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ReleaseServiceName is the fully-qualified name of the ReleaseService service.
	ReleaseServiceName = "svc.release.v1.ReleaseService"
)

// ReleaseServiceClient is a client for the svc.release.v1.ReleaseService service.
type ReleaseServiceClient interface {
	CreateReleaseChannel(context.Context, *connect_go.Request[v1.CreateReleaseChannelRequest]) (*connect_go.Response[v1.CreateReleaseChannelResponse], error)
	ListReleaseChannel(context.Context, *connect_go.Request[v1.ListReleaseChannelRequest]) (*connect_go.Response[v1.ListReleaseChannelResponse], error)
	PublishVersion(context.Context, *connect_go.Request[v1.PublishVersionRequest]) (*connect_go.Response[v1.PublishVersionResponse], error)
	UnpublishVersion(context.Context, *connect_go.Request[v1.UnpublishVersionRequest]) (*connect_go.Response[v1.UnpublishVersionResponse], error)
	CreateVersionArtifact(context.Context, *connect_go.Request[v1.CreateVersionArtifactRequest]) (*connect_go.Response[v1.CreateVersionArtifactResponse], error)
	DeleteVersionArtifact(context.Context, *connect_go.Request[v1.DeleteVersionArtifactRequest]) (*connect_go.Response[v1.DeleteVersionArtifactResponse], error)
	ListVersionArtifact(context.Context, *connect_go.Request[v1.ListVersionArtifactRequest]) (*connect_go.Response[v1.ListVersionArtifactResponse], error)
}

// NewReleaseServiceClient constructs a client for the svc.release.v1.ReleaseService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReleaseServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ReleaseServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &releaseServiceClient{
		createReleaseChannel: connect_go.NewClient[v1.CreateReleaseChannelRequest, v1.CreateReleaseChannelResponse](
			httpClient,
			baseURL+"/svc.release.v1.ReleaseService/CreateReleaseChannel",
			opts...,
		),
		listReleaseChannel: connect_go.NewClient[v1.ListReleaseChannelRequest, v1.ListReleaseChannelResponse](
			httpClient,
			baseURL+"/svc.release.v1.ReleaseService/ListReleaseChannel",
			opts...,
		),
		publishVersion: connect_go.NewClient[v1.PublishVersionRequest, v1.PublishVersionResponse](
			httpClient,
			baseURL+"/svc.release.v1.ReleaseService/PublishVersion",
			opts...,
		),
		unpublishVersion: connect_go.NewClient[v1.UnpublishVersionRequest, v1.UnpublishVersionResponse](
			httpClient,
			baseURL+"/svc.release.v1.ReleaseService/UnpublishVersion",
			opts...,
		),
		createVersionArtifact: connect_go.NewClient[v1.CreateVersionArtifactRequest, v1.CreateVersionArtifactResponse](
			httpClient,
			baseURL+"/svc.release.v1.ReleaseService/CreateVersionArtifact",
			opts...,
		),
		deleteVersionArtifact: connect_go.NewClient[v1.DeleteVersionArtifactRequest, v1.DeleteVersionArtifactResponse](
			httpClient,
			baseURL+"/svc.release.v1.ReleaseService/DeleteVersionArtifact",
			opts...,
		),
		listVersionArtifact: connect_go.NewClient[v1.ListVersionArtifactRequest, v1.ListVersionArtifactResponse](
			httpClient,
			baseURL+"/svc.release.v1.ReleaseService/ListVersionArtifact",
			opts...,
		),
	}
}

// releaseServiceClient implements ReleaseServiceClient.
type releaseServiceClient struct {
	createReleaseChannel  *connect_go.Client[v1.CreateReleaseChannelRequest, v1.CreateReleaseChannelResponse]
	listReleaseChannel    *connect_go.Client[v1.ListReleaseChannelRequest, v1.ListReleaseChannelResponse]
	publishVersion        *connect_go.Client[v1.PublishVersionRequest, v1.PublishVersionResponse]
	unpublishVersion      *connect_go.Client[v1.UnpublishVersionRequest, v1.UnpublishVersionResponse]
	createVersionArtifact *connect_go.Client[v1.CreateVersionArtifactRequest, v1.CreateVersionArtifactResponse]
	deleteVersionArtifact *connect_go.Client[v1.DeleteVersionArtifactRequest, v1.DeleteVersionArtifactResponse]
	listVersionArtifact   *connect_go.Client[v1.ListVersionArtifactRequest, v1.ListVersionArtifactResponse]
}

// CreateReleaseChannel calls svc.release.v1.ReleaseService.CreateReleaseChannel.
func (c *releaseServiceClient) CreateReleaseChannel(ctx context.Context, req *connect_go.Request[v1.CreateReleaseChannelRequest]) (*connect_go.Response[v1.CreateReleaseChannelResponse], error) {
	return c.createReleaseChannel.CallUnary(ctx, req)
}

// ListReleaseChannel calls svc.release.v1.ReleaseService.ListReleaseChannel.
func (c *releaseServiceClient) ListReleaseChannel(ctx context.Context, req *connect_go.Request[v1.ListReleaseChannelRequest]) (*connect_go.Response[v1.ListReleaseChannelResponse], error) {
	return c.listReleaseChannel.CallUnary(ctx, req)
}

// PublishVersion calls svc.release.v1.ReleaseService.PublishVersion.
func (c *releaseServiceClient) PublishVersion(ctx context.Context, req *connect_go.Request[v1.PublishVersionRequest]) (*connect_go.Response[v1.PublishVersionResponse], error) {
	return c.publishVersion.CallUnary(ctx, req)
}

// UnpublishVersion calls svc.release.v1.ReleaseService.UnpublishVersion.
func (c *releaseServiceClient) UnpublishVersion(ctx context.Context, req *connect_go.Request[v1.UnpublishVersionRequest]) (*connect_go.Response[v1.UnpublishVersionResponse], error) {
	return c.unpublishVersion.CallUnary(ctx, req)
}

// CreateVersionArtifact calls svc.release.v1.ReleaseService.CreateVersionArtifact.
func (c *releaseServiceClient) CreateVersionArtifact(ctx context.Context, req *connect_go.Request[v1.CreateVersionArtifactRequest]) (*connect_go.Response[v1.CreateVersionArtifactResponse], error) {
	return c.createVersionArtifact.CallUnary(ctx, req)
}

// DeleteVersionArtifact calls svc.release.v1.ReleaseService.DeleteVersionArtifact.
func (c *releaseServiceClient) DeleteVersionArtifact(ctx context.Context, req *connect_go.Request[v1.DeleteVersionArtifactRequest]) (*connect_go.Response[v1.DeleteVersionArtifactResponse], error) {
	return c.deleteVersionArtifact.CallUnary(ctx, req)
}

// ListVersionArtifact calls svc.release.v1.ReleaseService.ListVersionArtifact.
func (c *releaseServiceClient) ListVersionArtifact(ctx context.Context, req *connect_go.Request[v1.ListVersionArtifactRequest]) (*connect_go.Response[v1.ListVersionArtifactResponse], error) {
	return c.listVersionArtifact.CallUnary(ctx, req)
}

// ReleaseServiceHandler is an implementation of the svc.release.v1.ReleaseService service.
type ReleaseServiceHandler interface {
	CreateReleaseChannel(context.Context, *connect_go.Request[v1.CreateReleaseChannelRequest]) (*connect_go.Response[v1.CreateReleaseChannelResponse], error)
	ListReleaseChannel(context.Context, *connect_go.Request[v1.ListReleaseChannelRequest]) (*connect_go.Response[v1.ListReleaseChannelResponse], error)
	PublishVersion(context.Context, *connect_go.Request[v1.PublishVersionRequest]) (*connect_go.Response[v1.PublishVersionResponse], error)
	UnpublishVersion(context.Context, *connect_go.Request[v1.UnpublishVersionRequest]) (*connect_go.Response[v1.UnpublishVersionResponse], error)
	CreateVersionArtifact(context.Context, *connect_go.Request[v1.CreateVersionArtifactRequest]) (*connect_go.Response[v1.CreateVersionArtifactResponse], error)
	DeleteVersionArtifact(context.Context, *connect_go.Request[v1.DeleteVersionArtifactRequest]) (*connect_go.Response[v1.DeleteVersionArtifactResponse], error)
	ListVersionArtifact(context.Context, *connect_go.Request[v1.ListVersionArtifactRequest]) (*connect_go.Response[v1.ListVersionArtifactResponse], error)
}

// NewReleaseServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReleaseServiceHandler(svc ReleaseServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/svc.release.v1.ReleaseService/CreateReleaseChannel", connect_go.NewUnaryHandler(
		"/svc.release.v1.ReleaseService/CreateReleaseChannel",
		svc.CreateReleaseChannel,
		opts...,
	))
	mux.Handle("/svc.release.v1.ReleaseService/ListReleaseChannel", connect_go.NewUnaryHandler(
		"/svc.release.v1.ReleaseService/ListReleaseChannel",
		svc.ListReleaseChannel,
		opts...,
	))
	mux.Handle("/svc.release.v1.ReleaseService/PublishVersion", connect_go.NewUnaryHandler(
		"/svc.release.v1.ReleaseService/PublishVersion",
		svc.PublishVersion,
		opts...,
	))
	mux.Handle("/svc.release.v1.ReleaseService/UnpublishVersion", connect_go.NewUnaryHandler(
		"/svc.release.v1.ReleaseService/UnpublishVersion",
		svc.UnpublishVersion,
		opts...,
	))
	mux.Handle("/svc.release.v1.ReleaseService/CreateVersionArtifact", connect_go.NewUnaryHandler(
		"/svc.release.v1.ReleaseService/CreateVersionArtifact",
		svc.CreateVersionArtifact,
		opts...,
	))
	mux.Handle("/svc.release.v1.ReleaseService/DeleteVersionArtifact", connect_go.NewUnaryHandler(
		"/svc.release.v1.ReleaseService/DeleteVersionArtifact",
		svc.DeleteVersionArtifact,
		opts...,
	))
	mux.Handle("/svc.release.v1.ReleaseService/ListVersionArtifact", connect_go.NewUnaryHandler(
		"/svc.release.v1.ReleaseService/ListVersionArtifact",
		svc.ListVersionArtifact,
		opts...,
	))
	return "/svc.release.v1.ReleaseService/", mux
}

// UnimplementedReleaseServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedReleaseServiceHandler struct{}

func (UnimplementedReleaseServiceHandler) CreateReleaseChannel(context.Context, *connect_go.Request[v1.CreateReleaseChannelRequest]) (*connect_go.Response[v1.CreateReleaseChannelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.release.v1.ReleaseService.CreateReleaseChannel is not implemented"))
}

func (UnimplementedReleaseServiceHandler) ListReleaseChannel(context.Context, *connect_go.Request[v1.ListReleaseChannelRequest]) (*connect_go.Response[v1.ListReleaseChannelResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.release.v1.ReleaseService.ListReleaseChannel is not implemented"))
}

func (UnimplementedReleaseServiceHandler) PublishVersion(context.Context, *connect_go.Request[v1.PublishVersionRequest]) (*connect_go.Response[v1.PublishVersionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.release.v1.ReleaseService.PublishVersion is not implemented"))
}

func (UnimplementedReleaseServiceHandler) UnpublishVersion(context.Context, *connect_go.Request[v1.UnpublishVersionRequest]) (*connect_go.Response[v1.UnpublishVersionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.release.v1.ReleaseService.UnpublishVersion is not implemented"))
}

func (UnimplementedReleaseServiceHandler) CreateVersionArtifact(context.Context, *connect_go.Request[v1.CreateVersionArtifactRequest]) (*connect_go.Response[v1.CreateVersionArtifactResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.release.v1.ReleaseService.CreateVersionArtifact is not implemented"))
}

func (UnimplementedReleaseServiceHandler) DeleteVersionArtifact(context.Context, *connect_go.Request[v1.DeleteVersionArtifactRequest]) (*connect_go.Response[v1.DeleteVersionArtifactResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.release.v1.ReleaseService.DeleteVersionArtifact is not implemented"))
}

func (UnimplementedReleaseServiceHandler) ListVersionArtifact(context.Context, *connect_go.Request[v1.ListVersionArtifactRequest]) (*connect_go.Response[v1.ListVersionArtifactResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.release.v1.ReleaseService.ListVersionArtifact is not implemented"))
}
