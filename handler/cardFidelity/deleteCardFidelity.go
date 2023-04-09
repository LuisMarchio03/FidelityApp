package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Delete CardFidelity
// @Description Delete a cardFidelity
// @Tags CardFidelity
// @Accept json
// @Producer json
// @Param id query string true "CardFidelity identification"
// @Success 200 {object} DeleteCardFidelityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /cardFidelity [delete]
func DeleteCardFidelityHandler(ctx *gin.Context) {
	// Get id
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}

	cardFidelity := schemas.CardFidelity{}

	// Find
	if err := db.First(&cardFidelity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf(
			"cardFidelity with id: %s not found",
			id,
		))
		return
	}

	// Delete
	if err := db.Delete(&cardFidelity).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting cardFidelity with id: %s", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-cardFidelity", cardFidelity)
}
