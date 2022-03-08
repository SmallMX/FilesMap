package util

import (
	"path"
	"strings"
)

// FileFullName 获取文件全名
func FileFullName(filePath string) string {
	return path.Base(filePath)
}

// FileName 获取文件名
func FileName(filePath string) (prefixName string, suffixName string) {
	fileName := FileFullName(filePath)
	suffixName = path.Ext(filePath)
	prefixName = strings.TrimSuffix(fileName, suffixName)
	return
}