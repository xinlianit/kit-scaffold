package transport

import (
	"context"
	req "github.com/xinlianit/kit-scaffold/examples/http/repository/request"
	"net/http"
)

type IndexTransport struct {
}

func (t IndexTransport) HelloDecode(ctx context.Context, request *http.Request) (interface{}, error) {
	helloReq := req.HelloRequest{
		Message: "你好，世界！",
	}

	return helloReq, nil
}
