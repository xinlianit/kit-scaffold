package model

type Request struct {
	RequestId      string      `json:"request_id"`
	RequestReferer string      `json:"request_referer"`
	RequestTime    string      `json:"request_time"`
	Data           interface{} `json:"data"`
}

func (r Request) GetRequestId() string {
	return r.RequestId
}

func (r Request) GetRequestReferer() string {
	return r.RequestReferer
}

func (r Request) GetRequestTime() string {
	return r.RequestTime
}

func (r Request) GetData() interface{} {
	return r.Data
}
