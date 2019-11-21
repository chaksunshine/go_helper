package file

import (
	"errors"
	"fmt"
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
	if FolderExists(filepath.Dir(path), true) == false {
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
	if FolderExists(filepath.Dir(path), true) == false {
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
