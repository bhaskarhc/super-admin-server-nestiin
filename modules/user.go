package modules

import "gorm.io/gorm"

type UserData struct {
	gorm.Model
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	CreatedAt  string `json:"created_at"`
	UpdateAt   string `json:"update_at"`
	IsVerified bool   `json:"is_verified"`
}
