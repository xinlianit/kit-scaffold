package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/xinlianit/kit-scaffold/pb/transport/response"
)

// 服务连接点
type ServerEndpoint struct {

}

// 健康探针
func (e ServerEndpoint) Health() endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
		return &response.HealthResponse{
			Status: "SERVING",
		}, nil
	}
}
