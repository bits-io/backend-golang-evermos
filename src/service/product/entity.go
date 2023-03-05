package product

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id          int    `gorm:"not null;uniqueIndex;primary_key"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text;not null"`
	Ingredient  string `gorm:"type:text;not null"`
	Price       uint
	Stock       int
	Rate        float32
	Image       string `gorm:"size:255;not null"`
	TotalSold   uint
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Promotion   ProductPromotion
}

type ProductPromotion struct {
	Id                 int `gorm:"not null;uniqueIndex;primary_key"`
	ProductId          int
	Type               string
	IsActive           bool
	Price              uint
	PriceAfterDiscount uint
	Stock              int
	TotalSold          uint
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}
