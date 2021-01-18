package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/boot"
	"google.golang.org/grpc"
)

func main() {
	// 启动网关服务
	go gateway()

	// 创建 RPC 服务
	rpcServer := grpc.NewServer()

	// 注册服务
	boot.RegisterRpcServer(rpcServer)

	// 启动服务
	scaffold.RunRpcServer(rpcServer)
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
