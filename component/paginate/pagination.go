// paginate 分页组件

package paginate

import (
	"github.com/hongker/framework/util/number"
)

// Pagination 分页
type Pagination struct {
	// 当前页数
	PageNo int `json:"page_no"`

	// 每页行数
	Limit int `json:"limit"`

	// 总页数
	PageTotal int `json:"page_total"`
}

const (
	// 默认的每页行数
	defaultLimit       = 10
	// 默认的当前页数
	defaultCurrentPage = 1
)

// Paginate 通过总条数和当前页数，每页行数计算得到分页器
func Paginate(totalCount, currentPage, limit int) Pagination {
	pagination := Pagination{
		PageNo: currentPage,
		Limit: limit,
	}

	// 设置默认条数,后期改为可配置
	if pagination.Limit <= 0 {
		pagination.Limit = defaultLimit
	}

	// 设置默认当前页数
	if pagination.PageNo <= 0 {
		pagination.PageNo = defaultCurrentPage
	}

	// 计算总页数
	pagination.PageTotal = number.Div(totalCount, pagination.Limit)

	return pagination
}

// GetOffset 获取偏移量,用于数据库分页查询
func (p *Pagination) GetOffset() int {
	return (p.PageNo - 1) * p.Limit
}