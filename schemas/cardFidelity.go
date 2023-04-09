package schemas

import (
	"time"

	"gorm.io/gorm"
)

type CardFidelity struct {
	gorm.Model
	UserId      uint
	CompanyId   uint
	TotalPoints int
	Finished    bool
	Point       []Point
}

type CardFidelityResponse struct {
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	UserId      uint      `json:"userId"`
	CompanyId   uint      `json:"companyId"`
	TotalPoints int       `json:"totalPoints"`
	Finished    bool      `json:"finished"`
	Point       []PointResponse
}
