package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Point struct {
	gorm.Model
	CardFidelityId uint
	Point          int
}

type PointResponse struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
	CardFidelityId uint      `json:"cardFidelityId"`
	Point          int       `json:"point"`
}
