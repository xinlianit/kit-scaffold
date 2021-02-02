package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/pflag"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/examples/grpc/boot"
	"github.com/xinlianit/kit-scaffold/interceptor"
	"google.golang.org/grpc"
)

func main() {
	// 命令行参数解析
	commandLineParse()

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
	boot.RegisterRpcServer(rpcServer)

	// 启动服务
	scaffold.RunRpcServer(rpcServer)
}

// 命令行解析
func commandLineParse() {
	// 解析命令行参数
	pflag.String("env", "PRD","环境名称")
	pflag.String("server.host", "0.0.0.0", "服务地址")
	pflag.Int("server.port", 80, "服务端口")
	pflag.String("server.gateway.host", "0.0.0.0", "网关地址")
	pflag.Int("server.gateway.port", 8080, "网关端口")
	pflag.String("app.id", "", "应用ID")
	pflag.String("nacos.host", "127.0.0.1", "Nacos主机")
	pflag.Int("nacos.port", 8848, "Nacos端口")
	pflag.String("nacos.namespace", "", "Nacos名称空间")
	pflag.String("consul.host", "127.0.0.1", "Consul主机")
	pflag.Int("consul.port", 8500, "Consul端口")
	pflag.Parse()
	config.Config().BindPFlags(pflag.CommandLine)
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
