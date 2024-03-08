package table

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/iamajoe/templui"
)

//go:generate go run ../_generate/element.go -package=table -struct=Table
type Table struct {
	templui.Element

	HeaderRow *templ.Component
}

func (c Table) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithHeader(headerRow templ.Component) OptFn {
	return func(c *Table) {
		c.HeaderRow = &headerRow
	}
}

func New(opts ...OptFn) Table {
	var c Table
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
