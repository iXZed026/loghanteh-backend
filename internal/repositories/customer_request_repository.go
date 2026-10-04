package repositories

import (
	"context"

	"loghanteh-project/internal/models"

	"gorm.io/gorm"
)

type CustomerRequestRepository struct {
	db *gorm.DB
}

func NewCustomerRequestRepository(
	db *gorm.DB,
) *CustomerRequestRepository {
	return &CustomerRequestRepository{
		db: db,
	}
}

func (r *CustomerRequestRepository) CreateCustomerRequest(
	customerRequest *models.CustomerRequest,
	ctx context.Context,
) error {

	err := r.db.
		WithContext(ctx).
		Create(customerRequest).
		Error

	if err != nil {
		return err
	}

	return nil
}
