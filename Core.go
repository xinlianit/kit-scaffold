package scaffold

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/handler"
	"github.com/xinlianit/kit-scaffold/logger"
	"github.com/xinlianit/kit-scaffold/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// 配置初始化
	config.Init()

	// 命令行解析
	commandLineParse()

	// 日志初始化
	var baseFields []zap.Field
	logger.ZapLogger = logger.ZapInit(logger.NewDefaultZapConfig(), baseFields)
}

// 命令行解析
func commandLineParse() {
	// 解析命令行参数
	pflag.String("server.host", "0.0.0.0", "服务地址")
	pflag.Int("server.port", 80, "服务端口")
	pflag.String("gateway.host", "0.0.0.0", "网关地址")
	pflag.Int("gateway.port", 80, "网关端口")
	pflag.Parse()
	config.Config().BindPFlags(pflag.CommandLine)
}

// 运行 Http 服务
// @param handler http 处理器
func RunHttpServer(handler http.Handler) {
	// 服务地址
	address := server.GetServerAddress()

	httpServer := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  time.Duration(config.Config().GetInt("server.readTimeout")) * time.Millisecond,
		WriteTimeout: time.Duration(config.Config().GetInt("server.writeTimeout")) * time.Millisecond,
	}

	// 服务启动成功
	logger.ZapLogger.Info(fmt.Sprintf("Listening and serving HTTP on %s, PID: %d", address, os.Getpid()))

	go func() {
		// todo 服务注册

		// 启动并监听服务
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ZapLogger.Sugar().Errorf("Server run error: %v", err)

			// todo 服务注销

			panic("Server run error: " + err.Error())
		}
	}()

	httpServerGraceStop(httpServer)
}

// 运行 gRPC 服务
func RunRpcServer(grpcServer *grpc.Server) {
	// 服务地址
	address := server.GetServerAddress()

	// 是否在gRPC服务中注册reflection服务, 开启后支持grpcurl命令行工具
	if config.Config().GetBool("server.grpc.reflection.register") {
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)
	}

	// 监听端口
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.ZapLogger.Sugar().Errorf("Server listen error: %v", err)
		panic("Server listen error: " + err.Error())
	}

	// 服务启动成功
	logger.ZapLogger.Info(fmt.Sprintf("Listening and serving gRPC on %s, PID: %d", address, os.Getpid()))

	go func() {
		// todo 服务注册

		// 启动服务
		if err := grpcServer.Serve(lis); err != nil {
			logger.ZapLogger.Sugar().Errorf("Server run error: %v", err)

			// todo 注销服务

			panic("Server run error: " + err.Error())
		}
	}()

	gRpcServerGraceStop(grpcServer)
}

// 运行RPC代理服务
// @param handler http 处理器
func RunGatewayServer(handler http.Handler) {
	// 服务地址
	address := server.GetGatewayServerAddress()

	httpServer := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  time.Duration(config.Config().GetInt("server.readTimeout")) * time.Millisecond,
		WriteTimeout: time.Duration(config.Config().GetInt("server.writeTimeout")) * time.Millisecond,
	}

	// 启动Http服务器（gRPC服务代理）
	logger.ZapLogger.Info(fmt.Sprintf("Listening and serving Gateway HTTP on %s, PID: %d", address, os.Getpid()))

	go func() {
		// todo 注册网关服务

		// 启动并监听服务
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ZapLogger.Sugar().Errorf("Gateway server run error: %v", err)

			// todo 注销网关服务

			panic("Gateway server run error: " + err.Error())
		}
	}()

	httpServerGraceStop(httpServer)
}

// HTTP 服务停止
// @param server http 服务实例
func httpServerGraceStop(server *http.Server) {
	// 信号通道
	signalChan := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 信号转发到 signalChan
	sig := <-signalChan                                        // 阻塞等待接收上述两种信号时，往下执行服务关机
	logger.ZapLogger.Sugar().Infof("Get Signal: %d", sig)
	logger.ZapLogger.Info("Shutdown Server ...")

	// 5 秒超时自动取消(当执行一个go 协程时，超时自动取消协程)
	contextTimeout := config.Config().GetInt("server.contextTimeout")
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(contextTimeout)*time.Millisecond)
	defer cancelFunc()

	if err := server.Shutdown(ctx); err != nil {
		logger.ZapLogger.Sugar().Fatal("Server Shutdown: %v", err)
	}

	// todo 服务注销

	logger.ZapLogger.Info("Server exiting")
}

// RPC 服务停止
// @param server gRPC 服务实例
func gRpcServerGraceStop(server *grpc.Server) {
	// 信号通道
	signalChan := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 信号转发到 signalChan
	sig := <-signalChan                                        // 阻塞等待接收上述两种信号时，往下执行服务关机

	logger.ZapLogger.Sugar().Infof("Get Signal: %d", sig)
	logger.ZapLogger.Info("Shutdown Server ...")

	// todo 服务注销

	logger.ZapLogger.Info("gRPC Server exiting")
}

// 创建 Http 处理器
func NewHttpHandler() *handler.HttpHandler {
	return &handler.HttpHandler{}
}
