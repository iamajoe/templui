package ui

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestAnchor(t *testing.T) {
	comp := Anchor(
		AnchorOpts.WithTarget(AnchorTargetBlank),
		AnchorOpts.WithHref("https://google.com"),
		AnchorOpts.WithDisabled(),
		AnchorOpts.WithClasses("foo"),
		func(p *AnchorProps) error {
			p.ClassNames = append(p.ClassNames, "bar")
			return nil
		},
	)

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

	target := `a.foo.bar[href="https://google.com"][target="_blank"][disabled]`
	if doc.Find(target).Length() == 0 {
		t.Fatalf("expected anchor to render with attributes: %s", docHTML)
	}
}
