package reservation

import (
	"context"

	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"

	"gorm.io/gorm"
)

type SeatResourceProvider struct{}

func NewSeatResourceProvider() *SeatResourceProvider {
	return &SeatResourceProvider{}
}

func (p *SeatResourceProvider) ValidateAndLock(
	tx *gorm.DB,
	ctx context.Context,
	session *models.EventSession,
	resourceIDs []uint,
) error {
	if len(resourceIDs) == 0 {
		return appErrors.ErrNoResources
	}

	if hasDuplicateIDs(resourceIDs) {
		return appErrors.ErrDuplicateResources
	}

	var seats []models.Seat

	err := tx.
		WithContext(ctx).
		Where("seatid IN ?", resourceIDs).
		Where("hallid = ?", session.HallID).
		Find(&seats).
		Error
	if err != nil {
		return err
	}

	if len(seats) != len(resourceIDs) {
		return appErrors.ErrResourceNotFound
	}

	var bookedSeatIDs []uint

	err = tx.
		WithContext(ctx).
		Model(&models.BookingSeat{}).
		Where("sessionid = ?", session.SessionID).
		Where("seatid IN ?", resourceIDs).
		Pluck("seatid", &bookedSeatIDs).
		Error
	if err != nil {
		return err
	}

	if len(bookedSeatIDs) > 0 {
		return appErrors.ErrResourceAlreadyReserved
	}

	return nil
}

func (p *SeatResourceProvider) Reserve(
	tx *gorm.DB,
	ctx context.Context,
	bookingID uint,
	sessionID uint,
	resourceIDs []uint,
) error {
	bookingSeats := make([]models.BookingSeat, 0, len(resourceIDs))

	for _, seatID := range resourceIDs {
		bookingSeats = append(bookingSeats, models.BookingSeat{
			BookingID: bookingID,
			SessionID: sessionID,
			SeatID:    seatID,
		})
	}

	return tx.
		WithContext(ctx).
		Create(&bookingSeats).
		Error
}

func hasDuplicateIDs(ids []uint) bool {
	seen := make(map[uint]struct{}, len(ids))

	for _, id := range ids {
		if _, exists := seen[id]; exists {
			return true
		}

		seen[id] = struct{}{}
	}

	return false
}
