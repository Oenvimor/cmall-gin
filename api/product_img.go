package api

import (
	"cmall/service"
	"github.com/gin-gonic/gin"
)

// ShowProductImgs 商品详情接口
func ShowProductImgs(c *gin.Context) {
	service := service.ShowImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ShowInfoImgs 商品详情图片接口
func ShowInfoImgs(c *gin.Context) {
	service := service.ShowImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
