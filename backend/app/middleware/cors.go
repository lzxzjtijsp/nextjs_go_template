package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func Cors() gin.HandlerFunc {
	// CDNのパスを環境変数から取得
	bucketUrl := os.Getenv("BUCKET_URL")
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", bucketUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
