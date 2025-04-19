package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestCountryBuilder(t *testing.T) {
	var lang = "en"
	var text = "United States"

	country := xmltv.NewCountryBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if country.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, country.Lang)
	}
	if country.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, country.Text)
	}
}
