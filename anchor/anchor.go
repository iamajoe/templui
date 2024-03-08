package anchor

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

type Target string

const (
	TargetSelf   Target = "_self"
	TargetParent Target = "_parent"
	TargetTop    Target = "_top"
	TargetBlank  Target = "_blank"
)

//go:generate go run ../_generate/element.go -package=anchor -struct=Anchor
type Anchor struct {
	templui.Element

	Href   string
	Target Target
}

func (c Anchor) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithHref(href string) OptFn {
	return func(element *Anchor) {
		element.Href = href
	}
}

func WithTarget(target Target) OptFn {
	return func(element *Anchor) {
		element.Target = target
	}
}

func New(opts ...OptFn) Anchor {
	var c Anchor
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
