package models

type ActorBuilderInterface interface {
	NewActorBuilder() *ActorBuilder
	SetRole(role string) *ActorBuilder
	SetGuest(guest string) *ActorBuilder
	SetText(text string) *ActorBuilder
	SetImage(image *Image) *ActorBuilder
	SetURL(url *URL) *ActorBuilder
	Build() *Actor
}

type ActorBuilder struct {
	actor *Actor
}

func NewActorBuilder() *ActorBuilder {
	return &ActorBuilder{
		actor: &Actor{
			Image: nil,
			URL:   nil,
		},
	}
}

func (b *ActorBuilder) SetRole(role string) *ActorBuilder {
	b.actor.Role = role
	return b
}

func (b *ActorBuilder) SetGuest(guest string) *ActorBuilder {
	b.actor.Guest = guest
	return b
}

func (b *ActorBuilder) SetText(text string) *ActorBuilder {
	b.actor.Text = text
	return b
}

func (b *ActorBuilder) SetImage(image *Image) *ActorBuilder {
	b.actor.Image = image
	return b
}

func (b *ActorBuilder) SetURL(url *URL) *ActorBuilder {
	b.actor.URL = url
	return b
}

func (b *ActorBuilder) Build() *Actor {
	return b.actor
}
