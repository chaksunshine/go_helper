package str

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// 首字符大写转换成下滑线
func UpperToDownLine(val string, format byte) string {

	ret := make([]byte, 0)
	for key, value := range val {
		byteValue := byte(value)
		if byteValue >= 65 && byteValue <= 90 {
			if key != 0 {
				ret = append(ret, format)
			}
			ret = append(ret, byteValue+32)
		} else {
			ret = append(ret, byteValue)
		}
	}
	return string(ret)
}

// 将一个字符串md5加密
// @param str 要加密的字符串
func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}

// 将一个字符串basa64解码
// @param str 要加密的字符串
func Base64Decode(str string) string {
	bytes, _ := base64.StdEncoding.DecodeString(str)
	return string(bytes)
}

// 将一个字符串basa64编码
// @param str 要加密的字符串
func Base64Encode(str string) string {
	bytes := base64.StdEncoding.EncodeToString([]byte(str))
	return bytes
}

// 将一个字符串url解码
// @param str 需要解码的字符串
func UrlDecode(str string) string {
	val, _ := url.QueryUnescape(str)
	return val
}

// 将一个字符串url编码
// @param str 需要解码的字符串
func UrlEncode(str string) string {
	val := url.QueryEscape(str)
	return val
}

// json字符串编码
// @param str 需要解码的字符串
func JsonEncode(str interface{}) string {
	bytes, _ := json.Marshal(str)
	return string(bytes)
}
