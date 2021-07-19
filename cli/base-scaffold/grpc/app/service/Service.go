package service

import "github.com/xinlianit/kit-scaffold/examples/grpc/app/repository/dao"

// 创建商家信息服务
func NewBusinessInfoService() BusinessInfoService {
	return BusinessInfoService{
		businessDao: dao.NewBusinessDao(),
	}
}
