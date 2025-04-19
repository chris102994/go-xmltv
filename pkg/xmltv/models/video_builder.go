package models

type VideoBuilderInterface interface {
	NewVideoBuilder() *VideoBuilder
	SetPresent(present string) *VideoBuilder
	SetColour(colour string) *VideoBuilder
	SetAspect(aspect string) *VideoBuilder
	SetQuality(quality string) *VideoBuilder
	Build() *Video
}

type VideoBuilder struct {
	video *Video
}

func NewVideoBuilder() *VideoBuilder {
	return &VideoBuilder{
		video: &Video{},
	}
}

func (b *VideoBuilder) SetPresent(present string) *VideoBuilder {
	b.video.Present = present
	return b
}

func (b *VideoBuilder) SetColour(colour string) *VideoBuilder {
	b.video.Colour = colour
	return b
}

func (b *VideoBuilder) SetAspect(aspect string) *VideoBuilder {
	b.video.Aspect = aspect
	return b
}

func (b *VideoBuilder) SetQuality(quality string) *VideoBuilder {
	b.video.Quality = quality
	return b
}

func (b *VideoBuilder) Build() *Video {
	return b.video
}
