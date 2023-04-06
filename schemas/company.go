package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name              string
	CNPJ              string
	Address           string
	AddressNumber     string
	AddressComplement string
	AddressCity       string
	AddressState      string
	AddressZipCode    string
	Users             []User
}

type CompanyResponse struct {
	ID                uint      `json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         time.Time `json:"deletedAt,omitempty"`
	Name              string    `json:"name"`
	CNPJ              string    `json:"cnpj"`
	Address           string    `json:"address"`
	AddressNumber     string    `json:"addressNumber"`
	AddressComplement string    `json:"addressComplement"`
	AddressCity       string    `json:"addressCity"`
	AddressState      string    `json:"addressState"`
	AddressZipCode    string    `json:"addressZipCode"`
	Users             []UserResponse
}
