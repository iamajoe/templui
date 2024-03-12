package button

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

//go:generate go run ../_generate/element.go -package=button -struct=Button
type Button struct {
	templui.Element

	Disabled bool
	Kind     Kind
}

type Kind string

const (
	KindSubmit Kind = "submit"
	KindButton Kind = "button"
)

func (c Button) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithDisabled() OptFn {
	return func(element *Button) {
		element.Disabled = true
	}
}

func WithKind(kind Kind) OptFn {
	return func(element *Button) {
		element.Kind = kind
	}
}

func New(opts ...OptFn) Button {
	var c Button
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
