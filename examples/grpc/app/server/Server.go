package server

import (
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/transport"
)

// 创建服务实例
func NewBusinessInfoServer() BusinessInfoServer {
	return BusinessInfoServer{
		// 商家信息传输实例
		transport: transport.NewBusinessInfoTransport(),
		// 商家信息端点实例
		endpoint: endpoint.NewBusinessInfoEndpoint(),
	}
}
