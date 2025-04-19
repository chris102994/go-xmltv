package test

import (
	xmltv "github.com/chris102994/go-xmltv/pkg/xmltv/models"
	"testing"
)

func TestTVBuilder(t *testing.T) {
	var date = "2023-10-01"
	var sourceInfoURL = "http://example.com"
	var sourceInfoName = "Example Source"
	var sourceDataURL = "http://data.example.com"
	var generatorInfoName = "Example Generator"
	var generatorInfoURL = "http://generator.example.com"

	channel := &xmltv.Channel{
		ID: "channel1",
	}
	programme := &xmltv.Programme{
		Title: append(make([]*xmltv.Title, 0), &xmltv.Title{Text: "Example Programme"}),
	}

	tv := xmltv.NewTVBuilder().
		SetDate(date).
		SetSourceInfoURL(sourceInfoURL).
		SetSourceInfoName(sourceInfoName).
		SetSourceDataURL(sourceDataURL).
		SetGeneratorInfoName(generatorInfoName).
		SetGeneratorInfoURL(generatorInfoURL).
		AddChannel(channel).
		AddProgramme(programme).
		Build()

	if tv.Date != date {
		t.Errorf("expected Date to be '%s', got '%s'", date, tv.Date)
	}
	if tv.SourceInfoURL != sourceInfoURL {
		t.Errorf("expected SourceInfoURL to be '%s', got '%s'", sourceInfoURL, tv.SourceInfoURL)
	}
	if tv.SourceInfoName != sourceInfoName {
		t.Errorf("expected SourceInfoName to be '%s', got '%s'", sourceInfoName, tv.SourceInfoName)
	}
	if tv.SourceDataURL != sourceDataURL {
		t.Errorf("expected SourceDataURL to be '%s', got '%s'", sourceDataURL, tv.SourceDataURL)
	}
	if tv.GeneratorInfoName != generatorInfoName {
		t.Errorf("expected GeneratorInfoName to be '%s', got '%s'", generatorInfoName, tv.GeneratorInfoName)
	}
	if tv.GeneratorInfoURL != generatorInfoURL {
		t.Errorf("expected GeneratorInfoURL to be '%s', got '%s'", generatorInfoURL, tv.GeneratorInfoURL)
	}
	if len(tv.Channels) != 1 || tv.Channels[0] != channel {
		t.Errorf("expected Channels to only contain the channel: %v, got '%v'", channel, tv.Channels)
	}
	if len(tv.Programmes) != 1 || tv.Programmes[0] != programme {
		t.Errorf("expected Programmes to only contain the programme: %v, got '%v'", programme, tv.Programmes)
	}
}
