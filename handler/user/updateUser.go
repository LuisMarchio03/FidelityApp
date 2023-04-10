package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Update User
// @Description Update a new user
// @Tags User
// @Accept json
// @Producer json
// @Param id query string true "User identification"
// @Param request body UpdateUserRequest true "Request body"
// @Success 201 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [put]
func UpdateUserHandler(ctx *gin.Context) {
	// Get Request
	request := UpdateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Get id
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}

	user := schemas.User{}
	userAlreadyExists := schemas.User{}

	// Find
	if err := db.First(&user, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error update user with id: %s", id),
		)
	}

	// User Already Exists
	if err := db.Where("email = ?", request.Email).First(&userAlreadyExists).Error; err == nil {
		logger.Errorf("user already exists")
		sendError(ctx, http.StatusBadRequest, "user already exists")
		return
	}

	// Update
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Address != "" {
		user.Address = request.Address
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if request.CompanyId <= 0 {
		user.CompanyId = request.CompanyId
	}
	if request.Type != "" {
		user.Type = request.Type
	}

	// Save
	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-user", user)
}
