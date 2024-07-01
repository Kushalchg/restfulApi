package routes

import (
	"practice/restfulApi/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.GET("/posts", controllers.GetPost)
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
}

func UserRoutes(r *gin.Engine) {
	r.POST("/user/register", controllers.UserRegister)
}
