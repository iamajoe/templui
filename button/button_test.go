package button

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestButton(t *testing.T) {
	comp := New().
		WithKind(KindSubmit).
		WithDisabled().
		WithClasses("foo").
		WithCustomFn(func(p *Button) error {
			p.ClassNames = append(p.ClassNames, "bar")
			return nil
		})

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

	target := `button.foo.bar[type="submit"][disabled]`
	if doc.Find(target).Length() == 0 {
		t.Fatalf("expected button to render with attributes: %s", docHTML)
	}
}
