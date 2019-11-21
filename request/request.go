package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 格式化返回结果
// @param resp 响应结果
// @param err 错误信息
func formatResponse(resp *http.Response, err error) string {

	// 判断结果
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	// 获取全部内容
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// 发送POST请求
// @param url 请求参数
// @param contentType post请求方式
// @param params 请求参数
func SendSimplePost(url string, contentType string, params url.Values) string {
	rsp, err := http.Post(url, contentType, strings.NewReader(params.Encode()))
	return formatResponse(rsp, err)
}

// 发送GET请求
// @param url 请求参数
func SendSimpleGet(url string) string {
	resp, err := http.Get(url)
	return formatResponse(resp, err)
}
