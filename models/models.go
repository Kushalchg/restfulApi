package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `json:"title" gorm:"unique;not null"`
	Body  string `json:"body" gorm:"unique;not null"`
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
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
}
