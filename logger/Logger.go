package logger

import "github.com/xinlianit/kit-scaffold/config"

var (
	Config config.Logger
)

func Init()  {
	// 解析日志配置
	if err := config.Config().UnmarshalKey("logger", &Config); err != nil {
		panic(err)
	}
}
