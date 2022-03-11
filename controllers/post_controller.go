package controllers

import (
	"blogo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /posts
// Get all posts
func Index(con *gin.Context) {
	var posts []models.Posts
	models.DB.Find(&posts)

	con.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetHome(con *gin.Context) {
	con.JSON(http.StatusOK, gin.H{"data": "hello world"})
}
