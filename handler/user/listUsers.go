package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary List Users
// @Description List all Users
// @Tags User
// @Accept json
// @Producer json
// @Success 200 {object} ListUsersResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users [get]
func ListUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
	}

	sendSuccess(ctx, http.StatusOK, "list-users", users)
}
