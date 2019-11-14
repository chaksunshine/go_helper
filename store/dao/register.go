package dao

import "time"

// 添加序列化注册对象
// @author fuzeyu
type Register struct {
	Object   interface{}   // 需要序列化的对象
	FilePath string        // 序列化地址
	Second   time.Duration // 序列化间隔时间
}
