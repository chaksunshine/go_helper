package file

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 向一个路径中写入文件
// @param path 文件路径
// @param content 文件内容
func WriterFile(path string, content []byte) error {

	// 检查文件夹是否存在
	// 若不存在先创建文件夹
	exists, err := FolderExists(filepath.Dir(path), true)
	if err != nil {
		return err
	}
	if exists == false {
		return errors.New(fmt.Sprintf("create file fail is %v", path))
	}

	return ioutil.WriteFile(path, content, os.ModePerm)
}

// 向一个路径中追加写入文件
// @param path 文件路径
// @param content 文件内容
func WriterFileByAppend(path string, content []byte) error {

	// 检查文件夹是否存在
	// 若不存在先创建文件夹
	exists, err := FolderExists(filepath.Dir(path), true)
	if err != nil {
		return err
	}
	if exists == false {
		return errors.New(fmt.Sprintf("create file fail is %v", path))
	}

	// 以追加的方式打开文件夹
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {

		// 判断是否文件不存在
		if os.IsNotExist(err) == false {
			return err
		}
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// 检查一个文件是否存在
// @param filePath 文件路径
func FileExists(filePath string) (bool, error) {

	info, err := os.Stat(filePath)
	if err == nil && info != nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取一个文件的md5值
// @param path 文件路径
func FileMd5(path string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}
	bytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(bytes), nil
}
