package asset

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const translationPath = "translation/en.json"

var localizer *i18n.Localizer

func InitTranslations() error {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bytes, err := FS.ReadFile(translationPath)
	if err != nil {
		return fmt.Errorf("loading english translation: %w", err)
	}

	bundle.ParseMessageFileBytes(bytes, translationPath)
	localizer = i18n.NewLocalizer(bundle, language.English.String())

	return nil
}

// Translate a string into the current locale with optional template data
func TransData(str string, data interface{}) (string, error) {
	return localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    str,
		TemplateData: data,
	})
}

// Translate a string into the current locale, panicking if necessary
func Trans(str string) string {
	translation, err := TransData(str, nil)
	if err != nil {
		panic(err)
	}

	return translation
}
