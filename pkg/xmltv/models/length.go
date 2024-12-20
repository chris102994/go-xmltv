package models

import "encoding/xml"

type Length struct {
	XMLName xml.Name `xml:"length"`
	Units   string   `xml:"units,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
