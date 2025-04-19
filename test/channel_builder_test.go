package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestChannelBuilder(t *testing.T) {
	var id = "channel1"

	displayName := &xmltv.DisplayName{Text: "Example Channel"}
	icon := &xmltv.Icon{Src: "http://example.com/icon.png"}
	url := &xmltv.URL{Text: "http://example.com"}

	channel := xmltv.NewChannelBuilder().
		SetID(id).
		AddDisplayName(displayName).
		AddIcon(icon).
		AddURL(url).
		Build()

	if channel.ID != id {
		t.Errorf("expected ID to be '%s', got '%s'", id, channel.ID)
	}
	if len(channel.DisplayNames) != 1 || channel.DisplayNames[0] != displayName {
		t.Errorf("expected DisplayNames to contain one display name: '%v', got '%v'", displayName, channel.DisplayNames[0])
	}
	if len(channel.Icons) != 1 || channel.Icons[0] != icon {
		t.Errorf("expected Icons to contain one icon: '%v', got '%v'", icon, channel.Icons[0])
	}
	if len(channel.URLs) != 1 || channel.URLs[0] != url {
		t.Errorf("expected URLs to contain one URL: '%v', got '%v'", url, channel.URLs[0])
	}
}
