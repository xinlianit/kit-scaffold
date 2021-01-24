package interceptor

import (
	"context"
	"github.com/xinlianit/go-util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// 认证拦截器
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error){
	// 处理请求前
	log.Printf("before handling. Info: %+v", info)

	// 请求认证
	if err := auth(ctx); err != nil {
		// 认证不通过返回错误
		return nil, err
	}

	// 处理请求
	resp, err := handler(ctx, req)

	// 处理请求后
	log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

// 请求认证
func auth(ctx context.Context) error {
	metadataUtil := util.NewMetadataUtil()
	// 解析 metadata
	if md, ok := metadataUtil.ParseMetadata(ctx); ok {
		log.Print(md)

		// 获取凭证: app_id、app_secret
		appId := metadataUtil.GetStringValue("X-App-Id")
		appSecret := metadataUtil.GetStringValue("X-App-Secret")

		// 凭证验证
		if appId != "kit-scaffold.palm.http.api" || appSecret != "666666" {
			return status.Error(2, "客户端凭证验证失败")
		}
	}else{
		return status.Error(codes.OK, "客户端凭证验证失败")
	}

	return nil
}
