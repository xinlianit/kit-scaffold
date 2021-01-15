package endpoint

import (
	"context"
	"fmt"
	req "github.com/xinlianit/kit-scaffold/examples/http/repository/request"
	rsp "github.com/xinlianit/kit-scaffold/examples/http/repository/response"
	"github.com/xinlianit/kit-scaffold/examples/http/service"
	"github.com/xinlianit/kit-scaffold/server"
	"log"
	"time"
)

type IndexEndpoint struct {
	indexService service.IndexService
}

func (e IndexEndpoint) Hello(ctx context.Context, request interface{}) (response interface{}, err error) {

	fmt.Println(fmt.Sprintf("============requestObject: %#v", server.Request))
	fmt.Println("===========", server.Request.GetRequestId())

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

func (e IndexEndpoint) Test(ctx context.Context, request interface{}) (response interface{}, err error) {
	log.Println("======= Test")
	// 请求断言
	helloReq := request.(req.HelloRequest)

	// 调用服务
	helloEntity, _ := e.indexService.Hello(helloReq.GetId())

	// 返回响应
	response = rsp.HelloResponse{
		Data: helloEntity,
	}

	time.Sleep(time.Second * 3)

	return response, nil
}
