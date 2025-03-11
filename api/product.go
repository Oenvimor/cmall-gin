package api

import (
	"cmall/pkg/logging"
	"cmall/service"
	"github.com/gin-gonic/gin"
)

// ListProducts 商品列表接口
func ListProducts(c *gin.Context) {
	service := service.ListProductsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowProduct 商品详情接口
func ShowProduct(c *gin.Context) {
	service := service.ShowProductService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// SearchProducts 搜索商品的接口
func SearchProducts(c *gin.Context) {
	service := service.SearchProductsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
