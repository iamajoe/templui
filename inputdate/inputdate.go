package inputdate

import (
	"context"
	"io"
	"time"

	"github.com/iamajoe/templui/formelement"
)

//go:generate go run ../_generate/element.go -tmpl='formelement/elementopts.go' -tmpl-package=formelement -package=inputdate -struct=InputDate
type InputDate struct {
	formelement.Element

	Value *time.Time
	Min   *time.Time
	Max   *time.Time
}

func (c InputDate) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithValue(date time.Time) OptFn {
	return func(element *InputDate) {
		element.Value = &date
	}
}

func WithMin(date time.Time) OptFn {
	return func(element *InputDate) {
		element.Min = &date
	}
}

func WithMax(date time.Time) OptFn {
	return func(element *InputDate) {
		element.Max = &date
	}
}

func New(opts ...OptFn) InputDate {
	var c InputDate
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
