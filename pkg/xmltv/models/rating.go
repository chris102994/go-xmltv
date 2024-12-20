package models

import (
	"encoding/xml"
)

type Rating struct {
	XMLName xml.Name `xml:"rating"`
	System  string   `xml:"system,attr,omitempty"`
	Value   string   `xml:"value,omitempty"`
	Icon    *Icon    `xml:"icon,omitempty"`
}
