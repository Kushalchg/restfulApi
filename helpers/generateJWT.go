package helpers

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var (
	keyValue   []byte
	token      *jwt.Token
	tokenValue string
)

type MyClaims struct {
	Name   string `json:name`
	Role   string `json:role`
	UserId int    `json:user_id`
	jwt.StandardClaims
}

func GenerateAccess(claims MyClaims) (string, error) {

	keyValue = []byte(os.Getenv("KEY"))
	var err error
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenValue, err = token.SignedString(keyValue)

	return tokenValue, err

}

func GenerateRefresh(claims MyClaims) (string, error) {

	keyValue = []byte(os.Getenv("KEY"))
	var err error
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenValue, err = token.SignedString(keyValue)

	return tokenValue, err

}

func ParseAccessToken(accessToken string) *MyClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	return parsedAccessToken.Claims.(*MyClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
	ParseRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	return ParseRefreshToken.Claims.(*jwt.StandardClaims)

}
