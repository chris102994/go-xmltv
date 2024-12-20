package models

import (
	"encoding/xml"
)

type Title struct {
	XMLName xml.Name `xml:"title"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
