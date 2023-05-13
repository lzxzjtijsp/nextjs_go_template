package main

import (
	"backend/app/middleware"
	"github.com/gin-gonic/gin"
)

import "net/http"

func main() {
	engine := gin.Default()
	engine.Use(middleware.Cors())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":8080")
}
