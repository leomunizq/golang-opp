package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leomunizq/golang-opp/handler"
)

func initializeRoutes(router *gin.Engine) {
	// initialize handler
	handler.InitializeHandler()

	// Create a new group of routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)

		v1.POST("/opening", handler.CreateOpeninngHandler)

		v1.DELETE("/opening", handler.DeleteOpeninngHandler)

		v1.PUT("/opening", handler.UpdateOpeninngHandler)

		v1.GET("/openings", handler.ListOpeninngsHandler)
	}
}
