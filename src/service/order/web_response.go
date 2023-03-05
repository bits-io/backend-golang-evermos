package order

import (
	"encoding/json"
	"backend-golang-evermos/src/helper"
	"backend-golang-evermos/src/service/product"
	"time"
)

type FormatOrder struct {
	Id          int                  `json:"id"`
	Code        string               `json:"code"`
	Total       int                  `json:"total"`
	Status      string               `json:"status"`
	Description string               `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Products    []FormatOrderProduct `json:"products"`
	Histories   []FormatOrderHistory `json:"histories"`
}

type FormatOrderProduct struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Ingredients   []string  `json:"ingredients"`
	Type          string    `json:"type"`
	Price         uint      `json:"price"`
	OriginalPrice uint      `json:"original_price"`
	Stock         int       `json:"stock"`
	Image         string    `json:"image"`
	IsActive      bool      `json:"is_active"`
	TotalSold     uint      `json:"total_sold"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type FormatOrderHistory struct {
	Status    string    `json:"status"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

// OrderProductFormat is function for mapping format data order product before send
func OrderProductFormat(orderProduct OrderProduct) FormatOrderProduct {
	var productData product.Product
	err := json.Unmarshal([]byte(orderProduct.ProductInfo), &productData)
	helper.PanicIfError(err)

	productFormat := product.ProductFormat(productData)

	format := FormatOrderProduct{
		Id:            productFormat.Id,
		Name:          productFormat.Name,
		Description:   productFormat.Description,
		Ingredients:   productFormat.Ingredients,
		Type:          productFormat.Type,
		Price:         productFormat.Price,
		OriginalPrice: productFormat.OriginalPrice,
		Stock:         productFormat.Stock,
		Image:         productFormat.Image,
		IsActive:      productFormat.IsActive,
		TotalSold:     productFormat.TotalSold,
		CreatedAt:     productFormat.CreatedAt,
		UpdatedAt:     productFormat.UpdatedAt,
	}

	return format
}

// OrderProductsFromat is function for mapping format data order product more than 1 before send
func OrderProductsFromat(orderProducts []OrderProduct) []FormatOrderProduct {
	formats := []FormatOrderProduct{}

	for _, orderProduct := range orderProducts {
		format := OrderProductFormat(orderProduct)
		formats = append(formats, format)
	}

	return formats
}

// OrderHistoryFormat is function for mapping format data order history before send
func OrderHistoryFormat(history OrderHistory) FormatOrderHistory {
	return FormatOrderHistory{
		Status:    history.Status,
		CreatedBy: history.CreatedBy,
		CreatedAt: history.CreatedAt,
	}
}

// OrderHistoriesFormat is function for mapping format data history more than 1 before send
func OrderHistoriesFormat(histories []OrderHistory) []FormatOrderHistory {
	formats := []FormatOrderHistory{}

	for _, history := range histories {
		format := OrderHistoryFormat(history)
		formats = append(formats, format)
	}

	return formats
}

// OrderFormat is function to mapping format data order before send
func OrderFormat(order Order) FormatOrder {
	return FormatOrder{
		Id:          order.Id,
		Code:        order.Code,
		Total:       order.Total,
		Status:      order.Status,
		Description: order.Description,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		Products:    OrderProductsFromat(order.OrderProduct),
		Histories:   OrderHistoriesFormat(order.OrderHistory),
	}
}

// OrdersFormat is function to mapping format data order more than 1 before send
func OrdersFormat(orders []Order) []FormatOrder {
	formats := []FormatOrder{}

	for _, order := range orders {
		format := OrderFormat(order)
		formats = append(formats, format)
	}

	return formats
}
