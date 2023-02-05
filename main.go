package main

import (
	"github/souravdev-eng/resturant/database"
	// "github/souravdev-eng/resturant/middlewares"
	"github/souravdev-eng/resturant/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()

	router.Use(gin.Logger())
	routes.UserRoutes(router)
	// router.Use(middlewares.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	// routes.NoteRoutes(router)
	routes.UserRoutes(router)
	router.Run(":" + port)
}
