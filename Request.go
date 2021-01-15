package scaffold

import (
	"context"
	"github.com/google/uuid"
)

func Request(ctx context.Context) *request {
	return &request{ctx: ctx}
}

// 请求
type request struct {
	ctx context.Context
}

// 获取请求ID
func (r *request) GetRequestId() string {
	requestId, ok := r.ctx.Value("X-Request-Id").(string)
	if !ok || requestId == "" {
		id, _ := uuid.NewRandom()
		requestId = id.String()
	}
	context.WithValue(r.ctx, "X-Request-Id", requestId)

	return requestId
}

// 请求来源
func (r request) GetRequestReferer() string {
	requestReferer, ok := r.ctx.Value("X-Request-Referer").(string)
	if ok {
		return requestReferer
	}
	return ""
}
