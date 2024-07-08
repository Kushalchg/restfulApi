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

func ParseAccessToken(accessToken string) (*MyClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	// if err != nil {
	// 	fmt.Printf("error occur while parsing the jwt token %v \n", err)
	// 	// access token is not valid anymore
	// 	// generate new access Token and return the claims of that new access token
	// 	// newClaims := MyClaims{
	// 	// 	Name:   "kushal",
	// 	// 	Role:   "Developer",
	// 	// 	UserId: 2,
	// 	// 	StandardClaims: jwt.StandardClaims{
	// 	// 		IssuedAt:  time.Now().Unix(),
	// 	// 		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	// 	// 	},
	// 	// }
	// 	// newAccessToken, err := GenerateAccess(newClaims)
	// 	// if err != nil {
	// 	// 	log.Fatal("error creating new access Token")
	// 	// }

	// 	// return ParseAccessToken(newAccessToken)

	// }

	return parsedAccessToken.Claims.(*MyClaims), err
}

func ParseRefreshToken(refreshToken string) *MyClaims {
	//it retrun error on expire tocken so only return the value of token successful token
	ParseRefreshToken, err := jwt.ParseWithClaims(refreshToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	if err != nil {
		// refresh token is not valid anymore
		// since refesh token is not valid user need to logedin
		// make endpoint to check validation of refreshToken and if not valid sent refreshToken is not valid message
	}

	return ParseRefreshToken.Claims.(*MyClaims)

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
