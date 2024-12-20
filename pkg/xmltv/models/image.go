package models

import (
	"encoding/xml"
)

type Image struct {
	XMLName     xml.Name `xml:"image"`
	Type        string   `xml:"type,attr,omitempty"`
	Width       string   `xml:"width,attr,omitempty"`
	Height      string   `xml:"height,attr,omitempty"`
	Orientation string   `xml:"orientation,attr,omitempty"`
	Text        string   `xml:",chardata"`
}
