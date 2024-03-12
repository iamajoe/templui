package templui

import "github.com/a-h/templ"

type Element struct {
	ID         string
	ClassNames []string
	Attributes templ.Attributes

	CSS templ.CSSClass
}
