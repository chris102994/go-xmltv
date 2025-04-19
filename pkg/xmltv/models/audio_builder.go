package models

type AudioBuilderInterface interface {
	NewAudioBuilder() *AudioBuilder
	SetPresent(present string) *AudioBuilder
	SetStereo(stereo string) *AudioBuilder
	Build() *Audio
}

type AudioBuilder struct {
	audio *Audio
}

func NewAudioBuilder() *AudioBuilder {
	return &AudioBuilder{
		audio: &Audio{},
	}
}

func (b *AudioBuilder) SetPresent(present string) *AudioBuilder {
	b.audio.Present = present
	return b
}

func (b *AudioBuilder) SetStereo(stereo string) *AudioBuilder {
	b.audio.Stereo = stereo
	return b
}

func (b *AudioBuilder) Build() *Audio {
	return b.audio
}
