package button

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

type Opt interface {
	ButtonOpt()
}

type OptFn func(*Button)

func (OptFn) ButtonOpt() {}

type Kind string

const (
	KindSubmit Kind = "submit"
	KindButton Kind = "button"
)

func New(opts ...Opt) Button {
	var c Button
	for _, opt := range opts {
		switch fn := opt.(type) {
		case templui.OptFn:
			fn(&c.Element)
		case OptFn:
			fn(&c)
		}
	}

	return c
}

type Button struct {
	templui.Element

	Disabled bool
	Kind     Kind
}

func (c Button) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

var WithID = templui.WithID

var WithClasses = templui.WithClasses

var WithAttributes = templui.WithAttributes

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
