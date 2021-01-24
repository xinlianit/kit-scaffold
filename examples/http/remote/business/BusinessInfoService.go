package business

import (
	"context"
	"fmt"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	"google.golang.org/grpc/status"
)

type businessInfoService struct {
	client service.BusinessInfoServiceClient
}

func (s businessInfoService) GetBusinessInfo() *response.GetBusinessInfoResponse {
	req := &request.GetBusinessInfoRequest{
		BusinessId: 99,
	}
	rsp, err := s.client.GetBusinessInfo(context.Background(), req)

	if err != nil {
		fmt.Println(fmt.Sprintf("RPC 调用错误：%v", err))
		if rsp, ok := status.FromError(err); ok {
			fmt.Println(fmt.Sprintf("RPC 调用错误：%v", rsp))
			fmt.Println(fmt.Sprintf("code：%#v", rsp.Code().String()))
			fmt.Println(fmt.Sprintf("msg：%#v", rsp.Message()))
		}

		return nil
	}

	return rsp
}
