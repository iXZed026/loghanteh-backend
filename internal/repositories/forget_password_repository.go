package repositories

import (
	"context"
	"errors"
	"time"

	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"

	"gorm.io/gorm"
)

type ForgetPasswordVerificationRepository struct {
	db *gorm.DB
}

func NewForgetPasswordVerificationRepository(
	db *gorm.DB,
) *ForgetPasswordVerificationRepository {
	return &ForgetPasswordVerificationRepository{
		db: db,
	}
}

func (r *ForgetPasswordVerificationRepository) Create(
	ctx context.Context,
	verification *models.ForgetPasswordVerification,
) error {

	return r.db.
		WithContext(ctx).
		Create(verification).
		Error
}

func (r *ForgetPasswordVerificationRepository) FindValidToken(
	ctx context.Context,
	token string,
) (*models.ForgetPasswordVerification, error) {

	var verification models.ForgetPasswordVerification

	err := r.db.
		WithContext(ctx).
		Where(
			"tokenhash = ? AND verified = ? AND expiresat > ?",
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

func (r *ForgetPasswordVerificationRepository) FindVerifiedToken(
	ctx context.Context,
	token string,
) (*models.ForgetPasswordVerification, error) {

	var verification models.ForgetPasswordVerification

	err := r.db.
		WithContext(ctx).
		Where(
			"tokenhash = ? AND verified = ?",
			token,
			true,
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

func (r *ForgetPasswordVerificationRepository) MarkVerified(
	ctx context.Context,
	verification *models.ForgetPasswordVerification,
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
