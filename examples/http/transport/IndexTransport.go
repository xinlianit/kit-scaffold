package transport

import (
	"context"
	"encoding/json"
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

func (t IndexTransport) HelloEncode(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}
