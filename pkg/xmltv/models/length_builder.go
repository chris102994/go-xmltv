package models

type LengthBuilderInterface interface {
	NewLengthBuilder() *LengthBuilder
	SetUnits(units string) *LengthBuilder
	SetText(text string) *LengthBuilder
	Build() *Length
}

type LengthBuilder struct {
	length *Length
}

func NewLengthBuilder() *LengthBuilder {
	return &LengthBuilder{
		length: &Length{},
	}
}

func (b *LengthBuilder) SetUnits(units string) *LengthBuilder {
	b.length.Units = units
	return b
}

func (b *LengthBuilder) SetText(text string) *LengthBuilder {
	b.length.Text = text
	return b
}

func (b *LengthBuilder) Build() *Length {
	return b.length
}
