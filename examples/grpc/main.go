package main

import (
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/boot"
	"google.golang.org/grpc"
)

func main() {
	// 创建 RPC 服务
	rpcServer := grpc.NewServer()

	// 注册服务
	boot.RegisterRpcServer(rpcServer)

	// 启动服务
	scaffold.RunRpcServer(rpcServer)
}
