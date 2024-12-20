package models

import (
	"encoding/xml"
)

type LastChance struct {
	XMLName xml.Name `xml:"last-chance"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
