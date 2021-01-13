package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
)

func Test2Middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		log.Println("======== Test2Middleware =========")
		return next(ctx, request)
	}
}
