package models

import "encoding/xml"

type Language struct {
	XMLName xml.Name `xml:"language"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
