package main

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/gorilla/mux"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/boot"
	"github.com/xinlianit/kit-scaffold/examples/http/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/http/middleware"
	"github.com/xinlianit/kit-scaffold/examples/http/middleware/metrics"
	"github.com/xinlianit/kit-scaffold/examples/http/transport"
	"github.com/xinlianit/kit-scaffold/handler"
	"net/http"
)

func main() {
	// 初始化命令行
	//commandLineInit()

	httpHandler := NewHttpHandler()

	// 运行服务
	scaffold.RunHTTPServer(httpHandler)
}

// 命令行初始化
func commandLineInit() {
	// 命令行参数
	pflag.String("test.app.id", "","测试APP_ID")

	// 命令行解析
	scaffold.CommandLineParse()
}

func NewHttpHandler() http.Handler {
	labelNames := []string{"method"}
	// 请求计数
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "kit_scaffold",
		Subsystem: "IndexEndpoint",
		Name: "request_count",
		Help: "Number of requests received.",
	}, labelNames)

	// 请求耗时
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "kit_scaffold",
		Subsystem: "IndexEndpoint",
		Name: "request_latency",
		Help: "Total duration of requests in microseconds.",
	}, labelNames)

	httpHandler := handler.NewHttpHandler()
	httpHandler.Use(middleware.TestMiddleware, middleware.Test2Middleware)

	indexEndpoint := metrics.IndexMetricsMiddleware(requestCount, requestLatency)(endpoint.NewIndexEndpoint())
	//indexEndpoint := endpoint.NewIndexEndpoint()
	indexTransport := transport.NewIndexTransport()
	helloHandler := httpHandler.Server(indexEndpoint.Hello, indexTransport.HelloDecode)
	testHandler := httpHandler.Server(indexEndpoint.Test, indexTransport.HelloDecode)

	route := mux.NewRouter()

	// 注册基本服务
	boot.RegisterHTTPBaseServer(route, httpHandler)

	route.Methods(http.MethodGet).Path("/metrics").Handler(promhttp.Handler())
	route.Methods(http.MethodGet).Path("/index/hello").Handler(helloHandler)
	route.Methods(http.MethodGet).Path("/index/test").Handler(testHandler)

	return route
}
