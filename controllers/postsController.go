package controllers

import (
	"errors"
	"net/http"
	"test/backend-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// type validation post input
type ValidatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// type error message
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// function get error message
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

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

// store a post
func StorePost(c *gin.Context) {
	//validate input
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	//create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	//return response json
	c.JSON(201, gin.H{
		"success": true,
		"message": "Post Created Successfully",
		"data":    post,
	})
}
