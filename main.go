package main

import (
	"practice/restfulApi/initializers"
	"practice/restfulApi/routes"
	"practice/restfulApi/test"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	test.Test()
	r := gin.Default()
	// Registering routers
	routes.PostRoutes(r)
	routes.UserRoutes(r)

	r.Run()
}
