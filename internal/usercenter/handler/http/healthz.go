package http

import (
	"time"

	"github.com/gin-gonic/gin"
	ucv1 "github.com/ra1n6ow/opsx/pkg/api/usercenter/v1"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(c *gin.Context) {
	// 返回 JSON 响应
	c.JSON(200, &ucv1.HealthzResponse{
		Status:    ucv1.ServiceStatus_Healthy.Enum(),
		Timestamp: time.Now().Format(time.DateTime),
	})
}
