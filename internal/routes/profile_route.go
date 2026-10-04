package routes

import (
	"log/slog"

	"loghanteh-project/internal/auth"
	"loghanteh-project/internal/handlers"
	"loghanteh-project/internal/limiter"
	"loghanteh-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(
	api *gin.RouterGroup,
	logger *slog.Logger,
	profileHandler *handlers.ProfileHandler,
	jwtService *auth.JWTService,
	rateLimiter limiter.RateLimiter,
) {
	profile := api.Group(
		"/profile",
		middleware.AuthMiddleware(
			jwtService,
			logger,
		),
	)
	{
		profile.GET(
			"",
			middleware.RateLimiter(
				rateLimiter,
				60,
				logger,
			),
			profileHandler.Profile,
		)

		profile.PUT(
			"/edit-profile",
			middleware.RateLimiter(
				rateLimiter,
				3,
				logger,
			),
			profileHandler.EditProfile,
		)
	}
}
