package response

type HelloResponse struct {
	Code int32	`json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}
