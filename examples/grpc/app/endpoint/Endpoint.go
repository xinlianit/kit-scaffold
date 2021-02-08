package endpoint

import (
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/service"
)

// 服务实例
func NewServerEndpoint() ServerEndpoint {
	return ServerEndpoint{}
}

// 商户信息端点实例
func NewBusinessInfoEndpoint() BusinessInfoEndpoint {
	return BusinessInfoEndpoint{
		// 商家信息服务
		businessInfoService: service.NewBusinessInfoService(),
	}
}
