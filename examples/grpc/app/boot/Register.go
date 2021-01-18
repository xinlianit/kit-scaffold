package boot

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/server"
	"google.golang.org/grpc"
	"log"
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
func RegisterGatewayServer(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) []error {
	var errorList []error

	// 商家信息
	if err := service.RegisterBusinessInfoServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		// todo 记录错误日志
		log.Fatalf("error: %v", err)
		errorList = append(errorList, err)
	}

	return errorList
}
