package models

import (
	"encoding/xml"
)

type Country struct {
	XMLName xml.Name `xml:"country"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
