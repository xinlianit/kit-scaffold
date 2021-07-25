package interceptor

import (
	"context"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/common"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
)

// logChannel 日志通道
var logChannel = make(chan accessLog, 100)

// accessLog 访问日志
type accessLog struct {
	// 请求ID
	RequestId string
	// RequestMethod 请求方法
	RequestMethod string
	// RequestMetadata 请求元数据
	RequestMetadata map[string][]string
	// RequestTime 请求时间
	RequestTime string
	// ResponseTime 响应时间
	ResponseTime string
	// CostTime 请求耗时
	CostTime int64
	// Request 请求数据
	Request interface{}
	// Response 响应数据
	Response interface{}
	// ResponseError 响应错误
	ResponseError error
}

// AccessInterceptor 访问拦截器
func AccessInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 处理请求
	startTime := util.TimeUtil().GetCurrentMilliTime()

	// 请求实例
	request := common.Request(ctx)

	resp, err = handler(request.GetContext(), req)

	endTime := util.TimeUtil().GetCurrentMilliTime()

	// 是否开启日志
	if config.Config().GetBool("logger.access.enable") {
		// 异步记录日志
		go logRecord()

		// 记录日志
		logBody := accessLog{
			RequestId:       request.GetRequestId(),
			RequestMethod:   info.FullMethod,
			RequestMetadata: request.GetMetadata(),
			RequestTime:     util.TimeUtil().MillisecondToDateTime(startTime, constant.DateTimeMillisecondLayout),
			ResponseTime:    util.TimeUtil().MillisecondToDateTime(endTime, constant.DateTimeMillisecondLayout),
			CostTime:        endTime - startTime,
			Request:         req,
			Response:        resp,
			ResponseError:   err,
		}

		logChannel <- logBody
	}

	return resp, err
}

// logRecord 记录日志
func logRecord() {
	// 配置
	zapConfig := logger.NewDefaultZapConfig()
	// 禁用文件及行号
	zapConfig.RecordLineNumber = false

	// 日志基础字段
	hostName, _ := os.Hostname()
	hostIp := util.ServerUtil().GetServerIp()
	baseFields := []zap.Field{
		zap.String("host_name", hostName),
		zap.String("host_ip", hostIp),
	}

	accessLogFile := config.Config().GetString("logger.access.logFile")
	accessLogger := logger.NewLoggerOfZap(accessLogFile, zapConfig, baseFields)

	// 接收通道数据
	for logMsg := range logChannel {
		// 返回数据json序列化
		requestMetadata, err := util.SerializeUtil().JsonEncode(logMsg.RequestMetadata)
		if err != nil {
			logger.ZapLogger.Error(err.Error())
		}

		// 请求数据
		request, err := util.SerializeUtil().JsonEncode(logMsg.Request)
		if err != nil {
			logger.ZapLogger.Error(err.Error())
		}

		// 响应数据
		response, err := util.SerializeUtil().JsonEncode(logMsg.Response)
		if err != nil {
			logger.ZapLogger.Error(err.Error())
		}

		// 日志字段
		fields := []zap.Field{
			zap.String("request_id", logMsg.RequestId),
			zap.String("request_method", logMsg.RequestMethod),
			zap.String("request_metadata", requestMetadata),
			zap.String("request_time", logMsg.RequestTime),
			zap.String("response_time", logMsg.ResponseTime),
			zap.Int64("const_time", logMsg.CostTime),
			zap.String("request", request),
			zap.String("response", response),
		}

		// 记录响应错误
		if logMsg.ResponseError != nil {
			fields = append(fields, zap.String("response_error", logMsg.ResponseError.Error()))
		}

		// 记录日志
		accessLogger.Info("Access", fields...)
	}
}