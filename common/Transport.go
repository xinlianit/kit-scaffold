package common

import (
	"context"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/xinlianit/go-util/util"
	"github.com/xinlianit/kit-scaffold/model"
	"net/http"
)

// 请求解码
func RequestDecode(dec httpTransport.DecodeRequestFunc) httpTransport.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		// 请求数据
		requestData, err := dec(ctx, req)
		if err != nil {
			return nil, err
		}

		request := Request(ctx)

		//return &model.Request{
		//	RequestId:      request.GetRequestId(),
		//	RequestReferer: request.GetRequestReferer(),
		//	RequestTime:    util.TimeUtil().GetCurrentDateTime("2006-01-02 15:04:05"),
		//	RequestBody:    requestData,
		//}, nil

		requestBody := model.Request{
			RequestId:      request.GetRequestId(),
			RequestReferer: request.GetRequestReferer(),
			RequestTime:    util.TimeUtil().GetCurrentDateTime("2006-01-02 15:04:05"),
			RequestBody:    requestData,
		}

		return requestBody, nil
	}
}

// 响应编码
func ResponseEncode(ctx context.Context, w http.ResponseWriter, rsp interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	content, _ := util.SerializeUtil().JsonEncode(rsp)
	if _, err := w.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}
