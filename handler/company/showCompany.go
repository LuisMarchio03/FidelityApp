package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Show Company
// @Description Show a company
// @Tags Company
// @Accept json
// @Producer json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param id query string true "Company identification"
// @Success 200 {object} ShowCompanyResponse
// @Failure 404 {object} ErrorResponse
// @Router /company [get]
func ShowCompanyHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}
	company := schemas.Company{}
	users := []schemas.User{}

	if err := db.First(&company, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error show company with id: %s", id),
		)
	}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	company.Users = users

	sendSuccess(
		ctx,
		http.StatusOK,
		"show-company",
		company,
	)
}
