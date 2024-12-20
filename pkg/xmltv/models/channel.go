package models

import (
	"encoding/xml"
)

type Channel struct {
	XMLName      xml.Name       `xml:"channel"`
	ID           string         `xml:"id,attr"`
	DisplayNames []*DisplayName `xml:"display-name,omitempty"`
	Icons        []*Icon        `xml:"icon,omitempty"`
	URLs         []*URL         `xml:"url,omitempty"`
}
