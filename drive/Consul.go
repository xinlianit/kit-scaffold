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

	// 检测间隔时长
	interval := config.Config().GetInt("app.serviceCenter.healthCheck.grpc.interval")
	// 服务最大生存周期
	//maxLifeTime := config.Config().GetInt("app.serviceCenter.healthCheck.grpc.maxLifeTime")

	// 服务健康检查
	reg.Check = &consulApi.AgentServiceCheck{
		// 检测地址
		GRPC: fmt.Sprintf("%s:%d", reg.Address, reg.Port),
		// 是否启用TLS
		GRPCUseTLS: config.Config().GetBool("app.serviceCenter.healthCheck.grpc.tls.enable"),
		// 检测间隔
		Interval: (time.Millisecond * time.Duration(interval)).String(),
		// 注销时间，服务过期时间
		//DeregisterCriticalServiceAfter: (time.Millisecond * time.Duration(maxLifeTime)).String(),
	}

	return c.client.Agent().ServiceRegister(reg)
}

// 服务注销
func (c ConsulClient) DeregisterService(serviceId string) error {
	return c.client.Agent().ServiceDeregister(serviceId)
}
