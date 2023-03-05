package app

import (
	"backend-golang-evermos/src/helper"
	"backend-golang-evermos/src/service/order"
	"backend-golang-evermos/src/service/product"
	"backend-golang-evermos/src/service/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB is function connect to database and migration table
func NewDB(config Config) *gorm.DB {

	DB_USERNAME := config.Get("DB_USERNAME")
	DB_PASSWORD := config.Get("DB_PASSWORD")
	DB_HOST := config.Get("DB_HOST")
	DB_PORT := config.Get("DB_PORT")
	DB_NAME := config.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&product.ProductPromotion{})
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&order.OrderHistory{})
	db.AutoMigrate(&order.OrderProduct{})

	return db
}