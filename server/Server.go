package server

import (
	"github.com/xinlianit/kit-scaffold/endpoint"
	"github.com/xinlianit/kit-scaffold/model"
	"github.com/xinlianit/kit-scaffold/transport"
)

var Request model.Request

// 健康检查服务
func NewHealthServer() HealthServer {
	return HealthServer{}
}

// 服务实例
func NewServerServer() ServerServer {
	return ServerServer{
		transport: transport.NewServerTransport(),
		endpoint: endpoint.NewServerEndpoint(),
	}
}