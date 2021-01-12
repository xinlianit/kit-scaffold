package boot

func Init() {
	//// 配置初始化
	//config.Init()
	//
	//// 日志初始化
	//var baseFields []zap.Field
	//logger.ZapLogger = logger.ZapInit(logger.NewDefaultZapConfig(), baseFields)
	//
	//// 是否开启配置中心
	//configCenterEnable := config.Config().GetBool("app.configCenter.enable")
	//if configCenterEnable {
	//	// 配置中心类型
	//	configCenterType := config.Config().GetString("app.configCenter.type")
	//
	//	switch configCenterType {
	//	// nacos 配置中心
	//	case constant.ConfigCenterTypeNacos:
	//		// 配置中心初始化
	//		nacos.Init()
	//		// 监听并同步配置
	//		nacos.ListenSyncConfig()
	//		// 初始化动态配置
	//		config.InitDynamicConfig()
	//		break
	//	default:
	//		// 配置中心初始化
	//		nacos.Init()
	//		// 监听并同步配置
	//		nacos.ListenSyncConfig()
	//		// 初始化动态配置
	//		config.InitDynamicConfig()
	//		break
	//	}
	//}

	// 初始化故障转移
	//breaker.Init()
}
