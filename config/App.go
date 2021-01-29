package config

// 应用配置
type App struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ConfigCenter ClientConfig `json:"config_center"`
	ServiceCenter ServiceCenter `json:"service_center"`
}
