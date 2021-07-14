package boot

import (
	"github.com/xinlianit/kit-scaffold/app"
	"github.com/xinlianit/kit-scaffold/boot/nacos"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 初始化
func Init() {
	// 获取当前执行路径
	app.RootPath = getCurrentPathByExecutable()
	// 临时路径
	tempDir, _ := filepath.EvalSymlinks(os.TempDir())
	// 判断当前执行路径是否临时路径
	if strings.Contains(app.RootPath, tempDir) {
		// 临时路径获取调用路径
		app.RootPath = getCurrentPathByCaller()
	}

	app.ConfigPath = app.RootPath + "/config"
	app.LogPath = app.RootPath + "/logs"
	app.RuntimePath = app.RootPath + "/runtime"
	app.ResourcePath = app.RootPath + "/resource"
	app.CachePath = app.RootPath + "/cache"

	// 配置初始化
	config.Init()

	// 日志初始化
	logger.Init()
	var baseFields []zap.Field
	logger.ZapLogger = logger.ZapInit(logger.NewDefaultZapConfig(), baseFields)

	// 初始化故障转移
	//breaker.Init()
}

// 配置中心初始化
func ConfigCenterInit() {
	// 配置中心
	configCenterEnable := config.Config().GetBool("app.configCenter.enable")
	if configCenterEnable {
		// 配置中心类型
		configCenterType := config.Config().GetString("app.configCenter.type")

		switch configCenterType {
		case constant.ConfigCenterTypeNacos:
			nacosConfig()
		default:
			nacosConfig()
		}
	}
}

// nacos 配置中心
func nacosConfig() {
	// 配置中心初始化
	nacos.Init()
	// 监听并同步配置
	nacos.ListenSyncConfig()
	// 初始化动态配置
	config.InitDynamicConfig()
}

// getCurrentPathByExecutable 获取当前绝对路径 - 通过执行路径
// @return string 当前文件绝对路径
func getCurrentPathByExecutable() string {
	// 获取执行路径
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	// 获取真实路径（软链接会返回真实路径地址）
	realPath, _ := filepath.EvalSymlinks(execPath)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return path.Dir(realPath)
}

// getCurrentPathByCaller 获取当前绝路径 - 通过运行路径
// @return string 当前文件绝对路径
func getCurrentPathByCaller() string {
	var runPath string
	_, file, _, ok := runtime.Caller(4)
	if ok {
		runPath = path.Dir(file)
	}

	return runPath
}
