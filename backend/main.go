package main

import (
	"backend/app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"path/filepath"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api") // APIのルートを/apiに変更
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}

	router.NoRoute(func(c *gin.Context) { // 未知のルートに対するハンドラ
		dir, file := path.Split(c.Request.RequestURI)
		extension := filepath.Ext(file)

		// リクエストがファイルであることを確認する (".html", ".css", etc.)
		// リクエストがディレクトリの場合は "index.html"
		if file == "" || extension == "" {
			path := filepath.Join(dir, "index.html")
			c.File("../frontend/out" + path)
		} else {
			c.File("../frontend/out" + c.Request.RequestURI)
		}
	})

	router.Run(":8080")
}
