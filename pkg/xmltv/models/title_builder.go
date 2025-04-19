package models

type TitleBuilderInterface interface {
	NewTitleBuilder() *TitleBuilder
	SetLang(lang string) *TitleBuilder
	SetText(text string) *TitleBuilder
	Build() *Title
}

type TitleBuilder struct {
	title *Title
}

func NewTitleBuilder() *TitleBuilder {
	return &TitleBuilder{
		title: &Title{},
	}
}

func (b *TitleBuilder) SetLang(lang string) *TitleBuilder {
	b.title.Lang = lang
	return b
}

func (b *TitleBuilder) SetText(text string) *TitleBuilder {
	b.title.Text = text
	return b
}

func (b *TitleBuilder) Build() *Title {
	return b.title
}
