package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestDescBuilder(t *testing.T) {
	var lang = "en"
	var text = "This is a description"

	desc := xmltv.NewDescBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if desc.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, desc.Lang)
	}
	if desc.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, desc.Text)
	}
}
