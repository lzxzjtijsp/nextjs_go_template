package main

import (
	"backend/app/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	bucketUrl := os.Getenv("BUCKET_URL")

	api := router.Group("/api")
	api.Use(apiNotFoundHandler())
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

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to run server on port %s: %v", port, err)
	}
}

// apiNotFoundHandler returns a custom middleware for handling not found routes within the /api group.
func apiNotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				if e.Type == gin.ErrorTypePrivate {
					// Handle not found routes within /api group
					if c.Writer.Status() == http.StatusNotFound {
						c.JSON(http.StatusNotFound, gin.H{
							"message": "Not Found",
						})
					}
				}
			}
		}
	}
}
