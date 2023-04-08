package handler

import (
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Create Company
// @Description Create a new company
// @Tags Company
// @Accept json
// @Producer json
// @Param request body CreateCompanyRequest true "Request body"
// @Success 201 {object} CreateCompanyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /company [post]
func CreateCompanyHandler(ctx *gin.Context) {
	request := CreateCompanyRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	company := schemas.Company{
		Name:              request.Name,
		CNPJ:              request.CNPJ,
		Address:           request.Address,
		AddressNumber:     request.AddressNumber,
		AddressComplement: request.AddressComplement,
		AddressCity:       request.AddressCity,
		AddressState:      request.AddressState,
		AddressZipCode:    request.AddressZipCode,
	}

	if err := db.Create(&company).Error; err != nil {
		logger.Errorf("error creating company: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-company", company)
}