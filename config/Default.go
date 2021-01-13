package config

import (
	"github.com/gin-gonic/gin"
	"github.com/xinlianit/kit-scaffold/common"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"path/filepath"
	"strings"
)

var (
	Default map[string]interface{}
)

// 初始化默认值
func init() {
	envVal := common.GetEnv("APP_ENV", constant.EnvPrd)
	env, ok := envVal.(string)
	if !ok {
		env = constant.EnvPrd
	} else {
		env = strings.ToUpper(env)
	}

	// todo 添加配置默认项
	Default = make(map[string]interface{})

	// 环境变量
	Default["APP_ENV"] = env // 环境；DEV-开发、TEST-测试、PRE-预览、PRD-生产

	// 服务
	Default["server.readTimeout"] = 120    // 读超时时间(单位：秒)
	Default["server.writeTimeout"] = 120   // 写超时时间(单位：秒)
	Default["server.contextTimeout"] = 300 // 上下文超时时间(单位：秒)

	// 应用
	Default["app.host"] = "0.0.0.0"         // 应用地址
	Default["app.port"] = 21000             // 应用端口
	Default["app.appMod"] = gin.ReleaseMode // 环境模式: release-发布模式、test-测试模式、debug-开发调试模式
	Default["app.debug"] = false            // 是否开启Debug: true-开启、false-关闭

	// 应用日志
	Default["app.log.accessLogEnable"] = true                                // 访问日志记录：true-开启、false-关闭
	Default["app.log.rpcLogEnable"] = true                                   // RPC调用日志记录：true-开启、false-关闭s
	Default["app.log.runtimeLogFile"] = common.GetLogPath() + "/runtime.log" // 应用运行日志文件
	Default["app.log.accessLogFile"] = common.GetLogPath() + "/access.log"   // 访问日志文件
	Default["app.log.errorLogFile"] = common.GetLogPath() + "/error.log"     // 错误日志文件
	Default["app.log.rpcLogFile"] = common.GetLogPath() + "/rpc.log"         // RPC 调用日志

	// 配置中心
	Default["app.configCenter.enable"] = true                                                   // 是否启用配置中心; true-启用、false-关闭
	Default["app.configCenter.type"] = "nacos"                                                  // 配置中心类型
	Default["app.configCenter.nacosDefaultGroup"] = env                                         // nacos 默认分组
	Default["app.configCenter.configCacheDir"] = filepath.Join(common.GetCachePath(), "config") // 动态配置缓存目录

	// 服务中心
	Default["app.serviceCenter.register.enable"] = false // 是否注册服务; true-是、false-否
	Default["app.serviceCenter.type"] = "etcd"           // 服务中心类型: etcd、consul
}
