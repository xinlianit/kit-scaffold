package common

import (
	"github.com/google/uuid"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	// 可执行文件路径
	file, _ := exec.LookPath(os.Args[0])
	// 计算可执行文件绝对路径
	path, _ := filepath.Abs(file)
	// 获取字符串(路径分隔符)最后出现的索引位置
	index := strings.LastIndex(path, string(os.PathSeparator))
	// 截取字符串长度(并去除/bin)
	path = path[:index-4]

	return path
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
