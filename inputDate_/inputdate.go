package inputdate

import (
	"context"
	"io"
	"time"

	"github.com/a-h/templ"
)

type (
	OptsFn    func(*InputDate)
	InputDate struct {
		ID         string
		ClassNames []string
		Attributes templ.Attributes

		Disabled    bool
		Required    bool
		Name        string
		Placeholder string
		Value       *time.Time
		Min         *time.Time
		Max         *time.Time
	}
)

func (c InputDate) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithID(id string) OptsFn {
	return func(element *InputDate) {
		element.ID = id
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *InputDate) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptsFn {
	return func(element *InputDate) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}

func WithDisabled() OptsFn {
	return func(element *InputDate) {
		element.Disabled = true
	}
}

func WithName(name string) OptsFn {
	return func(element *InputDate) {
		element.Name = name
	}
}

func WithRequired() OptsFn {
	return func(element *InputDate) {
		element.Required = true
	}
}

func WithValue(date time.Time) OptsFn {
	return func(element *InputDate) {
		element.Value = &date
	}
}

func WithPlaceholder(placeholder string) OptsFn {
	return func(element *InputDate) {
		element.Placeholder = placeholder
	}
}

func WithMin(date time.Time) OptsFn {
	return func(element *InputDate) {
		element.Min = &date
	}
}

func WithMax(date time.Time) OptsFn {
	return func(element *InputDate) {
		element.Max = &date
	}
}

func New(opts ...OptsFn) InputDate {
	var c InputDate
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
