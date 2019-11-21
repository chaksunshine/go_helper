package internel

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/chaksunshine/go_helper/file"
	"github.com/chaksunshine/go_helper/store/dao"
	"github.com/chaksunshine/go_helper/str"
	"os"
	"path/filepath"
)

// 以文件的方式存储对象
// 将在当前程序的运行目录[	filepath.Abs(os.Args[0]) ]下建立store文件夹下进行存储
// @author fuzeyu
type StreamFile struct {
	dataFolder string // 文件夹路径
	dataName   string // 文件名称
}

// 创建文件夹
// @param name 对象名称
func (this *StreamFile) createFolder(name string) error {

	// 设置文件夹名称并创建文件
	this.dataFolder = fmt.Sprintf("%v/%v/%v/", filepath.Dir(os.Args[0]), groupSign, str.Md5(name))
	if file.FolderExists(this.dataFolder, true) == false {
		return errors.New(fmt.Sprintf("create folder fail %v", this.dataFolder))
	}

	this.dataName = name
	return nil
}

// 获取文件路径
func (this *StreamFile) filePath() string {
	return fmt.Sprintf("%v/%v.current.dat", this.dataFolder, this.dataName)
}

// 初始化
// @param register 注册对象
func (this *StreamFile) Init(register dao.Register) {

	// 初始化目录
	if err := this.createFolder(register.Name); err != nil {
		panic(err)
	}
}

// 写入文件
// @param register 注册对象
func (this *StreamFile) Writer(register dao.Register) error {

	// 创建打开文件
	open, err := os.OpenFile(this.filePath(), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer open.Close()

	// 序列化对象
	encoder := gob.NewEncoder(open)
	if err := encoder.Encode(register.Object); err != nil {
		return err
	}
	return nil
}

// 获取文件
// @param name 标志名
// @param obj 要反射的对象
func (this *StreamFile) Reader(name string, object interface{}) error {

	_ = this.createFolder(name)

	//// 判断文件是否存在
	filePath := this.filePath()
	//_, err := os.Stat(filePath)
	//if err != nil {
	//	if  {
	//
	//	}
	//}

	// 打开文件
	open, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer open.Close()

	// 获取反序列化内容
	_ = gob.NewDecoder(open).Decode(object)
	return nil
}
