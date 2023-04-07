package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary List Companies
// @Description List all Companies
// @Tags Opening
// @Accept json
// @Producer json
// @Success 200 {object} ListCompaniesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /companies [get]
func ListCompaniesHandler(ctx *gin.Context) {
	companies := []schemas.Company{}

	if err := db.Find(&companies).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing companies")
	}

	sendSuccess(ctx, http.StatusOK, "list-companies", companies)
}
