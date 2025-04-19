package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestLanguageBuilder(t *testing.T) {
	var lang = "en"
	var text = "English"

	language := xmltv.NewLanguageBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if language.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, language.Lang)
	}
	if language.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, language.Text)
	}
}
