package boot

import (
	"github.com/xinlianit/kit-scaffold/cli/common"
	"github.com/xinlianit/kit-scaffold/cli/util"
	"path/filepath"
)

func Init()  {
	// 当前路径
	currentPath := util.PathUtil().GetCurrentPath()
	// 根路径
	common.RootPath = filepath.Dir(currentPath)
	// 配置路径
	common.ConfigPath = common.RootPath + "/config"
}