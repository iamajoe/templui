package tablecell

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

//go:generate go run ../../_generate/element.go -package=tablecell -struct=TableCell
type TableCell struct {
	templui.Element

	IsHeader bool
}

func (c TableCell) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithIsHeader() OptFn {
	return func(c *TableCell) {
		c.IsHeader = true
	}
}

func New(opts ...OptFn) TableCell {
	var c TableCell
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
