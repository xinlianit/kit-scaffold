package dao

import (
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/column"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/entity"
	"github.com/xinlianit/kit-scaffold/lib/driver"
)

// 商家数据访问
type BusinessDao struct {
}

// 获取商家 - 通过商家ID
func (d *BusinessDao) GetBusinessById(businessId int32) (business entity.BusinessEntity) {
	// 查询字段
	columns := []string{
		column.BusinessColumn.BusinessId,
		column.BusinessColumn.BusinessName,
	}

	// 查询条件
	conditions := &entity.BusinessEntity{
		BusinessId: businessId,
	}

	// 查询数据
	driver.MysqlDB.Select(columns).Where(conditions).First(&business)

	//business.BusinessId = 100
	//business.BusinessName = "test_business==="

	return
}
