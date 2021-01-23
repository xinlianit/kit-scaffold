package remote

import (
	"github.com/xinlianit/kit-scaffold/examples/http/remote/business"
)

func Init() {
	// 商家服务连接
	business.Connect()
}

func Close() {
	// 关闭商家服务连接
	business.Close()
}
