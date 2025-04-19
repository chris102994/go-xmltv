package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestPreviouslyShownBuilder(t *testing.T) {
	var start = "2023-10-01T12:00:00Z"
	var channel = "channel1"

	previouslyShown := xmltv.NewPreviouslyShownBuilder().
		SetStart(start).
		SetChannel(channel).
		Build()

	if previouslyShown.Start != start {
		t.Errorf("expected Start to be '%s', got '%s'", start, previouslyShown.Start)
	}
	if previouslyShown.Channel != channel {
		t.Errorf("expected Channel to be '%s', got '%s'", channel, previouslyShown.Channel)
	}
}
