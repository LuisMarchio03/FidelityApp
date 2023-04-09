package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Update CardFidelity
// @Description Update a new cardFidelity
// @Tags CardFidelity
// @Accept json
// @Producer json
// @Param id query string true "CardFidelity identification"
// @Param request body UpdateCardFidelityRequest true "Request body"
// @Success 201 {object} UpdateCardFidelityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cardFidelity [put]
func UpdateCardFidelityHandler(ctx *gin.Context) {
	// Get Request
	request := UpdateCardFidelityRequest{}
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

	cardFidelity := schemas.CardFidelity{
		UserId:      request.UserId,
		CompanyId:   request.CompanyId,
		TotalPoints: request.TotalPoints,
		Finished:    request.Finished,
	}

	// Find
	if err := db.First(&cardFidelity, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error update cardFidelity with id: %s", id),
		)
		return
	}

	// Update
	if request.UserId > 0 {
		cardFidelity.UserId = request.UserId
	}
	if request.CompanyId > 0 {
		cardFidelity.CompanyId = request.CompanyId
	}
	if request.TotalPoints > 0 {
		cardFidelity.TotalPoints = request.TotalPoints
	}
	if request.Finished == true || request.Finished == false {
		cardFidelity.Finished = request.Finished
	}

	// Save
	if err := db.Save(&cardFidelity).Error; err != nil {
		logger.Errorf("error updating cardFidelity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating cardFidelity")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-cardFidelity", cardFidelity)
}
