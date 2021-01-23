package business

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"google.golang.org/grpc"
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
		grpc.WithInsecure(),
	}

	conn, err = grpc.Dial(serverAddress, dialOptions...)

	if err != nil {
		panic(err)
	}

	return conn
}

// 获取上下文
func getContext() (context.Context, context.CancelFunc) {
	// TODO metadata 运用
	return  context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond * 1000))
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
