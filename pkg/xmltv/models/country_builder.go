package models

type CountryBuilderInterface interface {
	NewCountryBuilder() *CountryBuilder
	SetLang(lang string) *CountryBuilder
	SetText(text string) *CountryBuilder
	Build() *Country
}

type CountryBuilder struct {
	country *Country
}

func NewCountryBuilder() *CountryBuilder {
	return &CountryBuilder{
		country: &Country{},
	}
}

func (b *CountryBuilder) SetLang(lang string) *CountryBuilder {
	b.country.Lang = lang
	return b
}

func (b *CountryBuilder) SetText(text string) *CountryBuilder {
	b.country.Text = text
	return b
}

func (b *CountryBuilder) Build() *Country {
	return b.country
}
