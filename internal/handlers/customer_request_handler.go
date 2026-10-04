package handlers

import (
	"log/slog"

	"loghanteh-project/internal/dto"
	"loghanteh-project/internal/response"
	"loghanteh-project/internal/services"

	"github.com/gin-gonic/gin"
)

type CustomerRequestHandler struct {
	customerRequestService *services.CustomerRequestService
	logger                 *slog.Logger
}

func NewCustomerRequestHandler(
	customerRequestService *services.CustomerRequestService,
	logger *slog.Logger,
) *CustomerRequestHandler {
	return &CustomerRequestHandler{
		customerRequestService: customerRequestService,
		logger:                 logger,
	}
}

func (h *CustomerRequestHandler) CreateCustomerRequest(
	c *gin.Context,
) {

	var req *dto.CustomerRequestReq

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"failed to get customer request req",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	if err := h.customerRequestService.
		CreateCustomerRequest(
			c.Request.Context(),
			req,
		); err != nil {

		h.logger.Error(
			"failed to create customer request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		"customer_request_created_successfully",
		nil,
	)
}
