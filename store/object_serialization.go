package store

import (
	"fmt"
	"github.com/chaksunshine/go_helper/store/dao"
	"github.com/chaksunshine/go_helper/store/enum"
	"github.com/chaksunshine/go_helper/store/internel"
	"time"
)

// 对象序列化
// @author fuzeyu
type objectSerialization struct {
}

// 获取一个存储接口
// @param rType 存储方式
func (this *objectSerialization) findType(rType int) internel.StreamStorage {
	switch rType {
	case enum.RegisterTypeFile:
		return new(internel.StreamFile)
	}

	panic(fmt.Sprintf("undefined rType in %v", rType))
}

// 注册一个需要存储的对象
// @param rType 存储方式
// @param register 注册对象
func (this *objectSerialization) Register(rType int, register dao.Register) {

	go func() {

		// 创建流并初始化
		stream := this.findType(rType)
		stream.Init(register)

		// 写入数据
		ticker := time.NewTicker(register.Second)
		for range ticker.C {
			err := stream.Writer(register)
			if err != nil {
				panic(err)
			}
		}
	}()
}

// 从指定流信息中加载一个对象
// @param rType 存储方式
// @param name 存储名
// @param obj 存储对象
func (this *objectSerialization) Find(rType int, name string, obj interface{}) interface{} {

	stream := this.findType(rType)
	if err := stream.Reader(name, obj); err != nil {
		panic(err)
	}
	return obj
}
