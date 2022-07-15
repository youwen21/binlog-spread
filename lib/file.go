package lib

import (
	"errors"
	"os"
	"path/filepath"
)

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetRootDir(dir string) (string, error) {
	if dir == "" {
		return "", errors.New("未能定位根目录")
	}

	if CheckFileIsExist(dir + "/.env") {
		return dir, nil
	}

	return GetRootDir(filepath.Dir(dir))
}
