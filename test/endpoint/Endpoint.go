package endpoint

import "github.com/xinlianit/kit-scaffold/test/service"

func NewIndexEndpoint() IndexEndpoint {
	return IndexEndpoint{
		indexService: service.NewIndexService(),
	}
}
