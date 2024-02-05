package router

import "github.com/gin-gonic/gin"


func Initialize(){

	// Creates a gin router with default configuration
	router := gin.Default()
	// Initialize the routes
	initializeRoutes(router)
	// Start serving the application
	router.Run() // listen and serve on 0.0.0.0:8080
}
