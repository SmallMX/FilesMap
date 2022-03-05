package util

import (
	"os"
)

// IsPathExist 返回给定路径是否存在
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir 路径是否为文件夹
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// IsFile 路径是否为文件
func IsFile(path string) bool {
	return  !IsDir(path)
}

//func AllSubPath(path string) []string  {
//	var paths = make([]string, 0)
//	if IsFile(path) {
//		return nil
//	}
//
//	pattern := strings.
//
//	files, err := filepath.Glob("*")
//
//}