package config

// Nacos 配置
type Nacos struct {
	ClientConfig ClientConfig   `json:"client_config"` // 客户端配置
	ServerConfig []ServerConfig `json:"server_config"` // 服务端配置
}

// 客户端配置
type ClientConfig struct {
	NamespaceId          string `json:"namespace_id"`            // 命名空间ID
	Timeout              uint64 `json:"timeout"`                 // 请求Nacos服务端超时时间;默认:10000 ms
	Endpoint             string `json:"endpoint"`                // 获取Nacos服务列表endpoint地址
	RegionId             string `json:"region_id"`               // kms的regionId，用于配置中心的鉴权
	AccessKey            string `json:"access_key"`              // kms的AccessKey，用于配置中心的鉴权
	SecretKey            string `json:"secret_key"`              // kms的SecretKey，用于配置中心的鉴权
	OpenKms              bool   `json:"open_kms"`                // 是否开启kms，默认不开启，kms可以参考文档 https: //help.aliyun.com/product/28933.html
	CacheDir             string `json:"cache_dir"`               // 缓存service信息目录，默认是当前运行目录
	UpdateThreadNum      int    `json:"update_thread_num"`       // 监听service变化并发数; 默认: 20
	NotLoadCacheAtStart  bool   `json:"not_load_cache_at_start"` // 在启动的时候不读取缓存在 CacheDir 的 service 信息
	UpdateCacheWhenEmpty bool   `json:"update_cache_when_empty"` // 当service返回的实例列表为空时，不更新缓存，用于推空保护
	UserName             string `json:"user_name"`               // Nacos服务端的API鉴权Username
	Password             string `json:"password"`                //  Nacos服务端的API鉴权Password
	LogDir               string `json:"log_dir"`                 // 日志存储路径
	RotateTime           string `json:"rotate_time"`             // 日志轮转周期; 比如：30m, 1h, 24h, 默认: 24h
	MaxAge               int64  `json:"max_age"`                 // 日志最大文件数，默认: 3
	LogLevel             string `json:"log_level"`               // 日志默认级别; 值必须是: debug,info,warn,error，默认值: info
}

// 服务端配置
type ServerConfig struct {
	Scheme      string `json:"scheme"`       // Nacos Scheme
	ContextPath string `json:"context_path"` // Nacos ContextPath
	IpAddr      string `json:"ip_addr"`      // Nacos服务地址
	Port        uint64 `json:"port"`         // Nacos服务端口
}
