package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary List CardFidelities
// @Description List all CardFidelities
// @Tags CardFidelity
// @Accept json
// @Producer json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} ListCardFidelitiesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /cardFidelities [get]
func ListCardFidelitiesHandler(ctx *gin.Context) {
	cardFidelities := []schemas.CardFidelity{}
	points := []schemas.Point{}

	if err := db.Find(&cardFidelities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error listing cardFidelities: %v", err))
		return
	}

	if err := db.Find(&points).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error listing points: %v", err))
		return
	}

	for i := 0; i < len(cardFidelities); i++ {
		for j := 0; j < len(points); j++ {
			if cardFidelities[i].ID == points[j].CardFidelityId {
				cardFidelities[i].Point = append(cardFidelities[i].Point, points[j])
			}
		}
	}

	sendSuccess(ctx, http.StatusOK, "list-cardFidelities", cardFidelities)
}
