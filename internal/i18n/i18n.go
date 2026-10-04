package i18n

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Messages struct {
	Errors  map[string]string `json:"errors"`
	Success map[string]string `json:"success"`
}

var translations = map[string]Messages{}

func init() {

	loadLanguage("en")
	loadLanguage("fa")
}

func loadLanguage(
	lang string,
) {

	path :=
		filepath.Join(
			"internal",
			"i18n",
			lang+".json",
		)

	data, err :=
		os.ReadFile(path)

	if err != nil {
		return
	}

	var messages Messages

	if err :=
		json.Unmarshal(
			data,
			&messages,
		); err != nil {
		return
	}

	translations[lang] = messages
}

func TranslateError(
	key string,
	lang string,
) string {

	lang =
		normalizeLanguage(lang)

	// Try requested language
	if messages, ok :=
		translations[lang]; ok {

		if message, ok :=
			messages.Errors[key]; ok {

			return message
		}
	}

	// Fallback to English
	if messages, ok :=
		translations["en"]; ok {

		if message, ok :=
			messages.Errors[key]; ok {

			return message
		}
	}

	// Last fallback
	return key
}

func TranslateSuccess(
	key string,
	lang string,
) string {

	lang =
		normalizeLanguage(lang)

	// Try requested language
	if messages, ok :=
		translations[lang]; ok {

		if message, ok :=
			messages.Success[key]; ok {

			return message
		}
	}

	// Fallback to English
	if messages, ok :=
		translations["en"]; ok {

		if message, ok :=
			messages.Success[key]; ok {

			return message
		}
	}

	// Last fallback
	return key
}

func normalizeLanguage(
	lang string,
) string {

	lang =
		strings.TrimSpace(
			strings.ToLower(lang),
		)

	if len(lang) >= 2 {
		return lang[:2]
	}

	return lang
}
