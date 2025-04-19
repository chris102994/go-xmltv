package models

type ImageBuilderInterface interface {
	NewImageBuilder() *ImageBuilder
	SetType(imageType string) *ImageBuilder
	SetWidth(width string) *ImageBuilder
	SetHeight(height string) *ImageBuilder
	SetOrientation(orientation string) *ImageBuilder
	SetText(text string) *ImageBuilder
	Build() *Image
}

type ImageBuilder struct {
	image *Image
}

func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{
		image: &Image{},
	}
}

func (b *ImageBuilder) SetType(imageType string) *ImageBuilder {
	b.image.Type = imageType
	return b
}

func (b *ImageBuilder) SetWidth(width string) *ImageBuilder {
	b.image.Width = width
	return b
}

func (b *ImageBuilder) SetHeight(height string) *ImageBuilder {
	b.image.Height = height
	return b
}

func (b *ImageBuilder) SetOrientation(orientation string) *ImageBuilder {
	b.image.Orientation = orientation
	return b
}

func (b *ImageBuilder) SetText(text string) *ImageBuilder {
	b.image.Text = text
	return b
}

func (b *ImageBuilder) Build() *Image {
	return b.image
}
