package inputtext

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

//go:generate go run ../_generate/element.go -package=inputtext -struct=InputText
type InputText struct {
	templui.Element

	Disabled    bool
	Required    bool
	Value       *int
	Placeholder string
	Name        string
}

func (c InputText) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithDisabled() OptFn {
	return func(element *InputText) {
		element.Disabled = true
	}
}

func WithName(name string) OptFn {
	return func(element *InputText) {
		element.Name = name
	}
}

func WithRequired() OptFn {
	return func(element *InputText) {
		element.Required = true
	}
}

func WithValue(value int) OptFn {
	return func(element *InputText) {
		element.Value = &value
	}
}

func WithPlaceholder(placeholder string) OptFn {
	return func(element *InputText) {
		element.Placeholder = placeholder
	}
}

func New(opts ...OptFn) InputText {
	var c InputText
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
