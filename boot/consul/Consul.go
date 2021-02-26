package consul

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"time"
)

// NewClient 创建Consul客户端
func NewClient() Client {
	// consul 地址
	var address string
	if address = config.Config().GetString("consul.address"); address == "" {
		address = config.Config().GetString("app.serviceCenter.consul.address")
	}

	// consul 配置
	cfg := consulApi.DefaultConfig()
	// consul 地址
	cfg.Address = address

	// 创建客户端
	client, err := consulApi.NewClient(cfg)
	if err != nil {
		logger.ZapLogger.Error(err.Error())
	}

	return Client{
		client: client,
	}
}

// Client Consul客户端
type Client struct {
	client *consulApi.Client
}

// RegisterService 注册服务
func (c Client) RegisterService(serviceID string) error {
	// 读取配置
	var cfg config.ServiceCenter
	if err := config.Config().UnmarshalKey("app.serviceCenter", &cfg); err != nil {
		logger.ZapLogger.Sugar().Errorf("Service center config unmarshal error: %v", err)
		return err
	}

	// 服务名称
	serviceName := cfg.Register.Name
	if serviceName == "" {
		serviceName = config.Config().GetString("app.id")
	}

	reg := &consulApi.AgentServiceRegistration{
		// 服务ID
		ID: serviceID,
		// 服务名称
		Name: serviceName,
		// 服务地址
		Address: util.ServerUtil().GetServerIp(),
		// 服务端口
		Port: config.Config().GetInt("server.port"),
	}


	// 服务标签
	if serviceTags := cfg.Register.Tags; serviceTags != nil {
		reg.Tags = serviceTags
	}

	// 服务健康检查
	if cfg.HealthCheck.Enable {
		// 健康检查地址
		checkAddress := cfg.HealthCheck.Address
		if checkAddress == "" {
			// 请求协议
			if cfg.HealthCheck.Protocol != "" {
				checkAddress = fmt.Sprintf("%s://%s:%d", cfg.HealthCheck.Protocol, reg.Address, reg.Port)
			}else{
				checkAddress = fmt.Sprintf("%s:%d", reg.Address, reg.Port)
			}

			// 健康检查路径
			if cfg.HealthCheck.Path != "" {
				checkAddress = fmt.Sprintf("%s/%s", checkAddress, cfg.HealthCheck.Path)
			}
		}

		reg.Check = &consulApi.AgentServiceCheck{
			// 检测间隔
			Interval: (time.Millisecond * time.Duration(cfg.HealthCheck.Interval)).String(),
			// 检测超时
			Timeout: (time.Millisecond * time.Duration(cfg.HealthCheck.Timeout)).String(),
			// 注销时间，服务过期时间
			DeregisterCriticalServiceAfter: (time.Millisecond * time.Duration(cfg.HealthCheck.MaxLifeTime)).String(),
		}

		// 检测ID
		if cfg.HealthCheck.Id != "" {
			reg.Check.CheckID = cfg.HealthCheck.Id
		}

		// 检测名称
		if cfg.HealthCheck.Name != "" {
			reg.Check.Name = cfg.HealthCheck.Name
		}

		// 健康检查类型
		switch cfg.HealthCheck.Type {
		case "tcp":
			// 检查地址
			reg.Check.TCP = checkAddress
			// 关闭tls验证
			//reg.Check.TLSSkipVerify
		case "grpc":
			// 检查地址
			reg.Check.GRPC = checkAddress
			// 是否启用TLS
			reg.Check.GRPCUseTLS = cfg.HealthCheck.Grpc.TlsEnable
		default:
			// 检测地址
			reg.Check.HTTP = checkAddress
			//reg.Check.Header
			// 检测请求方式
			reg.Check.Method = cfg.HealthCheck.Http.Method
		}
	}

	return c.client.Agent().ServiceRegister(reg)
}

// DeregisterService 服务注销
func (c Client) DeregisterService(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}

// DiscoverService 服务发现
func (c Client) DiscoverService(service, tag string, passingOnly bool, query *consulApi.QueryOptions) ([]*consulApi.ServiceEntry, error) {
	services, _, err := c.client.Health().Service(service, tag, passingOnly, query)
	if err != nil {
		return nil, err
	}

	return services, nil
}
