package models

import (
	"encoding/xml"
)

type Keyword struct {
	XMLName xml.Name `xml:"keyword"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
