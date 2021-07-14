package scaffold

import (
	"context"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/spf13/pflag"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/boot"
	"github.com/xinlianit/kit-scaffold/boot/consul"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 服务ID
var (
 	commandLineOnce sync.Once
	serviceRegisterID string
 	serviceGatewayRegisterID string
 	wg sync.WaitGroup
)

func Init() {
	// 框架初始化
	boot.Init()

	// 命令行初始化
	commandLineInit()
}

// 命令行初始化
func commandLineInit() {
	// 命令行参数
	pflag.String("env", "PRD","环境名称")
	pflag.String("server.host", "0.0.0.0", "服务地址")
	pflag.Int("server.port", 80, "服务端口")
	pflag.String("server.gateway.host", "0.0.0.0", "网关地址")
	pflag.Int("server.gateway.port", 8080, "网关端口")
	pflag.String("nacos.address", "", "Nacos地址")
	pflag.String("nacos.namespace", "", "Nacos名称空间")
	pflag.String("consul.address", "", "Consul地址")
}

// CommandLineParse 命令行参数解析
func CommandLineParse()  {
	commandLineOnce.Do(func() {
		pflag.Parse()
		config.Config().BindPFlags(pflag.CommandLine)
	})
}


// RunHTTPServer 运行 Http 服务
// @param handler http 处理器
func RunHTTPServer(handler http.Handler) {
	// 命令行参数解析
	CommandLineParse()

	// 配置中心初始化
	boot.ConfigCenterInit()

	// 服务监听地址
	listenAddress := fmt.Sprintf("%s:%d", config.Config().GetString("server.host"), config.Config().GetInt("server.port"))

	httpServer := &http.Server{
		Addr:         listenAddress,
		Handler:      handler,
		ReadTimeout:  time.Duration(config.ServerConfig.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.ServerConfig.WriteTimeout) * time.Millisecond,
	}

	// 服务启动成功
	logger.ZapLogger.Info(fmt.Sprintf("Listening and serving HTTP on %s, PID: %d", listenAddress, os.Getpid()))

	go func() {
		// 是否注册服务
		if config.AppConfig.ServiceCenter.Register.Enable {
			// 服务注册ID
			serviceRegisterID = fmt.Sprintf("%s-%s-%d", config.AppConfig.Id, util.ServerUtil().GetServerIp(), config.Config().GetInt("server.port"))

			// 服务注册
			if err := consul.NewClient().RegisterService(serviceRegisterID); err != nil {
				logger.ZapLogger.Sugar().Errorf("Http server register to consul error: %v", err)
			}else{
				logger.ZapLogger.Info("Http server register to consul successful")
			}
		}

		// 启动并监听服务
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ZapLogger.Sugar().Errorf("Http Server run error: %v", err)

			if config.AppConfig.ServiceCenter.Register.Enable {
				// 服务注销
				if deregisterErr := consul.NewClient().DeregisterService(serviceRegisterID); deregisterErr != nil {
					logger.ZapLogger.Sugar().Errorf("Http server deregister error from consul: %v", deregisterErr)
				}else{
					logger.ZapLogger.Info("Http server deregister successful from consul")
				}
			}

			panic("Http server run error: " + err.Error())
		}
	}()

	httpServerGraceStop(httpServer)
}

// RunRPCServer 运行 gRPC 服务
func RunRPCServer(grpcServer *grpc.Server) {
	// 命令行参数解析
	CommandLineParse()

	// 初始化
	Init()

	// 配置中心初始化
	boot.ConfigCenterInit()

	// 服务监听地址
	listenAddress := fmt.Sprintf("%s:%d", config.Config().GetString("server.host"), config.Config().GetInt("server.port"))

	// 是否在gRPC服务中注册reflection服务, 开启后支持grpcurl命令行工具
	if config.ServerConfig.Grpc.Reflection.Register {
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)
	}

	// 监听端口
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		logger.ZapLogger.Sugar().Errorf("gRPC server listen error: %v", err)
		panic("Server listen error: " + err.Error())
	}

	// 服务启动成功
	logger.ZapLogger.Info(fmt.Sprintf("Listening and serving gRPC on %s, PID: %d", listenAddress, os.Getpid()))

	go func() {
		// 是否注册服务
		if config.AppConfig.ServiceCenter.Register.Enable {
			// 服务注册ID
			serviceRegisterID = fmt.Sprintf("%s-%s-%d", config.AppConfig.Id, util.ServerUtil().GetServerIp(), config.Config().GetInt("server.port"))

			// 服务注册
			if err := consul.NewClient().RegisterService(serviceRegisterID); err != nil {
				logger.ZapLogger.Sugar().Errorf("gRPC server register to consul error: %v", err)
			}else{
				logger.ZapLogger.Info("gRPC server register to consul successful")
			}
		}

		// 启动服务
		if err := grpcServer.Serve(lis); err != nil {
			// 是否注册服务
			if config.AppConfig.ServiceCenter.Register.Enable {
				// 注销服务
				if deregisterErr := consul.NewClient().DeregisterService(serviceRegisterID); deregisterErr != nil {
					logger.ZapLogger.Sugar().Errorf("gRPC server deregister error from consul: %v", deregisterErr)
				}else{
					logger.ZapLogger.Info("gRPC server deregister successful from consul")
				}
			}

			logger.ZapLogger.Sugar().Panicf("gRPC server run error: %v", err)
		}
	}()

	gRPCServerGraceStop(grpcServer)
}

