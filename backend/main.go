package main

import (
	"backend/app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	port := os.Getenv("APP_PORT")
	bucketUrl := os.Getenv("BUCKET_URL")

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		var url string
		if c.Request.RequestURI == "/" {
			url = bucketUrl + "/index.html"
		} else {
			url = bucketUrl + strings.TrimSuffix(c.Request.RequestURI, "/") + ".html"
		}

		resp, err := http.Get(url)
		if err != nil || resp.StatusCode == 404 {
			// エラーが発生した場合、または404のステータスコードが返された場合は404.htmlを返す
			url = bucketUrl + "/404.html"
			resp, err = http.Get(url)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "failed to fetch the file",
					"error":   err.Error(),
				})
				return
			}
		}
		defer resp.Body.Close()

		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, map[string]string{})
	})

	router.Run(":" + port)
}
