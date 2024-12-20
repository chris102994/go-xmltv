package models

import (
	"encoding/xml"
)

type Credits struct {
	XMLName      xml.Name `xml:"credits"`
	Directors    []string `xml:"director,omitempty"`
	Actors       []*Actor `xml:"actor,omitempty"`
	Writers      []string `xml:"writer,omitempty"`
	Adapters     []string `xml:"adapter,omitempty"`
	Producers    []string `xml:"producer,omitempty"`
	Composers    []string `xml:"composer,omitempty"`
	Editors      []string `xml:"editor,omitempty"`
	Presenters   []string `xml:"presenter,omitempty"`
	Commentators []string `xml:"commentator,omitempty"`
	Guests       []string `xml:"guest,omitempty"`
}
