package handler

import (
	"fmt"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":     fmt.Sprintf("operation from handler: %s successfull", op),
		"successCode": code,
		"data":        data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateUserResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

type ListUsersResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.UserResponse `json:"data"`
}

type DeleteUserResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

type ShowUserResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

type UpdateUserResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

type LoginUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type SendCodeForMailResponse struct {
	Message string `json:"message"`
}

type RecoveryPasswordResponse struct {
	Message string `json:"message"`
}
