package business

import (
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
)

type businessInfoService struct {
	client service.BusinessInfoServiceClient
}

func (s businessInfoService) GetBusinessInfo() *response.GetBusinessInfoResponse {
	ctx, cancel := getContext()
	defer cancel()

	req := &request.GetBusinessInfoRequest{
		BusinessId: 99,
	}
	rsp, err := s.client.GetBusinessInfo(ctx, req)

	if err != nil {
		panic(err)
	}

	return rsp
}
