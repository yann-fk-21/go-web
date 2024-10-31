package types

import (
	"time"
)

type OrderStore interface {
	CreateOrder(*Order) error
}

type Order struct {
	ID int `json:"id"`
	Products []Product `json:"products"`
	TotalPrice float64 `json:"totalPrice"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
}

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"Image"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(email string)(*User, error)
	GetUserByID(ID int)(*User, error)
	CreateUser(u User) error
}

type User struct {
	ID        int `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}