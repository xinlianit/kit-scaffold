package endpoint

import (
	"context"
	"github.com/xinlianit/kit-scaffold/examples/http/remote/business"
	req "github.com/xinlianit/kit-scaffold/examples/http/repository/request"
	rsp "github.com/xinlianit/kit-scaffold/examples/http/repository/response"
	"github.com/xinlianit/kit-scaffold/examples/http/service"
	"log"
	"time"
)

type IndexEndpoint struct {
	indexService service.IndexService
}

func (e IndexEndpoint) Hello(ctx context.Context, request interface{}) (response interface{}, err error) {
	// 请求断言
	helloReq := request.(req.HelloRequest)

	businessInfoService := business.NewBusinessInfoService()
	defer business.Close()

	result, err := businessInfoService.GetBusinessInfo()

	if err != nil {
		log.Printf("RPC 逻辑错误: %v", err)
		log.Printf("RPC 逻辑错误: %+v", err)
		log.Printf("RPC 逻辑错误: %#v", err)
	}

	// 调用服务
	helloEntity, _ := e.indexService.Hello(helloReq.GetId())

	// 返回响应
	response = rsp.HelloResponse{
		Id:   int64(result.GetBusinessId()),
		Name: helloEntity.GetName(),
	}

	return response, nil
}

func (e IndexEndpoint) Test(ctx context.Context, request interface{}) (response interface{}, err error) {
	log.Println("======= Test")
	// 请求断言
	helloReq := request.(req.HelloRequest)

	// 调用服务
	helloEntity, _ := e.indexService.Hello(helloReq.GetId())

	// 返回响应
	response = rsp.HelloResponse{
		Id:   helloEntity.GetId(),
		Name: helloEntity.GetName(),
	}

	time.Sleep(time.Second * 3)

	return response, nil
}
