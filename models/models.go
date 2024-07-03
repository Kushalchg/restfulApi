package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type UserDetail struct {
	gorm.Model
	Name    string
	Age     int
	Role    string
	Mobile  int
	Address string
}
type User struct {
	gorm.Model
	Email    string `validate:"email,required"`
	Password string `validate:"required,min=8"`
}
