package routes

import (
	"log/slog"

	"loghanteh-project/internal/handlers"
	"loghanteh-project/internal/limiter"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRequestRoutes(
	api *gin.RouterGroup,
	logger *slog.Logger,
	customerRequestHandler *handlers.CustomerRequestHandler,
	rateLimiter limiter.RateLimiter,

) {

	customerRequests := api.Group("/customer-requests")
	{
		customerRequests.POST(
			"",
			customerRequestHandler.CreateCustomerRequest,
		)
	}
}
