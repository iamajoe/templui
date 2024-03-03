package anchor

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name         string
		comp         func() Anchor
		assertTarget string
	}{
		{
			name: "creates",
			comp: func() Anchor {
				return New(
					WithID("zing"),
					WithClasses("foo"),
					WithAttributes(map[string]any{"data-zed": "zung"}),
					WithTarget(TargetBlank),
					WithHref("https://google.com"),
					func(el *Anchor) {
						el.ClassNames = append(el.ClassNames, "bar")
					},
				)
			},
			assertTarget: `a#zing.foo.bar[href="https://google.com"][target="_blank"][data-zed="zung"]`,
		},
	}

	for _, test := range cases {
		comp := test.comp()

		r, w := io.Pipe()
		go func() {
			comp.Render(context.Background(), w)
			w.Close()
		}()

		doc, _ := goquery.NewDocumentFromReader(r)
		if doc.Find(test.assertTarget).Length() == 0 {
			body, _ := doc.Find("body").Html()
			t.Fatalf(
				"%s: expected target \"%s\" to be in html \"%s\"",
				test.name,
				test.assertTarget,
				body,
			)
		}
	}
}
