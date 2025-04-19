package models

type SubTitleBuilderInterface interface {
	NewSubTitleBuilder() *SubTitleBuilder
	SetLang(lang string) *SubTitleBuilder
	SetText(text string) *SubTitleBuilder
	Build() *SubTitle
}

type SubTitleBuilder struct {
	subTitle *SubTitle
}

func NewSubTitleBuilder() *SubTitleBuilder {
	return &SubTitleBuilder{
		subTitle: &SubTitle{},
	}
}

func (b *SubTitleBuilder) SetLang(lang string) *SubTitleBuilder {
	b.subTitle.Lang = lang
	return b
}

func (b *SubTitleBuilder) SetText(text string) *SubTitleBuilder {
	b.subTitle.Text = text
	return b
}

func (b *SubTitleBuilder) Build() *SubTitle {
	return b.subTitle
}
