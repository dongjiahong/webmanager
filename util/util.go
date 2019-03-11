package util

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var dataPath string

type Conf struct {
	MediaPath string `json:"media_path"`
}

func Init(conf *Conf) {
	dataPath = conf.MediaPath
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
	path := dataPath + fmt.Sprintf("%s/", kind)
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

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 下载文件到本地，并返回本地路径和错误信息
func DownloadWithUrl(url, filePath string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 客户端不校验省的x509错误
	}
	client := &http.Client{Transport: tr}

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("[DownloadWithUrl] new request err: %v, url: %s", err, url)
	}

	//处理返回结果
	resp, err := client.Do(reqest)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("[DownloadWithUrl] get err: %v, url: %s", err, url)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[DownloadWithUrl] read body err: %v", err)
	}

	if err := ioutil.WriteFile(filePath, body, 0666); err != nil {
		return fmt.Errorf("[DownloadWithUrl] write file err: %v", err)
	}
	return nil
}
