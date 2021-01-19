package boot

import (
	"context"
	"fmt"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/server"
	"github.com/xinlianit/kit-scaffold/logger"
	serverScaffold "github.com/xinlianit/kit-scaffold/server"
	"google.golang.org/grpc"
)

// RPC服务注册表
// @param rpcServer RPC服务实例
func RegisterRpcServer(rpcServer *grpc.Server) *grpc.Server {
	// 商家信息
	service.RegisterBusinessInfoServiceServer(rpcServer, server.NewBusinessInfoServer())
	return rpcServer
}

// Gateway网关服务注册
// @param ctx 上下文
// @param mux 多路器
// @param endpoint RPC服务连接地址
// @param opts RPC服务连接参数
func RegisterGatewayServer(ctx context.Context, mux *runtime.ServeMux) *runtime.ServeMux {
	// todo 连接参数
	opts := []grpc.DialOption{
		// 不启用TLS的认证
		grpc.WithInsecure(),
	}

	// 商家信息
	if err := service.RegisterBusinessInfoServiceHandlerFromEndpoint(ctx, mux, serverScaffold.GetServerAddress(), opts); err != nil {
		errMsg := fmt.Sprintf("Gateway server register error: %v", err)
		logger.ZapLogger.Error(errMsg)
		panic(errMsg)
	}

	return mux
}