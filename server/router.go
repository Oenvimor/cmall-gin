package server

import (
	"cmall/api"
	"cmall/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	// 路由
	v1 := r.Group("/api/v1")
	{
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("products/:id", api.ShowProduct)
		//商品图片操作
		v1.GET("imgs/:id", api.ShowProductImgs)
		//商品详情图片操作
		v1.GET("info-imgs/:id", api.ShowInfoImgs)
		//分类操作
		v1.GET("categories", api.ListCategories)
		//搜索操作
		v1.POST("searches", api.SearchProducts)
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//收藏夹操作
			authed.GET("favorites/:id", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites", api.DeleteFavorite)
			//订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("user/:id/orders", api.ListOrders)
			authed.GET("orders/:num", api.ShowOrder)
			authed.DELETE("orders/:num", api.DeleteOrder)
			//购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/:id", api.ShowCarts)
			authed.PUT("carts", api.UpdateCart)
			authed.DELETE("carts", api.DeleteCart)
			//收货地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.ShowAddresses)
			authed.PUT("addresses", api.UpdateAddress)
			authed.DELETE("addresses", api.DeleteAddress)
			//数量操作
			authed.GET("counts/:id", api.ShowCount)
		}
	}
	return r
}
