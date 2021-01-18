package entity

type entity struct {
	CreatedAt  string `json:"created_at"`  // 创建时间
	UpdatedAt  string `json:"updated_at"`  // 更新时间
	Sort       int    `json:"sort"`        // 排序
	DataStatus int    `json:"data_status"` // 数据状态
}

func (e entity) GetCreatedAt() string {
	return e.CreatedAt
}

func (e entity) GetUpdatedAt() string {
	return e.UpdatedAt
}

func (e entity) GetSort() int {
	return e.Sort
}

func (e entity) GetDataStatus() int {
	return e.DataStatus
}
