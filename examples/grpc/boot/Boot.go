package boot

import (
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/lib/driver"
)

// Init 应用初始化
func Init() {
	// 初始化动态配置
	config.InitDynamicConfig()
	// 初始化数据源
	driver.InitMySql()
}

// Destruct 应用销户
func Destruct()  {
	// 关闭清理数据库连接资源
	driver.SqlDB.Close()
}
