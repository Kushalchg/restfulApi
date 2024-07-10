package helpers

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

var (
	keyValue   []byte
	token      *jwt.Token
	tokenValue string
)

type MyClaims struct {
	Name   string
	Type   string
	Role   string
	UserId int
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

func ParseAccessToken(accessToken string) (*MyClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})

	return parsedAccessToken.Claims.(*MyClaims), err
}

func ParseRefreshToken(refreshToken string) (*MyClaims, error) {
	//it retrun error on expire tocken so only return the value of token successful token
	ParseRefreshToken, err := jwt.ParseWithClaims(refreshToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})

	return ParseRefreshToken.Claims.(*MyClaims), err

}

func ValidateAccessToken(accessToken string) string {
	accessClaims, err := ParseAccessToken(accessToken)

	if err != nil {
		log.Fatal("error occured while valideatin")
	}
	//access token is expired
	if accessClaims.Valid() != nil {
		log.Fatal("access token is expired generating new access Token")
		//creating new refresh token
		newAccessToken, err := GenerateAccess(*accessClaims)

		if err != nil {
			log.Fatal("error creating new access token")
		}

		return newAccessToken

	}
	return accessToken
}
