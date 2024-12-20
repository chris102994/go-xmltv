package models

import (
	"encoding/xml"
)

type SubTitle struct {
	XMLName xml.Name `xml:"sub-title"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
