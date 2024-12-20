package models

import (
	"encoding/xml"
)

type Desc struct {
	XMLName xml.Name `xml:"desc"`
	Lang    string   `xml:"lang,attr,omitempty"`
	Text    string   `xml:",chardata"`
}
