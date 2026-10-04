package routes

import (
	"log/slog"

	"loghanteh-project/internal/auth"
	"loghanteh-project/internal/handlers"
	"loghanteh-project/internal/limiter"
	"loghanteh-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTicketsRoutes(
	api *gin.RouterGroup,
	logger *slog.Logger,
	ticketHandler *handlers.TicketHandler,
	jwtService *auth.JWTService,
	rateLimiter limiter.RateLimiter,
) {
	tickets := api.Group("/tickets")
	{
		// --------------------------------------------------
		// Museum Tour
		// --------------------------------------------------

		tickets.GET(
			"/museum-tour",
			middleware.RateLimiter(
				rateLimiter,
				60,
				logger,
			),
			ticketHandler.GetMuseumTour,
		)

		// --------------------------------------------------
		// Events
		// --------------------------------------------------

		tickets.GET(
			"/events",
			middleware.RateLimiter(
				rateLimiter,
				60,
				logger,
			),
			ticketHandler.GetEvent,
		)

		tickets.GET(
			"/events/sessions/:sessionId",
			middleware.RateLimiter(
				rateLimiter,
				60,
				logger,
			),
			ticketHandler.GetEventSessionDetail,
		)

		// --------------------------------------------------
		// Courses
		// --------------------------------------------------

		tickets.GET(
			"/courses",
			middleware.RateLimiter(
				rateLimiter,
				60,
				logger,
			),
			ticketHandler.GetCourses,
		)

		tickets.GET(
			"/courses/sessions/:sessionId",
			middleware.RateLimiter(
				rateLimiter,
				60,
				logger,
			),
			ticketHandler.GetCourseSessionDetail,
		)

		// --------------------------------------------------
		// Cinema
		// --------------------------------------------------

		tickets.GET(
			"/cinema",
			middleware.RateLimiter(
				rateLimiter,
				30,
				logger,
			),
			ticketHandler.GetCinema,
		)

		// --------------------------------------------------
		// Theater
		// --------------------------------------------------

		tickets.GET(
			"/theater",
			middleware.RateLimiter(
				rateLimiter,
				30,
				logger,
			),
			ticketHandler.GetTheater,
		)

		// --------------------------------------------------
		// Bookings
		// --------------------------------------------------

		tickets.GET(
			"/bookings",
			middleware.AuthMiddleware(jwtService, logger),
			middleware.RateLimiter(rateLimiter, 60, logger),
			ticketHandler.GetUserBookings,
		)

		tickets.POST(
			"/bookings",
			middleware.AuthMiddleware(jwtService, logger),
			middleware.RateLimiter(rateLimiter, 10, logger),
			ticketHandler.CreateBooking,
		)

		tickets.GET(
			"/booking-seat/:sessionId",
			middleware.RateLimiter(rateLimiter, 60, logger),
			ticketHandler.GetSessionSeats,
		)

		tickets.GET(
			"/bookings/:bookingId",
			middleware.AuthMiddleware(jwtService, logger),
			middleware.RateLimiter(rateLimiter, 60, logger),
			ticketHandler.GetUserReservedSession,
		)

		tickets.POST(
			"/booking-seat",
			middleware.AuthMiddleware(
				jwtService,
				logger,
			),
			ticketHandler.CreateSeatBooking,
		)
	}
}
