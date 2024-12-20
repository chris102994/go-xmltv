package models

import "encoding/xml"

type Video struct {
	XMLName xml.Name `xml:"video"`
	Present string   `xml:"present,omitempty"`
	Colour  string   `xml:"colour,omitempty"`
	Aspect  string   `xml:"aspect,omitempty"`
	Quality string   `xml:"quality,omitempty"`
}
