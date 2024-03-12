package inputnumber

import (
	"context"
	"io"

	"github.com/iamajoe/templui/formelement"
)

//go:generate go run ../_generate/element.go -tmpl='formelement/elementopts.go' -tmpl-package=formelement -package=inputnumber -struct=InputNumber
type InputNumber struct {
	formelement.Element

	Value *int
	Min   *int
	Max   *int
	Step  *int
}

func (c InputNumber) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithValue(value int) OptFn {
	return func(element *InputNumber) {
		element.Value = &value
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
