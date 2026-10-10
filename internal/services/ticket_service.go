package services

import (
	"context"
	"time"

	"loghanteh-project/internal/dto"
	"loghanteh-project/internal/repositories"
	"loghanteh-project/internal/reservation"
)

const (
	MuseumTourEventTypeID uint = 1
	EventEventTypeID      uint = 2
	CourseEventTypeID     uint = 3
	CinemaEventTypeID     uint = 4
	TheaterEventTypeID    uint = 5
)

type TicketService struct {
	ticketRepository     *repositories.TicketRepository
	languageRepository   *repositories.LanguageRepository
	bookingRepository    *repositories.BookingRepository
	reservationService   *reservation.Service
	seatResourceProvider reservation.ResourceProvider
}

func NewTicketService(
	ticketRepository *repositories.TicketRepository,
	languageRepository *repositories.LanguageRepository,
	bookingRepository *repositories.BookingRepository,
	reservationService *reservation.Service,
	seatResourceProvider reservation.ResourceProvider,
) *TicketService {
	return &TicketService{
		ticketRepository:     ticketRepository,
		languageRepository:   languageRepository,
		bookingRepository:    bookingRepository,
		reservationService:   reservationService,
		seatResourceProvider: seatResourceProvider,
	}
}

// --------------------------------------------------
// Museum Tour
// --------------------------------------------------

func (s *TicketService) GetMuseumTourSessions(
	ctx context.Context,
	date time.Time,
	languageCode string,
) ([]dto.MuseumTourSessionResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(
		ctx,
		languageCode,
	)
	if err != nil {
		return nil, err
	}

	return s.ticketRepository.GetMuseumTourSessions(
		ctx,
		date,
		MuseumTourEventTypeID,
		languageID,
	)
}

// --------------------------------------------------
// Events
// --------------------------------------------------

func (s *TicketService) GetEventSessions(
	ctx context.Context,
	date time.Time,
	languageCode string,
) ([]dto.EventSessionListResponse, error) {
	return s.getSessionsByType(
		ctx,
		date,
		languageCode,
		EventEventTypeID,
		s.ticketRepository.GetEventSessions,
	)
}

// --------------------------------------------------
// Courses
// --------------------------------------------------

func (s *TicketService) GetCourseSessions(
	ctx context.Context,
	date time.Time,
	languageCode string,
) ([]dto.EventSessionListResponse, error) {
	return s.getSessionsByType(
		ctx,
		date,
		languageCode,
		CourseEventTypeID,
		s.ticketRepository.GetCourseSessions,
	)
}

// --------------------------------------------------
// Cinema
// --------------------------------------------------

func (s *TicketService) GetCinemaSessions(
	ctx context.Context,
	date time.Time,
	languageCode string,
) ([]dto.CinemaSessionListResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(
		ctx,
		languageCode,
	)
	if err != nil {
		return nil, err
	}

	return s.ticketRepository.GetCinemaSessions(
		ctx,
		date,
		CinemaEventTypeID,
		languageID,
	)
}

// --------------------------------------------------
// Theater
// --------------------------------------------------

func (s *TicketService) GetTheaterSessions(
	ctx context.Context,
	date time.Time,
	languageCode string,
) ([]dto.TheaterSessionListResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(
		ctx,
		languageCode,
	)
	if err != nil {
		return nil, err
	}

	return s.ticketRepository.GetTheaterSessions(
		ctx,
		date,
		TheaterEventTypeID,
		languageID,
	)
}

// --------------------------------------------------
// Generic session list helper
// --------------------------------------------------

type sessionListFunc func(
	context.Context,
	time.Time,
	uint,
	uint,
) ([]dto.EventSessionListResponse, error)

func (s *TicketService) getSessionsByType(
	ctx context.Context,
	date time.Time,
	languageCode string,
	eventTypeID uint,
	getSessions sessionListFunc,
) ([]dto.EventSessionListResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(
		ctx,
		languageCode,
	)
	if err != nil {
		return nil, err
	}

	return getSessions(
		ctx,
		date,
		eventTypeID,
		languageID,
	)
}

// --------------------------------------------------
// Event Detail
// --------------------------------------------------

func (s *TicketService) GetEventSessionDetail(
	ctx context.Context,
	sessionID uint,
	languageCode string,
) (*dto.EventSessionDetailResponse, error) {
	return s.getEventDetail(
		ctx,
		sessionID,
		languageCode,
		EventEventTypeID,
		s.ticketRepository.GetEventSessionDetail,
	)
}

