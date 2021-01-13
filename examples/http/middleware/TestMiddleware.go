package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
)

func TestMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		log.Println("======== TestMiddleware =========")
		return next(ctx, request)
	}
}
