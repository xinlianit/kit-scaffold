package drive

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"time"
)

// 创建Consul客户端
func NewConsulClient() ConsulClient {
	// consul 配置
	cfg := consulApi.DefaultConfig()
	// consul 地址
	cfg.Address = config.Config().GetString("app.serviceCenter.consul.address")

	// 创建客户端
	client, err := consulApi.NewClient(cfg)
	if err != nil {
		logger.ZapLogger.Error(err.Error())
	}

	return ConsulClient{
		client: client,
	}
}

// Consul客户端
type ConsulClient struct {
	client *consulApi.Client
}

// 注册服务
func (c ConsulClient) RegisterService(serviceId string) error {
	// 服务名称
	serviceName := config.Config().GetString("app.serviceCenter.register.name")
	if serviceName == "" {
		serviceName = config.Config().GetString("app.id")
	}

	reg := &consulApi.AgentServiceRegistration{
		// 服务ID
		ID: serviceId,
		// 服务名称
		Name: serviceName,
		// 服务地址
		Address: util.ServerUtil().GetServerIp(),
		// 服务端口
		Port: config.Config().GetInt("server.port"),
	}


	// 服务标签
	if serviceTags := config.Config().GetStringSlice("app.serviceCenter.register.tags"); serviceTags != nil {
		reg.Tags = serviceTags
	}

	// 服务健康检查
	if config.Config().GetBool("app.serviceCenter.healthCheck.enable") {
		// 检测类型
		checkType := config.Config().GetString("app.serviceCenter.healthCheck.type")
		// 检测间隔时长
		interval := config.Config().GetInt("app.serviceCenter.healthCheck.interval")
		// 检测超时
		timeout := config.Config().GetInt("app.serviceCenter.healthCheck.timeout")
		// 服务最大生存周期
		maxLifeTime := config.Config().GetInt("app.serviceCenter.healthCheck.maxLifeTime")
		// 健康检查地址
		checkAddress := config.Config().GetString("app.serviceCenter.healthCheck.address")
		if checkAddress == "" {
			// 请求协议
			if protocol := config.Config().GetString("app.serviceCenter.healthCheck.protocol"); protocol != "" {
				checkAddress = fmt.Sprintf("%s://%s:%d", protocol, reg.Address, reg.Port)
			}else{
				checkAddress = fmt.Sprintf("%s:%d", reg.Address, reg.Port)
			}

			// 健康检查路径
			if path := config.Config().GetString("app.serviceCenter.healthCheck.path"); path != "" {
				checkAddress = fmt.Sprintf("%s/%s", checkAddress, path)
			}
		}

		reg.Check = &consulApi.AgentServiceCheck{
			// 检测间隔
			Interval: (time.Millisecond * time.Duration(interval)).String(),
			// 检测超时
			Timeout: (time.Millisecond * time.Duration(timeout)).String(),
			// 注销时间，服务过期时间
			DeregisterCriticalServiceAfter: (time.Millisecond * time.Duration(maxLifeTime)).String(),
		}

		// 检测ID
		if checkID := config.Config().GetString("app.serviceCenter.healthCheck.id"); checkID != "" {
			reg.Check.CheckID = checkID
		}

		// 检测名称
		if checkName := config.Config().GetString("app.serviceCenter.healthCheck.name"); checkName != "" {
			reg.Check.Name = checkName
		}

		// 健康检查类型
		switch checkType {
		case "tcp":
			// 检查地址
			reg.Check.TCP = checkAddress
			// 关闭tls验证
			//reg.Check.TLSSkipVerify
		case "grpc":
			// 检查地址
			reg.Check.GRPC = checkAddress
			// 是否启用TLS
			reg.Check.GRPCUseTLS = config.Config().GetBool("app.serviceCenter.healthCheck.grpc.tls.enable")
			// 注销时间，服务过期时间
			reg.Check.DeregisterCriticalServiceAfter = "10s"
		default:
			// 检测地址
			reg.Check.HTTP = checkAddress
			//reg.Check.Header
			// 检测请求方式
			reg.Check.Method = config.Config().GetString("app.serviceCenter.healthCheck.http.method")
		}
	}

	return c.client.Agent().ServiceRegister(reg)
}

// 服务注销
func (c ConsulClient) DeregisterService(serviceId string) error {
	return c.client.Agent().ServiceDeregister(serviceId)
}
