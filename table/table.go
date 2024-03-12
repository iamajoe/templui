package table

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

//go:generate go run ../_generate/element.go -package=table -struct=Table
type Table struct {
	templui.Element

	HeaderRow     []string
	HeaderClasses []string

	Rows       [][]string
	RowClasses []string
}

func (c Table) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithHeader(headerRow ...string) OptFn {
	return func(c *Table) {
		c.HeaderRow = headerRow
	}
}

func WithHeaderClasses(classes ...string) OptFn {
	return func(c *Table) {
		c.HeaderClasses = classes
	}
}

func WithRows(rows ...[]string) OptFn {
	return func(c *Table) {
		c.Rows = rows
	}
}

func WithRowClasses(classes ...string) OptFn {
	return func(c *Table) {
		c.RowClasses = classes
	}
}

func New(opts ...OptFn) Table {
	var c Table
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
