package boot

import (
	"github.com/xinlianit/kit-scaffold/boot/nacos"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
)

// 初始化
func Init() {
	// 配置初始化
	config.Init()

	// 日志初始化
	var baseFields []zap.Field
	logger.ZapLogger = logger.ZapInit(logger.NewDefaultZapConfig(), baseFields)

	// 配置中心
	configCenterEnable := config.Config().GetBool("app.configCenter.enable")
	if configCenterEnable {
		// 配置中心类型
		configCenterType := config.Config().GetString("app.configCenter.type")

		switch configCenterType {
		case constant.ConfigCenterTypeNacos:
			nacosConfig()
		default:
			nacosConfig()
		}
	}

	// 服务注册发现

	// 初始化故障转移
	//breaker.Init()
}

// nacos 配置中心
func nacosConfig() {
	// 配置中心初始化
	nacos.Init()
	// 监听并同步配置
	nacos.ListenSyncConfig()
	// 初始化动态配置
	config.InitDynamicConfig()
}