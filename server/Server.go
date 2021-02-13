package server

import (
	"github.com/go-kit/kit/transport/grpc"
	"github.com/xinlianit/kit-scaffold/endpoint"
	"github.com/xinlianit/kit-scaffold/model"
	"github.com/xinlianit/kit-scaffold/transport"
)

var (
	Request model.Request
	// 服务处理器参数
	Options []grpc.ServerOption
)

func init()  {
	Options = DefaultServerOptions()
}

// 服务处理器默认参数设置
func DefaultServerOptions() []grpc.ServerOption {
	// 处理器参数设置
	options := []grpc.ServerOption{
		// 服务错误处理器
		grpc.ServerErrorHandler(transport.ErrorHandler{}),
	}

	return options
}

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