package models

import "encoding/xml"

type PreviouslyShown struct {
	XMLName xml.Name `xml:"previously-shown"`
	Start   string   `xml:"start,attr,omitempty"`
	Channel string   `xml:"channel,attr,omitempty"`
}
