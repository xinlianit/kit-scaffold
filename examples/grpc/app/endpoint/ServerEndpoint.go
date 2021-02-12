package endpoint

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	"github.com/go-kit/kit/endpoint"
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
