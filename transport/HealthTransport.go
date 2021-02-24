package transport

import (
	"context"
	req "github.com/xinlianit/kit-scaffold/repository/request"
	"net/http"
)

type HealthTransport struct {

}

func (t HealthTransport) CheckDecode(ctx context.Context, request *http.Request) (interface{}, error) {
	return req.HealthRequest{}, nil
}

func (t HealthTransport) CheckEncode()  {
	
}