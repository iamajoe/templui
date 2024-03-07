package templui

import (
	"github.com/a-h/templ"
)

type Element struct {
	ID         string
	ClassNames []string
	Attributes templ.Attributes
}

type OptFn func(*Element)

func (OptFn) AnchorOpt()      {}
func (OptFn) ButtonOpt()      {}
func (OptFn) InputDateOpt()   {}
func (OptFn) InputNumberOpt() {}

func WithID(id string) OptFn {
	return func(element *Element) {
		element.ID = id
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *Element) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]interface{}) OptFn {
	return func(element *Element) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]interface{})
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}
