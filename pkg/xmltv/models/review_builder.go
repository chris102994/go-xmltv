package models

type ReviewBuilderInterface interface {
	NewReviewBuilder() *ReviewBuilder
	SetType(reviewType string) *ReviewBuilder
	SetSource(source string) *ReviewBuilder
	SetReviewer(reviewer string) *ReviewBuilder
	SetLang(lang string) *ReviewBuilder
	SetText(text string) *ReviewBuilder
	Build() *Review
}

type ReviewBuilder struct {
	review *Review
}

func NewReviewBuilder() *ReviewBuilder {
	return &ReviewBuilder{
		review: &Review{},
	}
}

func (b *ReviewBuilder) SetType(reviewType string) *ReviewBuilder {
	b.review.Type = reviewType
	return b
}

func (b *ReviewBuilder) SetSource(source string) *ReviewBuilder {
	b.review.Source = source
	return b
}

func (b *ReviewBuilder) SetReviewer(reviewer string) *ReviewBuilder {
	b.review.Reviewer = reviewer
	return b
}

func (b *ReviewBuilder) SetLang(lang string) *ReviewBuilder {
	b.review.Lang = lang
	return b
}

func (b *ReviewBuilder) SetText(text string) *ReviewBuilder {
	b.review.Text = text
	return b
}

func (b *ReviewBuilder) Build() *Review {
	return b.review
}
