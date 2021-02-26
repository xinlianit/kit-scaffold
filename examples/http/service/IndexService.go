package service

import (
	"github.com/xinlianit/kit-scaffold/examples/http/repository/entity"
)

// IndexService index
type IndexService struct {
}

// Hello hello
func (s IndexService) Hello(id int64) (entity.HelloEntity, error) {
	helloEntity := entity.HelloEntity{}

	helloEntity.Id = id
	helloEntity.Name = "Hello"
	helloEntity.Desc = "测试Hello"

	return helloEntity, nil
}
