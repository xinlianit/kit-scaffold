package common

import (
	"context"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/xinlianit/go-util/util"
	"github.com/xinlianit/kit-scaffold/model"
	"log"
	"net/http"
)

func Transport() transport {
	return transport{}
}

type transport struct {
}

// 请求解码
func (t transport) RequestDecode(data interface{}) httpTransport.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		request := Request(ctx)

		requestBody := model.Request{
			RequestId:      request.GetRequestId(),
			RequestReferer: request.GetRequestReferer(),
			RequestTime:    util.TimeUtil().GetCurrentDateTime("2006-01-02 15:04:05"),
			Data:           data,
		}

		log.Println("----------------", requestBody)

		return requestBody, nil
	}
}

// 响应编码
func (t transport) ResponseEncode(ctx context.Context, w http.ResponseWriter, rsp interface{}) error {
	response := Response()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	content, _ := util.SerializeUtil().JsonEncode(response.Success(rsp))
	if _, err := w.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}
