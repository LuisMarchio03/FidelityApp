package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Login User
// @Description Login a user
// @Tags Auth
// @Accept json
// @Producer json
// @Param request body LoginUserRequest true "Request body"
// @Success 201 {object} LoginUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func LoginUserHandler(ctx *gin.Context) {
	// JSON Body Request
	request := LoginUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Find user on database
	user := schemas.User{}

	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		logger.Errorf("user not found")
		sendError(ctx, http.StatusBadRequest, "user not found")
		return
	}

	// Compare password
	if err := user.ComparePassword(request.Password); err != nil {
		logger.Errorf("invalid password")
		sendError(ctx, http.StatusBadRequest, "invalid password")
		return
	}

	// Generate token
	token, err := user.GenerateToken()
	if err != nil {
		logger.Errorf("error generating token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error generating token")
		return
	}

	sendSuccess(ctx, http.StatusOK, "login-user", LoginUserResponse{
		Message: "login successful",
		Token:   token,
	})
}
