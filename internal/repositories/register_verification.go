package repositories

import (
	"context"

	"gorm.io/gorm"

	"loghanteh-project/internal/models"
)

type RegisterVerificationRepository struct {
	db *gorm.DB
}

func NewRegisterVerificationRepository(
	db *gorm.DB,
) *RegisterVerificationRepository {
	return &RegisterVerificationRepository{
		db: db,
	}
}

func (r *RegisterVerificationRepository) Create(
	ctx context.Context,
	verification *models.RegisterVerification,
) error {

	return r.db.WithContext(ctx).
		Create(verification).
		Error
}

func (r *RegisterVerificationRepository) FindByToken(
	ctx context.Context,
	token string,
) (*models.RegisterVerification, error) {

	var verification models.RegisterVerification

	err := r.db.WithContext(ctx).
		Where("token = ?", token).
		First(&verification).
		Error

	if err != nil {
		return nil, err
	}

	return &verification, nil
}

func (r *RegisterVerificationRepository) MarkVerified(
	ctx context.Context,
	verification *models.RegisterVerification,
) error {

	return r.db.WithContext(ctx).
		Model(verification).
		Update("verified", true).
		Error
}
