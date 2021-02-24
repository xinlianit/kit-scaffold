package response

type HealthResponse struct {
	Status string `json:"status"`
}

func (r HealthResponse) GetStatus() string {
	return r.Status
}