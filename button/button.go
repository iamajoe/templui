package button

import (
	"context"
	"io"
)

type ButtonType string

const (
	ButtonTypeSubmit ButtonType = "submit"
	ButtonTypeButton ButtonType = "button"
)

type Button struct {
	Disabled   bool
	ClassNames []string
	Type       ButtonType
}

func (b *Button) WithDisabled() *Button {
	b.Disabled = true
	return b
}

func (b *Button) WithType(kind ButtonType) *Button {
	b.Type = kind
	return b
}

func (b *Button) WithClasses(classes ...string) *Button {
	b.ClassNames = append(b.ClassNames, classes...)
	return b
}

func (b *Button) WithCustomFn(fn func(*Button) error) *Button {
	// TODO: what to do with the error?
	_ = fn(b)
	return b
}

func (b *Button) Render(ctx context.Context, w io.Writer) error {
	return buttonTempl(b)
}

func New() Button {
	return Button{}
}
