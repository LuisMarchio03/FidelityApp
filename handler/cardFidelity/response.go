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

type CreateCardFidelityResponse struct {
	Message string                       `json:"message"`
	Data    schemas.CardFidelityResponse `json:"data"`
}

type ListCardFidelitiesResponse struct {
	Message string                         `json:"message"`
	Data    []schemas.CardFidelityResponse `json:"data"`
}

type DeleteCardFidelityResponse struct {
	Message string                       `json:"message"`
	Data    schemas.CardFidelityResponse `json:"data"`
}

type ShowCardFidelityResponse struct {
	Message string                       `json:"message"`
	Data    schemas.CardFidelityResponse `json:"data"`
}

type UpdateCardFidelityResponse struct {
	Message string                       `json:"message"`
	Data    schemas.CardFidelityResponse `json:"data"`
}
