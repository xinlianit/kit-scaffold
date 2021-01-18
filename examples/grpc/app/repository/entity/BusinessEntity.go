package entity

// 商家实体
type BusinessEntity struct {
	entity
	BusinessId      int32  `json:"business_id"`      // 商家ID
	BusinessName    string `json:"business_name"`    // 商家名称
	Contacts        string `json:"contacts"`         // 联系人
	Mobile          string `json:"mobile"`           // 手机号
	Address         string `json:"address"`          // 商家地址
	BusinessLicense string `json:"business_license"` // 营业执照
}

func (e BusinessEntity) GetBusinessId() int32 {
	return e.BusinessId
}

func (e BusinessEntity) GetBusinessName() string {
	return e.BusinessName
}

func (e BusinessEntity) GetContacts() string {
	return e.Contacts
}

func (e BusinessEntity) GetMobile() string {
	return e.Mobile
}

func (e BusinessEntity) GetAddress() string {
	return e.Address
}

func (e BusinessEntity) GetBusinessLicense() string {
	return e.BusinessLicense
}
