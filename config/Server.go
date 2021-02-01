package config

type Server struct {
	Host string `json:"host"`
	Port int `json:"port"`
	ReadTimeout int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
	ContextTimeout int `json:"context_timeout"`
	Gateway struct{
		Host string `json:"host"`
		Port int `json:"port"`
		ReadTimeout int `json:"read_timeout"`
		WriteTimeout int `json:"write_timeout"`
		ContextTimeout int `json:"context_timeout"`
	} `json:"gateway"`
	Grpc struct{
		Reflection struct{
			Register bool `json:"register"`
		} `json:"reflection"`
	} `json:"grpc"`
}
