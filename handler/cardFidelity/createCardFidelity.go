package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Create CardFidelity
// @Description Create a new cardFidelity
// @Tags CardFidelity
// @Accept json
// @Producer json
// @Param request body CreateCardFidelityRequest true "Request body"
// @Success 201 {object} CreateCardFidelityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cardFidelity [post]
func CreateCardFidelityHandler(ctx *gin.Context) {
	request := CreateCardFidelityRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cardFidelity := schemas.CardFidelity{
		UserId:      request.UserId,
		CompanyId:   request.CompanyId,
		TotalPoints: request.TotalPoints,
		Finished:    request.Finished,
	}

	if err := db.Create(&cardFidelity).Error; err != nil {
		logger.Errorf("error creating cardFidelity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating cardFidelity on database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-cardFidelity", cardFidelity)
}
