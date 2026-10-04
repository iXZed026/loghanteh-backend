package routes

import (
	"log/slog"

	"loghanteh-project/internal/handlers"
	"loghanteh-project/internal/limiter"
	"loghanteh-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(
	api *gin.RouterGroup,
	logger *slog.Logger,
	authHandler *handlers.AuthHandler,
	rateLimiter limiter.RateLimiter,
) {
	auth := api.Group("/auth")
	{
		auth.POST(
			"/register",
			middleware.RateLimiter(rateLimiter, 5, logger),
			authHandler.Register,
		)

		auth.POST(
			"/login",
			middleware.RateLimiter(rateLimiter, 10, logger),
			authHandler.Login,
		)

		auth.POST(
			"/login/verify",
			middleware.RateLimiter(rateLimiter, 5, logger),
			authHandler.VerifyLogin,
		)

		auth.POST(
			"/register/verify",
			middleware.RateLimiter(rateLimiter, 5, logger),
			authHandler.VerifyRegister,
		)

		auth.POST(
			"/refresh",
			middleware.RateLimiter(rateLimiter, 30, logger),
			authHandler.Refresh,
		)

		auth.GET(
			"/logout",
			middleware.RateLimiter(rateLimiter, 20, logger),
			authHandler.Logout,
		)

		auth.POST(
			"/login/forgot-password",
			middleware.RateLimiter(rateLimiter, 3, logger),
			authHandler.ForgetPassword,
		)

		auth.POST(
			"/login/forgot-password/verify",
			middleware.RateLimiter(rateLimiter, 5, logger),
			authHandler.VerifyForgetPassword,
		)

		auth.POST(
			"/login/forgot-password/verify/new-password",
			middleware.RateLimiter(rateLimiter, 5, logger),
			authHandler.NewPassword,
		)
	}
}
