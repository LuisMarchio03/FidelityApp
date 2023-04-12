package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Create Point
// @Description Create a new point
// @Tags Point
// @Accept json
// @Producer json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param request body CreatePointRequest true "Request body"
// @Success 201 {object} CreatePointResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /point [post]
func CreatePointHandler(ctx *gin.Context) {
	request := CreatePointRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	point := schemas.Point{
		CardFidelityId: request.CardFidelityId,
		Point:          1,
	}

	if err := db.Create(&point).Error; err != nil {
		logger.Errorf("error creating point: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating point on database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-point", point)
}
