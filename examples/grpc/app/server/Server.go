package server

import (
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/transport"
)

var (
	// 服务处理器默认参数
	defaultServerOptions []grpcTransport.ServerOption
)

func init() {
	// 初始化服务处理器默认参数
	defaultServerOptions = DefaultServerOptions()
}

// 服务处理器默认参数设置
func DefaultServerOptions() []grpcTransport.ServerOption {
	// 处理器参数设置
	options := []grpcTransport.ServerOption{
		// 服务错误处理器
		grpcTransport.ServerErrorHandler(transport.ErrorHandler{}),
	}

	return options
}

// 服务实例
func NewServerServer() ServerServer {
	return ServerServer{
		transport: transport.NewServerTransport(),
		endpoint: endpoint.NewServerEndpoint(),
	}
}

// 创建服务实例
func NewBusinessInfoServer() BusinessInfoServer {
	return BusinessInfoServer{
		// 商家信息传输实例
		transport: transport.NewBusinessInfoTransport(),
		// 商家信息端点实例
		endpoint: endpoint.NewBusinessInfoEndpoint(),
	}
}
