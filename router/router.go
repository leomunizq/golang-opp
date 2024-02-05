package router

import "github.com/gin-gonic/gin"


func Initialize(){

	// Creates a gin router with default configuration
	router := gin.Default()
	// Define a route handler for the GET /ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Start serving the application
	router.Run() // listen and serve on 0.0.0.0:8080
}
