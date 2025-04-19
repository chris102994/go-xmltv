package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestIconBuilder(t *testing.T) {
	var src = "http://example.com/icon.png"
	var width = "100"
	var height = "200"

	icon := xmltv.NewIconBuilder().
		SetSrc(src).
		SetWidth(width).
		SetHeight(height).
		Build()

	if icon.Src != src {
		t.Errorf("expected Src to be '%s', got '%s'", src, icon.Src)
	}
	if icon.Width != width {
		t.Errorf("expected Width to be '%s', got '%s'", width, icon.Width)
	}
	if icon.Height != height {
		t.Errorf("expected Height to be '%s', got '%s'", height, icon.Height)
	}
}
