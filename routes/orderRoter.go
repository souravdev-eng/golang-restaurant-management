package routes

import (
	controller "github/souravdev-eng/resturant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.GET("/orders", controller.GetOrders())
	r.GET("/orders/:order_id", controller.GetOrder())
	r.POST("/orders", controller.CreateOrder())
	r.PATCH("/orders/:order_id", controller.UpdateOrder())
}
