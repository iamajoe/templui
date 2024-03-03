package button

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestButton(t *testing.T) {
	cases := [](struct {
		comp         func() *Button
		assertTarget string
	}){
		{
			comp: func() *Button {
				return New().
					WithKind(KindSubmit).
					WithDisabled().
					WithClasses("foo").
					WithFunction(func(p *Button) error {
						p.ClassNames = append(p.ClassNames, "bar")
						return nil
					})
			},
			assertTarget: `button.foo.bar[type="submit"][disabled]`,
		}, {
			comp: func() *Button {
				return New(
					WithKind(KindSubmit),
					WithDisabled(),
					WithClasses("foo"),
					func(p *Button) error {
						p.ClassNames = append(p.ClassNames, "bar")
						return nil
					},
				)
			},
			assertTarget: `button.foo.bar[type="submit"][disabled]`,
		},
	}

	for _, test := range cases {
		comp := test.comp()

		r, w := io.Pipe()
		go func() {
			comp.Render(context.Background(), w)
			w.Close()
		}()

		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			t.Fatalf("failed to read template: %v", err)
		}

		docHTML, err := doc.Html()
		if err != nil {
			t.Fatalf("failed to return html: %v", err)
		}

		if doc.Find(test.assertTarget).Length() == 0 {
			t.Fatalf("expected button to render with attributes: %s", docHTML)
		}
	}
}
