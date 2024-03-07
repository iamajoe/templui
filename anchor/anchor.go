package anchor

import (
	"context"
	"io"

	"github.com/iamajoe/templui"
)

type Opt interface {
	AnchorOpt()
}

type OptFn func(*Anchor)

func (OptFn) AnchorOpt() {}

type Target string

const (
	TargetSelf   Target = "_self"
	TargetParent Target = "_parent"
	TargetTop    Target = "_top"
	TargetBlank  Target = "_blank"
)

func New(opts ...Opt) Anchor {
	var c Anchor
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

type Anchor struct {
	templui.Element

	Href   string
	Target Target
}

func (c Anchor) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

var WithID = templui.WithID

var WithClasses = templui.WithClasses

var WithAttributes = templui.WithAttributes

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
