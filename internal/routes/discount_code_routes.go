package routes

import (
	"log/slog"

	"loghanteh-project/internal/auth"
	"loghanteh-project/internal/handlers"
	"loghanteh-project/internal/limiter"
	"loghanteh-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupDiscountCodeRoutes(
	api *gin.RouterGroup,
	logger *slog.Logger,
	discountCodeHandler *handlers.DiscountCodeHandler,
	jwtService *auth.JWTService,
	rateLimiter limiter.RateLimiter,
) {
	discountCodes := api.Group("/discount-codes")

	// Temporary public endpoint for testing.
	discountCodes.GET(
		"",
		middleware.RateLimiter(rateLimiter, 60, logger),
		discountCodeHandler.GetAll,
	)

	// Existing endpoint: authentication removed temporarily for testing.
	discountCodes.GET(
		"/:code",
		middleware.RateLimiter(rateLimiter, 60, logger),
		discountCodeHandler.GetByCode,
	)
}
