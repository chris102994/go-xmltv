package models

type KeywordBuilderInterface interface {
	NewKeywordBuilder() *KeywordBuilder
	SetLang(lang string) *KeywordBuilder
	SetText(text string) *KeywordBuilder
	Build() *Keyword
}

type KeywordBuilder struct {
	keyword *Keyword
}

func NewKeywordBuilder() *KeywordBuilder {
	return &KeywordBuilder{
		keyword: &Keyword{},
	}
}

func (b *KeywordBuilder) SetLang(lang string) *KeywordBuilder {
	b.keyword.Lang = lang
	return b
}

func (b *KeywordBuilder) SetText(text string) *KeywordBuilder {
	b.keyword.Text = text
	return b
}

func (b *KeywordBuilder) Build() *Keyword {
	return b.keyword
}
