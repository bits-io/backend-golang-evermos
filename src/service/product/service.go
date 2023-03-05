package product

import (
	"errors"
	"backend-golang-evermos/src/helper"
	"time"
)

type Service interface {
	Save(input CreateProductRequest) Product
	FindAll() (products []Product)
	FindById(Id int) (product Product, err error)
	Update(product Product) Product
	GetPriceProduct(product Product) uint
	CreatePromotion(product Product, input CreateProductRequest) ProductPromotion
	UpdatePromotion(promotion ProductPromotion) ProductPromotion
	Delete(product Product)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

// Save is function to create data product with promotion if product promotion
func (s *service) Save(input CreateProductRequest) Product {

	ingredients := helper.ConvertDataToJsonString(input.Ingredient)
	product := Product{
		Name:        input.Name,
		Description: input.Description,
		Ingredient:  ingredients,
		Price:       input.Price,
		Stock:       int(input.Stock),
		TotalSold:   0,
		Rate:        0,
		Image:       input.Image,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	product = s.repository.Create(product)

	if input.Type != TYPE_NORMAL {
		_ = s.CreatePromotion(product, input)
	}
	return product
}

// FindAll is function to get all data products
func (s *service) FindAll() (products []Product) {
	products = s.repository.FindAll()
	return
}

// FindById is function to get data detail product
func (s *service) FindById(Id int) (product Product, err error) {
	product = s.repository.FindById(Id)
	if product.Id == 0 {
		return product, errors.New("product not found")
	}
	return
}

// GetPriceProduct is function to get price product
func (s *service) GetPriceProduct(product Product) uint {
	price := product.Price
	if product.Promotion.Id != 0 {
		price = product.Promotion.PriceAfterDiscount
	}
	return price
}

// Update is function to update data product
func (s *service) Update(product Product) Product {
	product = s.repository.Update(product)
	return product
}

// CreatePromotion is function to create data product promotion
func (s *service) CreatePromotion(product Product, input CreateProductRequest) ProductPromotion {
	promotion := ProductPromotion{
		ProductId:          product.Id,
		Type:               input.Type,
		IsActive:           true,
		Price:              product.Price,
		PriceAfterDiscount: input.PriceAfterDiscount,
		Stock:              int(input.StockPromotion),
		TotalSold:          0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	promotion = s.repository.CreatePromotion(promotion)
	return promotion
}

// UpdatePromotion is function to update data product promotion
func (s *service) UpdatePromotion(promotion ProductPromotion) ProductPromotion {
	promotion = s.repository.UpdatePromotion(promotion)
	return promotion
}

// Delete is function to delete data product
func (s *service) Delete(product Product) {
	s.repository.Delete(product)

	return
}
