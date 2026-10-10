package reservation

import (
	"context"
	"errors"
	"time"

	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"
	"loghanteh-project/internal/repositories"

	"github.com/google/uuid"
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
	discountCodeID *uint,
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

		var discountAmount decimal.Decimal

		if discountCodeID != nil {
			discountCode, err := s.bookingRepository.GetDiscountCodeForBooking(
				tx,
				ctx,
				*discountCodeID,
			)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return appErrors.ErrDiscountCodeNotFound
				}

				return err
			}

			if !discountCode.IsActive {
				return appErrors.ErrDiscountCodeInactive
			}

			now := time.Now()

			if now.Before(discountCode.StartAt) {
				return appErrors.ErrDiscountCodeNotStarted
			}

			if now.After(discountCode.ExpiresAt) {
				return appErrors.ErrDiscountCodeExpired
			}

			if discountCode.UsageLimit > 0 &&
				discountCode.UsedCount >= discountCode.UsageLimit {
				return appErrors.ErrDiscountCodeUsageLimitReached
			}

			discountAmount = totalPrice.
				Mul(discountCode.DiscountPercent).
				Div(decimal.NewFromInt(100))

			if discountCode.MaxDiscountAmount.GreaterThan(decimal.Zero) &&
				discountAmount.GreaterThan(discountCode.MaxDiscountAmount) {
				discountAmount = discountCode.MaxDiscountAmount
			}

			if discountAmount.GreaterThan(totalPrice) {
				discountAmount = totalPrice
			}
		}

		booking = models.Booking{
			UserID:         userID,
			SessionID:      sessionID,
			Quantity:       quantity,
			TotalPrice:     totalPrice,
			DiscountCodeID: discountCodeID,
			DiscountAmount: discountAmount,
			TicketToken:    uuid.New(),
			TicketIsValid:  true,
		}

		if err := s.bookingRepository.Create(
			tx,
			ctx,
			&booking,
		); err != nil {
			return err
		}

		if discountCodeID != nil {
			if err := s.bookingRepository.IncrementDiscountUsage(
				tx,
				ctx,
				*discountCodeID,
			); err != nil {
				return err
			}
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
