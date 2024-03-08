package inputnumber

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

//go:generate go run ../_generate/element.go -package=inputnumber -struct=InputNumber
type InputNumber struct {
	templui.Element

	Disabled    bool
	Required    bool
	Value       *int
	Placeholder string
	Name        string
	Min         *int
	Max         *int
	Step        *int
}

func (c InputNumber) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithDisabled() OptFn {
	return func(element *InputNumber) {
		element.Disabled = true
	}
}

func WithName(name string) OptFn {
	return func(element *InputNumber) {
		element.Name = name
	}
}

func WithRequired() OptFn {
	return func(element *InputNumber) {
		element.Required = true
	}
}

func WithValue(value int) OptFn {
	return func(element *InputNumber) {
		element.Value = &value
	}
}

func WithPlaceholder(placeholder string) OptFn {
	return func(element *InputNumber) {
		element.Placeholder = placeholder
	}
}

func WithMin(min int) OptFn {
	return func(element *InputNumber) {
		element.Min = &min
	}
}

func WithMax(max int) OptFn {
	return func(element *InputNumber) {
		element.Max = &max
	}
}

func WithStep(step int) OptFn {
	return func(element *InputNumber) {
		element.Step = &step
	}
}

func New(opts ...OptFn) InputNumber {
	var c InputNumber
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
