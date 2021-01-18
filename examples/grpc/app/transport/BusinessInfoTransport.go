package transport

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/status"
)

// 传输实例
func NewBusinessInfoTransport() BusinessInfoTransport {
	return BusinessInfoTransport{}
}

// 商家信息传输
type BusinessInfoTransport struct {
}

// 获取商家信息请求解码
func (t BusinessInfoTransport) DecodeGetBusinessInfoRequest() grpcTransport.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		// 请求断言
		getBusinessInfoReq, ok := req.(*request.GetBusinessInfoRequest)
		if !ok {
			return nil, status.Errorf(100, "请求协议错误: %#v", req)
		}

		// todo 请求过滤及包装

		// 返回请求
		return getBusinessInfoReq, nil
	}
}

// 获取商家信息响应编码
func (t BusinessInfoTransport) EncodeGetBusinessInfoResponse() grpcTransport.EncodeResponseFunc {
	return func(ctx context.Context, rsp interface{}) (interface{}, error) {
		// 请求断言
		getBusinessInfoRsp, ok := rsp.(*response.GetBusinessInfoResponse)
		if !ok {
			return nil, status.Errorf(100, "响应协议错误: %#v", rsp)
		}

		// todo 响应过滤及包装

		// 返回响应
		return getBusinessInfoRsp, nil
	}
}