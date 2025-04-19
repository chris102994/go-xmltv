package models

type TVBuilderInterface interface {
	NewTVBuilder() *TVBuilder
	SetDate(date string) *TVBuilder
	SetSourceInfoURL(url string) *TVBuilder
	SetSourceInfoName(name string) *TVBuilder
	SetSourceDataURL(url string) *TVBuilder
	SetGeneratorInfoName(name string) *TVBuilder
	SetGeneratorInfoURL(url string) *TVBuilder
	AddChannel(channel *Channel) *TVBuilder
	AddProgramme(programme *Programme) *TVBuilder
	Build() *TV
}

type TVBuilder struct {
	tv *TV
}

func NewTVBuilder() *TVBuilder {
	return &TVBuilder{
		tv: &TV{
			Channels:   make([]*Channel, 0),
			Programmes: make([]*Programme, 0),
		},
	}
}

func (b *TVBuilder) SetDate(date string) *TVBuilder {
	b.tv.Date = date
	return b
}

func (b *TVBuilder) SetSourceInfoURL(url string) *TVBuilder {
	b.tv.SourceInfoURL = url
	return b
}

func (b *TVBuilder) SetSourceInfoName(name string) *TVBuilder {
	b.tv.SourceInfoName = name
	return b
}

func (b *TVBuilder) SetSourceDataURL(url string) *TVBuilder {
	b.tv.SourceDataURL = url
	return b
}

func (b *TVBuilder) SetGeneratorInfoName(name string) *TVBuilder {
	b.tv.GeneratorInfoName = name
	return b
}

func (b *TVBuilder) SetGeneratorInfoURL(url string) *TVBuilder {
	b.tv.GeneratorInfoURL = url
	return b
}

func (b *TVBuilder) AddChannel(channel *Channel) *TVBuilder {
	b.tv.Channels = append(b.tv.Channels, channel)
	return b
}

func (b *TVBuilder) AddProgramme(programme *Programme) *TVBuilder {
	b.tv.Programmes = append(b.tv.Programmes, programme)
	return b
}

func (b *TVBuilder) Build() *TV {
	return b.tv
}
