package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Delete Company
// @Description Delete a company
// @Tags Company
// @Accept json
// @Producer json
// @Param id query string true "Company identification"
// @Success 200 {object} DeleteCompanyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /company [delete]
func DeleteCompanyHandler(ctx *gin.Context) {
	// Get id
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}

	company := schemas.Company{}

	// Find
	if err := db.First(&company, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf(
			"company with id: %s not found",
			id,
		))
	}

	// Delete
	if err := db.Delete(&company).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting company with id: %s", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-company", company)
}
