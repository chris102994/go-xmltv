package models

import (
	"encoding/xml"
)

type StarRating struct {
	XMLName xml.Name `xml:"star-rating"`
	System  string   `xml:"system,attr,omitempty"`
	Value   string   `xml:"value,omitempty"`
	Icon    *Icon    `xml:"icon,omitempty"`
}
