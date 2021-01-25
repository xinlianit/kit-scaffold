package dao

import (
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/entity"
)

// 商家数据访问
type BusinessDao struct {
	businessEntity entity.BusinessEntity
}

// 获取商家 - 通过商家ID
func (d BusinessDao) GetBusinessById(businessId int32) entity.BusinessEntity {

	//return d.businessEntity

	// todo 数据库查询
	d.businessEntity.BusinessId = 100
	d.businessEntity.BusinessName = "test_business==="

	return d.businessEntity
}
