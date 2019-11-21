package date

import "time"

// 获取一个月份的第一秒时间
// @param timestamp 时间戳
func FirstMonthSecond(timestamp int64) time.Time {
	un := time.Unix(timestamp, 0).Format("2006-01") + "-01 00:00:00"
	return ParseDateString("2006-01-02 00:00:00", un)
}

// 获取一个月份的最后一秒时间
// @param timestamp 时间戳
func LastMonthSecond(timestamp int64) time.Time {
	unix := FirstMonthSecond(timestamp).AddDate(0, 1, 0).Unix()
	return time.Unix(unix-1, 0)
}
