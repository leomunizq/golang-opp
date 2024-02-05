package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET openings",
	})
}

func CreateOpeninngHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST opening",
	})
}

func DeleteOpeninngHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}

func UpdateOpeninngHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}

func ListOpeninngsHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List opening",
	})
}