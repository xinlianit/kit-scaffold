package common

import (
	"github.com/google/uuid"
	"os"
)

// GetEnv 获取环境变量
func GetEnv(key string, defaultValue interface{}) interface{} {
	appEnv := os.Getenv(key)
	if appEnv == "" {
		return defaultValue
	}

	return appEnv
}

// GenerateUUID 生成UUID
func GenerateUUID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
