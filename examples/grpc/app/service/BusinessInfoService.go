package service

import (
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/dao"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/entity"
	"github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/model"
)

// 商家信息服务
type BusinessInfoService struct {
	businessDao dao.BusinessDao
}

func (s BusinessInfoService) BusinessInfo(businessId int32) (model.BusinessModel, error) {
	// 获取商家信息
	businessEntity := s.businessDao.GetBusinessById(businessId)

	// todo 获取商家扩展信息
	businessExtendEntity := entity.BusinessExtendEntity{}
	businessExtendEntity.BusinessDesc = "商家描述"

	return model.BusinessModel{
		BusinessEntity:       businessEntity,
		BusinessExtendEntity: businessExtendEntity,
	}, nil
}
