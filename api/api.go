package api

import (
	"Projects/Car24/car24_api_gateway/api/handlers"
	"Projects/Car24/car24_api_gateway/config"

	_ "Projects/Car24/car24_api_gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {

	//user
	r.POST("/user", h.CreateClient)
	r.GET("/user/:id", h.GetClientByID)
	r.GET("/user", h.GetClientList)
	r.PUT("user/:id", h.UpdateClient)
	r.DELETE("/user/:id", h.DeleteClient)
	r.PATCH("/user/:id", h.UpdatePatchClient)

	//order
	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetOrderByID)
	r.GET("/order", h.GetListOrder)
	r.PUT("order/:id", h.UpdateOrder)
	r.DELETE("order/:id", h.DeleteOrder)
	r.PATCH("order/:id", h.UpdatePatchOrder)

	// otp
	r.POST("/check", h.CreateUserOTP)
	r.GET("/check", h.VerifyUserOTP)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
