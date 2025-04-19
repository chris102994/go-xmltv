package models

type StarRatingBuilderInterface interface {
	NewStarRatingBuilder() *StarRatingBuilder
	SetSystem(system string) *StarRatingBuilder
	SetValue(value string) *StarRatingBuilder
	SetIcon(icon *Icon) *StarRatingBuilder
	Build() *StarRating
}

type StarRatingBuilder struct {
	starRating *StarRating
}

func NewStarRatingBuilder() *StarRatingBuilder {
	return &StarRatingBuilder{
		starRating: &StarRating{},
	}
}

func (b *StarRatingBuilder) SetSystem(system string) *StarRatingBuilder {
	b.starRating.System = system
	return b
}

func (b *StarRatingBuilder) SetValue(value string) *StarRatingBuilder {
	b.starRating.Value = value
	return b
}

func (b *StarRatingBuilder) SetIcon(icon *Icon) *StarRatingBuilder {
	b.starRating.Icon = icon
	return b
}

func (b *StarRatingBuilder) Build() *StarRating {
	return b.starRating
}
