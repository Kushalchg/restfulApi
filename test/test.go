package test

import (
	"fmt"
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
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	// generating  access token
	value, err := helpers.GenerateAccess(userClaims)

	if err != nil {
		fmt.Printf("error while generating jwt")
	}

	fmt.Printf("value is %s\n", value)
	//Testing for parsing the access token
	result := helpers.ParseAccessToken(value)
	fmt.Printf("parsed value is %v \n", result)

}
