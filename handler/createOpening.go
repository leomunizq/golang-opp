package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeninngHandler(ctx *gin.Context) {
	request := CreateOpeninngRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("error validating request: %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errf("error create opening: %v", err.Error())
		return
	}
}
