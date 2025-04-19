package models

type EpisodeNumBuilderInterface interface {
	NewEpisodeNumBuilder() *EpisodeNumBuilder
	SetSystem(system string) *EpisodeNumBuilder
	SetText(text string) *EpisodeNumBuilder
	Build() *EpisodeNum
}

type EpisodeNumBuilder struct {
	episodeNum *EpisodeNum
}

func NewEpisodeNumBuilder() *EpisodeNumBuilder {
	return &EpisodeNumBuilder{
		episodeNum: &EpisodeNum{},
	}
}

func (b *EpisodeNumBuilder) SetSystem(system string) *EpisodeNumBuilder {
	b.episodeNum.System = system
	return b
}

func (b *EpisodeNumBuilder) SetText(text string) *EpisodeNumBuilder {
	b.episodeNum.Text = text
	return b
}

func (b *EpisodeNumBuilder) Build() *EpisodeNum {
	return b.episodeNum
}
