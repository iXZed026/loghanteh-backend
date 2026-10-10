package repositories

import (
	"context"
	"time"

	"loghanteh-project/internal/dto"
	"loghanteh-project/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(
	db *gorm.DB,
) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) GetSessionForBooking(
	tx *gorm.DB,
	ctx context.Context,
	sessionID uint,
) (*models.EventSession, error) {

	var session models.EventSession

	err := tx.
		WithContext(ctx).
		Clauses(
			clause.Locking{
				Strength: "UPDATE",
			},
		).
		Where(
			"sessionid = ?",
			sessionID,
		).
		First(&session).
		Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *BookingRepository) Create(
	tx *gorm.DB,
	ctx context.Context,
	booking *models.Booking,
) error {

	return tx.
		WithContext(ctx).
		Create(booking).
		Error
}

// --------------------------------------------------
// Discount Code
// --------------------------------------------------

func (r *BookingRepository) GetDiscountCodeForBooking(
	tx *gorm.DB,
	ctx context.Context,
	discountCodeID uint,
) (*models.DiscountCode, error) {

	var discountCode models.DiscountCode

	err := tx.
		WithContext(ctx).
		Clauses(
			clause.Locking{
				Strength: "UPDATE",
			},
		).
		Where(
			"discountcodeid = ?",
			discountCodeID,
		).
		First(&discountCode).
		Error

	if err != nil {
		return nil, err
	}

	return &discountCode, nil
}

func (r *BookingRepository) IncrementDiscountUsage(
	tx *gorm.DB,
	ctx context.Context,
	discountCodeID uint,
) error {

	return tx.
		WithContext(ctx).
		Model(&models.DiscountCode{}).
		Where(
			"discountcodeid = ?",
			discountCodeID,
		).
		UpdateColumn(
			"usedcount",
			gorm.Expr("usedcount + ?", 1),
		).
		Error
}

// --------------------------------------------------
// Seat validation
// --------------------------------------------------

func (r *BookingRepository) ValidateSeatsForSession(
	tx *gorm.DB,
	ctx context.Context,
	sessionID uint,
	seatIDs []uint,
) (bool, error) {

	var count int64

	err := tx.
		WithContext(ctx).
		Table("seats AS s").
		Joins(`
			INNER JOIN eventsession AS es
				ON es.hallid = s.hallid
		`).
		Where(
			"es.sessionid = ?",
			sessionID,
		).
		Where(
			"s.seatid IN ?",
			seatIDs,
		).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count == int64(len(seatIDs)), nil
}

// --------------------------------------------------
// Booking seats
// --------------------------------------------------

func (r *BookingRepository) CreateBookingSeats(
	tx *gorm.DB,
	ctx context.Context,
	bookingID uint,
	sessionID uint,
	seatIDs []uint,
) error {

	bookingSeats := make(
		[]models.BookingSeat,
		0,
		len(seatIDs),
	)

	for _, seatID := range seatIDs {

		bookingSeats = append(
			bookingSeats,
			models.BookingSeat{
				BookingID: bookingID,
				SessionID: sessionID,
				SeatID:    seatID,
			},
		)
	}

	return tx.
		WithContext(ctx).
		Create(&bookingSeats).
		Error
}

// --------------------------------------------------
// Capacity
// --------------------------------------------------

func (r *BookingRepository) DecreaseSessionCapacity(
	tx *gorm.DB,
	ctx context.Context,
	sessionID uint,
	quantity int,
) error {

	result := tx.
		WithContext(ctx).
		Model(&models.EventSession{}).
		Where(
			"sessionid = ? AND capacity >= ?",
			sessionID,
			quantity,
		).
		UpdateColumn(
			"capacity",
			gorm.Expr(
				"capacity - ?",
				quantity,
			),
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrInvalidData
	}

	return nil
}

// --------------------------------------------------
// User bookings
// --------------------------------------------------

func (r *BookingRepository) GetBookingCount(
	ctx context.Context,
	bookingID uint,
	userID uint,
) (*dto.BookingCountResponse, error) {
	var count int64

	err := r.db.
		WithContext(ctx).
		Model(&models.Booking{}).
		Select("quantity").
		Where("bookingid = ? AND userid = ?", bookingID, userID).
		Scan(&count).
		Error
	if err != nil {
		return nil, err
	}

	return &dto.BookingCountResponse{Count: count}, nil
}

func (r *BookingRepository) GetNextBooking(
	ctx context.Context,
	userID uint,
	languageID uint,
) (*dto.NextBookingResponse, error) {
	var nextBooking dto.NextBookingResponse

	err := r.db.
		WithContext(ctx).
		Table("bookings AS b").
		Select(`
			b.bookingid AS "bookingId",
			COALESCE(ets.name, '') AS "name",
			b.tickettoken AS "ticketToken",
			es.startat AS "startAt"
		`).
		Joins(`
			INNER JOIN eventsession AS es
				ON es.sessionid = b.sessionid
		`).
		Joins(`
			INNER JOIN events AS e
				ON e.eventid = es.eventid
		`).
		Joins(`
			LEFT JOIN eventstranslations AS ets
				ON ets.eventid = e.eventid
				AND ets.languagesid = ?
		`, languageID).
		Where("b.userid = ?", userID).
		Where("es.startat > ?", time.Now()).
		Order("es.startat ASC").
		Limit(1).
		Scan(&nextBooking).
		Error
	if err != nil {
		return nil, err
	}
	if nextBooking.StartAt.IsZero() {
		return nil, gorm.ErrRecordNotFound
	}

	return &nextBooking, nil
}

func (r *BookingRepository) GetUserBookings(
	ctx context.Context,
	userID uint,
	languageID uint,
) ([]dto.UserBookingResponse, error) {

	var bookings []dto.UserBookingResponse

	err := r.db.
		WithContext(ctx).
		Table("bookings AS b").
		Select(`
			b.bookingid AS "bookingId",
			et.eventtypeid AS "eventTypeId",
			COALESCE(ett.name, et.eventtypename) AS "eventTypeName",
			e.eventid AS "eventId",
			es.sessionid AS "sessionId",
			COALESCE(ets.name, '') AS "name",
			COALESCE(ets.description, '') AS "description",
			es.startat AS "startAt",
			es.duration AS "duration",
			b.quantity AS "quantity",
			(
				b.totalprice - COALESCE(b.discountamount, 0)
			) AS "totalPrice",
			b.purchasedat AS "purchasedAt",
			b.discountcodeid AS "discountCodeId",
			COALESCE(b.discountamount, 0) AS "discountAmount",
			b.tickettoken AS "ticketToken",
			b.ticketisvalid AS "ticketIsValid"
		`).
		Joins(`
			INNER JOIN eventsession AS es
				ON es.sessionid = b.sessionid
		`).
		Joins(`
			INNER JOIN events AS e
				ON e.eventid = es.eventid
		`).
		Joins(`
			INNER JOIN eventtypes AS et
				ON et.eventtypeid = e.eventtypeid
		`).
		Joins(`
			LEFT JOIN eventstranslations AS ets
				ON ets.eventid = e.eventid
				AND ets.languagesid = ?
		`, languageID).
		Joins(`
			LEFT JOIN eventtypetranslations AS ett
				ON ett.eventtypeid = et.eventtypeid
				AND ett.languagesid = ?
		`, languageID).
		Where(
			"b.userid = ?",
			userID,
		).
		Order(
			"b.purchasedat DESC",
		).
		Scan(&bookings).
		Error

	if err != nil {
		return nil, err
	}

	return bookings, nil
}
