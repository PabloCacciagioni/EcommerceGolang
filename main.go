package main

import (
	"log"
	"os"

	"github.com/PabloCacciagioni/EcommerceGolang/controllers"
	"github.com/PabloCacciagioni/EcommerceGolang/database"
	"github.com/PabloCacciagioni/EcommerceGolang/middleware"
	"github.com/PabloCacciagioni/EcommerceGolang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcehckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
