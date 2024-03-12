package formelement

import "github.com/iamajoe/templui"

type Element struct {
	templui.Element

	Disabled    bool
	Required    bool
	Name        string
	Placeholder string
}
