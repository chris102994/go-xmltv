package models

type PreviouslyShownBuilderInterface interface {
	NewPreviouslyShownBuilder() *PreviouslyShownBuilder
	SetStart(start string) *PreviouslyShownBuilder
	SetChannel(channel string) *PreviouslyShownBuilder
	Build() *PreviouslyShown
}

type PreviouslyShownBuilder struct {
	previouslyShown *PreviouslyShown
}

func NewPreviouslyShownBuilder() *PreviouslyShownBuilder {
	return &PreviouslyShownBuilder{
		previouslyShown: &PreviouslyShown{},
	}
}

func (b *PreviouslyShownBuilder) SetStart(start string) *PreviouslyShownBuilder {
	b.previouslyShown.Start = start
	return b
}

func (b *PreviouslyShownBuilder) SetChannel(channel string) *PreviouslyShownBuilder {
	b.previouslyShown.Channel = channel
	return b
}

func (b *PreviouslyShownBuilder) Build() *PreviouslyShown {
	return b.previouslyShown
}
