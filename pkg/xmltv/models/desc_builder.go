package models

type DescBuilderInterface interface {
	NewDescBuilder() *DescBuilder
	SetLang(lang string) *DescBuilder
	SetText(text string) *DescBuilder
	Build() *Desc
}

type DescBuilder struct {
	desc *Desc
}

func NewDescBuilder() *DescBuilder {
	return &DescBuilder{
		desc: &Desc{},
	}
}

func (b *DescBuilder) SetLang(lang string) *DescBuilder {
	b.desc.Lang = lang
	return b
}

func (b *DescBuilder) SetText(text string) *DescBuilder {
	b.desc.Text = text
	return b
}

func (b *DescBuilder) Build() *Desc {
	return b.desc
}
