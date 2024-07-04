package controllers

import (
	"net/http"
	"practice/restfulApi/initializers"
	"practice/restfulApi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

}

func UserRegister(c *gin.Context) {
	var body struct {
		Email    string `validate:"email,required"`
		Password string `validate:"required,min=8"`
	}

	if err := Validate.Struct(body); err != nil {
		// fmt.Printf("validation Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "missmatch with required format",
			"detail": err,
		})
		return

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
