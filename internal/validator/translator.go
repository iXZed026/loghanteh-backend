package validator

import (
	"github.com/go-playground/validator/v10"
)

func TranslateErrors(
	err error,
	lang string,
) string {

	languageMessages, ok := messages[lang]

	if !ok {
		languageMessages = messages["en"]
	}

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok || len(validationErrors) == 0 {
		return err.Error()
	}

	e := validationErrors[0]

	fieldMessages, ok := languageMessages[e.Field()]

	if !ok {
		return e.Error()
	}

	message, ok := fieldMessages[e.Tag()]

	if !ok {
		return e.Error()
	}

	return message
}
