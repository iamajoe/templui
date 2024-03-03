package button

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type (
	OptsFn func(*Button)
	Button struct {
		ID         string
		ClassNames []string
		Attributes templ.Attributes

		Disabled bool
		Kind     Kind
	}
	Kind string
)

const (
	KindSubmit Kind = "submit"
	KindButton Kind = "button"
)

func (c Button) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithID(id string) OptsFn {
	return func(element *Button) {
		element.ID = id
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *Button) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptsFn {
	return func(element *Button) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}

func WithDisabled() OptsFn {
	return func(element *Button) {
		element.Disabled = true
	}
}

func WithKind(kind Kind) OptsFn {
	return func(element *Button) {
		element.Kind = kind
	}
}

func New(opts ...OptsFn) Button {
	var c Button
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
