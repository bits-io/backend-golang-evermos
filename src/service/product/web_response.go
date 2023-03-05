package product

import (
	"encoding/json"
	"backend-golang-evermos/src/helper"
	"time"
)

type FormatProduct struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Ingredients   []string  `json:"ingredients"`
	Type          string    `json:"type"`
	Price         uint      `json:"price"`
	OriginalPrice uint      `json:"original_price"`
	Stock         int       `json:"stock"`
	Rate          float32   `json:"rate"`
	Image         string    `json:"image"`
	IsActive      bool      `json:"is_active"`
	TotalSold     uint      `json:"total_sold"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ProductFormat is function for mapping format data product before send
func ProductFormat(product Product) FormatProduct {
	ingredients := []string{}
	err := json.Unmarshal([]byte(product.Ingredient), &ingredients)
	helper.PanicIfError(err)

	productType := TYPE_NORMAL
	productPrice := product.Price
	ProductOriginalPrice := product.Price
	productStock := product.Stock
	productTotalSold := product.TotalSold

	if product.Promotion.Id != 0 {
		productType = product.Promotion.Type
		productPrice = product.Promotion.PriceAfterDiscount
		ProductOriginalPrice = product.Promotion.Price
		productStock = product.Promotion.Stock
		productTotalSold = product.Promotion.TotalSold
	}
	return FormatProduct{
		Id:            product.Id,
		Name:          product.Name,
		Description:   product.Description,
		Ingredients:   ingredients,
		Price:         productPrice,
		OriginalPrice: ProductOriginalPrice,
		Type:          productType,
		Stock:         productStock,
		Rate:          product.Rate,
		Image:         product.Image,
		IsActive:      product.IsActive,
		TotalSold:     productTotalSold,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}
}

// ProductsFormat is function for mapping format data product more than 1 before send
func ProductsFormat(products []Product) []FormatProduct {
	formats := []FormatProduct{}

	for _, product := range products {
		format := ProductFormat(product)
		formats = append(formats, format)
	}

	return formats
}
