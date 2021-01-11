package transport

import (
	"context"
	"encoding/json"
	httpTransport "github.com/go-kit/kit/transport/http"
	req "github.com/xinlianit/kit-scaffold/test/repository/request"
	"net/http"
)

type IndexTransport struct {
}

func (t IndexTransport) HelloDecode() httpTransport.DecodeRequestFunc {
	return func(ctx context.Context, request *http.Request) (interface{}, error) {
		helloReq := req.HelloRequest{
			Message: "你好，世界！",
		}

		return helloReq, nil
	}
}

func (t IndexTransport) HelloEncode() httpTransport.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")

		return json.NewEncoder(w).Encode(response)
	}
}
