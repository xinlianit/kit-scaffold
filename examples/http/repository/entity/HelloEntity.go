package entity

type HelloEntity struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (e HelloEntity) GetId() int64 {
	return e.Id
}

func (e HelloEntity) GetName() string {
	return e.Name
}

func (e HelloEntity) GetDesc() string {
	return e.Desc
}
