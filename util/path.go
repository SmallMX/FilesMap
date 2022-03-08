package util

import (
	"os"
	"path/filepath"
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

func AllSubPath(path string) []string  {
	if !IsPathExist(path) {
		return nil
	}

	if IsFile(path) {
		return nil
	}

	pattern := path + "/*"

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil
	}
	return  files
}