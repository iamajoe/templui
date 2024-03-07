package inputdate

import (
	"context"
	"io"
	"time"

	"github.com/iamajoe/templui"
)

type Opt interface {
	InputDateOpt()
}

type OptFn func(*InputDate)

func (OptFn) InputDateOpt() {}

func New(opts ...Opt) InputDate {
	var c InputDate
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

type InputDate struct {
	templui.Element

	Disabled    bool
	Required    bool
	Name        string
	Placeholder string
	Value       *time.Time
	Min         *time.Time
	Max         *time.Time
}

func (c InputDate) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

var WithID = templui.WithID

var WithClasses = templui.WithClasses

var WithAttributes = templui.WithAttributes

func WithDisabled() OptFn {
	return func(element *InputDate) {
		element.Disabled = true
	}
}

func WithName(name string) OptFn {
	return func(element *InputDate) {
		element.Name = name
	}
}

func WithRequired() OptFn {
	return func(element *InputDate) {
		element.Required = true
	}
}

func WithValue(date time.Time) OptFn {
	return func(element *InputDate) {
		element.Value = &date
	}
}

func WithPlaceholder(placeholder string) OptFn {
	return func(element *InputDate) {
		element.Placeholder = placeholder
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
