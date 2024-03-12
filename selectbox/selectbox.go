package selectbox

import (
	"context"
	"io"

	"github.com/iamajoe/templui/formelement"
)

//go:generate go run ../_generate/element.go -tmpl='formelement/elementopts.go' -tmpl-package=formelement -package=selectbox -struct=SelectBox

// TODO: can we use WithValue instead WithSelected?!

type SelectItem struct {
	Value    string
	Label    string
	Selected bool
}

type SelectBox struct {
	formelement.Element

	Items []SelectItem
	Value string

	// TODO: implement with a js lib like choices
	// TODO: how to change choices styling here?!
	// enableHTMLLabel bool
	// enableSearch    bool
}

func (c SelectBox) Render(ctx context.Context, w io.Writer) error {
	return render(c).Render(ctx, w)
}

func WithItems(items []SelectItem) OptFn {
	return func(element *SelectBox) {
		if len(element.Value) == 0 {
			element.Items = items
			return
		}

		// we want to make sure the selected is in tandem, as such
		// use the Selected on the actual component as the source
		// of truth
		newItems := []SelectItem{}
		for _, item := range items {
			item.Selected = item.Value == element.Value
			newItems = append(newItems, item)
		}
		element.Items = newItems
	}
}

func WithValue(value string) OptFn {
	return func(element *SelectBox) {
		element.Value = value

		// make the items selected in tandem with the key
		fn := WithItems(element.Items)
		fn(element)
	}
}

func New(opts ...OptFn) SelectBox {
	var c SelectBox
	for _, opt := range opts {
		opt(&c)
	}

	return c
}
