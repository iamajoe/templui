package inputtext

import (
	"context"
	"io"

	"github.com/iamajoe/templui/formelement"
)

//go:generate go run ../_generate/element.go -tmpl='formelement/elementopts.go' -tmpl-package=formelement -package=inputtext -struct=InputText
type InputText struct {
	formelement.Element

	Value *int
}

func (c InputText) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithValue(value int) OptFn {
	return func(element *InputText) {
		element.Value = &value
	}
}

func New(opts ...OptFn) InputText {
	var c InputText
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
