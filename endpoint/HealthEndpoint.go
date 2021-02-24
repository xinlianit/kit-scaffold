package endpoint

import (
	"context"
	"github.com/xinlianit/kit-scaffold/repository/response"
)

// HealthEndpoint 健康检查连接点
type HealthEndpoint struct {
	
}

// Check 健康检查
func (e HealthEndpoint) Check(ctx context.Context, request interface{}) (interface{}, error) {
	return response.HealthResponse{
		Status: "SERVING",
	}, nil
}
