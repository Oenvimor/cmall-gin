package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowProductService 商品详情的服务
type ShowProductService struct {
}

// Show 商品
func (service *ShowProductService) Show(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
