package models

type URLBuilderInterface interface {
	NewURLBuilder() *URLBuilder
	SetSystem(system string) *URLBuilder
	SetText(text string) *URLBuilder
	Build() *URL
}

type URLBuilder struct {
	url *URL
}

func NewURLBuilder() *URLBuilder {
	return &URLBuilder{
		url: &URL{},
	}
}

func (b *URLBuilder) SetSystem(system string) *URLBuilder {
	b.url.System = system
	return b
}

func (b *URLBuilder) SetText(text string) *URLBuilder {
	b.url.Text = text
	return b
}

func (b *URLBuilder) Build() *URL {
	return b.url
}
