package anchor

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type (
	OptsFn func(*Anchor)
	Anchor struct {
		ID         string
		ClassNames []string
		Attributes templ.Attributes

		Href   string
		Target Target
	}
	Target string
)

const (
	TargetSelf   Target = "_self"
	TargetParent Target = "_parent"
	TargetTop    Target = "_top"
	TargetBlank  Target = "_blank"
)

func (c Anchor) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithID(id string) OptsFn {
	return func(element *Anchor) {
		element.ID = id
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *Anchor) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptsFn {
	return func(element *Anchor) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}

func WithHref(href string) OptsFn {
	return func(element *Anchor) {
		element.Href = href
	}
}

func WithTarget(target Target) OptsFn {
	return func(element *Anchor) {
		element.Target = target
	}
}

func New(opts ...OptsFn) *Anchor {
	c := &Anchor{}
	for _, opt := range opts {
		opt(c)
	}

	return c
}
