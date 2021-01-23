package business

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/examples/http/remote/business/credential"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

var (
	conn *grpc.ClientConn
	err  error
)

func connect() *grpc.ClientConn {
	// 创建连接
	serverAddress := "127.0.0.1:8080"

	dialOptions := []grpc.DialOption{
		// 忽略TLS验证
		grpc.WithInsecure(),
		// grpc 客户端连接默认为异步连接，若需要同步则调用 WithBlock()，完成状态: Ready (确保握手成功)
		grpc.WithBlock(),
		// 应用凭证
		grpc.WithPerRPCCredentials(credential.AppCredential{}),
	}

	conn, err = grpc.Dial(serverAddress, dialOptions...)

	if err != nil {
		panic(err)
	}

	return conn
}

// 获取上下文
func getContext() (context.Context, context.CancelFunc) {
	// 创建 context
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*1000))

	metadataUtil := util.NewMetadataUtil()
	// 设置 metadata 到 context
	mdKvs := map[string]interface{}{
		"X-Request-Id": "202101231700123", // 请求ID
	}
	metadataUtil.SetMetadata(mdKvs)
	ctx = metadata.NewOutgoingContext(ctx, metadataUtil.GetMetadata())
	return ctx, cancel
}

// 商家信息服务实例
func NewBusinessInfoService() businessInfoService {
	return businessInfoService{
		client: service.NewBusinessInfoServiceClient(connect()),
	}
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}
