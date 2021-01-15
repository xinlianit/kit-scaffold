package model

// 响应
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r Response) GetCode() int {
	return r.Code
}

func (r Response) GetMsg() string {
	return r.Msg
}

func (r Response) GetData() interface{} {
	return r.Data
}
