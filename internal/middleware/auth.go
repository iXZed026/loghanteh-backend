package middleware

import (
	"log/slog"
	"strings"

	appAuth "loghanteh-project/internal/auth"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(
	jwtService *appAuth.JWTService,
	logger *slog.Logger,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			response.HandleError(
				c,
				appErrors.ErrUnauthorized,
				logger,
			)

			c.Abort()

			return
		}

		parts := strings.SplitN(
			authHeader,
			" ",
			2,
		)

		if len(parts) != 2 ||
			parts[0] != "Bearer" ||
			parts[1] == "" {

			response.HandleError(
				c,
				appErrors.ErrUnauthorized,
				logger,
			)

			c.Abort()

			return
		}

		token, err := jwtService.ValidateToken(parts[1])

		if err != nil {

			response.HandleError(
				c,
				appErrors.ErrUnauthorized,
				logger,
			)

			c.Abort()

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {

			response.HandleError(
				c,
				appErrors.ErrUnauthorized,
				logger,
			)

			c.Abort()

			return
		}

		userID, ok := claims["user_id"].(string)

		if !ok || userID == "" {

			response.HandleError(
				c,
				appErrors.ErrUnauthorized,
				logger,
			)

			c.Abort()

			return
		}

		c.Set(
			"user_id",
			userID,
		)

		c.Next()
	}
}
