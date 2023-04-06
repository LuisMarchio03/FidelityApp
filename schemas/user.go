package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Address   string
	Phone     string
	Email     string
	Password  string
	CompanyId uint
	Type      string
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CompanyId uint      `json:"companyId"`
	Type      string    `json:"type"`
}
