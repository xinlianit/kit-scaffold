package model

import (
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/entity"
)

// 商家模型
type BusinessModel struct {
	entity.BusinessEntity       // 商家基础信息
	entity.BusinessExtendEntity // 商家扩展信息
}
