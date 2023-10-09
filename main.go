package main

import (
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/middleware"
	"ecommerce/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	app := controllers.NewApplications(
		database.ProductData(database.Client, "Products"), 
		database.UserData(database.Client, "Users"),
	)

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addcart", app.AddToCart())
	router.GET("/removecart", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}