package breaker

import (
	"gitee.com/plam-bfa/bfa-gateway-api/service/fallback"
	"github.com/afex/hystrix-go/hystrix"
	"reflect"
	"time"
)

// 定义 hystrix 命令
const (
	CommandGoodsGetGoods = "Goods.GetGoods"
)

var (
	commandConfig hystrix.CommandConfig
	fallbackServiceList = make(map[string]reflect.Type)
	fallbackServiceApiList = make(map[string]string)
)

// 初始化
func Init()  {
	// 初始化 hystrix 配置
	commandConfig = hystrix.CommandConfig{
		Timeout:                int(time.Second * 3),	// 执行超时时间(单位：毫秒), 默认: 1s (当请求时间超过超时时间，触发熔断)
		MaxConcurrentRequests:  100, 	// 最大并发请求(QPS), 默认: 10 (当并发调用超过最大并发数，触发熔断)

		// RequestVolumeThreshold: 单位时间内(10s), 当请求调用次数达到或超过20次, ErrorPercentThreshold: 且错误率超过50%，触发熔断
		RequestVolumeThreshold: 30,	// 单位时间内(10s), 触发熔断的最低请求次数, 默认: 20
		ErrorPercentThreshold:  50,	// 触发熔断要错误率, 默认: 50%
		SleepWindow:            int(time.Second * 5),	// 休眠时间窗(单位：毫秒)，默认: 5s (当熔断开启时，经过休眠时间后，再次检测是否需要开启熔断)
	}

	// 配置 hystrix 命令
	//hystrix.Configure()  // 配置多个命令
	hystrix.ConfigureCommand(CommandGoodsGetGoods, commandConfig) // 获取商品

	// 注册本地降级服务
	RegisterFallbackService()
	RegisterFallbackServiceApi()
}

// 注册本地降级服务
func RegisterFallbackService()  {
	// todo: 服务名配置化
	fallbackServiceList["com.palm.service.goods"] = reflect.TypeOf(fallback.GoodsServiceFallback{})
}

// 注册本地降级服务接口
func RegisterFallbackServiceApi()  {
	// todo: 服务名配置化
	fallbackServiceApiList["com.palm.service.goods.Goods.GetGoods"] = "GetGoods"
}

// 获取本地降级服务注册表
func GetRegisterFallbackService() map[string]reflect.Type {
	return fallbackServiceList
}

// 获取本地降级服务接口注册表
func GetRegisterFallbackServiceApi() map[string]string {
	return fallbackServiceApiList
}