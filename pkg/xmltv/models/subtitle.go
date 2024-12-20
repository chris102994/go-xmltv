package models

import (
	"encoding/xml"
)

type Subtitle struct {
	XMLName xml.Name  `xml:"subtitles"`
	Type    string    `xml:"type,attr,omitempty"`
	Lang    *Language `xml:"language,omitempty"`
}
