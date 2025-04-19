package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestImageBuilder(t *testing.T) {
	var imageType = "icon"
	var width = "100"
	var height = "200"
	var orientation = "landscape"
	var text = "example.png"

	image := xmltv.NewImageBuilder().
		SetType(imageType).
		SetWidth(width).
		SetHeight(height).
		SetOrientation(orientation).
		SetText(text).
		Build()

	if image.Type != imageType {
		t.Errorf("expected Type to be '%s', got '%s'", imageType, image.Type)
	}
	if image.Width != width {
		t.Errorf("expected Width to be '%s', got '%s'", width, image.Width)
	}
	if image.Height != height {
		t.Errorf("expected Height to be '%s', got '%s'", height, image.Height)
	}
	if image.Orientation != orientation {
		t.Errorf("expected Orientation to be '%s', got '%s'", orientation, image.Orientation)
	}
	if image.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, image.Text)
	}
}
