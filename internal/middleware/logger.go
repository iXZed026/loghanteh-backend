package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(
	logger *slog.Logger,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		duration := time.Since(start)

		logger.Info(
			"HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration,
			"ip", c.ClientIP(),
		)

	}
}
