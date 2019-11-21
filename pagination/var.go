package pagination

import "regexp"

// 分页字段
const pageFields = "page"

// 每页分页长度
const constPageSize = 10

// 正则表达式
var (
	regexpRemovePage, _ = regexp.Compile(pageFields + "=\\d+")
)
