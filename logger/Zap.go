package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/xinlianit/go-util"
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
	LogFormatter     string
}

// 默认配置
func NewDefaultZapConfig() ZapConfig {
	// 日志切割类型
	rotateType := util.RotateTypeDate
	if Config.Rotate.Type == "size" {
		rotateType = util.RotateTypeSize
	}

	// 最低日志级别
	lowestLevel := zapcore.DebugLevel
	switch Config.LowestLevel {
	case "info":
		lowestLevel = zapcore.InfoLevel
	case "warn":
		lowestLevel = zapcore.WarnLevel
	case "error":
		lowestLevel = zapcore.ErrorLevel
	case "panic":
		lowestLevel = zapcore.PanicLevel
	case "fatal":
		lowestLevel = zapcore.FatalLevel
	default:
		lowestLevel = zapcore.DebugLevel
	}

	return ZapConfig{
		RotateEnable:     Config.Rotate.Enable,
		RotateType:       rotateType,
		MaxSize:          Config.Rotate.Size.MaxSize,
		MaxBackups:       Config.Rotate.Size.MaxBackups,
		Compress:         Config.Rotate.Size.Compress,
		Extend:           Config.Rotate.Date.Extend,
		MaxAge:           Config.MaxAge,
		LogFile:          Config.RuntimeLogFile,
		ErrorLogFile:     Config.ErrorLogFile,
		LowestLevel:      lowestLevel,
		RecordLineNumber: Config.RecordLineNumber,
		LogFormatter:     Config.LogFormatter,
	}
}

// Zap 日志记录器
// @param cfg 初始化配置
// @param baseFields 日志基础字段
func ZapInit(cfg ZapConfig, baseFields []zap.Field) *zap.Logger {
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
	// 编码器配置
	encoderConfig := util.ZapUtil().NewDefaultEncoderConfig()
	// 日志格式
	if cfg.LogFormatter == "text" {
		// 文本编码
		zapUtilConfig.LogFormatter = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		// JSON编码
		zapUtilConfig.LogFormatter = zapcore.NewJSONEncoder(encoderConfig)
	}
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
