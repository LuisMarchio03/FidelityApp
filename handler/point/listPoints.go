package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary List Points
// @Description List all Points
// @Tags Point
// @Accept json
// @Producer json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} ListPointsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /points [get]
func ListPointsHandler(ctx *gin.Context) {
	points := []schemas.Point{}

	if err := db.Find(&points).Error; err != nil {

		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error listing points: %v", err))
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-points", points)
}
