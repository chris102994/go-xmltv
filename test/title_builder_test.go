package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestTitleBuilder(t *testing.T) {
	var lang = "en"
	var text = "Example Title"

	title := xmltv.NewTitleBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if title.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, title.Lang)
	}
	if title.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, title.Text)
	}
}
