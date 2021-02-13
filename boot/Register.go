package boot

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/xinlianit/kit-scaffold/logger"
	"github.com/xinlianit/kit-scaffold/pb/service"
	"github.com/xinlianit/kit-scaffold/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// 基础服务注册
// @param rpcServer RPC服务实例
func RegisterBaseServer(rpcServer *grpc.Server) *grpc.Server {
	// 健康检查
	grpc_health_v1.RegisterHealthServer(rpcServer, server.NewHealthServer())
	// 服务信息
	service.RegisterServerServiceServer(rpcServer, server.NewServerServer())

	return rpcServer
}

// Gateway 网关基础服务注册
// @param ctx 上下文
// @param mux 多路器
// @param endpoint RPC服务连接地址
// @param opts RPC服务连接参数
func RegisterGatewayBaseServer(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) *runtime.ServeMux {
	// 服务信息
	if err := service.RegisterServerServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		errMsg := fmt.Sprintf("Gateway server register error: %v", err)
		logger.ZapLogger.Error(errMsg)
		panic(errMsg)
	}

	return mux
}
