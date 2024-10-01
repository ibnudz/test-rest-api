package controllers

import (
	"test/backend-api/models"

	"github.com/gin-gonic/gin"
)

// get all posts
func FindPosts(c *gin.Context) {

	// get data from database using model
	var posts []models.Post
	models.DB.Find(&posts)

	// return json
	c.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Posts",
		"data":    posts,
	})
}
