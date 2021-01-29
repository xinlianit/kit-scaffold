package config

// 服务中心
type ServiceCenter struct {
	Type string `json:"type"`
	Consul Consul `json:"consul"`
	Register Register `json:"register"`
	HealthCheck HealthCheck `json:"health_check"`
}

type Consul struct{
	Address string `json:"address"`
}

type Register struct{
	Enable bool `json:"enable"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
}

type HealthCheck struct{
	Enable bool `json:"enable"`
	Type string `json:"type"`
	Id string `json:"id"`
	Name string `json:"name"`
	Interval int `json:"interval"`
	Timeout int `json:"timeout"`
	MaxLifeTime int `json:"max_life_time"`
	Protocol string `json:"protocol"`
	Path string `json:"path"`
	Address string `json:"address"`
	Http struct{
		TlsSkipVerify bool `json:"tls_skip_verify"`
		Method string `json:"method"`
	} `json:"http"`
	Grpc struct{
		TlsEnable bool `json:"tls_enable"`
	} `json:"grpc"`
	Gateway struct{
		Id string `json:"id"`
		Name string `json:"name"`
		Interval int `json:"interval"`
		Timeout int `json:"timeout"`
		Method string `json:"method"`
		MaxLifeTime int `json:"max_life_time"`
		Address string `json:"address"`
	} `json:"gateway"`
}