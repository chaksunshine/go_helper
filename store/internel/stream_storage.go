package internel

import "github.com/chaksunshine/go_helper/store/dao"

// 流式存储接口
// @author fuzeyu
type StreamStorage interface {
	Init(register dao.Register)
	Writer(register dao.Register) error
	Reader(string, interface{}) error
}
