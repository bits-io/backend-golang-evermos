package product

import (
	"backend-golang-evermos/src/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Create(product Product) Product
	FindAll() (products []Product)
	FindById(Id int) (product Product)
	Update(product Product) Product
	Delete(product Product)
	CreatePromotion(promotion ProductPromotion) ProductPromotion
	UpdatePromotion(promotion ProductPromotion) ProductPromotion
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// Create is function to create data product
func (repo *repository) Create(product Product) Product {
	err := repo.db.Create(&product).Error
	helper.PanicIfError(err)

	return product
}

// FindAll is function to get data all product with Promotion data if product promotion
func (repo *repository) FindAll() (products []Product) {
	err := repo.db.Preload("Promotion").Find(&products).Error
	helper.PanicIfError(err)

	return
}

// FindById is function to get detail data product with param product id
func (repo *repository) FindById(Id int) (product Product) {
	err := repo.db.Preload("Promotion", "product_promotions.is_active = true").Where("id = ?", Id).Find(&product).Error
	helper.PanicIfError(err)

	return
}

// Update is function to update data product
func (repo *repository) Update(product Product) Product {
	err := repo.db.Save(&product).Error
	helper.PanicIfError(err)

	return product
}

// Delete is function to delete data product
func (repo *repository) Delete(product Product) {
	err := repo.db.Delete(&product).Error
	helper.PanicIfError(err)

	return
}

// CreatePromotion is function to create data promotion product
func (repo *repository) CreatePromotion(promotion ProductPromotion) ProductPromotion {
	err := repo.db.Create(&promotion).Error
	helper.PanicIfError(err)

	return promotion
}

// UpdatePromotion is function to update data promotion product
func (repo *repository) UpdatePromotion(promotion ProductPromotion) ProductPromotion {
	err := repo.db.Save(&promotion).Error
	helper.PanicIfError(err)

	return promotion
}
