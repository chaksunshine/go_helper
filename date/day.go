package date

import (
	"time"
)

// 获取一个日期的第一秒时间
// @param timestamp 时间戳
func FirstDaySecond(timestamp int64) time.Time {
	un := time.Unix(timestamp, 0).Format("2006-01-02") + " 00:00:00"
	return ParseDateString("2006-01-02 00:00:00", un)
}

// 将一个时间文本转换成时间对象
// @param layout 时间
// @param timestamp 日期字符串
func ParseDateString(layout string, timestamp string) time.Time {
	location, e := time.ParseInLocation(layout, timestamp, zhLocalPack)
	if e != nil {
		panic(e)
	}
	return location
}
