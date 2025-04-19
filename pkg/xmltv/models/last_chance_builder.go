package models

type LastChanceBuilderInterface interface {
	NewLastChanceBuilder() *LastChanceBuilder
	SetLang(lang string) *LastChanceBuilder
	SetText(text string) *LastChanceBuilder
	Build() *LastChance
}

type LastChanceBuilder struct {
	lastChance *LastChance
}

func NewLastChanceBuilder() *LastChanceBuilder {
	return &LastChanceBuilder{
		lastChance: &LastChance{},
	}
}

func (b *LastChanceBuilder) SetLang(lang string) *LastChanceBuilder {
	b.lastChance.Lang = lang
	return b
}

func (b *LastChanceBuilder) SetText(text string) *LastChanceBuilder {
	b.lastChance.Text = text
	return b
}

func (b *LastChanceBuilder) Build() *LastChance {
	return b.lastChance
}
