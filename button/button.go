package button

import (
	"context"
	"io"
)

type Kind string

const (
	KindSubmit Kind = "submit"
	KindButton Kind = "button"
)

type Button struct {
	Disabled   bool
	ClassNames []string
	Kind       Kind
}

func (b *Button) WithDisabled() *Button {
	b.Disabled = true
	return b
}

func (b *Button) WithKind(kind Kind) *Button {
	b.Kind = kind
	return b
}

func (b *Button) WithClasses(classes ...string) *Button {
	b.ClassNames = append(b.ClassNames, classes...)
	return b
}

func (b *Button) WithCustom(fn func(*Button) error) *Button {
	// TODO: what to do with the error?
	_ = fn(b)
	return b
}

func (b *Button) Render(ctx context.Context, w io.Writer) error {
	return render(*b).Render(ctx, w)
}

func New() *Button {
	return &Button{}
}
