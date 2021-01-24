package business

import (
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"github.com/xinlianit/kit-scaffold/examples/http/remote/business/credential"
	"github.com/xinlianit/kit-scaffold/examples/http/remote/business/interceptor"
	"google.golang.org/grpc"
)

var (
	conn *grpc.ClientConn
	err  error
)

func connect() *grpc.ClientConn {
	// 连接器
	interceptors := []grpc.UnaryClientInterceptor{
		// 请求拦截器
		interceptor.RequestInterceptor,
	}

	// 连接参数
	dialOptions := []grpc.DialOption{
		// 忽略TLS验证
		grpc.WithInsecure(),
		// grpc 客户端连接默认为异步连接，若需要同步则调用 WithBlock()，完成状态: Ready (确保握手成功)
		grpc.WithBlock(),
		// 应用凭证
		grpc.WithPerRPCCredentials(credential.AppCredential{}),
		// 注册拦截器
		grpc.WithChainUnaryInterceptor(interceptors...),
	}

	// 创建连接
	serverAddress := "127.0.0.1:8080"
	conn, err = grpc.Dial(serverAddress, dialOptions...)

	if err != nil {
		panic(err)
	}

	return conn
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
