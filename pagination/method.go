package pagination

// 获取redis分页对象
// @param total 总页数
// @param pageIndex 当前页数
// @param pageSize 页码长度
func NewPagination(total int, pageIndex int, pageSize ...int) Pagination {

	// 判断页码长度
	var size = constPageSize
	if len(pageSize) >= 1 {
		size = pageSize[0]
	}

	p := &Pagination{
		pageIndex: pageIndex,
		pageSize:  size,
		total:     total,
	}
	p.init()
	return *p
}
