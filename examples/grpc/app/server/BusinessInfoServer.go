package server

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/transport"
	"google.golang.org/grpc/status"
)

// 商家详情服务
type BusinessInfoServer struct {
	// 服务传输实例
	transport transport.BusinessInfoTransport
	// 服务端点实例
	endpoint endpoint.BusinessInfoEndpoint
}

// 获取商户信息
func (s BusinessInfoServer) GetBusinessInfo(ctx context.Context, request *request.GetBusinessInfoRequest) (*response.GetBusinessInfoResponse, error) {
	// 创建服务处理器
	handler := grpcTransport.NewServer(
		s.endpoint.GetBusinessInfo(),                // 连接点
		s.transport.DecodeGetBusinessInfoRequest(),  // 请求解码器
		s.transport.EncodeGetBusinessInfoResponse(), // 响应编码器
		defaultServerOptions...,                     // 处理器参数
	)

	// 绑定 grpc 处理器
	ctx, rpcRsp, err := handler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	// 响应断言
	rsp, ok := rpcRsp.(*response.GetBusinessInfoResponse)
	if !ok {
		return nil, status.Errorf(100, "响应协议错误")
	}

	return rsp, nil
}

// 获取商户基本信息
func (s BusinessInfoServer) GetBusinessInfoBase(ctx context.Context, request *request.GetBusinessInfoBaseRequest) (*response.GetBusinessInfoBaseResponse, error) {
	// 创建服务处理器 todo
	handler := grpcTransport.NewServer(
		s.endpoint.GetBusinessInfo(),                // 连接点
		s.transport.DecodeGetBusinessInfoRequest(),  // 请求解码器
		s.transport.EncodeGetBusinessInfoResponse(), // 响应编码器
		defaultServerOptions...,                     // 处理器参数
	)

	ctx, rpcRsp, err := handler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	// 断言响应
	rsp, ok := rpcRsp.(*response.GetBusinessInfoBaseResponse)
	if !ok {
		return nil, status.Errorf(100, "响应协议错误")
	}

	return rsp, nil
}
