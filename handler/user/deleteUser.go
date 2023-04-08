package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Delete User
// @Description Delete a user
// @Tags User
// @Accept json
// @Producer json
// @Param id query string true "User identification"
// @Success 200 {object} DeleteUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /user [delete]
func DeleteUserHandler(ctx *gin.Context) {
	// Get id
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}

	user := schemas.User{}

	// Find
	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf(
			"user with id: %s not found",
			id,
		))
	}

	// Delete
	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id: %s", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-user", user)
}
