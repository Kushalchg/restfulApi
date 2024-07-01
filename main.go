package main

import (
	"practice/restfulApi/initializers"
	"practice/restfulApi/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	r := gin.Default()

	// initializing routes
	routes.PostRoutes(r)
	routes.UserRoutes(r)

	r.Run()
}
