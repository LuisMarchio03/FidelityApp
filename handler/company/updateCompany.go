package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/acim-backend/schemas"
	"github.com/gin-gonic/gin"
)

// @BaseParh /api/v1

// @Summary Update Company
// @Description Update a new company
// @Tags Company
// @Accept json
// @Producer json
// @Param id query string true "Company identification"
// @Param request body UpdateCompanyRequest true "Request body"
// @Success 201 {object} UpdateCompanyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /company [put]
func UpdateCompanyHandler(ctx *gin.Context) {
	// Get Request
	request := UpdateCompanyRequest{}
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

	company := schemas.Company{}
	companyAlredyExists := schemas.Company{}

	// Find
	if err := db.First(&company, id).Error; err != nil {
		sendError(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("error update company with id: %s", id),
		)
	}

	// Company Alredy Exists
	if err := db.Where("cnpj = ?", request.CNPJ).First(&companyAlredyExists).Error; err == nil {
		logger.Errorf("company already exists")
		sendError(ctx, http.StatusBadRequest, "company already exists")
		return
	}

	// Update
	if request.Name != "" {
		company.Name = request.Name
	}
	if request.Address != "" {
		company.Address = request.Address
	}
	if request.AddressCity != "" {
		company.AddressCity = request.AddressCity
	}
	if request.AddressComplement != "" {
		company.AddressComplement = request.AddressComplement
	}
	if request.AddressNumber != "" {
		company.AddressNumber = request.AddressNumber
	}
	if request.AddressState != "" {
		company.AddressState = request.AddressState
	}
	if request.AddressZipCode != "" {
		company.AddressZipCode = request.AddressZipCode
	}
	if request.CNPJ != "" {
		company.CNPJ = request.CNPJ
	}

	// Save
	if err := db.Save(&company).Error; err != nil {
		logger.Errorf("error updating company: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating company")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-company", company)
}
