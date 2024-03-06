package inputNumber

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type (
	OptsFn      func(*InputNumber)
	InputNumber struct {
		ID         string
		ClassNames []string
		Attributes templ.Attributes

		Disabled    bool
		Required    bool
		Value       *int
		Placeholder string
		Name        string
		Min         *int
		Max         *int
		Step        *int
	}
)

func (c InputNumber) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithID(id string) OptsFn {
	return func(element *InputNumber) {
		element.ID = id
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *InputNumber) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptsFn {
	return func(element *InputNumber) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}

func WithDisabled() OptsFn {
	return func(element *InputNumber) {
		element.Disabled = true
	}
}

func WithName(name string) OptsFn {
	return func(element *InputNumber) {
		element.Name = name
	}
}

func WithRequired() OptsFn {
	return func(element *InputNumber) {
		element.Required = true
	}
}

func WithValue(value int) OptsFn {
	return func(element *InputNumber) {
		element.Value = &value
	}
}

func WithPlaceholder(placeholder string) OptsFn {
	return func(element *InputNumber) {
		element.Placeholder = placeholder
	}
}

func WithMin(min int) OptsFn {
	return func(element *InputNumber) {
		element.Min = &min
	}
}

func WithMax(max int) OptsFn {
	return func(element *InputNumber) {
		element.Max = &max
	}
}

func WithStep(step int) OptsFn {
	return func(element *InputNumber) {
		element.Step = &step
	}
}

func New(opts ...OptsFn) InputNumber {
	var c InputNumber
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
