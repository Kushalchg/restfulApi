package controllers

import (
	"net/http"
	"practice/restfulApi/initializers"
	"practice/restfulApi/models"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body)
	user := models.User{Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)
	// result := initializers.DB.Exec("INSERT INTO users (Email,Password) VALUES (?,?)", body.Email, body.Password)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "error occur while user register",
			"detail": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"value": "user created successfully",
		"data":  user,
	})

}
