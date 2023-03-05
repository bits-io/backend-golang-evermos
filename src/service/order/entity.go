package order

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Id           int    `gorm:"not null;uniqueIndex;primary_key"`
	UserInfo     string `gorm:"type:text;not null"`
	UserId       int
	Code         string
	Total        int
	Status       string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	OrderProduct []OrderProduct
	OrderHistory []OrderHistory
}

type OrderProduct struct {
	Id          int `gorm:"not null;uniqueIndex;primary_key"`
	OrderId     int
	ProductInfo string `gorm:"type:text;not null"`
	ProductId   int
	Quantity    int
	SubTotal    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type OrderHistory struct {
	Id        int `gorm:"not null;uniqueIndex;primary_key"`
	OrderId   int
	Status    string
	CreatedBy string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
