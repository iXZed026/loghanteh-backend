package i18n

import (
	"embed"
	"encoding/json"
	"strings"
)

//go:embed en.json fa.json
var translationFiles embed.FS

type Messages struct {
	Errors  map[string]string `json:"errors"`
	Success map[string]string `json:"success"`
}

var translations = map[string]Messages{}

func init() {
	loadLanguage("en")
	loadLanguage("fa")
}

func loadLanguage(lang string) {
	data, err := translationFiles.ReadFile(lang + ".json")
	if err != nil {
		panic("failed to read translation file: " + lang + ".json: " + err.Error())
	}

	var messages Messages
	if err := json.Unmarshal(data, &messages); err != nil {
		panic("failed to parse translation file: " + lang + ".json: " + err.Error())
	}

	if messages.Errors == nil {
		messages.Errors = make(map[string]string)
	}

	if messages.Success == nil {
		messages.Success = make(map[string]string)
	}

	translations[lang] = messages
}

func TranslateError(key string, lang string) string {
	lang = normalizeLanguage(lang)

	if messages, ok := translations[lang]; ok {
		if message, ok := messages.Errors[key]; ok {
			return message
		}
	}

	if messages, ok := translations["en"]; ok {
		if message, ok := messages.Errors[key]; ok {
			return message
		}
	}

	return key
}

func TranslateSuccess(key string, lang string) string {
	lang = normalizeLanguage(lang)

	if messages, ok := translations[lang]; ok {
		if message, ok := messages.Success[key]; ok {
			return message
		}
	}

	if messages, ok := translations["en"]; ok {
		if message, ok := messages.Success[key]; ok {
			return message
		}
	}

	return key
}

func normalizeLanguage(lang string) string {
	lang = strings.TrimSpace(strings.ToLower(lang))

	if lang == "" {
		return "en"
	}

	// Accept values such as "fa-IR,fa;q=0.9,en;q=0.8".
	if commaIndex := strings.Index(lang, ","); commaIndex >= 0 {
		lang = lang[:commaIndex]
	}

	if semicolonIndex := strings.Index(lang, ";"); semicolonIndex >= 0 {
		lang = lang[:semicolonIndex]
	}

	lang = strings.TrimSpace(lang)

	if hyphenIndex := strings.Index(lang, "-"); hyphenIndex >= 0 {
		lang = lang[:hyphenIndex]
	}

	if len(lang) >= 2 {
		return lang[:2]
	}

	return lang
}
