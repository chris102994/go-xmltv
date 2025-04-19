package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestDisplayNameBuilder(t *testing.T) {
	var lang = "en"
	var text = "Example Channel"

	displayName := xmltv.NewDisplayNameBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if displayName.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, displayName.Lang)
	}
	if displayName.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, displayName.Text)
	}
}
