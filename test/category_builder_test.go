package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestCategoryBuilder(t *testing.T) {
	var lang = "en"
	var text = "News"

	category := xmltv.NewCategoryBuilder().
		SetLang(lang).
		SetText(text).
		Build()

	if category.Lang != lang {
		t.Errorf("expected Lang to be '%s', got '%s'", lang, category.Lang)
	}
	if category.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, category.Text)
	}
}
