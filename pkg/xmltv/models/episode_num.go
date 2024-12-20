package models

import (
	"encoding/xml"
)

type EpisodeNum struct {
	XMLName xml.Name `xml:"episode-num"`
	System  string   `xml:"system,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
