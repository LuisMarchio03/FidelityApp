package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateUser

type CreateUserRequest struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CompanyId uint   `json:"companyId"`
	Type      string `json:"type"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" {
		return errParamsIsRequired("name", "string")
	}
	if r.Address == "" {
		return errParamsIsRequired("address", "string")
	}
	if r.Phone == "" {
		return errParamsIsRequired("phone", "string")
	}
	if r.Email == "" {
		return errParamsIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamsIsRequired("password", "string")
	}
	if r.CompanyId <= 0 {
		return errParamsIsRequired("companyId", "uint")
	}
	if r.Type == "" {
		return errParamsIsRequired("type", "string")
	}
	if r.Type != "user" && r.Type != "employee" {
		return fmt.Errorf("type must be user or employee")
	}
	return nil
}

// UpdateUser

type UpdateUserRequest struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CompanyId uint   `json:"companyId"`
	Type      string `json:"type"`
}

func (r *UpdateUserRequest) Validate() error {
	// if any field is provider, validation is truthy
	if r.Name != "" && r.Address != "" && r.Phone != "" && r.Email != "" && r.Password != "" && r.CompanyId > 0 && r.Type != "" {
		return nil
	}
	// If none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}

// Login User

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() error {
	if r.Email == "" {
		return errParamsIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamsIsRequired("password", "string")
	}
	return nil
}
