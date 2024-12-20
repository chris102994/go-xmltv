package models

import (
	"encoding/xml"
)

type TV struct {
	XMLName           xml.Name     `xml:"tv"`
	Date              string       `xml:"date,attr,omitempty"`
	SourceInfoURL     string       `xml:"source-info-url,attr,omitempty"`
	SourceInfoName    string       `xml:"source-info-name,attr,omitempty"`
	SourceDataURL     string       `xml:"source-data-url,attr,omitempty"`
	GeneratorInfoName string       `xml:"generator-info-name,attr,omitempty"`
	GeneratorInfoURL  string       `xml:"generator-info-url,attr,omitempty"`
	Channels          []*Channel   `xml:"channel,omitempty"`
	Programmes        []*Programme `xml:"programme,omitempty"`
}
