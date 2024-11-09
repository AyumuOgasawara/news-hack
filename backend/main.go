package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのルーターを作成
	router := gin.Default()

	// /health エンドポイントを設定
	router.GET("/health", func(c *gin.Context) {
		// JSONレスポンスを返す
		c.JSON(http.StatusOK, gin.H{
			"Message": "hello, okutama",
			"status":  200,
		})
	})

	// サーバーを起動
	router.Run(":8080") // デフォルトでポート8080でリッスン
}
