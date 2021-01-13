package common

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

// 获取环境变量
func GetEnv(key string, defaultValue interface{}) interface{} {
	appEnv := os.Getenv(key)
	if appEnv == "" {
		return defaultValue
	}

	return appEnv
}

// 获取应用根路径
func GetAppRootPath() string {
	appPath, _ := os.Getwd()
	return appPath
}

// 获取应用根路径
func GetLogPath() string {
	return filepath.Join(GetAppRootPath(), "logs")
}

// 获取配置路径
func GetConfigPath() string {
	return filepath.Join(GetAppRootPath(), "config")
}

// 获取资源路径
func GetResourcePath() string {
	return filepath.Join(GetAppRootPath(), "resource")
}

// 获取运行时路径
func GetRuntimePath() string {
	return filepath.Join(GetAppRootPath(), "runtime")
}

// 获取缓存路径
func GetCachePath() string {
	return filepath.Join(GetRuntimePath(), "cache")
}

// 生成UUID
func GenerateUUID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
