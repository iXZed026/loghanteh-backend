package response

import (
	"errors"
	"log/slog"
	"net/http"

	appErr "loghanteh-project/internal/errors"
	appI18n "loghanteh-project/internal/i18n"
	appValidator "loghanteh-project/internal/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleError(
	c *gin.Context,
	err error,
	logger *slog.Logger,
) {

	logger.Error(
		err.Error(),
	)

	// Handle validation errors
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {

		lang := c.GetHeader("Accept-Language")

		message :=
			appValidator.TranslateErrors(
				validationErrors,
				lang,
			)

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": message,
			},
		)

		return
	}

	// Handle application errors
	var e *appErr.AppError

	if errors.As(err, &e) {

		lang := c.GetHeader("Accept-Language")

		message :=
			appI18n.TranslateError(
				e.Key,
				lang,
			)

		c.JSON(
			e.Code,
			gin.H{
				"success": false,
				"message": message,
			},
		)

		return
	}

	// Unknown errors
	lang := c.GetHeader("Accept-Language")

	message :=
		appI18n.TranslateError(
			"internal_server_error",
			lang,
		)

	c.JSON(
		http.StatusInternalServerError,
		gin.H{
			"success": false,
			"message": message,
		},
	)
}
