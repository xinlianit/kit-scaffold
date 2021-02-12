package server

import (
	"github.com/xinlianit/kit-scaffold/model"
)

var Request model.Request

// 健康检查服务
func NewHealthServer() HealthServer {
	return HealthServer{}
}