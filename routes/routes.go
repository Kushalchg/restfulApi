package routes

import (
	"practice/restfulApi/controllers"
	"practice/restfulApi/helpers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.GET("/posts", helpers.AuthMiddleware(), controllers.GetPost)
	r.POST("/posts", helpers.AuthMiddleware(), controllers.PostCreate)
	r.GET("/posts/:id", helpers.AuthMiddleware(), controllers.GetSinglePost)
	r.PATCH("/posts/:id", helpers.AuthMiddleware(), controllers.UpdatePost)
	r.DELETE("/posts/:id", helpers.AuthMiddleware(), controllers.DeletePost)
}

func UserRoutes(r *gin.Engine) {
	r.POST("/user/register", controllers.UserRegister)
	r.POST("/user/login", controllers.UserLogin)
}
