package scaffold

import (
	"fmt"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func init() {
	// 配置初始化
	config.Init()

	// 日志初始化
	var baseFields []zap.Field
	logger.ZapLogger = logger.ZapInit(logger.NewDefaultZapConfig(), baseFields)
}

// 运行 Http 服务
func RunHttpServer(address string, handler http.Handler) {
	httpServer := http.Server{
		Addr:    address,
		Handler: handler,
	}

	log.Println(fmt.Sprintf("test server address: %v", address))

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

// 创建 Http 处理器
func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}
