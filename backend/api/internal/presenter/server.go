package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/news-hack/backend/api/internal/controller/system"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// health check
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
