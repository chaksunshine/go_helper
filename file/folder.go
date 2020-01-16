package file

import (
	"io/ioutil"
	"os"
)

// 获取一个路径下的全部文件
// @param path 读取路径
// @param isRecursive 是否递归读取
func ReadFolder(path string, isRecursive bool) []string {

	var ret = make([]string, 0)

	// 读取全部目录
	catalog, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, value := range catalog {

		// 保存当前目录
		crtPath := path + string(os.PathSeparator) + value.Name()
		ret = append(ret, crtPath)

		// 递归获取
		if isRecursive && value.IsDir() {
			ret = append(ret, ReadFolder(crtPath, isRecursive)...)
		}
	}

	return ret
}

// 检查一个文件夹是否存在
// 若不存在运行自行创建
// @param path 文件夹路径
// @param isCreate 若文件夹不存在，是否自动创建
func FolderExists(path string, isCreate bool) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) != false {
			return false, err
		}

		// 不允许创建
		// 直接返回false
		if !isCreate {
			return false, nil
		}

		// 创建文件
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
