package pool

import "time"

// 协程池
// @author fuzeyu
type GroutingPool struct {
	total   int
	mapping chan byte
	limit   time.Duration
}

// 初始化
// @param total 协程池总长度
// @param limit 下次分发协程的间隔时间
func (this *GroutingPool) Init(total int, limit time.Duration) *GroutingPool {
	this.total = total
	this.limit = limit
	this.mapping = make(chan byte, total)
	this.startProduce()
	return this
}

// 归还协程
func (this *GroutingPool) remandGrouting() {
	this.mapping <- 1
}

// 开始循环循环产生协程池内容
func (this *GroutingPool) startProduce() {
	for crt := 0; crt < this.total; crt++ {
		this.remandGrouting()
	}
}

// 获取下一次协程池
func (this *GroutingPool) Next() {
	<-this.mapping
}

// 归还协程
func (this *GroutingPool) Remand() {
	go func() {
		time.Sleep(this.limit)
		this.remandGrouting()
	}()
}
