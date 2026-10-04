package services

import (
	"context"
	"errors"

	"loghanteh-project/internal/models"
	"loghanteh-project/internal/reservation"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReservationService struct {
	db *gorm.DB
}

func NewReservationService(
	db *gorm.DB,
) *ReservationService {
	return &ReservationService{
		db: db,
	}
}

func (s *ReservationService) Reserve(
	ctx context.Context,
	bookingID uint,
	sessionID uint,
	resourceIDs []uint,
	provider reservation.ResourceProvider,
) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var session models.EventSession

		err := tx.
			WithContext(ctx).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&session, sessionID).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return gorm.ErrRecordNotFound
			}

			return err
		}

		if err := provider.ValidateAndLock(
			tx,
			ctx,
			&session,
			resourceIDs,
		); err != nil {
			return err
		}

		return provider.Reserve(
			tx,
			ctx,
			bookingID,
			sessionID,
			resourceIDs,
		)
	})
}
