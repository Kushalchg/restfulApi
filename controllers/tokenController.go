package controllers

import (
	"net/http"
	"practice/restfulApi/global"
	"practice/restfulApi/helpers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateNewTokens(c *gin.Context) {
	var body struct {
		Refresh string `json:"refresh" validate:"required"`
	}

	c.Bind(&body)
	if len(body.Refresh) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Your provided refresh token is empty",
		})
		return
	}
	global.Logger.Print(" body struct from token controller", body.Refresh)
	global.Logger.Print(" refresh token from token controller", body.Refresh)
	mc, err := helpers.ParseAccessToken(body.Refresh)

	global.Logger.Print("value of claims and err respectively from token controller", mc, err)
	global.Logger.Print("log from create New refresh token")
	// while creating new access token if the refresh token is also expired
	// response code will be 400
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "Refresh token is expired or not valid ",
			"detail": err.Error(),
		})
		return
	}
	if mc.Valid() != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "  Refresh token is expired or not valid ",
		})
		return
	}

	if mc.Type != "refresh" {

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Remember you have to pass refresh token",
		})
		return
	}

	newAccessClaims := helpers.MyClaims{
		Name:   mc.Name,
		Type:   "access",
		Role:   mc.Role,
		UserId: mc.UserId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
		},
	}
	newRefreshClaims := helpers.MyClaims{
		Name:   mc.Name,
		Type:   "refresh",
		Role:   mc.Role,
		UserId: mc.UserId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 50).Unix(),
		},
	}

	newAccessToken, err := helpers.GenerateAccess(newAccessClaims)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{

			"error": "error generating new access token",
		})
		return
	}

	newRefreshToken, err := helpers.GenerateRefresh(newRefreshClaims)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{

			"error": "error generating new refresh token",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"access":  newAccessToken,
		"refresh": newRefreshToken,
	})

}
