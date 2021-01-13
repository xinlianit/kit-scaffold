package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xinlianit/go-util/util"
	"github.com/xinlianit/kit-scaffold/config"
	"time"
)

// Logrus 日志记录器
func LogrusInit() {
	logger := util.LoggerUtil()
	logger.Init(util.LoggerUtilConfig{
		// 日志文件
		LogFile: config.Config().GetString("app.log.runtimeLogFile"),
		// 日志格式: json - logrus.JSONFormatter{}、text: logrus.xtFormatter{}
		LogFormatter: &logrus.JSONFormatter{
			// 时间格式
			TimestampFormat: "2006-01-02 15:04:05",
		},
		// 记录日志最低级别
		LowestLevel: logrus.DebugLevel,
		// 开启日志切割
		RotateEnable: true,
		// 切割日志后缀
		RotateExtend: ".%Y%m%d",
		// 日志切割配置
		RotateOptions: []rotatelogs.Option{
			rotatelogs.WithLinkName(config.Config().GetString("app.log.runtimeLogFile")), // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(7 * 24 * time.Hour),                                    // 设置最大保存时间(7天)
			rotatelogs.WithRotationTime(24 * time.Hour),                                  // 设置日志切割时间间隔(1天)
		},
	})
}
