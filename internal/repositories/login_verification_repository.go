package repositories

import (
	"context"
	"errors"
	"time"

	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"

	"gorm.io/gorm"
)

type LoginVerificationRepository struct {
	db *gorm.DB
}

func NewLoginVerificationRepository(
	db *gorm.DB,
) *LoginVerificationRepository {
	return &LoginVerificationRepository{
		db: db,
	}
}

func (r *LoginVerificationRepository) Create(
	ctx context.Context,
	verification *models.LoginVerification,
) error {

	return r.db.
		WithContext(ctx).
		Create(verification).
		Error
}

func (r *LoginVerificationRepository) FindValidToken(
	ctx context.Context,
	token string,
) (*models.LoginVerification, error) {

	var verification models.LoginVerification

	err := r.db.
		WithContext(ctx).
		Where(
			"token = ? AND verified = ? AND expiresat > ?",
			token,
			false,
			time.Now(),
		).
		First(&verification).
		Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrInvalidInput
		}

		return nil, err
	}

	return &verification, nil
}

func (r *LoginVerificationRepository) MarkVerified(
	ctx context.Context,
	verification *models.LoginVerification,
) error {

	return r.db.
		WithContext(ctx).
		Model(verification).
		Update(
			"verified",
			true,
		).
		Error
}
