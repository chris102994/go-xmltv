package models

import (
	"encoding/xml"
)

type Icon struct {
	XMLName xml.Name `xml:"icon"`
	Src     string   `xml:"src,attr"`
	Width   string   `xml:"width,attr,omitempty"`
	Height  string   `xml:"height,attr,omitempty"`
}
