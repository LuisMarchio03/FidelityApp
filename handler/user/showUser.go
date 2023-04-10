package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Show User
// @Description Show a user
// @Tags User
// @Accept json
// @Producer json
// @Param id query string true "User identification"
// @Success 200 {object} ShowUserResponse
// @Failure 404 {object} ErrorResponse
// @Router /user [get]
func ShowUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}
	user := schemas.User{}

	if err := db.First(&user, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error show user with id: %s", id),
		)
		return
	}

	sendSuccess(
		ctx,
		http.StatusOK,
		"show-user",
		user,
	)
}
