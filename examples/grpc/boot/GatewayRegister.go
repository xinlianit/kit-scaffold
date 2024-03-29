package boot

import (
	"context"
	"fmt"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/xinlianit/kit-scaffold/boot"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"google.golang.org/grpc"
)

// RegisterGatewayServer Gateway网关服务注册
// @param ctx 上下文
// @param mux 多路器
// @param endpoint RPC服务连接地址
// @param opts RPC服务连接参数
func RegisterGatewayServer(ctx context.Context, mux *runtime.ServeMux) *runtime.ServeMux {
	// 连接点
	endpoint := fmt.Sprintf("%s:%d", "127.0.0.1", config.Config().GetInt("server.port"))

	// todo 连接参数
	opts := []grpc.DialOption{
		// 不启用TLS的认证
		grpc.WithInsecure(),
	}

	// 注册基础服务
	boot.RegisterGatewayBaseServer(ctx, mux, endpoint, opts)

	// TODO: 注册服务处理器
	// 商家信息
	if err := service.RegisterBusinessInfoServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		errMsg := fmt.Sprintf("Gateway server register error: %v", err)
		logger.ZapLogger.Error(errMsg)
		panic(errMsg)
	}

	return mux
}
