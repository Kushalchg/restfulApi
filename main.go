package main

import (
	"practice/restfulApi/initializers"
	"practice/restfulApi/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {

	Validate = validator.New(validator.WithRequiredStructEnabled())
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	// test.Test()
	r := gin.Default()
	// Registering routers
	routes.PostRoutes(r)
	routes.UserRoutes(r)

	r.Run()
}
