package system

import "github.com/gin-gonic/gin"

type SystemHandler struct{}

// HealthCheck godoc
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/health [get]
func (h *SystemHandler) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Status":  "ok",
		"Message": "hello okutama",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
