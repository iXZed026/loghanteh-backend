package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"loghanteh-project/internal/dto"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/repositories"

	"gorm.io/gorm"
)

type DiscountCodeService struct {
	discountCodeRepository *repositories.DiscountCodeRepository
}

func NewDiscountCodeService(
	discountCodeRepository *repositories.DiscountCodeRepository,
) *DiscountCodeService {
	return &DiscountCodeService{
		discountCodeRepository: discountCodeRepository,
	}
}

// --------------------------------------------------
// Get All Discount Codes
// --------------------------------------------------

func (s *DiscountCodeService) GetAll(
	ctx context.Context,
) ([]dto.DiscountCodeResponse, error) {
	return s.discountCodeRepository.GetAll(ctx)
}

// --------------------------------------------------
// Get Discount Code By Code
// --------------------------------------------------

func (s *DiscountCodeService) GetByCode(
	ctx context.Context,
	code string,
) (*dto.DiscountCodeResponse, error) {
	code = strings.TrimSpace(code)

	if code == "" {
		return nil, appErrors.ErrEmptyValue
	}

	discountCode, err := s.discountCodeRepository.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrDiscountCodeNotFound
		}

		return nil, err
	}

	now := time.Now()

	if !discountCode.IsActive {
		return nil, appErrors.ErrDiscountCodeInactive
	}

	if now.Before(discountCode.StartAt) {
		return nil, appErrors.ErrDiscountCodeNotStarted
	}

	if !now.Before(discountCode.ExpiresAt) {
		return nil, appErrors.ErrDiscountCodeExpired
	}

	if discountCode.UsageLimit > 0 &&
		discountCode.UsedCount >= discountCode.UsageLimit {
		return nil, appErrors.ErrDiscountCodeUsageLimitReached
	}

	return discountCode, nil
}
