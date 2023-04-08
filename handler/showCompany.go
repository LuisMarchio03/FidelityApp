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

	if err := db.First(&company, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error show company with id: %s", id),
		)
	}

	sendSuccess(
		ctx,
		http.StatusOK,
		"show-company",
		company,
	)
}
