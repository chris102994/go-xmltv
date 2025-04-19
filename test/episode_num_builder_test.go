package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestEpisodeNumBuilder(t *testing.T) {
	var system = "system1"
	var text = "episode1"

	episodeNum := xmltv.NewEpisodeNumBuilder().
		SetSystem(system).
		SetText(text).
		Build()

	if episodeNum.System != system {
		t.Errorf("expected System to be '%s', got '%s'", system, episodeNum.System)
	}
	if episodeNum.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, episodeNum.Text)
	}
}
