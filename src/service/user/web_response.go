package user

import "time"

type FormatUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserFormat is function for mapping data user before send
func UserFormat(user User) FormatUser {
	return FormatUser{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		City:      user.City,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// UsersFormat is function for mapping data user more than 1 before send
func UsersFormat(users []User) []FormatUser {
	formats := []FormatUser{}

	for _, user := range users {
		format := UserFormat(user)
		formats = append(formats, format)
	}

	return formats
}
