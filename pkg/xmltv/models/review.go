package models

import (
	"encoding/xml"
)

type Review struct {
	XMLName  xml.Name `xml:"review"`
	Type     string   `xml:"type,attr,omitempty"`
	Source   string   `xml:"source,attr,omitempty"`
	Reviewer string   `xml:"reviewer,attr,omitempty"`
	Lang     string   `xml:"lang,attr,omitempty"`
	Text     string   `xml:",chardata"`
}
