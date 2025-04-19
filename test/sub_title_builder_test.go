package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestSubTitleBuilder(t *testing.T) {
	var lang = "en"
	var text = "Example Subtitle"

	subTitle := xmltv.NewSubTitleBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if subTitle.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, subTitle.Lang)
	}
	if subTitle.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, subTitle.Text)
	}
}
