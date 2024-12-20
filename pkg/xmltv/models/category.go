package models

import (
	"encoding/xml"
)

type Category struct {
	XMLName xml.Name `xml:"category"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
