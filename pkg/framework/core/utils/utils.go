package utils

import (
	"os"

	"github.com/jinzhu/copier"
)

// PathExists 路径文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// Copy struct to strcut
func Copy(fromValue any, toValue any) error {
	return copier.Copy(toValue, fromValue)
}
