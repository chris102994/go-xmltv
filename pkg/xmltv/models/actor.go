package models

import (
	"encoding/xml"
)

type Actor struct {
	XMLName xml.Name `xml:"actor"`
	Role    string   `xml:"role,attr,omitempty"`
	Guest   string   `xml:"guest,attr,omitempty"`
	Text    string   `xml:",chardata"`
	Image   *Image   `xml:",omitempty"`
	URL     *URL     `xml:",omitempty"`
}
