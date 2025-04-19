package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestLengthBuilder(t *testing.T) {
	var units = "minutes"
	var text = "120"

	length := xmltv.NewLengthBuilder().
		SetUnits(units).
		SetText(text).
		Build()
	if length.Units != units {
		t.Errorf("expected Units to be '%s', got '%s'", units, length.Units)
	}

	if length.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, length.Text)
	}
}
