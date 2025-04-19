package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestKeywordBuilder(t *testing.T) {
	var lang = "en"
	var text = "example keyword"

	keyword := xmltv.NewKeywordBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if keyword.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, keyword.Lang)
	}
	if keyword.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, keyword.Text)
	}
}
