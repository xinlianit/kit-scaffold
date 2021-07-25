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

	// 时间格式
	DateLayout = "2006-01-02"
	TimeLayout = "15:04:05"
	DateTimeLayout = "2006-01-02 15:04:05"
	DateTimeMillisecondLayout = "2006-01-02 15:04:05.000"
)
