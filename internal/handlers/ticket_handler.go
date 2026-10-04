package handlers

import (
	"log/slog"
	"strconv"
	"time"

	"loghanteh-project/internal/dto"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/response"
	"loghanteh-project/internal/services"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ticketService *services.TicketService
	logger        *slog.Logger
}

func NewTicketHandler(
	ticketService *services.TicketService,
	logger *slog.Logger,
) *TicketHandler {
	return &TicketHandler{
		ticketService: ticketService,
		logger:        logger,
	}
}

// --------------------------------------------------
// Language Helper
// --------------------------------------------------

func getLanguage(c *gin.Context) string {
	language := c.GetHeader("Accept-Language")

	if len(language) >= 2 {
		language = language[:2]
	}

	if language == "" {
		language = "en"
	}

	return language
}

// --------------------------------------------------
// Museum Tour
// --------------------------------------------------

func (h *TicketHandler) GetMuseumTour(
	c *gin.Context,
) {
	dateString := c.Query("date")

	if dateString == "" {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	date, err := time.Parse(
		"2006-01-02",
		dateString,
	)

	if err != nil {
		h.logger.Error(
			"failed to parse museum tour date",
			"date", dateString,
			"error", err,
		)

		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	sessions, err := h.ticketService.GetMuseumTourSessions(
		c.Request.Context(),
		date,
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get museum tour sessions",
			"date", dateString,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"museum tour sessions fetched successfully",
		sessions,
	)
}

// --------------------------------------------------
// Events
// --------------------------------------------------

func (h *TicketHandler) GetEvent(
	c *gin.Context,
) {
	dateString := c.Query("date")

	if dateString == "" {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	date, err := time.Parse(
		"2006-01-02",
		dateString,
	)

	if err != nil {
		h.logger.Error(
			"failed to parse events date",
			"date", dateString,
			"error", err,
		)

		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	sessions, err := h.ticketService.GetEventSessions(
		c.Request.Context(),
		date,
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get event sessions",
			"date", dateString,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"event sessions fetched successfully",
		sessions,
	)
}

func (h *TicketHandler) GetEventSessionDetail(
	c *gin.Context,
) {
	sessionID, err := strconv.ParseUint(
		c.Param("sessionId"),
		10,
		32,
	)

	if err != nil {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	session, err := h.ticketService.GetEventSessionDetail(
		c.Request.Context(),
		uint(sessionID),
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get event session detail",
			"sessionId", sessionID,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"event session fetched successfully",
		session,
	)
}

// --------------------------------------------------
// Courses
// --------------------------------------------------

func (h *TicketHandler) GetCourses(
	c *gin.Context,
) {
	dateString := c.Query("date")

	if dateString == "" {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	date, err := time.Parse(
		"2006-01-02",
		dateString,
	)

	if err != nil {
		h.logger.Error(
			"failed to parse courses date",
			"date", dateString,
			"error", err,
		)

		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	sessions, err := h.ticketService.GetCourseSessions(
		c.Request.Context(),
		date,
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get course sessions",
			"date", dateString,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"course sessions fetched successfully",
		sessions,
	)
}

func (h *TicketHandler) GetCourseSessionDetail(
	c *gin.Context,
) {
	sessionID, err := strconv.ParseUint(
		c.Param("sessionId"),
		10,
		32,
	)

	if err != nil {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	session, err := h.ticketService.GetCourseSessionDetail(
		c.Request.Context(),
		uint(sessionID),
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get course session detail",
			"sessionId", sessionID,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"course session fetched successfully",
		session,
	)
}

// --------------------------------------------------
// Cinema
// --------------------------------------------------

func (h *TicketHandler) GetCinema(
	c *gin.Context,
) {
	dateString := c.Query("date")

	if dateString == "" {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	date, err := time.Parse(
		"2006-01-02",
		dateString,
	)

	if err != nil {
		h.logger.Error(
			"failed to parse cinema date",
			"date", dateString,
			"error", err,
		)

		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	sessions, err := h.ticketService.GetCinemaSessions(
		c.Request.Context(),
		date,
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get cinema sessions",
			"date", dateString,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"cinema sessions fetched successfully",
		sessions,
	)
}

// --------------------------------------------------
// Theater
// --------------------------------------------------

func (h *TicketHandler) GetTheater(
	c *gin.Context,
) {
	dateString := c.Query("date")

	if dateString == "" {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	date, err := time.Parse(
		"2006-01-02",
		dateString,
	)

	if err != nil {
		h.logger.Error(
			"failed to parse theater date",
			"date", dateString,
			"error", err,
		)

		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	sessions, err := h.ticketService.GetTheaterSessions(
		c.Request.Context(),
		date,
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get theater sessions",
			"date", dateString,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"theater sessions fetched successfully",
		sessions,
	)
}

// --------------------------------------------------
// Booking
// --------------------------------------------------

func (h *TicketHandler) CreateBooking(
	c *gin.Context,
) {
	var req dto.CreateBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(
			c,
			appErrors.ErrInvalidInput,
			h.logger,
		)
		return
	}

	userIDValue, exists := c.Get("user_id")

	if !exists {
		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)
		return
	}

	userIDString, ok := userIDValue.(string)

	if !ok {
		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)
		return
	}

	userID64, err := strconv.ParseUint(
		userIDString,
		10,
		64,
	)

	if err != nil {
		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)
		return
	}

	booking, err := h.ticketService.CreateBooking(
		c.Request.Context(),
		uint(userID64),
		req,
	)

	if err != nil {
		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"booking_created_successfully",
		booking,
	)
}

// --------------------------------------------------
// User Bookings
// --------------------------------------------------

func (h *TicketHandler) GetUserBookings(
	c *gin.Context,
) {
	userIDValue, exists := c.Get("user_id")

	if !exists {
		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)
		return
	}

	userIDString, ok := userIDValue.(string)

	if !ok {
		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)
		return
	}

	userID64, err := strconv.ParseUint(
		userIDString,
		10,
		64,
	)

	if err != nil {
		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)
		return
	}

	language := getLanguage(c)

	bookings, err := h.ticketService.GetUserBookings(
		c.Request.Context(),
		uint(userID64),
		language,
	)

	if err != nil {
		h.logger.Error(
			"failed to get user bookings",
			"userID", userID64,
			"language", language,
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)
		return
	}

	response.Success(
		c,
		"user_bookings_fetched_successfully",
		bookings,
	)
}

// --------------------------------------------------
// Seats
// --------------------------------------------------

func (h *TicketHandler) GetSessionSeats(c *gin.Context) {
	sessionID, err := strconv.ParseUint(c.Param("sessionId"), 10, 64)
	if err != nil || sessionID == 0 {
		response.HandleError(c, appErrors.ErrInvalidInput, h.logger)
		return
	}

	seats, err := h.ticketService.GetSessionSeats(
		c.Request.Context(),
		uint(sessionID),
	)
	if err != nil {
		response.HandleError(c, err, h.logger)
		return
	}

	response.Success(
		c,
		"session_seats_fetched_successfully",
		seats,
	)
}

func (h *TicketHandler) CreateSeatBooking(c *gin.Context) {
	var req dto.CreateSeatBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, appErrors.ErrInvalidInput, h.logger)
		return
	}

	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil || userID == 0 {
		response.HandleError(c, appErrors.ErrUnauthorized, h.logger)
		return
	}

	booking, err := h.ticketService.CreateSeatBooking(
		c.Request.Context(),
		uint(userID),
		req,
	)
	if err != nil {
		response.HandleError(c, err, h.logger)
		return
	}

	response.Success(
		c,
		"booking_created_successfully",
		booking,
	)
}

// --------------------------------------------------
// User Reserved Session
// --------------------------------------------------

func (h *TicketHandler) GetUserReservedSession(c *gin.Context) {
	bookingID, err := strconv.ParseUint(
		c.Param("bookingId"),
		10,
		64,
	)

	if err != nil || bookingID == 0 {
		response.HandleError(c, appErrors.ErrInvalidInput, h.logger)
		return
	}

	userID, err := strconv.ParseUint(
		c.GetString("user_id"),
		10,
		64,
	)

	if err != nil || userID == 0 {
		response.HandleError(c, appErrors.ErrUnauthorized, h.logger)
		return
	}

	language := getLanguage(c)

	session, err := h.ticketService.GetUserReservedSession(
		c.Request.Context(),
		uint(userID),
		uint(bookingID),
		language,
	)
	if err != nil {
		h.logger.Error(
			"failed to get user reserved session",
			"userID", userID,
			"bookingId", bookingID,
			"language", language,
			"error", err,
		)

		response.HandleError(c, err, h.logger)
		return
	}

	response.Success(
		c,
		"user_reserved_session_fetched_successfully",
		session,
	)
}
