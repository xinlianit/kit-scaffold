package constant

const (
	// 环境配置
	EnvDev  = "DEV"
	EnvTest = "TEST"
	EnvPre  = "PRE"
	EnvPrd  = "PRD"

	// 配置中心类型
	ConfigCenterTypeNacos = "nacos"

	// 服务中心类型
	ServiceCenterTypeEtcd   = "etcd"
	ServiceCenterTypeConsul = "consul"
	ServiceCenterTypeNacos  = "nacos"

	// 默认时间格式
	DefaultTimeLayout = "2006-01-02 15:04:05"
)
