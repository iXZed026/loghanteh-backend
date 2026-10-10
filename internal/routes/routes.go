package routes

import (
	"log/slog"

	"loghanteh-project/internal/auth"
	"loghanteh-project/internal/handlers"
	"loghanteh-project/internal/limiter"
	"loghanteh-project/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	logger *slog.Logger,
	authHandler *handlers.AuthHandler,
	profileHandler *handlers.ProfileHandler,
	ticketHandler *handlers.TicketHandler,
	customerRequestHandler *handlers.CustomerRequestHandler,
	discountCodeHandler *handlers.DiscountCodeHandler,
	jwtService *auth.JWTService,
	rateLimiter limiter.RateLimiter,
) {

	r.Use(
		middleware.RecoveryMiddleware(logger),
	)

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://loghanteh-frontend.vercel.app",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"Accept-Language",
		},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		SetupAuthRoutes(
			api,
			logger,
			authHandler,
			rateLimiter,
		)

		SetupProfileRoutes(
			api,
			logger,
			profileHandler,
			jwtService,
			rateLimiter,
		)

		SetupTicketsRoutes(
			api,
			logger,
			ticketHandler,
			jwtService,
			rateLimiter,
		)

		SetupCustomerRequestRoutes(
			api,
			logger,
			customerRequestHandler,
			rateLimiter,
		)

		SetupDiscountCodeRoutes(
			api,
			logger,
			discountCodeHandler,
			jwtService,
			rateLimiter,
		)
	}
}
