package server

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/transport"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/status"
)

type ServerServer struct {
	transport transport.ServerTransport
	endpoint endpoint.ServerEndpoint
}

// 健康探针
func (s ServerServer) Health(ctx context.Context, request *request.HealthRequest) (*response.HealthResponse, error) {
	handler := grpcTransport.NewServer(
		s.endpoint.Health(),
		s.transport.DecodeHealthRequest(),
		s.transport.EncodeHealthResponse(),
		defaultServerOptions...,
	)

	ctx, resp, err := handler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	rsp, ok := resp.(*response.HealthResponse)

	if !ok {
		return nil, status.Error(100, "响应协议错误")
	}

	return rsp, nil
}
