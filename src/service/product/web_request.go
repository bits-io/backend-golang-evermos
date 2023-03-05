package product

type CreateProductRequest struct {
	Name               string   `form:"name" binding:"required"`
	Description        string   `form:"description" binding:"required"`
	Ingredient         []string `form:"ingredient[]" binding:"required"`
	Price              uint     `form:"price" binding:"required"`
	Type               string   `form:"type" binding:"required"`
	Image              string   `form:"image" binding:"required"`
	Stock              uint     `form:"stock" binding:"required"`
	StockPromotion     uint     `form:"stock_promotion" binding:"required"`
	PriceAfterDiscount uint     `form:"price_after_discount" binding:"required"`
}

type DetailProductRequest struct {
	Id int `uri:"id" binding:"required"`
}
