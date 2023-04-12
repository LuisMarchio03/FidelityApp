package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary List Companies
// @Description List all Companies
// @Tags Company
// @Accept json
// @Producer json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} ListCompaniesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /companies [get]
func ListCompaniesHandler(ctx *gin.Context) {
	companies := []schemas.Company{}
	users := []schemas.User{}

	if err := db.Find(&companies).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing companies")
	}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
	}

	for i := range companies {
		for j := range users {
			if companies[i].ID == users[j].CompanyId {
				companies[i].Users = append(companies[i].Users, users[j])
			}
		}
	}

	sendSuccess(ctx, http.StatusOK, "list-companies", companies)
}
