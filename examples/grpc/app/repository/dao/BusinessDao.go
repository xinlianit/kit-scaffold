package dao

import (
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/entity"
)

// 商家数据访问
type BusinessDao struct {
}

// 获取商家 - 通过商家ID
func (d BusinessDao) GetBusinessById(businessId int32) entity.BusinessEntity {
	// todo 数据库查询
	return entity.BusinessEntity{
		BusinessId:   100,
		BusinessName: "test_business",
	}
}
