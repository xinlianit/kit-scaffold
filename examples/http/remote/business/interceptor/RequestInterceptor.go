package interceptor

import (
	"context"
	"github.com/xinlianit/go-util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

// 请求拦截器
func RequestInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 设置请求超时时间
	var cancel context.CancelFunc
	ctx, cancel = context.WithDeadline(ctx, time.Now().Add(time.Millisecond*1000))
	defer cancel()

	metadataUtil := util.NewMetadataUtil()
	// 设置 metadata 到 context
	mdKvs := map[string]interface{}{
		"X-Request-Id": "202101231700123666666", // 请求ID
	}
	metadataUtil.SetMetadata(mdKvs)
	ctx = metadata.NewOutgoingContext(ctx, metadataUtil.GetMetadata())

	log.Printf("before invoker. method: %+v, request:%+v, ctx: %#v, ctx-2: %+v", method, req, ctx, ctx)

	// 引用方法
	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Printf("after invoker. reply: %+v", reply)
	return err
}
