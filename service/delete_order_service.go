package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
	"github.com/jinzhu/gorm"
)

// DeleteOrderService 删除订单服务
type DeleteOrderService struct{}

// Delete 根据订单号删除订单
func (service *DeleteOrderService) Delete(num string) serializer.Response {
	var order model.Order
	code := e.SUCCESS

	// 根据订单号查找订单
	if err := model.DB.Where("order_num = ?", num).First(&order).Error; err != nil {
		// 如果订单不存在，返回错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_ORDER
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 删除订单
	if err := model.DB.Delete(&order).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
