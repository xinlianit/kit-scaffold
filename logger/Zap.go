package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/xinlianit/go-util/util"
	"github.com/xinlianit/kit-scaffold/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	ZapLogger *zap.Logger
)

// Zap 配置
type ZapConfig struct {
	RotateEnable     bool
	RotateType       util.RotateType
	MaxSize          int
	MaxBackups       int
	MaxAge           int
	Compress         bool
	Extend           string
	LogFile          string
	ErrorLogFile     string
	LowestLevel      zapcore.Level
	RecordLineNumber bool
}

// 默认配置
func NewDefaultZapConfig() ZapConfig {
	// todo 读取yml配置
	return ZapConfig{
		RotateEnable:     true,
		RotateType:       util.RotateTypeDate,
		MaxSize:          10,
		MaxBackups:       100,
		MaxAge:           30,
		Compress:         true,
		Extend:           ".%Y%m%d",
		LogFile:          config.Config().GetString("app.log.runtimeLogFile"),
		ErrorLogFile:     config.Config().GetString("app.log.errorLogFile"),
		LowestLevel:      zapcore.DebugLevel,
		RecordLineNumber: true,
	}
}

// Zap 日志记录器
// @param cfg 初始化配置
// @param baseFields 日志基础字段
func ZapInit(cfg ZapConfig, baseFields []zap.Field) *zap.Logger {
	// 编码器配置
	encoderConfig := util.ZapUtil().NewDefaultEncoderConfig()

	// 大小切割配置
	rotateSizeConfig := util.ZapUtil().NewDefaultRotateSizeConfig()
	rotateSizeConfig.MaxSize = cfg.MaxSize       // 在进行切割之前，日志文件的最大大小（以MB为单位)
	rotateSizeConfig.MaxBackups = cfg.MaxBackups // 保留旧文件的最大个数
	rotateSizeConfig.MaxAge = cfg.MaxAge         // 保留旧文件的最大天数
	rotateSizeConfig.Compress = cfg.Compress     // 是否压缩/归档旧文件

	// 日期切割配置
	rotateDateConfig := util.ZapUtil().NewDefaultRotateDateConfig()
	rotateDateConfig.Extend = cfg.Extend // 日志切割后缀扩展名
	rotateDateConfig.Options = []rotatelogs.Option{
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge) * 24 * time.Hour), // 设置最大保存时间(30天)
		rotatelogs.WithRotationTime(24 * time.Hour),                       // 设置日志切割时间间隔(1天)
	}

	// 默认配置
	zapUtilConfig := util.ZapUtil().NewDefaultConfig()
	// 日志文件
	zapUtilConfig.LogFile = cfg.LogFile
	// 错误日志文件
	zapUtilConfig.ErrorLogFile = cfg.ErrorLogFile
	// 日志格式
	zapUtilConfig.LogFormatter = zapcore.NewJSONEncoder(encoderConfig)
	// 最低记录日志级别
	zapUtilConfig.LowestLevel = cfg.LowestLevel
	// 是否记录行号
	zapUtilConfig.RecordLineNumber = cfg.RecordLineNumber
	// 日志基础字段
	zapUtilConfig.BaseFields = baseFields
	// 日志切割
	zapUtilConfig.RotateEnable = cfg.RotateEnable
	// 切割类型
	zapUtilConfig.RotateType = cfg.RotateType
	// 大小切割配置
	zapUtilConfig.RotateSizeConfig = rotateSizeConfig
	// 日期切割配置
	zapUtilConfig.RotateDateConfig = rotateDateConfig

	// 日志初始化
	return util.ZapUtil().Init(zapUtilConfig)
}

// 创建Zap日志记录器
// @param cfg 初始化配置
// @param fileName 文件名
// @param baseFields 日志基础字段
func NewLoggerOfZap(fileName string, cfg ZapConfig, baseFields []zap.Field) *zap.Logger {
	// 日志记录单文件
	cfg.LogFile = fileName
	cfg.ErrorLogFile = ""

	return ZapInit(cfg, baseFields)
}
