package models

import (
	"encoding/xml"
)

type DisplayName struct {
	XMLName xml.Name `xml:"display-name"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
