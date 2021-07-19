package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/cli/base-scaffold/grpc/boot"
	"github.com/xinlianit/kit-scaffold/interceptor"
	"google.golang.org/grpc"
)

func main() {
	// 启动初始化
	boot.Init()

	// 启动网关服务
	go gateway()

	// 拦截器
	interceptors := interceptor.DefaultUnaryServerInterceptor()
	// 凭证认证拦截器
	//interceptors = append(interceptors, interceptor.AuthInterceptor)

	// 创建 RPC 服务
	opts := []grpc.ServerOption{
		// 注册拦截器
		grpc.ChainUnaryInterceptor(interceptors...),
	}
	rpcServer := grpc.NewServer(opts...)

	// 注册服务
	//boot.RegisterRpcServer(rpcServer)

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
	//boot.RegisterGatewayServer(ctx, gatewayMux)

	// 启动网关服务
	scaffold.RunGatewayServer(gatewayMux)
}
