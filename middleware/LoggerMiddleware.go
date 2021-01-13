package middleware

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/xinlianit/go-util/util"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
	"log"
)

// 日志通道
var logChannel = make(chan logBody, 100)

// 日志实体
type logBody struct {
	RequestId          string      `json:"request_id"`           // 请求ID
	RequestTime        int64       `json:"request_time"`         // 请求时间
	ResponseTime       int64       `json:"response_time"`        // 响应时间
	CostTime           int64       `json:"cost_time"`            // 耗时
	ResponseCode       int         `json:"response_code"`        // 响应状态码
	RequestProto       string      `json:"request_proto"`        // 请求协议
	RequestMethod      string      `json:"request_method"`       // 请求方法
	RequestReferer     string      `json:"request_referer"`      // 请求来源
	RequestUri         string      `json:"request_uri"`          // 请求URI
	RequestContextType string      `json:"request_context_type"` // 请求体类型
	RequestBody        string      `json:"request_body"`         // 请求体
	ClientIp           string      `json:"client_ip"`            // 客户端IP
	UserAgent          string      `json:"user_agent"`           // User Agent
	ReturnCode         int         `json:"return_code"`          // 返回状态码
	ReturnMsg          string      `json:"return_msg"`           // 返回信息
	ReturnData         interface{} `json:"return_data"`          // 返回数据
}

// 日志字段
var logFields = struct {
	RequestId          string
	RequestTime        string
	ResponseTime       string
	CostTime           string
	ResponseCode       string
	RequestProto       string
	RequestMethod      string
	RequestReferer     string
	RequestUri         string
	RequestContextType string
	RequestBody        string
	ClientIp           string
	UserAgent          string
	ReturnCode         string
	ReturnMsg          string
	ReturnData         string
}{
	RequestId:          "request_id",
	RequestTime:        "request_time",
	CostTime:           "cost_time",
	ResponseCode:       "response_code",
	RequestProto:       "request_proto",
	RequestMethod:      "request_method",
	RequestReferer:     "request_referer",
	RequestUri:         "request_uri",
	RequestContextType: "request_context_type",
	RequestBody:        "request_body",
	ClientIp:           "client_ip",
	UserAgent:          "user_agent",
	ReturnCode:         "return_code",
	ReturnMsg:          "return_msg",
	ReturnData:         "return_data",
}

// 日志中间件
func LoggerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	// 接收并记录日志
	go receiveAndRecordLog()

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println(config.Config().Get("test"))
		log.Println("============LoggerMiddleware")
		// todo 记录日志
		log := logBody{
			RequestId:   getRequestId(ctx),
			RequestTime: util.TimeUtil().GetCurrentMilliTime(),
		}

		// 执行请求
		response, err = next(ctx, request)

		// 响应数据
		log.ResponseTime = util.TimeUtil().GetCurrentMilliTime()
		// 耗时(RT)
		log.CostTime = log.ResponseTime - log.RequestTime

		// 发送日志通道
		logChannel <- log
		return
	}
}

// 获取请求ID
func getRequestId(ctx context.Context) string {
	requestId, ok := ctx.Value("X-Request-Id").(string)
	if !ok || requestId == "" {
		id, _ := uuid.NewRandom()
		requestId = id.String()
	}
	context.WithValue(ctx, "X-Request-Id", requestId)

	return requestId
}

// 接收并记录日志
func receiveAndRecordLog() {
	// 配置
	zapConfig := logger.NewDefaultZapConfig()
	// 禁用文件及行号
	zapConfig.RecordLineNumber = false

	// 日志基础字段
	var baseFields []zap.Field

	//accessLogFile := config.Config().GetString("app.log.accessLogFile")
	accessLogFile := "./logs/access.log"
	accessLogger := logger.NewLoggerOfZap(accessLogFile, zapConfig, baseFields)

	for log := range logChannel {
		fmt.Println(log.RequestId)

		// 日志字段
		fields := []zap.Field{
			zap.String(logFields.RequestId, log.RequestId),
			zap.Int64(logFields.RequestTime, log.RequestTime),
			zap.Int64(logFields.ResponseTime, log.ResponseTime),
			zap.Int64(logFields.CostTime, log.CostTime),
		}

		accessLogger.Info("", fields...)
	}
}
