package entity

type entity struct {
	CreatedAt  string `json:"created_at"`  // 创建时间
	UpdatedAt  string `json:"updated_at"`  // 更新时间
	Sort       uint32 `json:"sort"`        // 排序
	DataStatus uint8  `json:"data_status"` // 数据状态
}

func (e entity) GetCreatedAt() string {
	return e.CreatedAt
}

func (e entity) GetUpdatedAt() string {
	return e.UpdatedAt
}

func (e entity) GetSort() uint32 {
	return e.Sort
}

func (e entity) GetDataStatus() uint8 {
	return e.DataStatus
}
