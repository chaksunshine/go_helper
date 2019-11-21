package dao

import "time"

// 添加序列化注册对象
// @author fuzeyu
type Register struct {
	Name   string        // 序列化名称
	Object interface{}   // 需要序列化的对象
	Second time.Duration // 序列化间隔时间
}
