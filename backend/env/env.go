package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// 設定ファイル読み込み環境変数
const targetEnvName = "local"

func Load() {
	// 環境未設定の場合はローカルを設定
	if "" == os.Getenv(targetEnvName) {
		_ = os.Setenv(targetEnvName, "local")
	}
	err := godotenv.Load(fmt.Sprintf("env/%s.env", os.Getenv(targetEnvName)))
	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv(targetEnvName))
	}
}
