package store

import (
	"encoding/gob"
	. "github.com/chaksunshine/go_helper/store/dao"
	"os"
	"time"
)

// 将对象序列化
// @author fuzeyu
type objectSerialization struct {
}

// 将一个对象序列化到文件中
// @param filePath 路径地址
// @param object 序列化对象
func (this *objectSerialization) encode(filePath string, object interface{}) (bool, error) {

	// 创建打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// 序列化对象
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(object); err != nil {
		return false, err
	}
	return true, nil
}

// 将对象信息写入到文件中
// @param r 注册对象
func (this *objectSerialization) writerFile(r Register) {

	ticker := time.NewTicker(r.Second)
	for {
		<-ticker.C
		this.encode(r.FilePath, r.Object)
	}

}

// 添加序列化对象
// @param r 注册对象
func (this *objectSerialization) Register(r Register) {
	go this.writerFile(r)
}

// 从文件中反序列化对象
// @param filePath 文件地址
// @param object 序列化地址
func (this *objectSerialization) Find(filePath string, object interface{}) interface{} {

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return object
	}
	defer file.Close()

	// 获取反序列化内容
	gob.NewDecoder(file).Decode(object)
	return object
}
