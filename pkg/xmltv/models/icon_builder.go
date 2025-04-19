package models

type IconBuilderInterface interface {
	NewIconBuilder() *IconBuilder
	SetSrc(src string) *IconBuilder
	SetWidth(width string) *IconBuilder
	SetHeight(height string) *IconBuilder
	Build() *Icon
}

type IconBuilder struct {
	icon *Icon
}

func NewIconBuilder() *IconBuilder {
	return &IconBuilder{
		icon: &Icon{},
	}
}

func (b *IconBuilder) SetSrc(src string) *IconBuilder {
	b.icon.Src = src
	return b
}

func (b *IconBuilder) SetWidth(width string) *IconBuilder {
	b.icon.Width = width
	return b
}

func (b *IconBuilder) SetHeight(height string) *IconBuilder {
	b.icon.Height = height
	return b
}

func (b *IconBuilder) Build() *Icon {
	return b.icon
}
