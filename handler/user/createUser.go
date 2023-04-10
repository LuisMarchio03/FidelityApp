package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Create User
// @Description Create a new user
// @Tags User
// @Accept json
// @Producer json
// @Param request body CreateUserRequest true "Request body"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Name:      request.Name,
		Email:     request.Email,
		Address:   request.Address,
		Phone:     request.Phone,
		Password:  request.Password,
		CompanyId: request.CompanyId,
		Type:      request.Type,
	}

	if err := db.Where("email = ?", request.Email).First(&user).Error; err == nil {
		logger.Errorf("user already exists")
		sendError(ctx, http.StatusBadRequest, "user already exists")
		return
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user on database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-user", user)
}
