package dto

type PageForm struct {
	OrderBy     string `form:"orderBy" json:"orderBy" `        // 字段
	OrderDirect string `form:"orderDirect" json:"orderDirect"` // 方式
	Page        int    `form:"page,default=1" json:"page"`
	Psize       int    `form:"psize,default=20" json:"psize"`
}

//PageParam 分页控制
type PageParam struct {
	PageNumber int `json:"pageNumber,default=1" form:"pageNumber,default=1"`
	PageSize   int `json:"pageSize,default=10" form:"pageSize,default=10"`
	OrderBy    []string
}

//Limit 获取每页记录数
func (p PageParam) Limit() int {
	return p.PageSize
}

//Offset 获取当前页开始ID
func (p PageParam) Offset() int {
	if p.PageNumber == 0 {
		return 0
	}

	return (p.PageNumber - 1) * p.PageSize
}
