package util

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var (
	pathUtilInstance *Path
	pathUtilOnce     sync.Once
)

// PathUtil 逻辑工具
func PathUtil() *Path {
	pathUtilOnce.Do(func() {
		pathUtilInstance = new(Path)
	})

	return pathUtilInstance
}

type Path struct {
}

// GetCurrentPath 获取当前绝对路径
func (u Path) GetCurrentPath() string {
	// 获取当前执行路径
	execPath := u.GetCurrentPathByExecutable()
	// 临时路径
	tempDir, _ := filepath.EvalSymlinks(os.TempDir())

	if strings.Contains(execPath, tempDir) {
		return u.GetCurrentByCaller()
	}

	return execPath
}

// GetCurrentPathByExecutable 获取当前绝对路径 - 通过执行路径
func (u Path) GetCurrentPathByExecutable() string {
	// 获取执行路径
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		return ""
	}

	// 获取真实路径（软链接会返回真实路径地址）
	realPath, _ := filepath.EvalSymlinks(execPath)
	if err != nil {
		log.Fatal(err)
	}

	return realPath
}

// GetCurrentByCaller 获取当前绝路径 - 通过运行路径
func (u Path) GetCurrentByCaller() string {
	var runPath string
	_, file, _, ok := runtime.Caller(2)
	if ok {
		runPath = path.Dir(file)
	}

	return runPath
}
