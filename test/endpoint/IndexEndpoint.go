package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	req "github.com/xinlianit/kit-scaffold/test/repository/request"
	rsp "github.com/xinlianit/kit-scaffold/test/repository/response"
	"github.com/xinlianit/kit-scaffold/test/service"
)

type IndexEndpoint struct {
	indexService service.IndexService
}

func (e IndexEndpoint) Hello() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// 请求断言
		helloReq := request.(req.HelloRequest)

		// 调用服务
		helloEntity, _ := e.indexService.Hello(helloReq.GetId())

		// 返回响应
		response = rsp.HelloResponse{
			Data: helloEntity,
		}

		return response, nil
	}
}
