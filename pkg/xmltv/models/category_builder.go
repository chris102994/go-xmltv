package models

type CategoryBuilderInterface interface {
	NewCategoryBuilder() *CategoryBuilder
	SetLang(lang string) *CategoryBuilder
	SetText(text string) *CategoryBuilder
	Build() *Category
}

type CategoryBuilder struct {
	category *Category
}

func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{
		category: &Category{},
	}
}

func (b *CategoryBuilder) SetLang(lang string) *CategoryBuilder {
	b.category.Lang = lang
	return b
}

func (b *CategoryBuilder) SetText(text string) *CategoryBuilder {
	b.category.Text = text
	return b
}

func (b *CategoryBuilder) Build() *Category {
	return b.category
}
