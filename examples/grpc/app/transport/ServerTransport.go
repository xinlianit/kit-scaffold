package transport

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/status"
)

func NewServerTransport() ServerTransport {
	return ServerTransport{}
}

// 服务传输
type ServerTransport struct {

}

func (t ServerTransport) DecodeHealthRequest() grpcTransport.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		// 请求断言
		healthReq, ok := req.(*request.HealthRequest)
		if !ok {
			return nil, status.Errorf(100, "请求协议错误: %#v", req)
		}

		// 返回请求
		return healthReq, nil
	}
}

func (t ServerTransport) EncodeHealthResponse() grpcTransport.EncodeResponseFunc {
	return func(ctx context.Context, rsp interface{}) (interface{}, error) {
		// 请求断言
		healthRsp, ok := rsp.(*response.HealthResponse)
		if !ok {
			return nil, status.Errorf(100, "响应协议错误: %#v", rsp)
		}

		// 返回响应
		return healthRsp, nil
	}
}