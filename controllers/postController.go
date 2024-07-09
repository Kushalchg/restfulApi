package controllers

import (
	"net/http"
	"practice/restfulApi/global"
	"practice/restfulApi/initializers"
	"practice/restfulApi/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// var body models.Post
	var body struct {
		Title string `json:"title" validate:"required"`
		Body  string `json:"body" validate:"required"`
	}

	c.Bind(&body)

	// validate the requested body
	// requested body must contain title and body
	if err := global.Validate.Struct(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "the requested body must contain both Title and Body field",
			"detail": err.Error(),
		})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	// post := models.Post{Title: "title", Body: "body"}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"posts": post,
	})

}
func GetPost(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.IndentedJSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func GetSinglePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func UpdatePost(c *gin.Context) {

	var post models.Post

	//get the id from param
	id := c.Param("id")

	//get the body data
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := c.Bind(&body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// it reterive  the post with given id and store on address of post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//update that post and what store on address of post
	updateResult := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": updateResult.Error.Error(),
		})
	}

	// response with updated data

	c.IndentedJSON(http.StatusOK, gin.H{
		"data":        "successfully updated data",
		"recentValue": body.Title,
		"value":       post,
	})

}

func DeletePost(c *gin.Context) {

	// delete the post with id
	// var post models.Post
	id := c.Param("id")
	// deleteResult := initializers.DB.Delete(&models.Post{}, id)
	var posts []models.Post
	deleteResult := initializers.DB.Exec("DELETE FROM posts where id=?", id)

	if deleteResult.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": deleteResult.Error.Error(),
		})
		return
	}

	if deleteResult.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	// get all post
	postValue := initializers.DB.Find(&posts)
	if postValue.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": postValue.Error.Error(),
		})
		return
	}
	// return remaining posts
	c.IndentedJSON(http.StatusOK, gin.H{
		"detail":       "successfully deleted",
		"posts":        posts,
		"rowsAffected": deleteResult.RowsAffected,
	})

}
