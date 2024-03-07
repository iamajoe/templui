package inputnumber

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

type Opt interface {
	InputNumberOpt()
}

type OptFn func(*InputNumber)

func (OptFn) InputNumberOpt() {}

func New(opts ...Opt) InputNumber {
	var c InputNumber
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

var WithID = templui.WithID

var WithClasses = templui.WithClasses

var WithAttributes = templui.WithAttributes

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
