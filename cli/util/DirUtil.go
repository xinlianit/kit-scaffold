package util

import (
	"io/ioutil"
	"os"
	"sync"
)

var (
	dirUtilInstance *Dir
	dirUtilOnce     sync.Once
)

// DirUtil 文件工具
func DirUtil() *Dir {
	dirUtilOnce.Do(func() {
		dirUtilInstance = new(Dir)
	})

	return dirUtilInstance
}

type Dir struct {
}

// IsDir 判断目录是否存在
// @param dirPath 	目录路径
// @return bool 	是否目录; true-是、false-否
func (u Dir) IsDir(dirPath string) bool {
	isDir, _ := u.IsDirWithError(dirPath)
	return isDir
}

// IsDirWithError 判断目录是否存在
// @param dirPath 	目录路径
// @return isDir 	是否目录; true-是、false-否
// @return err
func (u Dir) IsDirWithError(dirPath string) (isDir bool, err error) {
	DirInfo, err := os.Stat(dirPath)

	if err != nil {
		return false, err
	}

	return DirInfo.IsDir(), nil
}

// IsEmptyDir 判断是否为空目录
// @param dirPath 	目录路径
// @return bool 	是否空目录; true-是、false-否
func (u Dir) IsEmptyDir(dirPath string) bool {
	isEmpty, _ := u.IsEmptyDirWithError(dirPath)
	return isEmpty
}

// IsEmptyDirWithError 判断是否为空目录
// @param dirPath 	目录路径
// @return isEmpty 	是否空目录; true-是、false-否
// @return err
func (u Dir) IsEmptyDirWithError(dirPath string) (isEmpty bool, err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return false, err
	}

	if len(dir) == 0 {
		return true, nil
	}

	return false, nil
}

// CreateDir 创建目录
// @param dirPath 	目录路径
// @param recursive 是否递归创建目录; true-是、false-否
// @return err
func (u Dir) CreateDir(dirPath string, recursive bool) error {
	if !u.IsDir(dirPath) {
		if recursive {
			return os.MkdirAll(dirPath, os.ModePerm)
		} else {
			return os.Mkdir(dirPath, os.ModePerm)
		}
	}
	return nil
}
