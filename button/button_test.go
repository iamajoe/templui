package button

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name         string
		comp         func() Button
		assertTarget string
	}{
		{
			name: "creates",
			comp: func() Button {
				return New(
					WithID("zing"),
					WithClasses("foo"),
					WithAttributes(map[string]any{"data-zed": "zung"}),
					WithKind(KindSubmit),
					WithDisabled(),
					OptFn(func(el *Button) {
						el.ClassNames = append(el.ClassNames, "bar")
					}),
				)
			},
			assertTarget: `button#zing.foo.bar[type="submit"][disabled][data-zed="zung"]`,
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
