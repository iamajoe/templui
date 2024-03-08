package element

import "github.com/a-h/templ"

//go:generate go run generate/element.go

type Element struct {
	ID         string
	ClassNames []string
	Attributes templ.Attributes
}

func (element *Element) SetID(id string) {
	element.ID = id
}

func (element *Element) SetClasses(classes ...string) {
	element.ClassNames = append(element.ClassNames, classes...)
}

func (element *Element) SetAttributes(attributes map[string]any) {
	if element.Attributes == nil {
		element.Attributes = make(map[string]any)
	}

	for k, v := range attributes {
		element.Attributes[k] = v
	}
}
