package middleware

import (
	"log/slog"

	appErr "loghanteh-project/internal/errors"
	"loghanteh-project/internal/limiter"
	"loghanteh-project/internal/response"

	"github.com/gin-gonic/gin"
)

func RateLimiter(
	rateLimiter limiter.RateLimiter,
	max int,
	logger *slog.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		route := c.FullPath()
		method := c.Request.Method

		key := method + ":" + route + ":" + ip

		allowed, err := rateLimiter.Allow(
			c.Request.Context(),
			key,
			max,
		)

		if err != nil {
			logger.Error(
				"rate limiter error",
				"error", err,
				"ip", ip,
				"method", method,
				"route", route,
			)

			response.HandleError(
				c,
				appErr.ErrInternalServer,
				logger,
			)

			c.Abort()

			return
		}

		if !allowed {
			logger.Warn(
				"rate limit exceeded",
				"ip", ip,
				"method", method,
				"route", route,
				"max", max,
			)

			response.HandleError(
				c,
				appErr.ErrTooManyRequests,
				logger,
			)

			c.Abort()

			return
		}

		c.Next()
	}
}
