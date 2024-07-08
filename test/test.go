package test

import (
	"fmt"
	"log"
	"practice/restfulApi/helpers"
	"time"

	"github.com/golang-jwt/jwt"
)

func Test() {
	userClaims := helpers.MyClaims{
		Name:   "kushal",
		Role:   "Developer",
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 15).Unix(),
		},
	}
	// generating  access token
	value, err := helpers.GenerateAccess(userClaims)

	if err != nil {
		fmt.Printf("error while generating jwt")
	}

	fmt.Printf("value is %s\n", value)
	//Testing for parsing the access token
	// result := helpers.ParseAccessToken(value)

	// fmt.Printf("parsed value is %v \n", result)
	staticToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoia3VzaGFsIiwiUm9sZSI6IkRldmVsb3BlciIsIlVzZXJJZCI6MiwiZXhwIjoxNzIwMTM1NzMyLCJpYXQiOjE3MjAwODE3MzJ9.J5dKx4WWO6HnSKMIuqhcM1gkzptbe9yiRMptcoUmQpo"

	// validating accessToken
	token, err := helpers.ParseAccessToken(staticToken)
	if err != nil {
		log.Fatal("error occured while parsing")
	}
	actualTokenValue, err := helpers.GenerateAccess(*token)

	if err != nil {
		// error generating new token
		fmt.Printf("error from generating new token is %v\n", err)

	}
	fmt.Printf("new token is   %v \n", actualTokenValue)

}
