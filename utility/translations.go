package utility

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var localizer *i18n.Localizer

func InitTranslations() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("translations/en.json")
	localizer = i18n.NewLocalizer(bundle, language.English.String())
}

// Translate a string into the current locale with optional template data
func TransData(str string, data interface{}) (string, error) {
	config := i18n.LocalizeConfig{MessageID: str, TemplateData: data}
	return localizer.Localize(&config)
}

// Translate a string into the current locale, panicking if necessary
func Trans(str string) string {
	translation, err := TransData(str, nil)
	if err != nil {
		panic(err)
	}
	return translation
}
