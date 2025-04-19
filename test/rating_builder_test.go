package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestRatingBuilder(t *testing.T) {
	var system = "MPAA"
	var value = "PG-13"
	icon := &xmltv.Icon{Src: "http://example.com/icon.png"}

	rating := xmltv.NewRatingBuilder().
		SetSystem(system).
		SetValue(value).
		SetIcon(icon).
		Build()

	if rating.System != system {
		t.Errorf("expected System to be '%s', got '%s'", system, rating.System)
	}
	if rating.Value != value {
		t.Errorf("expected Value to be '%s', got '%s'", value, rating.Value)
	}
	if rating.Icon != icon {
		t.Errorf("expected Icon to be '%v', got '%v'", icon, rating.Icon)
	}
}
