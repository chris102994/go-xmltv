package models

type CreditsBuuilderInterface interface {
	NewCreditsBuilder() *CreditsBuilder
	AddDirector(director string) *CreditsBuilder
	AddActor(actor *Actor) *CreditsBuilder
	AddWriter(writer string) *CreditsBuilder
	AddAdapter(adapter string) *CreditsBuilder
	AddProducer(producer string) *CreditsBuilder
	AddComposer(composer string) *CreditsBuilder
	AddEditor(editor string) *CreditsBuilder
	AddPresenter(presenter string) *CreditsBuilder
	AddCommentator(commentator string) *CreditsBuilder
	AddGuest(guest string) *CreditsBuilder
	Build() *Credits
}

type CreditsBuilder struct {
	credits *Credits
}

func NewCreditsBuilder() *CreditsBuilder {
	return &CreditsBuilder{
		credits: &Credits{
			Directors:    make([]string, 0),
			Actors:       make([]*Actor, 0),
			Writers:      make([]string, 0),
			Adapters:     make([]string, 0),
			Producers:    make([]string, 0),
			Composers:    make([]string, 0),
			Editors:      make([]string, 0),
			Presenters:   make([]string, 0),
			Commentators: make([]string, 0),
			Guests:       make([]string, 0),
		},
	}
}

func (b *CreditsBuilder) AddDirector(director string) *CreditsBuilder {
	b.credits.Directors = append(b.credits.Directors, director)
	return b
}

func (b *CreditsBuilder) AddActor(actor *Actor) *CreditsBuilder {
	b.credits.Actors = append(b.credits.Actors, actor)
	return b
}

func (b *CreditsBuilder) AddWriter(writer string) *CreditsBuilder {
	b.credits.Writers = append(b.credits.Writers, writer)
	return b
}

func (b *CreditsBuilder) AddAdapter(adapter string) *CreditsBuilder {
	b.credits.Adapters = append(b.credits.Adapters, adapter)
	return b
}

func (b *CreditsBuilder) AddProducer(producer string) *CreditsBuilder {
	b.credits.Producers = append(b.credits.Producers, producer)
	return b
}

func (b *CreditsBuilder) AddComposer(composer string) *CreditsBuilder {
	b.credits.Composers = append(b.credits.Composers, composer)
	return b
}

func (b *CreditsBuilder) AddEditor(editor string) *CreditsBuilder {
	b.credits.Editors = append(b.credits.Editors, editor)
	return b
}

func (b *CreditsBuilder) AddPresenter(presenter string) *CreditsBuilder {
	b.credits.Presenters = append(b.credits.Presenters, presenter)
	return b
}

func (b *CreditsBuilder) AddCommentator(commentator string) *CreditsBuilder {
	b.credits.Commentators = append(b.credits.Commentators, commentator)
	return b
}

func (b *CreditsBuilder) AddGuest(guest string) *CreditsBuilder {
	b.credits.Guests = append(b.credits.Guests, guest)
	return b
}

func (b *CreditsBuilder) Build() *Credits {
	return b.credits
}
