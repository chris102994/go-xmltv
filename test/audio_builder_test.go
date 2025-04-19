package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestAudioBuilder(t *testing.T) {
	var present = "yes"
	var stereo = "stereo"

	audio := xmltv.NewAudioBuilder().
		SetPresent(present).
		SetStereo(stereo).
		Build()

	if audio.Present != present {
		t.Errorf("expected Present to be '%s', got '%s'", present, audio.Present)
	}
	if audio.Stereo != stereo {
		t.Errorf("expected Stereo to be '%s', got '%s'", stereo, audio.Stereo)
	}
}
