package boot

import (
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"github.com/xinlianit/kit-scaffold/boot"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/server"
	"google.golang.org/grpc"
)

// RegisterRpcServer RPC服务注册表
// @param rpcServer RPC服务实例
func RegisterRpcServer(rpcServer *grpc.Server) *grpc.Server {
	// 注册基础服务
	boot.RegisterRPCBaseServer(rpcServer)

	// TODO: 注册服务处理器
	// 商家信息
	service.RegisterBusinessInfoServiceServer(rpcServer, server.NewBusinessInfoServer())

	return rpcServer
}
