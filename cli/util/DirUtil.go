package util

import (
	"github.com/xinlianit/go-util"
	"io/ioutil"
	"os"
	"path"
	"strings"
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

// LsDir 列出目录文件
// @param dirPath 		目录路径
// @param readChildDir 	读取子目录
// @return fileList 	文件列表
// @return err
func (u Dir) LsDir(dirPath string, readChildDir bool) (fileList []string, err error) {

	rd, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return fileList, err
	}

	for _, fi := range rd {
		if fi.IsDir() {
			// 递归读取子目录
			if readChildDir {
				childFileList, err := u.LsDir(strings.TrimRight(dirPath, "/")+"/"+fi.Name(), readChildDir)
				if err != nil {
					return fileList, err
				}

				for _, childFile := range childFileList {
					fileList = append(fileList, fi.Name()+"/"+childFile)
				}
			}
		} else {
			fileList = append(fileList, fi.Name())
		}
	}

	return fileList, nil
}

// CopyDir 递归拷贝目录
// @param src 源目录
// @param dst 目标目录
// @return error
func (u Dir) CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcInfo os.FileInfo

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcFp := path.Join(src, fd.Name())
		dstFp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = u.CopyDir(srcFp, dstFp); err != nil {
				return err
			}
		} else {
			// TODO: 需确认是否修改同包下的FileUtil
			if err = util.FileUtil().CopyFile(srcFp, dstFp); err != nil {
				return err
			}
		}
	}
	return nil
}

// CopyDirWithFix 递归拷贝目录
// @param src 			源目录
// @param dst 			目标目录
// @param filePrefix 	文件前缀
// @return error
func (u Dir) CopyDirWithFix(src string, dst string, filePrefix string) error {
	var err error
	var fds []os.FileInfo
	var srcInfo os.FileInfo

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}

	for _, fd := range fds {
		srcFp := path.Join(src, fd.Name())
		dstFp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = u.CopyDirWithFix(srcFp, dstFp, filePrefix); err != nil {
				return err
			}
		} else {
			// TODO: 需确认是否修改同包下的FileUtil
			if err = util.FileUtil().CopyFileWithFix(srcFp, dstFp, filePrefix); err != nil {
				return err
			}
		}
	}
	return nil
}