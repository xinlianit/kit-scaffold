package interceptor

import (
	"context"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/common/enum"
	"github.com/xinlianit/kit-scaffold/common/exception"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"google.golang.org/grpc"
)

// https://golang2.eddycjy.com/posts/ch3/09-grpc-metadata-creds/

// AuthInterceptor 认证拦截器
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 处理请求前
	authorizeEnable := config.DynamicConfig().GetBool("authorize.enable")
	var authorizeSwitch string
	if authorizeEnable {
		authorizeSwitch = "开启"
	}else{
		authorizeSwitch = "关闭"
	}
	logger.ZapLogger.Sugar().Debugf("应用授权认证: %v", authorizeSwitch)

	// 是否开启应用授权认证
	if config.DynamicConfig().GetBool("authorize.enable") {
		// 请求认证
		if err := auth(ctx); err != nil {
			// 认证不通过返回错误
			return nil, err
		}
	}

	// 处理请求
	resp, err := handler(ctx, req)

	// 处理请求后
	return resp, err
}

// auth 请求认证
func auth(ctx context.Context) error {
	metadataUtil := util.NewMetadataUtil()
	// 解析 metadata
	if _, ok := metadataUtil.ParseMetadata(ctx); ok {
		// 认证结果
		authorized := false

		//// 获取凭证: X-App-Key、X-App-Secret
		appId := metadataUtil.GetStringValue("X-App-Id")
		appKey := metadataUtil.GetStringValue("X-App-Key")
		appSecret := metadataUtil.GetStringValue("X-App-Secret")

		if appId == "" || appKey == "" || appSecret == "" {
			logger.ZapLogger.Sugar().Infof("无效认证参数: X-App-Id: %v, X-App-Key: %v, X-App-Secret: %v", appId, appKey, appSecret)
			return exception.NewExceptionByEnum(enum.CodeInvalidParams)
		}

		// 从配置中心获取应用凭证
		certificates := config.DynamicConfig().Get("authorize.certificate." + appId)
		if certificates == nil {
			return exception.NewExceptionByEnum(enum.CodeUnauthorized)
		}

		// 解析应用授权信息
		certificateList, ok := certificates.([]interface{})
		if ok {
			// 遍历应用授权信息，验证授权是否合法
			for _, certificate := range certificateList {
				certificateData, ok := certificate.(map[interface{}]interface{})
				if ok {
					// 验证成功，终止遍历
					if certificateData["app_key"] == appKey && certificateData["app_secret"] == appSecret {
						authorized = true
						break
					}
				}
			}
		}

		// 凭证验证
		if !authorized {
			logger.ZapLogger.Sugar().Infof("认证失败: X-App-Id: %v, X-App-Key: %v, X-App-Secret: %v", appId, appKey, appSecret)
			return exception.NewExceptionByEnum(enum.CodeAuthorizeFail)
		}
	} else {
		return exception.NewExceptionByEnum(enum.CodeInvalidParams)
	}

	return nil
}