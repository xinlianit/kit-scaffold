package entity

// 商家扩展信息
type BusinessExtendEntity struct {
	entity
	BusinessDesc string `json:"business_desc"` // 商家简介
}

func (e BusinessExtendEntity) GetBusinessDesc() string {
	return e.BusinessDesc
}
