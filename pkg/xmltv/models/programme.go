package models

import (
	"encoding/xml"
)

type Programme struct {
	XMLName         xml.Name         `xml:"programme"`
	Start           string           `xml:"start,attr,omitempty"`
	Stop            string           `xml:"stop,attr,omitempty"`
	PdcStart        string           `xml:"pdc-start,attr,omitempty"`
	VpsStart        string           `xml:"vps-start,attr,omitempty"`
	Showview        string           `xml:"showview,attr,omitempty"`
	Videoplus       string           `xml:"videoplus,attr,omitempty"`
	Channel         string           `xml:"channel,attr,omitempty"`
	Clumpidx        string           `xml:"clumpidx,attr,omitempty"`
	Title           []*Title         `xml:"title,omitempty"`
	SubTitle        []*SubTitle      `xml:"sub-title,omitempty"`
	Desc            []*Desc          `xml:"desc,omitempty"`
	Credits         *Credits         `xml:"credits,omitempty"`
	Date            string           `xml:"date,omitempty"`
	Category        []*Category      `xml:"category,omitempty"`
	Keyword         []*Keyword       `xml:"keyword,omitempty"`
	Language        string           `xml:"language,omitempty"`
	OrigLanguage    string           `xml:"orig-language,omitempty"`
	Length          *Length          `xml:"length,omitempty"`
	Icons           []*Icon          `xml:"icon,omitempty"`
	URLs            []*URL           `xml:"url,omitempty"`
	Country         []*Country       `xml:"country,omitempty"`
	EpisodeNum      []*EpisodeNum    `xml:"episode-num,omitempty"`
	Video           *Video           `xml:"video,omitempty"`
	Audio           *Audio           `xml:"audio,omitempty"`
	PreviouslyShown *PreviouslyShown `xml:"previously-shown,omitempty"`
	Premiere        string           `xml:"premiere,omitempty"`
	LastChance      []*LastChance    `xml:"last-chance,omitempty"`
	New             string           `xml:"new,omitempty"`
	Subtitles       []*Subtitle      `xml:"subtitles,omitempty"`
	Rating          []*Rating        `xml:"rating,omitempty"`
	StarRating      []*StarRating    `xml:"star-rating,omitempty"`
	Review          []*Review        `xml:"review,omitempty"`
	Image           []*Image         `xml:"image,omitempty"`
}
