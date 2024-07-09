package controllers

import (
	"net/http"
	"practice/restfulApi/helpers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateNewTokens(c *gin.Context) {
	var body struct {
		Refresh string `josn:"refresh" validate:"required"`
	}

	c.Bind(&body)
	mc, err := helpers.ParseRefreshToken(body.Refresh)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "can't create now access and refresh token",
			"detail": err.Error(),
		})
		return
	}
	if mc.Valid() != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "can't create now access and refresh token",
		})
		return
	}
	newAccessClaims := helpers.MyClaims{
		Name:   "kushl",
		Role:   "admin",
		UserId: 20,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
		},
	}

	newAccessToken, err := helpers.GenerateAccess(newAccessClaims)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{

			"error": "error generating new access token",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": newAccessToken,
	})

}
