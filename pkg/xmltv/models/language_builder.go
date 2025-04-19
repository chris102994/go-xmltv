package models

type LanguageBuilderInterface interface {
	NewLanguageBuilder() *LanguageBuilder
	SetLang(lang string) *LanguageBuilder
	SetText(text string) *LanguageBuilder
	Build() *Language
}

type LanguageBuilder struct {
	language *Language
}

func NewLanguageBuilder() *LanguageBuilder {
	return &LanguageBuilder{
		language: &Language{},
	}
}

func (b *LanguageBuilder) SetLang(lang string) *LanguageBuilder {
	b.language.Lang = lang
	return b
}

func (b *LanguageBuilder) SetText(text string) *LanguageBuilder {
	b.language.Text = text
	return b
}

func (b *LanguageBuilder) Build() *Language {
	return b.language
}
