package config

// 配置中心
type ConfigCenter struct {
	Enable bool `json:"enable"`
	Type string `json:"type"`
	SyncConfigDataIds string `json:"sync_config_data_ids"` // TODO: 字符串改切片
	ConfigCacheDir string `json:"config_cache_dir"`
}