// --------------------------------------------------
// Course Detail
// --------------------------------------------------

func (s *TicketService) GetCourseSessionDetail(
	ctx context.Context,
	sessionID uint,
	languageCode string,
) (*dto.EventSessionDetailResponse, error) {
	return s.getEventDetail(
		ctx,
		sessionID,
		languageCode,
		CourseEventTypeID,
		s.ticketRepository.GetCourseSessionDetail,
	)
}

// --------------------------------------------------
// Generic Event Detail Helper
// --------------------------------------------------

type sessionDetailFunc func(
	context.Context,
	uint,
	uint,
	uint,
) (*dto.EventSessionDetailResponse, error)

func (s *TicketService) getEventDetail(
	ctx context.Context,
	sessionID uint,
	languageCode string,
	eventTypeID uint,
	getDetail sessionDetailFunc,
) (*dto.EventSessionDetailResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(
		ctx,
		languageCode,
	)
	if err != nil {
		return nil, err
	}

	return getDetail(
		ctx,
		sessionID,
		eventTypeID,
		languageID,
	)
}

// --------------------------------------------------
// Booking
// --------------------------------------------------

func (s *TicketService) CreateBooking(
	ctx context.Context,
	userID uint,
	req dto.CreateBookingRequest,
) (*dto.CreateBookingResponse, error) {
	booking, err := s.reservationService.CreateBooking(
		ctx,
		userID,
		req.SessionID,
		req.Quantity,
		nil,
		nil,
		req.DiscountCodeID,
	)
	if err != nil {
		return nil, err
	}

	return &dto.CreateBookingResponse{
		BookingID:   booking.BookingID,
		SessionID:   booking.SessionID,
		Quantity:    booking.Quantity,
		TotalPrice:  booking.TotalPrice,
		PurchasedAt: booking.PurchasedAt,
	}, nil
}

// --------------------------------------------------
// User Bookings
// --------------------------------------------------

func (s *TicketService) GetUserBookings(
	ctx context.Context,
	userID uint,
	languageCode string,
) ([]dto.UserBookingResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(
		ctx,
		languageCode,
	)
	if err != nil {
		return nil, err
	}

	return s.bookingRepository.GetUserBookings(
		ctx,
		userID,
		languageID,
	)
}

func (s *TicketService) GetBookingCount(
	ctx context.Context,
	bookingID uint,
	userID uint,
) (*dto.BookingCountResponse, error) {
	return s.bookingRepository.GetBookingCount(ctx, bookingID, userID)
}

func (s *TicketService) GetNextBooking(
	ctx context.Context,
	userID uint,
	languageCode string,
) (*dto.NextBookingResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(ctx, languageCode)
	if err != nil {
		return nil, err
	}

	return s.bookingRepository.GetNextBooking(ctx, userID, languageID)
}

// --------------------------------------------------
// Seats
// --------------------------------------------------

func (s *TicketService) GetSessionSeats(
	ctx context.Context,
	sessionID uint,
) ([]dto.SessionSeatResponse, error) {
	return s.ticketRepository.GetSessionSeats(
		ctx,
		sessionID,
	)
}

func (s *TicketService) CreateSeatBooking(
	ctx context.Context,
	userID uint,
	req dto.CreateSeatBookingRequest,
) (*dto.CreateBookingResponse, error) {
	quantity := len(req.SeatIDs)

	booking, err := s.reservationService.CreateBooking(
		ctx,
		userID,
		req.SessionID,
		quantity,
		s.seatResourceProvider,
		req.SeatIDs,
		req.DiscountCodeID,
	)
	if err != nil {
		return nil, err
	}

	return &dto.CreateBookingResponse{
		BookingID:   booking.BookingID,
		SessionID:   booking.SessionID,
		Quantity:    booking.Quantity,
		TotalPrice:  booking.TotalPrice,
		PurchasedAt: booking.PurchasedAt,
	}, nil
}

// --------------------------------------------------
// User Reserved Session
// --------------------------------------------------

func (s *TicketService) GetUserReservedSession(
	ctx context.Context,
	userID uint,
	bookingID uint,
	languageCode string,
) (*dto.UserReservedSessionResponse, error) {
	languageID, err := s.languageRepository.GetIDByCode(ctx, languageCode)
	if err != nil {
		return nil, err
	}

	return s.ticketRepository.GetUserReservedSession(
		ctx,
		userID,
		bookingID,
		languageID,
	)
}
