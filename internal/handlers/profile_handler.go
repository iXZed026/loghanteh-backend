package handlers

import (
	"fmt"
	"log/slog"
	"strconv"

	"loghanteh-project/internal/dto"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/i18n"
	"loghanteh-project/internal/response"
	"loghanteh-project/internal/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService *services.ProfileService
	logger         *slog.Logger
}

func NewProfileHandler(
	profileService *services.ProfileService,
	logger *slog.Logger,
) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
		logger:         logger,
	}
}

func (h *ProfileHandler) Profile(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {

		h.logger.Error(
			"failed to get user id from context",
			"error", "user_id is missing",
		)

		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)

		return
	}

	userID, ok := userIDValue.(string)

	if !ok {

		h.logger.Error(
			"invalid user id in context",
			"type", fmt.Sprintf("%T", userIDValue),
		)

		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)

		return
	}

	id, err := strconv.ParseUint(
		userID,
		10,
		64,
	)

	if err != nil {

		h.logger.Error(
			"failed to parse user id",
			"user_id", userID,
			"error", err,
		)

		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)

		return
	}

	user, err := h.profileService.GetProfile(
		c.Request.Context(),
		uint(id),
	)

	if err != nil {

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		i18n.TranslateSuccess(
			"profile_fetched_successfully",
			c.GetHeader("Accept-Language"),
		),
		user,
	)
}

func (h *ProfileHandler) EditProfile(
	c *gin.Context,
) {

	userIDValue, exists := c.Get("user_id")

	if !exists {

		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)

		return
	}

	userIDString, ok := userIDValue.(string)

	if !ok {

		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)

		return
	}

	userID, err := strconv.ParseUint(
		userIDString,
		10,
		64,
	)

	if err != nil {

		response.HandleError(
			c,
			appErrors.ErrUnauthorized,
			h.logger,
		)

		return
	}

	var req dto.EditProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"failed to bind edit profile request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	profile, err := h.profileService.EditProfile(
		c.Request.Context(),
		uint(userID),
		req,
	)

	if err != nil {

		h.logger.Error(
			"failed to edit profile",
			"error", err,
			"userID", userID,
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
		i18n.TranslateSuccess(
			"profile_updated_successfully",
			c.GetHeader("Accept-Language"),
		),
		profile,
	)
}
