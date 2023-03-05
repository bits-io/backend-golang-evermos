package order

import (
	"backend-golang-evermos/src/helper"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	WithTrx(trxHandle *gorm.DB) *repository
	Create(order Order) Order
	FindAll() (orders []Order)
	FindById(Id int) (order Order)
	FindByProductId(productId int) (orders []Order)
	CreateOrderProduct(product OrderProduct) OrderProduct
	CreateOrderHistory(history OrderHistory) OrderHistory
	Delete(order Order)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// WithTrx is function for connection database with transaction
func (repo *repository) WithTrx(trxHandle *gorm.DB) *repository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return repo
	}
	repo.db = trxHandle
	return repo
}

// Create is function create to table orders
func (repo *repository) Create(order Order) Order {
	err := repo.db.Create(&order).Error
	helper.PanicIfError(err)

	return order
}

// FindAll is function to get all data orders with orderProducts and orderHistories
func (repo *repository) FindAll() (orders []Order) {
	err := repo.db.Preload("OrderProduct").Preload("OrderHistory").Find(&orders).Error
	helper.PanicIfError(err)

	return
}

// FindById is function to get detail order with data order products and order histories
func (repo *repository) FindById(Id int) (order Order) {
	err := repo.db.Preload("OrderProduct").Preload("OrderHistory").Where("id = ?", Id).Find(&order).Error
	helper.PanicIfError(err)

	return
}

// FindByProductId is function to get all data orders with condition where product_id at order_products
func (repo *repository) FindByProductId(productId int) (orders []Order) {
	err := repo.db.Preload("OrderProduct", "order_products.product_id = ?", productId).Find(&orders).Error
	helper.PanicIfError(err)

	return
}

// CreateOrderProduct is function for create order product
func (repo *repository) CreateOrderProduct(product OrderProduct) OrderProduct {
	err := repo.db.Create(&product).Error
	helper.PanicIfError(err)

	return product
}

// CreateOrderHistory is function to create order history
func (repo *repository) CreateOrderHistory(history OrderHistory) OrderHistory {
	err := repo.db.Create(&history).Error
	helper.PanicIfError(err)

	return history
}

// Delete is function to delete data order
func (repo *repository) Delete(order Order) {
	err := repo.db.Delete(&order).Error
	helper.PanicIfError(err)

	return
}
