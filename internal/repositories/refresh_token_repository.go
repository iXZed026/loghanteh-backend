package repositories

import (
	"context"
	"errors"
	"time"

	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(
	db *gorm.DB,
) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Create(
	ctx context.Context,
	refreshToken *models.RefreshToken,
) error {

	return r.db.
		WithContext(ctx).
		Create(refreshToken).
		Error
}

func (r *RefreshTokenRepository) FindValidToken(
	ctx context.Context,
	token string,
) (*models.RefreshToken, error) {

	var refreshToken models.RefreshToken

	err := r.db.
		WithContext(ctx).
		Where(
			"token = ? AND revoked = ? AND expiresat > ?",
			token,
			false,
			time.Now(),
		).
		First(&refreshToken).
		Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrInvalidCredentials
		}

		return nil, err
	}

	return &refreshToken, nil
}

func (r *RefreshTokenRepository) Revoke(
	ctx context.Context,
	token string,
) error {

	result := r.db.
		WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("token = ?", token).
		Update("revoked", true)

	return result.Error
}

func (r *RefreshTokenRepository) GetAllRefreshTokens(
	ctx context.Context,
) ([]models.RefreshToken, error) {

	var tokens []models.RefreshToken

	result := r.db.
		WithContext(ctx).
		Find(&tokens)

	if result.Error != nil {
		return nil, result.Error
	}

	return tokens, nil
}

func (r *LoginVerificationRepository) FindByToken(
	ctx context.Context,
	token string,
) (*models.LoginVerification, error) {

	var verification models.LoginVerification

	err := r.db.
		WithContext(ctx).
		Where("token = ?", token).
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
