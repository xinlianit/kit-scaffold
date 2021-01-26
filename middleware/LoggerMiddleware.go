package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/common"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
)

// 日志通道
var logChannel = make(chan logBody, 100)

// 日志实体
type logBody struct {
	RequestId        string `json:"request_id"`         // 请求ID
	RequestTime      int64  `json:"request_time"`       // 请求时间
	ResponseTime     int64  `json:"response_time"`      // 响应时间
	CostTime         int64  `json:"cost_time"`          // 耗时
	RequestReferer   string `json:"request_referer"`    // 请求来源
	RequestRefererIp string `json:"request_referer_ip"` // 请求来源IP
	RequestTarget    string `json:"request_target"`     // 请求目标
	RequestBody      string `json:"request_body"`       // 请求体
	ResponseBody     string `json:"response_body"`      // 响应体
	ReturnCode       int    `json:"return_code"`        // 返回状态码
	ReturnMsg        string `json:"return_msg"`         // 返回信息
	ReturnData       string `json:"return_data"`        // 返回数据
}

// 日志字段
var logFields = struct {
	RequestId        string
	RequestTime      string
	ResponseTime     string
	CostTime         string
	RequestReferer   string
	RequestRefererIp string
	RequestTarget    string
	RequestBody      string
	ResponseBody     string
	ReturnCode       string
	ReturnMsg        string
	ReturnData       string
	ClientIp         string
}{
	RequestId:        "request_id",
	RequestTime:      "request_time",
	CostTime:         "cost_time",
	RequestReferer:   "request_referer",
	RequestRefererIp: "request_referer_ip",
	RequestTarget:    "request_target",
	RequestBody:      "request_body",
	ResponseBody:     "response_body",
	ReturnCode:       "return_code",
	ReturnMsg:        "return_msg",
	ReturnData:       "return_data",
	ClientIp:         "client_ip",
}

// 日志中间件
func LoggerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	// 访问日志记录
	if !config.Config().GetBool("logger.access.enable") {
		return next
	}

	// 接收并记录日志
	go receiveAndRecordLog()

	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
		// 处理请求前
		logrus.Printf("[before] 日志中间件")

		// 请求对象
		request := common.Request(ctx)

		// 请求体
		reqBody, err := util.SerializeUtil().JsonEncode(req)
		if err != nil {
			logger.ZapLogger.Sugar().Errorf("request json encode error: %v", err.Error())
		}

		log := logBody{
			RequestId:        request.GetRequestId(),
			RequestTime:      util.TimeUtil().GetCurrentMilliTime(),
			RequestReferer:   request.GetRequestReferer(),
			RequestRefererIp: "", // todo 请求来源IP
			RequestTarget:    "", // todo 请求目标
			RequestBody:      reqBody,
			ReturnCode:       0,  // todo 返回状态码
			ReturnMsg:        "", // todo 返回状态信息
			ReturnData:       "", // todo 返回数据
		}

		// 执行请求
		rsp, err = next(ctx, req)

		// 处理请求后
		logrus.Printf("[after] 日志中间件")

		// 响应数据
		log.ResponseTime = util.TimeUtil().GetCurrentMilliTime()
		// 耗时(RT)
		log.CostTime = log.ResponseTime - log.RequestTime

		// 响应体
		rspBody, err := util.SerializeUtil().JsonEncode(rsp)
		if err != nil {
			logger.ZapLogger.Sugar().Errorf("response json encode error: %v", err.Error())
		}
		log.ResponseBody = rspBody

		// 发送日志通道
		logChannel <- log
		return
	}
}

// 接收并记录日志
func receiveAndRecordLog() {
	// 配置
	zapConfig := logger.NewDefaultZapConfig()
	// 禁用文件及行号
	zapConfig.RecordLineNumber = false

	// 日志基础字段
	var baseFields []zap.Field

	accessLogFile := config.Config().GetString("logger.access.logFile")
	accessLogger := logger.NewLoggerOfZap(accessLogFile, zapConfig, baseFields)

	for log := range logChannel {
		// 日志字段
		fields := []zap.Field{
			zap.String(logFields.RequestId, log.RequestId),
			zap.String(logFields.RequestReferer, log.RequestReferer),
			zap.String(logFields.RequestRefererIp, log.RequestRefererIp),
			zap.Int64(logFields.RequestTime, log.RequestTime),
			zap.Int64(logFields.ResponseTime, log.ResponseTime),
			zap.Int64(logFields.CostTime, log.CostTime),
			zap.String(logFields.RequestTarget, log.RequestTarget),
			zap.String(logFields.RequestBody, log.RequestBody),
			zap.String(logFields.ResponseBody, log.ResponseBody),
			zap.Int(logFields.ReturnCode, log.ReturnCode),
			zap.String(logFields.ReturnMsg, log.ReturnMsg),
			zap.String(logFields.ReturnData, log.ReturnData),
		}

		accessLogger.Info("", fields...)
	}
}
