package handlers

import (
	"log/slog"
	"strings"

	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/response"
	"loghanteh-project/internal/services"

	"github.com/gin-gonic/gin"
)

type DiscountCodeHandler struct {
	discountCodeService *services.DiscountCodeService
	logger              *slog.Logger
}

func NewDiscountCodeHandler(
	discountCodeService *services.DiscountCodeService,
	logger *slog.Logger,
) *DiscountCodeHandler {
	return &DiscountCodeHandler{
		discountCodeService: discountCodeService,
		logger:              logger,
	}
}

// --------------------------------------------------
// Get All Discount Codes
// --------------------------------------------------

func (h *DiscountCodeHandler) GetAll(c *gin.Context) {
	discountCodes, err := h.discountCodeService.GetAll(
		c.Request.Context(),
	)
	if err != nil {
		h.logger.Error(
			"failed to get all discount codes",
			"error", err,
		)

		response.HandleError(c, err, h.logger)
		return
	}

	response.Success(
		c,
		"discount_codes_fetched_successfully",
		discountCodes,
	)
}

// --------------------------------------------------
// Get Discount Code By Code
// --------------------------------------------------

func (h *DiscountCodeHandler) GetByCode(c *gin.Context) {
	code := strings.TrimSpace(c.Param("code"))

	if code == "" {
		response.HandleError(c, appErrors.ErrEmptyValue, h.logger)
		return
	}

	if len(code) > 50 {
		response.HandleError(c, appErrors.ErrInvalidInput, h.logger)
		return
	}

	discountCode, err := h.discountCodeService.GetByCode(
		c.Request.Context(),
		code,
	)
	if err != nil {
		h.logger.Error(
			"failed to get discount code",
			"code", code,
			"error", err,
		)

		response.HandleError(c, err, h.logger)
		return
	}

	response.Success(
		c,
		"discount_code_applied_successfully",
		discountCode,
	)
}
