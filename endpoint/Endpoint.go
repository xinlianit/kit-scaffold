package endpoint

// NewServerEndpoint 服务实例
func NewServerEndpoint() ServerEndpoint {
	return ServerEndpoint{}
}

// NewHealthEndpoint 健康检查
func NewHealthEndpoint() HealthEndpoint {
	return HealthEndpoint{}
}
