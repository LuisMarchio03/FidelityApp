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

type CreatePointResponse struct {
	Message string                `json:"message"`
	Data    schemas.PointResponse `json:"data"`
}

type ListPointsResponse struct {
	Message string                  `json:"message"`
	Data    []schemas.PointResponse `json:"data"`
}

type DeletePointResponse struct {
	Message string                `json:"message"`
	Data    schemas.PointResponse `json:"data"`
}

type ShowPointResponse struct {
	Message string                `json:"message"`
	Data    schemas.PointResponse `json:"data"`
}

type UpdatePointResponse struct {
	Message string                `json:"message"`
	Data    schemas.PointResponse `json:"data"`
}
