package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestLastChanceBuilder(t *testing.T) {
	var lang = "en"
	var text = "Last chance to watch!"

	lastChance := xmltv.NewLastChanceBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if lastChance.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, lastChance.Lang)
	}
	if lastChance.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, lastChance.Text)
	}
}
