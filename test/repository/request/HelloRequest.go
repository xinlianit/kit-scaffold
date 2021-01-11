package request

type HelloRequest struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

func (r HelloRequest) GetId() int64 {
	return r.Id
}

func (r HelloRequest) GetMessage() string {
	return r.Message
}
