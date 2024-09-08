package main

import (
	"JWTgo/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatalln("Port is not found in environment")
	}
	fmt.Println("Port : ", portString)

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for API-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for API-2"})
	})

	router.Run(":" + portString)
}
