package routes

import (
	controller "github/souravdev-eng/resturant/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.Engine) {
	r.GET("/foods", controller.GetFoods())
	r.GET("/foods/:food_id", controller.GetFood())
	r.POST("/foods", controller.CreateFood())
	r.PATCH("/foods/:food_id", controller.UpdateFood())
}
