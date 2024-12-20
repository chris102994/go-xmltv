package models

import "encoding/xml"

type Audio struct {
	XMLName xml.Name `xml:"audio"`
	Present string   `xml:"present,omitempty"`
	Stereo  string   `xml:"stereo,omitempty"`
}
