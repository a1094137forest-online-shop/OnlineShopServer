package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	v1 "OnlineShopServer/controller/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default())

	apiv1 := router.Group("api/v1")

	onlineShopAPI := apiv1.Group("/onlineShop")
	{
		userAPI := onlineShopAPI.Group("/user")
		{
			userAPI.GET("", v1.GetUser)

			userAPI.POST("", v1.CreateUser)

			//userAPI.PUT("", v1.UpdateUser)
		}

		shopAPI := onlineShopAPI.Group("/shop")
		{
			shopAPI.GET("/productsList", v1.GetProductsList)
			shopAPI.POST("/product", v1.CreateProduct)
			shopAPI.GET("/product", v1.GetProduct)
			shopAPI.PUT("/product", v1.UpdateProduct)
			shopAPI.DELETE("/product", v1.DeleteProduct)
		}
	}

	return router
}
