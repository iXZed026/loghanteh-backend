package reservation

import (
	"context"

	"loghanteh-project/internal/models"

	"gorm.io/gorm"
)

type ResourceProvider interface {
	ValidateAndLock(
		tx *gorm.DB,
		ctx context.Context,
		session *models.EventSession,
		resourceIDs []uint,
	) error

	Reserve(
		tx *gorm.DB,
		ctx context.Context,
		bookingID uint,
		sessionID uint,
		resourceIDs []uint,
	) error
}
