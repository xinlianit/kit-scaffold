package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/interceptor"
	"github.com/xinlianit/kit-scaffold/examples/grpc/boot"
	"google.golang.org/grpc"
)

func main() {
	// 初始化脚手架
	scaffold.Init()

	// ---------------------------  应用逻辑 ----------------------

	// 启动初始化
	boot.Init()
	defer boot.Destruct()

	// 启动网关服务
	go gateway()

	// 创建 RPC 服务
	opts := []grpc.ServerOption{
		// 注册一元拦截器
		grpc.ChainUnaryInterceptor(interceptor.UnaryServerInterceptor()...),
	}
	rpcServer := grpc.NewServer(opts...)

	// 注册服务
	boot.RegisterRpcServer(rpcServer)

	// ---------------------------  应用逻辑 ----------------------

	// 启动服务
	scaffold.RunRPCServer(rpcServer)
}

// 网关服务
func gateway() {
	// todo 是否设置连接超时
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 网关多路复用器
	gatewayMux := runtime.NewServeMux()

	// 注册网关服务
	boot.RegisterGatewayServer(ctx, gatewayMux)

	// 启动网关服务
	scaffold.RunGatewayServer(gatewayMux)
}
