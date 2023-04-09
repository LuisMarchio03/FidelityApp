package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Delete Point
// @Description Delete a point
// @Tags Point
// @Accept json
// @Producer json
// @Param id query string true "Point identification"
// @Success 200 {object} DeletePointResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /point [delete]
func DeletePointHandler(ctx *gin.Context) {
	// Get id
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}

	point := schemas.Point{}

	// Find
	if err := db.First(&point, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf(
			"point with id: %s not found",
			id,
		))
		return
	}

	// Delete
	if err := db.Delete(&point).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting point with id: %s", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-point", point)
}
