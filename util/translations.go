package util

import (
	"encoding/json"
	"fmt"

	"github.com/leviceccato/clarity/asset"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
)

func InitTranslations() error {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	path := "translation/en.json"
	bytes, err := asset.FS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("loading english translation: %s", err)
	}
	bundle.ParseMessageFileBytes(bytes, path)
	localizer = i18n.NewLocalizer(bundle, language.English.String())
	return nil
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
