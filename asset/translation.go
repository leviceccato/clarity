package asset

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator struct {
	Language   language.Tag
	localizers map[language.Tag]*i18n.Localizer
}

func NewTranslator() *Translator {
	t := &Translator{}
	t.localizers = map[language.Tag]*i18n.Localizer{}
	return t
}

// Add a new localizer, accepts a json file path and a language
func (t *Translator) AddLocalizer(path string, lang language.Tag) error {
	// Load translation file
	file, err := FS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("loading '%s' translation: %s", lang.String(), err)
	}

	bundle := i18n.NewBundle(lang)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.ParseMessageFileBytes(file, path)

	t.localizers[lang] = i18n.NewLocalizer(bundle, lang.String())

	return nil
}

func (t Translator) TransData(str string, data interface{}) (string, error) {
	translation, err := t.localizers[t.Language].Localize(&i18n.LocalizeConfig{
		MessageID:    str,
		TemplateData: data,
	})
	if err != nil {
		return "", fmt.Errorf("translating string '%s' with data: %s", str, err)
	}

	return translation, nil
}

func (t Translator) Trans(str string) string {
	translation, err := t.TransData(str, nil)
	if err != nil {
		panic(fmt.Sprintf("translating string '%s': %s", str, err))
	}

	return translation
}
