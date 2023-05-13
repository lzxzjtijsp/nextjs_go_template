package middleware

import (
	"backend/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func Cors() gin.HandlerFunc {
	env.Load()
	// CDNのパスを環境変数から取得
	staticFilePath := os.Getenv("STATIC_FILE_PATH")
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", staticFilePath},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
