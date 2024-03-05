package selectbox

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

// TODO: need to setup tests

type (
	OptsFn     func(*SelectBox)
	SelectItem struct {
		Value    string
		Label    string
		Selected bool
	}
	SelectBox struct {
		ID         string
		Name       string
		ClassNames []string
		Attributes templ.Attributes

		Items    []SelectItem
		Selected string

		Disabled    bool
		Required    bool
		Placeholder string

		// TODO: implement with a js lib like choices
		// TODO: how to change choices styling here?!
		enableHTMLLabel bool
		enableSearch    bool
	}
)

func (c SelectBox) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithID(id string) OptsFn {
	return func(element *SelectBox) {
		element.ID = id
	}
}

func WithName(name string) OptsFn {
	return func(element *SelectBox) {
		element.Name = name
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *SelectBox) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptsFn {
	return func(element *SelectBox) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}

func WithDisabled() OptsFn {
	return func(element *SelectBox) {
		element.Disabled = true
	}
}

func WithRequired() OptsFn {
	return func(element *SelectBox) {
		element.Required = true
	}
}

func WithPlaceholder(placeholder string) OptsFn {
	return func(element *SelectBox) {
		element.Placeholder = placeholder
	}
}

func WithItems(items []SelectItem) OptsFn {
	return func(element *SelectBox) {
		element.Items = items
	}
}

func WithSelected(selected string) OptsFn {
	return func(element *SelectBox) {
		element.Selected = selected
	}
}

func New(opts ...OptsFn) SelectBox {
	var c SelectBox
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
