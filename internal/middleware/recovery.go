package middleware

import (
	"log/slog"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {

				value, _ := c.Get("request_id")

				logger.Error(
					"panic recovered",
					"request_id", value,
					"error", err, //Panic error
					"stack", string(debug.Stack()), //main.go 50 ...
				)
				c.Abort()
				c.JSON(
					500,
					gin.H{
						"success": false,
						"message": "internal server error",
					},
				)
				return
			}
		}()
		c.Next()

	}
}
