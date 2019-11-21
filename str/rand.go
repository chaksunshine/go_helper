package str

import (
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"
)

// 获取一个指定范围的随机数
func GetRand(min int, max int) int {
	maxItem, minItem := int64(max), int64(min)
	for true {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)

		next := rand.Int63n(maxItem)
		if next >= minItem && next <= maxItem {
			return int(next)
		}
	}
	return 0
}

// 获取一个随机位数的字符串
func GetShuffle(length int) string {
	var ret = ""
	loop := length / 32
	for i := 0; i <= loop; i++ {
		stringMd5 := Md5(fmt.Sprintf("%v-%v", GetRand(length*10, length*100), time.Now().UnixNano()))
		ret += base64.StdEncoding.EncodeToString([]byte(stringMd5))
	}
	ret = regexpAZHaving.ReplaceAllString(ret, "")
	return ret[:length]
}

// 获取一个不会重复的唯一随机数
func GetUniqueStr() string {
	return fmt.Sprintf("%v%v", GetShuffle(32), time.Now().UnixNano())
}

// 将一个字符串转换成随机数字
// @param val 字符串
func StringToNumber(val string) int {
	v := int(crc32.ChecksumIEEE([]byte(val)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
