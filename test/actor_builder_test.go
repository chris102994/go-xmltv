package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestActorBuilder(t *testing.T) {
	var role = "Main Actor"
	var guest = "Guest Actor"
	var text = "John Doe"
	var image = &xmltv.Image{
		Text: "http://example.com/image.jpg",
	}
	var url = &xmltv.URL{
		Text: "http://example.com",
	}

	actor := xmltv.NewActorBuilder().
		SetRole(role).
		SetGuest(guest).
		SetText(text).
		SetImage(image).
		SetURL(url).
		Build()

	if actor.Role != role {
		t.Errorf("expected Role to be '%s', got '%s'", role, actor.Role)
	}
	if actor.Guest != guest {
		t.Errorf("expected Guest to be '%s', got '%s'", guest, actor.Guest)
	}
	if actor.Text != text {
		t.Errorf("expected Text to be '%s', got '%s'", text, actor.Text)
	}
	if actor.Image.Text != image.Text {
		t.Errorf("expected Image Text to be '%s', got '%s'", image.Text, actor.Image.Text)
	}
	if actor.URL.Text != url.Text {
		t.Errorf("expected URL Text to be '%s', got '%s'", url.Text, actor.URL.Text)
	}
}
