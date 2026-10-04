package reservation

import (
	"context"
	"errors"

	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"
	"loghanteh-project/internal/repositories"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Service struct {
	db                *gorm.DB
	bookingRepository *repositories.BookingRepository
}

func NewService(
	db *gorm.DB,
	bookingRepository *repositories.BookingRepository,
) *Service {
	return &Service{
		db:                db,
		bookingRepository: bookingRepository,
	}
}

func (s *Service) CreateBooking(
	ctx context.Context,
	userID uint,
	sessionID uint,
	quantity int,
	provider ResourceProvider,
	resourceIDs []uint,
) (*models.Booking, error) {
	if quantity < 1 {
		return nil, appErrors.ErrInvalidInput
	}

	if provider != nil && len(resourceIDs) != quantity {
		return nil, appErrors.ErrInvalidInput
	}

	var booking models.Booking

	err := s.db.Transaction(func(tx *gorm.DB) error {
		session, err := s.bookingRepository.GetSessionForBooking(
			tx,
			ctx,
			sessionID,
		)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return appErrors.ErrNotFound
			}

			return err
		}

		if quantity > session.Capacity {
			return appErrors.ErrInsufficientCapacity
		}

		if provider != nil {
			if err := provider.ValidateAndLock(
				tx,
				ctx,
				session,
				resourceIDs,
			); err != nil {
				return err
			}
		}

		var totalPrice decimal.Decimal

		if session.Price != nil {
			totalPrice = session.Price.Mul(
				decimal.NewFromInt(
					int64(quantity),
				),
			)
		}

		booking = models.Booking{
			UserID:     userID,
			SessionID:  sessionID,
			Quantity:   quantity,
			TotalPrice: totalPrice,
		}

		if err := s.bookingRepository.Create(
			tx,
			ctx,
			&booking,
		); err != nil {
			return err
		}

		if provider != nil {
			if err := provider.Reserve(
				tx,
				ctx,
				booking.BookingID,
				sessionID,
				resourceIDs,
			); err != nil {
				return err
			}
		}

		if err := s.bookingRepository.DecreaseSessionCapacity(
			tx,
			ctx,
			sessionID,
			quantity,
		); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &booking, nil
}
