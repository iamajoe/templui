package button

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type Kind string

const (
	KindSubmit Kind = "submit"
	KindButton Kind = "button"
)

type (
	OptsFn func(*Button) error
	Button struct {
		ID         string
		ClassNames []string
		Disabled   bool
		Kind       Kind

		Opts []OptsFn
	}
)

func WithID(id string) OptsFn {
	return func(element *Button) error {
		element.ID = id
		return nil
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *Button) error {
		element.ClassNames = append(element.ClassNames, classes...)
		return nil
	}
}

func WithDisabled() OptsFn {
	return func(element *Button) error {
		element.Disabled = true
		return nil
	}
}

func WithKind(kind Kind) OptsFn {
	return func(element *Button) error {
		element.Kind = kind
		return nil
	}
}

func (b *Button) WithFunction(fn func(*Button) error) *Button {
	b.Opts = append(b.Opts, fn)
	return b
}

func (b *Button) WithID(id string) *Button {
	b.Opts = append(b.Opts, WithID(id))
	return b
}

func (b *Button) WithClasses(classes ...string) *Button {
	b.Opts = append(b.Opts, WithClasses(classes...))
	return b
}

func (b *Button) WithDisabled() *Button {
	b.Opts = append(b.Opts, WithDisabled())
	return b
}

func (b *Button) WithKind(kind Kind) *Button {
	b.Opts = append(b.Opts, WithDisabled())
	return b
}

func (b *Button) Render(ctx context.Context, w io.Writer) error {
	for _, opt := range b.Opts {
		if err := opt(b); err != nil {
			// TODO: need to figure how to handle errors
			return templ.NopComponent.Render(ctx, w)
		}
	}

	return render(*b).Render(ctx, w)
}

func New(opts ...OptsFn) *Button {
	return &Button{
		Opts: opts,
	}
}
