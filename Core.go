package scaffold

import (
	"context"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/boot"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/drive"
	"github.com/xinlianit/kit-scaffold/logger"
	"github.com/xinlianit/kit-scaffold/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 服务ID
var serviceId string
var serviceGatewayId string

func init() {
	// 框架初始化
	boot.Init()

	// 服务ID
	appId := config.Config().GetString("app.id")	// 服务appId
	serverAddress := util.ServerUtil().GetServerIp()	// 服务地址
	serverPort := config.Config().GetInt("server.port")	// 服务端口
	serviceId = fmt.Sprintf("%s-%s:%d", appId, serverAddress, serverPort)

	// 网关服务ID
	serverGatewayPort := config.Config().GetInt("server.gateway.port")	// 服务端口
	serviceGatewayId = fmt.Sprintf("%s-%s-%s:%d", appId, "gateway", serverAddress, serverGatewayPort)
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
		// consul 客户端
		consulClient := drive.NewConsulClient()

		// 服务注册
		if err := consulClient.RegisterService(serviceId); err != nil {
			logger.ZapLogger.Sugar().Errorf("Server register to consul error: %v", err)
		}

		// 启动并监听服务
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ZapLogger.Sugar().Errorf("Server run error: %v", err)

			// 服务注销
			if deregisterErr := consulClient.DeregisterService(serviceId); deregisterErr != nil {
				logger.ZapLogger.Sugar().Errorf("Service deregister error from consul: %v", deregisterErr)
			}

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
		// consul 客户端
		consulClient := drive.NewConsulClient()

		// 服务注册
		if err := consulClient.RegisterService(serviceId); err != nil {
			logger.ZapLogger.Sugar().Errorf("Server register to consul error: %v", err)
		}

		// 启动服务
		if err := grpcServer.Serve(lis); err != nil {
			// 注销服务
			if deregisterErr := consulClient.DeregisterService(serviceId); deregisterErr != nil {
				logger.ZapLogger.Sugar().Errorf("Service deregister error from consul: %v", deregisterErr)
			}

			logger.ZapLogger.Sugar().Panicf("Server run error: %v", err)
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
		// consul 客户端
		cfg := consulApi.DefaultConfig()
		// consul 地址
		cfg.Address = config.Config().GetString("app.serviceCenter.consul.address")
		consulClient, err := consulApi.NewClient(cfg)
		if err != nil {
			logger.ZapLogger.Sugar().Errorf("Consul client initialize error: %v", err)
		}else{
			// 服务名称
			serviceName := config.Config().GetString("app.serviceCenter.register.name")
			if serviceName == "" {
				serviceName = config.Config().GetString("app.id")
			}

			// 服务注册信息
			reg := &consulApi.AgentServiceRegistration{
				// 服务ID
				ID: serviceGatewayId,
				// 服务名称
				Name: serviceName + "-gateway",
				// 服务地址
				Address: util.ServerUtil().GetServerIp(),
				// 服务端口
				Port: config.Config().GetInt("server.gateway.port"),
			}

			// 服务标签
			if serviceTags := config.Config().GetStringSlice("app.serviceCenter.register.tags"); serviceTags != nil {
				var newTags []string

				for _, tagItem := range serviceTags {
					newTags = append(newTags, tagItem + "-gateway")
				}

				reg.Tags = newTags
			}

			// 健康检查
			interval := config.Config().GetInt("app.serviceCenter.healthCheck.gateway.interval")
			timeout := config.Config().GetInt("app.serviceCenter.healthCheck.gateway.timeout")
			reg.Check = &consulApi.AgentServiceCheck{
				// 检测间隔
				Interval: (time.Millisecond * time.Duration(interval)).String(),
				// 检测超时
				Timeout: (time.Millisecond * time.Duration(timeout)).String(),
				// 检测地址
				HTTP: fmt.Sprintf("http://%s:%d/%s", reg.Address, reg.Port, "health"),
				// 检测请求方式
				Method: config.Config().GetString("app.serviceCenter.healthCheck.gateway.method"),
			}

			// 检测项名称
			if checkName := config.Config().GetString("app.serviceCenter.healthCheck.gateway.name"); checkName != "" {
				reg.Check.Name = checkName
			}

			// 注册网关服务
			if err := consulClient.Agent().ServiceRegister(reg); err != nil {
				logger.ZapLogger.Sugar().Errorf("Server register to consul error: %v", err)
			}
		}

		// 启动并监听服务
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ZapLogger.Sugar().Errorf("Gateway server run error: %v", err)

			// 注销网关服务
			if deregisterErr := consulClient.Agent().ServiceDeregister(serviceGatewayId); deregisterErr != nil {
				logger.ZapLogger.Sugar().Errorf("Service deregister error from consul: %v", deregisterErr)
			}

			panic("Gateway server run error: " + err.Error())
		}
	}()

	gatewayServerGraceStop(httpServer)
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

	// 服务注销
	if deregisterErr := drive.NewConsulClient().DeregisterService(serviceId); deregisterErr != nil {
		logger.ZapLogger.Sugar().Errorf("Service deregister error from consul: %v", deregisterErr)
	}

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

	// 服务注销
	if deregisterErr := drive.NewConsulClient().DeregisterService(serviceId); deregisterErr != nil {
		logger.ZapLogger.Sugar().Errorf("Service deregister error from consul: %v", deregisterErr)
	}

	logger.ZapLogger.Info("gRPC Server exiting")
}

// gateway 服务停止
// @param server http 服务实例
func gatewayServerGraceStop(server *http.Server) {
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

	// 服务注销
	if deregisterErr := drive.NewConsulClient().DeregisterService(serviceGatewayId); deregisterErr != nil {
		logger.ZapLogger.Sugar().Errorf("Service deregister error from consul: %v", deregisterErr)
	}

	logger.ZapLogger.Info("Server exiting")
}
