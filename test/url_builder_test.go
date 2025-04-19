package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestURLBuilder(t *testing.T) {
	var system = "system1"
	var text = "http://example.com"

	url := xmltv.NewURLBuilder().
		SetSystem(system).
		SetText(text).
		Build()

	if url.System != system {
		t.Errorf("expected System to be '%s', got '%s'", system, url.System)
	}
	if url.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, url.Text)
	}
}
