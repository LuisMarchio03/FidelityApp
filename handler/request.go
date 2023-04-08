package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateCompany

type CreateCompanyRequest struct {
	Name              string `json:"name"`
	CNPJ              string `json:"cnpj"`
	Address           string `json:"address"`
	AddressNumber     string `json:"addressNumber"`
	AddressComplement string `json:"addressComplement"`
	AddressCity       string `json:"addressCity"`
	AddressState      string `json:"addressState"`
	AddressZipCode    string `json:"addressZipCode"`
}

func (r *CreateCompanyRequest) Validate() error {
	if r.Name == "" {
		return errParamsIsRequired("name", "string")
	}
	if r.CNPJ == "" {
		return errParamsIsRequired("cnpj", "string")
	}
	if r.Address == "" {
		return errParamsIsRequired("address", "string")
	}
	if r.AddressNumber == "" {
		return errParamsIsRequired("addressNumber", "string")
	}
	if r.AddressCity == "" {
		return errParamsIsRequired("addressCity", "string")
	}
	if r.AddressState == "" {
		return errParamsIsRequired("addressState", "string")
	}
	if r.AddressZipCode == "" {
		return errParamsIsRequired("addressZipCode", "string")
	}
	return nil
}
