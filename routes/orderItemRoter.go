package routes

import (
	controller "github/souravdev-eng/resturant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(r *gin.Engine) {
	r.GET("/orderItems", controller.GetOrderItems())
	r.GET("/orderItems/:orderItems_id", controller.GetOrderItem())
	r.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	r.POST("/orderItems", controller.CreateOrderItem())
	r.PATCH("/orderItems/:orderItems_id", controller.UpdateOrderItem())
}
