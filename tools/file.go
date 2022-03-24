package tools

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/zhangshanwen/shard/initialize/conf"
)

func SaveFile(fileName, fileBody, filePath string) (err error) {
	if filePath == "" {
		filePath = conf.C.File.Path
	}
	s, err := os.Stat(filePath)
	if err != nil {
		return os.MkdirAll(filePath, os.ModePerm)
	}
	if !s.IsDir() {
		return errors.New("path is a file")
	}

	var decodeData []byte
	if decodeData, err = base64.StdEncoding.DecodeString(fileBody); err != nil {
		return
	}

	fileName = path.Join(filePath, fileName)
	var file *os.File
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write(decodeData)
	return
}

func FileToBase64(fileName, filePath string) (s string, err error) {
	if filePath == "" {
		filePath = conf.C.File.Path
	}
	fileName = path.Join(filePath, fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
