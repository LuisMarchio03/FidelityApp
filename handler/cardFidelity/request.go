package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateCardFidelity
type CreateCardFidelityRequest struct {
	UserId      uint `json:"userId"`
	CompanyId   uint `json:"companyId"`
	TotalPoints int  `json:"totalPoints"`
	Finished    bool `json:"finished"`
}

func (r *CreateCardFidelityRequest) Validate() error {
	if r.UserId <= 0 {
		return errParamsIsRequired("userId", "uint")
	}
	if r.CompanyId <= 0 {
		return errParamsIsRequired("companyId", "uint")
	}
	if r.TotalPoints <= 0 {
		return errParamsIsRequired("totalPoints", "int")
	}
	if r.Finished == false {
		return errParamsIsRequired("finished", "bool")
	}
	return nil
}

// UpdateCardFidelity
type UpdateCardFidelityRequest struct {
	UserId      uint `json:"userId"`
	CompanyId   uint `json:"companyId"`
	TotalPoints int  `json:"totalPoints"`
	Finished    bool `json:"finished"`
}

func (r *UpdateCardFidelityRequest) Validate() error {
	// if any field is provider, validation is truthy
	if r.UserId > 0 && r.CompanyId > 0 && r.TotalPoints > 0 && r.Finished == true || r.Finished == false {
		return nil
	}
	// If none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