// RunGatewayServer 运行RPC代理服务
// @param handler http 处理器
func RunGatewayServer(handler http.Handler) {
	// 增加等待计数器
	wg.Add(1)

	// 命令行参数解析
	CommandLineParse()

	// 网关监听地址
	gatewayListenAddress := fmt.Sprintf("%s:%d",
		config.Config().GetString("server.gateway.host"),
		config.Config().GetInt("server.gateway.port"),
		)

	httpServer := &http.Server{
		Addr:         gatewayListenAddress,
		Handler:      handler,
		ReadTimeout:  time.Duration(config.ServerConfig.Gateway.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.ServerConfig.Gateway.WriteTimeout) * time.Millisecond,
	}

	// 启动Http服务器（gRPC服务代理）
	logger.ZapLogger.Info(fmt.Sprintf("Listening and serving Gateway HTTP on %s, PID: %d", gatewayListenAddress, os.Getpid()))

	go func() {
		// 是否注册服务
		if config.AppConfig.ServiceCenter.Register.Enable {
			// 服务注册ID
			serviceGatewayRegisterID = fmt.Sprintf("%s-%s-%d",
				config.AppConfig.Id,
				util.ServerUtil().GetServerIp(),
				config.Config().GetInt("server.gateway.port"),
			)

			// 服务中心
			serviceCenterCfg := config.AppConfig.ServiceCenter

			// consul 客户端
			cfg := consulApi.DefaultConfig()
			// consul 地址
			cfg.Address = serviceCenterCfg.Consul.Address
			consulClient, err := consulApi.NewClient(cfg)
			if err != nil {
				logger.ZapLogger.Sugar().Errorf("Consul client initialize error: %v", err)
			}else{
				// 服务名称
				serviceName := serviceCenterCfg.Register.Name
				if serviceName == "" {
					serviceName = config.AppConfig.Id
				}

				// 服务注册信息
				reg := &consulApi.AgentServiceRegistration{
					// 服务ID
					ID: serviceGatewayRegisterID,
					// 服务名称
					Name: serviceName,
					// 服务地址
					Address: util.ServerUtil().GetServerIp(),
					// 服务端口
					Port: config.ServerConfig.Gateway.Port,
				}

				// 服务标签
				reg.Tags = append(serviceCenterCfg.Register.Tags, "http")

				// 服务健康检查
				if serviceCenterCfg.HealthCheck.Enable {
					// 健康检查
					checkAddress := serviceCenterCfg.HealthCheck.Gateway.Address
					if checkAddress == "" {
						checkAddress = fmt.Sprintf("%s://%s:%d/%s", serviceCenterCfg.HealthCheck.Gateway.Protocol, reg.Address, reg.Port, serviceCenterCfg.HealthCheck.Gateway.Path)
					}
					reg.Check = &consulApi.AgentServiceCheck{
						// 检测间隔
						Interval: (time.Millisecond * time.Duration(serviceCenterCfg.HealthCheck.Gateway.Interval)).String(),
						// 检测超时
						Timeout: (time.Millisecond * time.Duration(serviceCenterCfg.HealthCheck.Gateway.Timeout)).String(),
						// 检测地址
						HTTP: checkAddress,
						// 检测请求方式
						Method: serviceCenterCfg.HealthCheck.Gateway.Method,
						// 注销时间，服务过期时间
						DeregisterCriticalServiceAfter: (time.Millisecond * time.Duration(serviceCenterCfg.HealthCheck.Gateway.MaxLifeTime)).String(),
					}

					// 检测项名称
					if serviceCenterCfg.HealthCheck.Gateway.Name != "" {
						reg.Check.Name = serviceCenterCfg.HealthCheck.Gateway.Name
					}
				}

				// 注册网关服务
				if err := consulClient.Agent().ServiceRegister(reg); err != nil {
					logger.ZapLogger.Sugar().Errorf("Gateway server register to consul error: %v", err)
				}else{
					logger.ZapLogger.Info("Gateway server register to consul successful")
				}
			}
		}

		// 启动并监听服务
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ZapLogger.Sugar().Errorf("Gateway server run error: %v", err)

			// 是否注册服务
			if config.AppConfig.ServiceCenter.Register.Enable {
				// 注销网关服务
				if deregisterErr := consul.NewClient().DeregisterService(serviceGatewayRegisterID); deregisterErr != nil {
					logger.ZapLogger.Sugar().Errorf("Gateway server deregister error from consul: %v", deregisterErr)
				}else{
					logger.ZapLogger.Info("Gateway server deregister successful from consul")
				}
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
	logger.ZapLogger.Info("Http Shutdown Server ...")

	// 5 秒超时自动取消(当执行一个go 协程时，超时自动取消协程)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.ServerConfig.ContextTimeout)*time.Millisecond)
	defer cancelFunc()

	if err := server.Shutdown(ctx); err != nil {
		logger.ZapLogger.Sugar().Fatal("Http Server Shutdown: %v", err)
	}

	if config.AppConfig.ServiceCenter.Register.Enable {
		// 服务注销
		if deregisterErr := consul.NewClient().DeregisterService(serviceRegisterID); deregisterErr != nil {
			logger.ZapLogger.Sugar().Errorf("Http server deregister error from consul: %v", deregisterErr)
		}else{
			logger.ZapLogger.Info("Http server deregister successful from consul")
		}
	}

	logger.ZapLogger.Info("Http Server exiting")
}

// RPC 服务停止
// @param server gRPC 服务实例
func gRPCServerGraceStop(server *grpc.Server) {
	// 信号通道
	signalChan := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 信号转发到 signalChan
	sig := <-signalChan                                        // 阻塞等待接收上述两种信号时，往下执行服务关机

	logger.ZapLogger.Sugar().Infof("Get Signal: %d", sig)
	logger.ZapLogger.Info("gRPC Shutdown Server ...")

	if config.AppConfig.ServiceCenter.Register.Enable {
		// 服务注销
		if deregisterErr := consul.NewClient().DeregisterService(serviceRegisterID); deregisterErr != nil {
			logger.ZapLogger.Sugar().Errorf("gRPC server deregister error from consul: %v", deregisterErr)
		}else{
			logger.ZapLogger.Info("gRPC server deregister successful from consul")
		}
	}

	// 优雅关机
	server.GracefulStop()

	logger.ZapLogger.Info("gRPC server exiting")

	// 等待所有协程退出
	wg.Wait()
}

// gateway 服务停止
// @param server http 服务实例
func gatewayServerGraceStop(server *http.Server) {
	// 完成退出，减少等待计数器
	defer wg.Done()
	// 信号通道
	signalChan := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 信号转发到 signalChan
	sig := <-signalChan                                        // 阻塞等待接收上述两种信号时，往下执行服务关机
	logger.ZapLogger.Sugar().Infof("Get Signal: %d", sig)
	logger.ZapLogger.Info("Shutdown Gateway Server ...")

	// 5 秒超时自动取消(当执行一个go 协程时，超时自动取消协程)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.ServerConfig.ContextTimeout)*time.Millisecond)
	defer cancelFunc()

	if err := server.Shutdown(ctx); err != nil {
		logger.ZapLogger.Sugar().Fatal("Gateway Server Shutdown: %v", err)
	}

	if config.AppConfig.ServiceCenter.Register.Enable {
		// 服务注销
		if deregisterErr := consul.NewClient().DeregisterService(serviceGatewayRegisterID); deregisterErr != nil {
			logger.ZapLogger.Sugar().Errorf("Gateway server deregister error from consul: %v", deregisterErr)
		}else{
			logger.ZapLogger.Info("Gateway server deregister successful from consul")
		}
	}

	logger.ZapLogger.Info("Gateway Server exiting")
}
