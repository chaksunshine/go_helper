package date

// 时间范围结构体
type DateRange struct {
	Start int64 // 开始日期时间戳
	End   int64 // 结束日期时间戳
}

// 获取一个时间范围对象
func InitDateRange(start int64, end int64) DateRange {
	return DateRange{
		End:   end,
		Start: start,
	}
}
