package tablerow

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

//go:generate go run ../../_generate/element.go -package=tablerow -struct=TableRow
type TableRow struct {
	templui.Element
}

func (c TableRow) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func New(opts ...OptFn) TableRow {
	var c TableRow
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
