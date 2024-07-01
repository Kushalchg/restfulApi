package controllers

import (
	"net/http"
	"practice/restfulApi/initializers"
	"practice/restfulApi/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// var body models.Post
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
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
		Title string
		Body  string
	}

	c.Bind(&body)

	//update that post
	updateResult := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	// updateQuery := `
	// UPDATE users
	// SET
	//     Title= COALESCE(NULLIF($1, ''), Title),
	//     Body= COALESCE(NULLIF($2, ''), Body)
	// WHERE id = $3`
	// updateResult := initializers.DB.Exec(updateQuery, body.Title, body.Body, id)

	//find the post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": updateResult.Error.Error(),
		})
	}

	// response with updated data

	c.IndentedJSON(http.StatusOK, gin.H{
		"data":  "successfully updated data",
		"value": post,
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
