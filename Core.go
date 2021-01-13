package scaffold

import (
	"context"
	"fmt"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	httpServer := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  time.Duration(config.Config().GetInt("server.readTimeout")) * time.Millisecond,
		WriteTimeout: time.Duration(config.Config().GetInt("server.writeTimeout")) * time.Millisecond,
	}

	// 服务启动成功
	fmt.Println("Listening and serving HTTP on " + address + ", PID: " + fmt.Sprintf("%d", os.Getpid()))

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("server run error: " + err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt) // 指定中断信号(Interrupt)转发到 signalChan
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	// 3 秒超时自动取消(当执行一个go 协程时，超时自动取消协程)
	contextTimeout := config.Config().GetInt("server.contextTimeout")
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(contextTimeout)*time.Millisecond)
	defer cancelFunc()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// 创建 Http 处理器
func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}
