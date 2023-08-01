package driver

import (
	"encoding/json"
	"fmt"
	src "github.com/mindwingx/go-clean-arch-boilerplate"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

type LocaleAbstraction interface {
	InitLocale()
	Get(key string) string
	Plural(key string, params map[string]string) string
	FormatNumber(number int64) string
	FormatDate(date time.Time) string
	FormatCurrency(value float64, cur currency.Unit) string
}

type locale struct {
	bundle *i18n.Bundle
	Lang   string // Lang tag, like "en-US"
}

func NewLocale(registry RegistryAbstraction) LocaleAbstraction {
	lang := new(locale)
	registry.Parse(&lang)
	lang.bundle = i18n.NewBundle(language.English)

	return lang
}

func (l *locale) InitLocale() {
	l.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	// todo: use the below code snippet to register more translations of many other languages
	l.bundle.MustLoadMessageFile(fmt.Sprintf("%s/config/locale/%s", src.Root(), "en-US.json"))
}

func (l *locale) Get(key string) string {
	localizer := i18n.NewLocalizer(l.bundle, l.Lang)

	localizedMessage, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: key, // other fields are available with the i18n.Message struct
		},
	})

	return localizedMessage
}

func (l *locale) Plural(key string, params map[string]string) string {
	localizer := i18n.NewLocalizer(l.bundle, l.Lang)
	data := make(map[string]string)

	for localizerKey, localizerValue := range params {
		data[localizerKey] = localizerValue
	}

	formattedLocalizer := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: key,
		},
		TemplateData: data,
	})

	return formattedLocalizer
}

func (l *locale) FormatNumber(number int64) string {
	lang, _ := language.Parse(l.Lang)
	p := message.NewPrinter(lang)
	return p.Sprintf("%d", number)
}

func (l *locale) FormatDate(date time.Time) string {
	lang, _ := language.Parse(l.Lang)
	p := message.NewPrinter(lang)
	return p.Sprintf(
		"%s, %s %d, %d",
		date.Weekday(), date.Month(), date.Day(), date.Year(),
	)
}

func (l *locale) FormatCurrency(value float64, cur currency.Unit) string {
	lang, _ := language.Parse(l.Lang)
	p := message.NewPrinter(lang)
	return p.Sprintf("%s %.2f", currency.Symbol(cur), value)
}
