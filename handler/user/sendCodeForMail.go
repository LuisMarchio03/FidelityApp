package handler

import (
	"net/http"
	"os/exec"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/LuisMarchio03/acim-backend/service"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Send Code For Mail
// @Description Send Code For Mail
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body SendCodeForMailRequest true "Request body"
// @Success 200 {object} SendCodeForMailResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/send-code-for-mail [post]
func SendCodeForMailHandler(ctx *gin.Context) {
	// JSON Body Request
	request := SendCodeForMailRequest{}
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

	// Send code for email
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		logger.Errorf("error generating code: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error generating code")
		return
	}

	code := string(newUUID)
	if err := service.SendEmail(user.Email, code); err != nil {
		logger.Errorf("error sending email: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error sending email")
		return
	}

	// Include code in database user
	if code != "" {
		user.Code = code
	}

	// Save
	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user code: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user code")
		return
	}

	sendSuccess(ctx, http.StatusOK, "send-code-for-mail", code)
}
