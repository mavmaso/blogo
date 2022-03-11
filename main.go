package main

import (
	"blogo/controllers"
	"blogo/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/", controllers.GetHome)
	router.GET("/posts", controllers.Index)

	router.Run()
}
