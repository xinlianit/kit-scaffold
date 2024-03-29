package config

import (
	"github.com/xinlianit/kit-scaffold/app"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"path/filepath"
	"strings"
)

var (
	Default map[string]interface{}
)

// 初始化默认值
func init() {
	// todo 添加配置默认项
	Default = make(map[string]interface{})

	// 环境变量
	defaultEnv := strings.ToUpper(constant.EnvPrd)
	Default["env"] = defaultEnv // 环境；DEV-开发、TEST-测试、PRE-预览、PRD-生产

	// 服务
	Default["server.host"] = "0.0.0.0"                // 服务地址
	Default["server.port"] = 80                       // 服务端口
	Default["server.readTimeout"] = 3000              // 请求读超时(单位：毫秒): Accept + Wait + TLSHandshake + Read Request Head + Read Request Body
	Default["server.writeTimeout"] = 5000             // 请求写超时(单位：毫秒): Read Request Head + Read Request Body + Response Write
	Default["server.contextTimeout"] = 5000           // 协程超时(单位：毫秒): 超时自动取消协程
	Default["server.gateway.host"] = "0.0.0.0"        // 网关地址
	Default["server.gateway.port"] = 8080             // 网关端口
	Default["server.grpc.reflection.register"] = true // 是否在gRPC服务中注册reflection服务; true-启用、false-禁用（启用后支持grpcurl命令行工具）

	// 应用
	Default["app.debug"] = false // 是否开启Debug: true-开启、false-关闭

	// 日志记录
	Default["logger.lowestLevel"] = "debug"                                 // 最低记录日志级别; debug、info、warn、error、panic、fatal（级别从低到高）
	Default["logger.recordLineNumber"] = false                              // 是否记录行号; true-是、false-否
	Default["logger.logFormatter"] = "text"                                 // 日志格式; text-文本格式、json-JSON格式
	Default["logger.maxAge"] = 30                                           // 保留旧文件的最大天数
	Default["logger.runtimeLogFile"] = app.RuntimePath + "/runtime.log" 	// 应用运行日志文件
	Default["logger.errorLogFile"] = app.LogPath + "/error.log"     		// 错误日志文件
	Default["logger.rotate.enable"] = true                                  // 是否开启日志切割
	Default["logger.rotate.type"] = "date"                                  // 日志切割类型; size-按大小切割、date-按日期切割
	Default["logger.rotate.size.maxSize"] = 10                              // 在进行切割之前，日志文件的最大大小（以MB为单位)
	Default["logger.rotate.size.maxBackups"] = 100                          // 保留旧文件的最大个数
	Default["logger.rotate.size.compress"] = true                           // 是否压缩/归档旧文件
	Default["logger.rotate.date.extend"] = ".%Y%m%d"                        // 切割后缀
	Default["logger.access.enable"] = true                                  // 访问日志记录：true-开启、false-关闭
	Default["logger.access.logFile"] = app.LogPath + "/access.log"  // 访问日志文件
	Default["logger.rpc.enable"] = true                                     // RPC调用日志记录：true-开启、false-关闭s
	Default["logger.rpc.logFile"] = app.LogPath + "/rpc.log"        // RPC 调用日志

	// 配置中心
	Default["app.configCenter.enable"] = false                                                  // 是否启用配置中心; true-启用、false-关闭
	Default["app.configCenter.type"] = "nacos"                                                  // 配置中心类型
	Default["app.configCenter.configCacheDir"] = filepath.Join(app.CachePath, "config") // 动态配置缓存目录

	// 服务中心
	Default["app.serviceCenter.register.enable"] = false                      // 是否注册服务; true-是、false-否
	Default["app.serviceCenter.type"] = "consul"                              // 服务中心类型: consul、etcd
	// 健康检测
	Default["app.serviceCenter.healthCheck.enable"] = true // 是否启用健康检查；true-是、false-否
	Default["app.serviceCenter.healthCheck.type"] = "http" // 健康检查类型；http、tcp、grpc
	Default["app.serviceCenter.healthCheck.method"] = "GET" // 检测请求方式；GET、POST...
	Default["app.serviceCenter.healthCheck.interval"] = 3000 // 检测间隔（单位：毫秒）
	Default["app.serviceCenter.healthCheck.timeout"] = 1000 // 检查超时（单位：毫秒）
	Default["app.serviceCenter.healthCheck.maxLifeTime"] = 60000 // 服务最大生存周期（单位：毫秒）
	Default["app.serviceCenter.healthCheck.grpc.tls.enable"] = false // 是否启用TLS；true-是、false-否
	Default["app.serviceCenter.healthCheck.gateway.interval"] = 3000 // 检测间隔（单位：毫秒）
	Default["app.serviceCenter.healthCheck.gateway.timeout"] = 1000 // 检测超时（单位：毫秒）
	Default["app.serviceCenter.healthCheck.gateway.method"] = "GET" // 检测请求方式；默认：GET
	// consul
	Default["app.serviceCenter.consul.address"] = "127.0.0.1:8500" // 服务中心地址
}
