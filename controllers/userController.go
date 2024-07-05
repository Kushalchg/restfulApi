package controllers

import (
	"fmt"
	"log"
	"net/http"
	"practice/restfulApi/initializers"
	"practice/restfulApi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

}

func UserRegister(c *gin.Context) {

	var body struct {
		Email           string `validate:"required,email"`
		Password        string `validate:"required,min=8"`
		ConformPassword string `validate:"required,eqfield=Password"`
	}
	c.Bind(&body)

	// check the validataion
	// email must be in email format
	// password must contain min 8 letters
	// conform password must match password
	if err := Validate.Struct(&body); err != nil {
		fmt.Printf("validation Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "required format is not met!",
			"detail": err.Error(),
		})
		return

	}

	// create hash password
	// hash password in []byte type
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		// error occured on creating hash password
		log.Fatal("error on creating hash password")

	}

	//It takes ConformPassword from user but doesn't upload ot the database
	// ConformPassword is there to prevent user to enter unintended password.
	user := models.User{Email: body.Email, Password: string(hashPassword)}

	result := initializers.DB.Create(&user)

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
