package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestVideoBuilder(t *testing.T) {
	var present = "yes"
	var colour = "color"
	var aspect = "16:9"
	var quality = "high"

	video := xmltv.NewVideoBuilder().
		SetPresent(present).
		SetColour(colour).
		SetAspect(aspect).
		SetQuality(quality).
		Build()

	if video.Present != present {
		t.Errorf("expected Present to be '%s', got '%s'", present, video.Present)
	}
	if video.Colour != colour {
		t.Errorf("expected Colour to be '%s', got '%s'", colour, video.Colour)
	}
	if video.Aspect != aspect {
		t.Errorf("expected Aspect to be '%s', got '%s'", aspect, video.Aspect)
	}
	if video.Quality != quality {
		t.Errorf("expected Quality to be '%s', got '%s'", quality, video.Quality)
	}
}
