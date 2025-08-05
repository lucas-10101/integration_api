package services

import (
	"api/static"
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
)

const (
	LANG_EN string = "en_us"
	LANG_PT string = "pt_br"
)

var (
	AvailableLanguages = []string{
		LANG_EN,
		LANG_PT,
	}
)

type I18nLanguagePack struct {
	translations map[string]string
}

func (pack *I18nLanguagePack) Init(translations map[string]string) {
	pack.translations = translations
}

func (pack *I18nLanguagePack) T(message string) string {
	if translation, ok := pack.translations[message]; ok {
		return translation
	}
	return ""
}

type I18nService struct {
	translations map[string]*I18nLanguagePack
}

func (instance *I18nService) Init() (err error) {

	instance.translations = make(map[string]*I18nLanguagePack)

	var entries []fs.DirEntry
	if entries, err = static.TranslationsFS.ReadDir("translations"); err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			return fmt.Errorf("%s is not a localization file", entry.Name())
		}

		file, err := static.TranslationsFS.Open(fmt.Sprintf("%s/%s", "translations", entry.Name()))

		if err != nil {
			return err
		}
		defer file.Close()

		var messages = make(map[string]string)
		if err = json.NewDecoder(file).Decode(&messages); err != nil {
			return err
		}

		language := strings.SplitN(entry.Name(), ".", 2)[0]

		languagePack := &I18nLanguagePack{}
		languagePack.Init(messages)

		instance.translations[language] = languagePack
	}

	return nil

}

func (instance *I18nService) GetLanguagePack(lang string) *I18nLanguagePack {
	if languagePack, ok := instance.translations[lang]; ok {
		return languagePack
	}
	return nil
}
