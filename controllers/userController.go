package controllers

import (
	"log"
	"net/http"
	"practice/restfulApi/global"
	"practice/restfulApi/helpers"
	"practice/restfulApi/initializers"
	"practice/restfulApi/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type responseData struct {
	Refresh string
	Access  string
}

func UserRegister(c *gin.Context) {

	var body struct {
		Email           string `json:"email" validate:"required,email"  `
		Password        string `json:"password" validate:"required,min=8"`
		ConformPassword string `json:"conform_password" validate:"required,eqfield=Password"`
	}

	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Error occured while register try again",
			"detail": err,
		})
		return
	}

	// check the validataion
	// email must be in email format
	// password must contain min 8 letters
	// conform password must match password
	if err := global.Validate.Struct(&body); err != nil {
		global.Logger.Printf("validation Failed %s \n", err)
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
	// user := models.User{Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)

	global.Logger.Printf("the result and error is  %v and %v \n", result, result.Error)
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

func UserLogin(c *gin.Context) {
	var user models.User
	var body struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Unable to signin with error",
			"detail": err,
		})
	}
	// check validation of requested body

	if err := global.Validate.Struct(&body); err != nil {
		global.Logger.Printf("validation Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "required format is not met!",
			"detail": err.Error(),
		})
		return

	}

	// check if the user provided email is present or not?
	// here email must be unique and there is only one accout for one email
	result := initializers.DB.Find(&user, " Email= ? ", body.Email).First(&user)

	// result := initializers.DB.Where(&user{Email: body.Email, Password: body.Password})
	global.Logger.Printf("the result is %v \n", result)
	global.Logger.Printf("the result is %v \n", user)
	global.Logger.Printf("rows affected is  %v \n", result.RowsAffected)
	// log.Fatal("the result is ii ")

	if result.Error != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Unable to signin with error",
			"detail": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Unable to signin",
			"detail": result.Error,
		})
		return
	}

	// check whether the user provided password is right or wrong
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":  "Credentail is wrong",
			"Detail": err,
		})
		return
	}

	// generate access token and refresh token and send on response of login
	accessClaims := helpers.MyClaims{
		Name:   user.Email,
		Type:   "access",
		Role:   "admin",
		UserId: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 15).Unix(),
		},
	}

	refreshClaims := helpers.MyClaims{
		Name:   user.Email,
		Type:   "refresh",
		Role:   "admin",
		UserId: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 500).Unix(),
		},
	}
	accessToken, err := helpers.GenerateAccess(accessClaims)
	if err != nil {

		c.IndentedJSON(http.StatusOK, gin.H{
			"Error":  "Technical Error occured",
			"Detail": err,
		})
		return
	}

	refreshToken, err := helpers.GenerateRefresh(refreshClaims)
	if err != nil {

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":  "Technical Error occured",
			"Detail": err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"success": "Login successful",
		"data": responseData{
			Refresh: refreshToken,
			Access:  accessToken,
		},
	})

}

func UserLogout(c *gin.Context) {
	// request should have valid access token
	// if access token is valid is make access token unvalid.
}
