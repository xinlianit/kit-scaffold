package service

import "github.com/xinlianit/kit-scaffold/examples/http/repository/entity"

type IndexService struct {
}

func (s IndexService) Hello(id int64) (entity.HelloEntity, error) {
	helloEntity := entity.HelloEntity{}

	helloEntity.Id = id
	helloEntity.Name = "Hello"
	helloEntity.Desc = "测试Hello"

	return helloEntity, nil
}
