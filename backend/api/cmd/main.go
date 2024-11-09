package main

import (
	"context"

	"github.com/news-hack/backend/api/internal/presenter"
)

// @title ユーザー管理サービスAPI
// @version v0.1.0
// @description
// @host localhost:8080
func main() {
	srv := presenter.NewServer()
	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}
