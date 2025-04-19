package models

type DisplayNameBuilderInterface interface {
	NewDisplayNameBuilder() *DisplayNameBuilder
	SetLang(lang string) *DisplayNameBuilder
	SetText(text string) *DisplayNameBuilder
	Build() *DisplayName
}

type DisplayNameBuilder struct {
	displayName *DisplayName
}

func NewDisplayNameBuilder() *DisplayNameBuilder {
	return &DisplayNameBuilder{
		displayName: &DisplayName{},
	}
}

func (b *DisplayNameBuilder) SetLang(lang string) *DisplayNameBuilder {
	b.displayName.Lang = lang
	return b
}

func (b *DisplayNameBuilder) SetText(text string) *DisplayNameBuilder {
	b.displayName.Text = text
	return b
}

func (b *DisplayNameBuilder) Build() *DisplayName {
	return b.displayName
}
