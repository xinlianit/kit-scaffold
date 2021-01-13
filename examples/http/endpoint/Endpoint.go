package endpoint

import "github.com/xinlianit/kit-scaffold/examples/http/service"

func NewIndexEndpoint() IndexEndpoint {
	return IndexEndpoint{
		indexService: service.NewIndexService(),
	}
}
