package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestStarRatingBuilder(t *testing.T) {
	var system = "MPAA"
	var value = "PG-13"
	icon := &xmltv.Icon{Src: "http://example.com/icon.png"}

	starRating := xmltv.NewStarRatingBuilder().
		SetSystem(system).
		SetValue(value).
		SetIcon(icon).
		Build()

	if starRating.System != system {
		t.Errorf("expected System to be '%s', got '%s'", system, starRating.System)
	}
	if starRating.Value != value {
		t.Errorf("expected Value to be '%s', got '%s'", value, starRating.Value)
	}
	if starRating.Icon != icon {
		t.Errorf("expected Icon to be '%v', got '%v'", icon, starRating.Icon)
	}
}
