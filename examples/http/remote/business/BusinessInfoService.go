package business

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	"google.golang.org/grpc/status"
	"log"
)

type businessInfoService struct {
	client service.BusinessInfoServiceClient
}

func (s businessInfoService) GetBusinessInfo() (*response.GetBusinessInfoResponse, error) {
	req := &request.GetBusinessInfoRequest{
		BusinessId: 99,
	}
	rsp, err := s.client.GetBusinessInfo(context.Background(), req)

	if err != nil {
		if rsp, ok := status.FromError(err); ok {
			log.Printf("RPC 错误: code: %d, message: %s", rsp.Proto().GetCode(), rsp.Proto().GetMessage())
		}

		return nil, err
	}

	return rsp, nil
}
