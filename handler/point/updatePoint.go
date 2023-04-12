package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Update Point
// @Description Update a new point
// @Tags Point
// @Accept json
// @Producer json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param id query string true "Point identification"
// @Param request body UpdatePointRequest true "Request body"
// @Success 201 {object} UpdatePointResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /point [put]
func UpdatePointHandler(ctx *gin.Context) {
	// Get Request
	request := UpdatePointRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Get id
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}

	point := schemas.Point{
		CardFidelityId: request.CardFidelityId,
		Point:          1,
	}

	// Find
	if err := db.First(&point, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error update point with id: %s", id),
		)
	}

	// Update
	if request.CardFidelityId <= 0 {
		point.CardFidelityId = request.CardFidelityId
	}

	// Save
	if err := db.Save(&point).Error; err != nil {
		logger.Errorf("error updating point: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating point")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-point", point)
}
