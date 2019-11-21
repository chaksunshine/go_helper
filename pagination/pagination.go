package pagination

import (
	"encoding/json"
	"math"
	"net/url"
	"strconv"
	"strings"
)

// redis分页工具
type Pagination struct {
	pageIndex int // 页码
	pageSize  int // 分页步长

	total     int // 数据总数
	totalPage int // 总页数
}

// 计算出当前查询条件的总页数
func (this *Pagination) queryTotal() {

	// 计算总页数
	f := float64(this.total) / float64(this.pageSize)
	this.totalPage = int(math.Ceil(f))
}

// 检查传递的页码是否合法
// 若超过最大页数则取最大页数
func (this *Pagination) formatPageIndex() {

	// 最小
	if this.pageIndex < 1 {
		this.pageIndex = 1
	}

	// 最大
	if this.pageIndex > this.totalPage {
		this.pageIndex = this.totalPage
	}
}

// 分析分页数据
func (this *Pagination) init() {
	this.queryTotal()
	this.formatPageIndex()
}

// 获取查询数据库分页限制记录
func (this *Pagination) Limit() int {
	return this.pageSize
}

// 获取查询数据库分页起始记录
func (this *Pagination) Offset() int {
	return (this.pageIndex - 1) * this.pageSize
}

// 获取当前页码
func (this *Pagination) PageIndex() int {
	return this.pageIndex
}

// 获取分页步长
func (this *Pagination) PageSize() int {
	return this.pageSize
}

// 获取总页数
func (this *Pagination) TotalPage() int {
	return this.totalPage
}

// 将分页信息转换成json字符串
// @param path 当前请求连接，不包含page字段
func (this *Pagination) Json(uri *url.URL) string {

	// 兼容地址栏中的问号
	crtUri := uri.String()
	if strings.Contains(crtUri, "?") == false {
		crtUri += "?"
	}

	bytes, _ := json.Marshal(map[string]string{
		"total":     strconv.Itoa(this.total),
		"totalPage": strconv.Itoa(this.totalPage),
		"pageIndex": strconv.Itoa(this.pageIndex),
		"pageSize":  strconv.Itoa(this.pageSize),
		"uri":       regexpRemovePage.ReplaceAllString(crtUri, ""),
	})
	return string(bytes)
}
