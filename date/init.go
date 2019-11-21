package date

import "time"

func init() {

	// 初始化加载时间包
	r, e := time.LoadLocation("Asia/Shanghai")
	if e != nil {
		panic(e)
	}
	zhLocalPack = r
}
