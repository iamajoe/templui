package anchor

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type Target string

const (
	TargetSelf   Target = "_self"
	TargetParent Target = "_parent"
	TargetTop    Target = "_top"
	TargetBlank  Target = "_blank"
)

type (
	OptsFn func(*Anchor) error
	Anchor struct {
		ID         string
		ClassNames []string
		Attributes templ.Attributes

		Href   string
		Target Target

		Opts []OptsFn
	}
)

func WithID(id string) OptsFn {
	return func(element *Anchor) error {
		element.ID = id
		return nil
	}
}

func WithClasses(classes ...string) OptsFn {
	return func(element *Anchor) error {
		element.ClassNames = append(element.ClassNames, classes...)
		return nil
	}
}

func WithAttributes(attributes map[string]any) OptsFn {
	return func(element *Anchor) error {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}

		return nil
	}
}

func WithHref(href string) OptsFn {
	return func(element *Anchor) error {
		element.Href = href
		return nil
	}
}

func WithTarget(target Target) OptsFn {
	return func(element *Anchor) error {
		element.Target = target
		return nil
	}
}

func (c *Anchor) WithFunction(fn func(*Anchor) error) *Anchor {
	c.Opts = append(c.Opts, fn)
	return c
}

func (c *Anchor) WithID(id string) *Anchor {
	c.Opts = append(c.Opts, WithID(id))
	return c
}

func (c *Anchor) WithClasses(classes ...string) *Anchor {
	c.Opts = append(c.Opts, WithClasses(classes...))
	return c
}

func (c *Anchor) WithAttributes(attributes map[string]any) *Anchor {
	c.Opts = append(c.Opts, WithAttributes(attributes))
	return c
}

func (c *Anchor) WithHref(href string) *Anchor {
	c.Opts = append(c.Opts, WithHref(href))
	return c
}

func (c *Anchor) WithTarget(target Target) *Anchor {
	c.Opts = append(c.Opts, WithTarget(target))
	return c
}

func (c *Anchor) Render(ctx context.Context, w io.Writer) error {
	for _, opt := range c.Opts {
		if err := opt(c); err != nil {
			// TODO: need to figure how to handle errors
			return templ.NopComponent.Render(ctx, w)
		}
	}

	return render(*c).Render(ctx, w)
}

func New(opts ...OptsFn) *Anchor {
	return &Anchor{
		Opts: opts,
	}
}
