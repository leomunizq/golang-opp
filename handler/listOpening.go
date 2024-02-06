package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leomunizq/golang-opp/schemas"
)

func ListOpeninngsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing openings")
		return
	}
	sendSuccess(ctx, "List", openings)
}
