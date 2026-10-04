package repositories

import (
	"context"

	"loghanteh-project/internal/errors"

	"gorm.io/gorm"
)

type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(
	db *gorm.DB,
) *LanguageRepository {
	return &LanguageRepository{
		db: db,
	}
}

func (r *LanguageRepository) GetIDByCode(
	ctx context.Context,
	code string,
) (uint, error) {

	var languageID uint

	err := r.db.
		WithContext(ctx).
		Table("languages").
		Select("languagesid").
		Where("LOWER(code) = LOWER(?)", code).
		Scan(&languageID).
		Error

	if err != nil {
		return 0, err
	}

	if languageID == 0 {
		return 0, errors.ErrInvalidInput
	}

	return languageID, nil
}
