package models

type ChannelBuilderInterface interface {
	NewChannelBuilder() *ChannelBuilder
	SetID(id string) *ChannelBuilder
	AddDisplayName(displayName *DisplayName) *ChannelBuilder
	AddIcon(icon *Icon) *ChannelBuilder
	AddURL(url *URL) *ChannelBuilder
	Build() *Channel
}

type ChannelBuilder struct {
	channel *Channel
}

func NewChannelBuilder() *ChannelBuilder {
	return &ChannelBuilder{
		channel: &Channel{
			DisplayNames: make([]*DisplayName, 0),
			Icons:        make([]*Icon, 0),
			URLs:         make([]*URL, 0),
		},
	}
}

func (b *ChannelBuilder) SetID(id string) *ChannelBuilder {
	b.channel.ID = id
	return b
}

func (b *ChannelBuilder) AddDisplayName(displayName *DisplayName) *ChannelBuilder {
	b.channel.DisplayNames = append(b.channel.DisplayNames, displayName)
	return b
}

func (b *ChannelBuilder) AddIcon(icon *Icon) *ChannelBuilder {
	b.channel.Icons = append(b.channel.Icons, icon)
	return b
}

func (b *ChannelBuilder) AddURL(url *URL) *ChannelBuilder {
	b.channel.URLs = append(b.channel.URLs, url)
	return b
}

func (b *ChannelBuilder) Build() *Channel {
	return b.channel
}
