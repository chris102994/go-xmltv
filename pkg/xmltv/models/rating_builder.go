package models

type RatingBuilderInterface interface {
	NewRatingBuilder() *RatingBuilder
	SetSystem(system string) *RatingBuilder
	SetValue(value string) *RatingBuilder
	SetIcon(icon *Icon) *RatingBuilder
	Build() *Rating
}

type RatingBuilder struct {
	rating *Rating
}

func NewRatingBuilder() *RatingBuilder {
	return &RatingBuilder{
		rating: &Rating{},
	}
}

func (b *RatingBuilder) SetSystem(system string) *RatingBuilder {
	b.rating.System = system
	return b
}

func (b *RatingBuilder) SetValue(value string) *RatingBuilder {
	b.rating.Value = value
	return b
}

func (b *RatingBuilder) SetIcon(icon *Icon) *RatingBuilder {
	b.rating.Icon = icon
	return b
}

func (b *RatingBuilder) Build() *Rating {
	return b.rating
}
