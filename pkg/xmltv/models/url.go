package models

import (
	"encoding/xml"
)

type URL struct {
	XMLName xml.Name `xml:"url"`
	System  string   `xml:"system,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
