package user

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	City     string `json:"city" binding:"required"`
}

type DetailUserRequest struct {
	Id int `uri:"id" binding:"required"`
}
