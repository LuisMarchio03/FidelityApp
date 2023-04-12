package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Show CardFidelity
// @Description Show a cardFidelity
// @Tags CardFidelity
// @Accept json
// @Producer json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param id query string true "CardFidelity identification"
// @Success 200 {object} ShowCardFidelityResponse
// @Failure 404 {object} ErrorResponse
// @Router /cardFidelity [get]
func ShowCardFidelityHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}
	cardFidelity := schemas.CardFidelity{}
	points := []schemas.Point{}

	if err := db.First(&cardFidelity, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error show cardFidelity with id: %s", id),
		)
		return
	}

	if err := db.Find(&points).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing points")
		return
	}

	cardFidelity.Point = points

	sendSuccess(
		ctx,
		http.StatusOK,
		"show-cardFidelity",
		cardFidelity,
	)
}
