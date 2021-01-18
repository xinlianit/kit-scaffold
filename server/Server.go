package server

import (
	"fmt"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/model"
)

var Request model.Request

// 获取服务地址
func GetServerAddress() string {
	return fmt.Sprintf("%s:%d", config.Config().GetString("server.host"), config.Config().GetInt("server.port"))
}

// 获取网关服务地址
func GetGatewayServerAddress() string {
	return fmt.Sprintf("%s:%d", config.Config().GetString("server.gateway.host"), config.Config().GetInt("server.gateway.port"))
}
