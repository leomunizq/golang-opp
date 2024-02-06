package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leomunizq/golang-opp/schemas"
)

func UpdateOpeninngHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	// Update the opening with the new values
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	//save the changes
	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("Error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error updating opening")
	}

	sendSuccess(ctx, "Update", opening)
}
