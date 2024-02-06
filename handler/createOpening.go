package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leomunizq/golang-opp/schemas"
)

func CreateOpeninngHandler(ctx *gin.Context) {
	request := CreateOpeninngRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error create opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(ctx, "create-opening-success", opening)
}
