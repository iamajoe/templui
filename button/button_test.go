package button

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestNew(t *testing.T) {
	cases := [](struct {
		name         string
		comp         func() *Button
		assertTarget string
	}){
		{
			name: "through chain",
			comp: func() *Button {
				return New().
					WithID("zing").
					WithKind(KindSubmit).
					WithDisabled().
					WithClasses("foo").
					WithFunction(func(p *Button) error {
						p.ClassNames = append(p.ClassNames, "bar")
						return nil
					})
			},
			assertTarget: `button#zing.foo.bar[type="submit"][disabled]`,
		}, {
			name: "through argument spread",
			comp: func() *Button {
				return New(
					WithID("zing"),
					WithKind(KindSubmit),
					WithDisabled(),
					WithClasses("foo"),
					func(p *Button) error {
						p.ClassNames = append(p.ClassNames, "bar")
						return nil
					},
				)
			},
			assertTarget: `button#zing.foo.bar[type="submit"][disabled]`,
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
