package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

var dataPath string

func init() {
	dataPath = "/Users/lele/tmp/go/webmanager/"
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // 存在
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetCommonPath(kind string) string {
	path := dataPath + fmt.Sprintf("%s_data/", kind)
	if !IsExist(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Println("[GetCommonPath] err: ", err)
		}
	}
	return path
}

func WriteFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}

	return err
}
