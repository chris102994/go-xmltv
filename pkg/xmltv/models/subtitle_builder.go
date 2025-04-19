package models

type SubtitleBuilderInterface interface {
	NewSubtitleBuilder() *SubtitleBuilder
	SetType(subtitleType string) *SubtitleBuilder
	SetLanguage(language *Language) *SubtitleBuilder
	Build() *Subtitle
}

type SubtitleBuilder struct {
	subtitle *Subtitle
}

func NewSubtitleBuilder() *SubtitleBuilder {
	return &SubtitleBuilder{
		subtitle: &Subtitle{},
	}
}

func (b *SubtitleBuilder) SetType(subtitleType string) *SubtitleBuilder {
	b.subtitle.Type = subtitleType
	return b
}

func (b *SubtitleBuilder) SetLanguage(language *Language) *SubtitleBuilder {
	b.subtitle.Lang = language
	return b
}

func (b *SubtitleBuilder) Build() *Subtitle {
	return b.subtitle
}
