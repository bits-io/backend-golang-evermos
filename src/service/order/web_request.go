package order

type CreateOrderRequest struct {
	UserId    int `json:"user_id" binding:"required"`
	ProductId int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type DetailOrderRequest struct {
	Id int `uri:"id" binding:"required"`
}
