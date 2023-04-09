package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreatePoint
type CreatePointRequest struct {
	CardFidelityId uint `json:"cardFidelityId"`
}

func (r *CreatePointRequest) Validate() error {
	if r.CardFidelityId <= 0 {
		return errParamsIsRequired("cardFidelityId", "uint")
	}
	return nil
}

// UpdatePoint
type UpdatePointRequest struct {
	CardFidelityId uint `json:"cardFidelityId"`
}

func (r *UpdatePointRequest) Validate() error {
	// if any field is provider, validation is truthy
	if r.CardFidelityId > 0 {
		return nil
	}
	// If none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
