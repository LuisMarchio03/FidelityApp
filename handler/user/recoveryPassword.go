package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Recovery Password
// @Description Recovery Password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RecoveryPasswordRequest true "Request body"
// @Success 200 {object} RecoveryPasswordResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/recovery-password [post]
func RecoveryPasswordHandler(ctx *gin.Context) {
	request := RecoveryPasswordRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Find user on database
	user := schemas.User{}
	if err := db.Where("code = ?", request.Code).First(&user).Error; err != nil {
		logger.Errorf("user not found")
		sendError(ctx, http.StatusBadRequest, "user not found")
		return
	}

	// Compare password
	if err := user.ComparePassword(request.Password); err != nil {
		logger.Errorf("invalid credentials")
		sendError(ctx, http.StatusBadRequest, "invalid credentials")
		return
	}

	// Alter Password
	if request.NewPassword != request.PasswordConfirm {
		logger.Errorf("passwords do not match")
		sendError(ctx, http.StatusBadRequest, "passwords do not match")
		return
	}

	passwordHash, err := schemas.HashPassword(request.NewPassword)
	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error hashing password")
		return
	}

	if request.NewPassword != "" {
		user.Password = passwordHash
		user.Code = ""
	}

	// Save
	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user password")
		return
	}
}
