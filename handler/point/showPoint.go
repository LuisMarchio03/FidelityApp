package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Show Point
// @Description Show a point
// @Tags Point
// @Accept json
// @Producer json
// @Param id query string true "Point identification"
// @Success 200 {object} ShowPointResponse
// @Failure 404 {object} ErrorResponse
// @Router /point [get]
func ShowPointHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}
	point := schemas.Point{}

	if err := db.First(&point, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error show point with id: %s", id),
		)
	}

	sendSuccess(
		ctx,
		http.StatusOK,
		"show-point",
		point,
	)
}
