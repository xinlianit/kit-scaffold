package server

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/xinlianit/kit-scaffold/endpoint"
	"github.com/xinlianit/kit-scaffold/pb/transport/request"
	"github.com/xinlianit/kit-scaffold/pb/transport/response"
	"github.com/xinlianit/kit-scaffold/transport"
	"google.golang.org/grpc/status"
)

type ServerServer struct {
	transport transport.ServerTransport
	endpoint endpoint.ServerEndpoint
}

// 健康探针
func (s ServerServer) Health(ctx context.Context, request *request.HealthRequest) (*response.HealthResponse, error) {
	handler := grpc.NewServer(
		s.endpoint.Health(),
		s.transport.DecodeHealthRequest(),
		s.transport.EncodeHealthResponse(),
		Options...,
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
