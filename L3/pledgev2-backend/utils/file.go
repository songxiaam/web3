package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func MkDir(dir string) error {
	rootDir, err := filepath.Abs("./")
	if err != nil {
		return errors.New("获取根目录失败," + err.Error())
	}
	dir = rootDir + "/" + dir
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
