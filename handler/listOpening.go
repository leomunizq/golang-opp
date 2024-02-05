package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeninngsHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List opening",
	})
}