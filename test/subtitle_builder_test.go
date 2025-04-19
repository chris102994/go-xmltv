package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestSubtitleBuilder(t *testing.T) {
	subtitleType := "subtitles"
	language := &xmltv.Language{Text: "en"}

	subtitle := xmltv.NewSubtitleBuilder().
		SetType(subtitleType).
		SetLanguage(language).
		Build()

	if subtitle.Type != subtitleType {
		t.Errorf("expected Type to be '%s', got '%s'", subtitleType, subtitle.Type)
	}
	if subtitle.Lang.Text != language.Text {
		t.Errorf("expected Language to be '%s', got '%s'", language.Text, subtitle.Lang.Text)
	}
}
