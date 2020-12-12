package main

import (
	"gin_demo/controllers"
	"gin_demo/models"
	"github.com/gin-gonic/gin"
)

func registerUrl(r *gin.Engine) {
	r.GET("/books", controllers.ListBook)
	r.POST("/books", controllers.CreateBook)
}

func main() {
	r := gin.Default()
	models.Init()
	registerUrl(r)

	r.Run(":8081")
}
