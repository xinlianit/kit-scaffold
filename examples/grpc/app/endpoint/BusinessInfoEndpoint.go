package endpoint

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	"github.com/go-kit/kit/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/service"
)

// 商家详情端点
type BusinessInfoEndpoint struct {
	businessInfoService service.BusinessInfoService
}

// 获取商家详情
func (e BusinessInfoEndpoint) GetBusinessInfo() endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
		// 请求断言
		getBusinessInfoReq := req.(*request.GetBusinessInfoRequest)

		// 调用服务
		business, err := e.businessInfoService.BusinessInfo(getBusinessInfoReq.GetBusinessId())
		if err != nil {

		}

		// 返回响应数据
		return &response.GetBusinessInfoResponse{
			BusinessId:   business.GetBusinessId(),
			BusinessName: business.GetBusinessName(),
		}, nil
	}
}
