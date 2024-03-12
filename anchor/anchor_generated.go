package anchor

// NOTE: Don't change this file. It is auto generated! You will lose any change

import "github.com/a-h/templ"

type OptFn func(*Anchor)

func WithID(id string) OptFn {
	return func(element *Anchor) {
		element.ID = id
	}
}

func WithCSS(css templ.CSSClass) OptFn {
	return func(element *Anchor) {
		element.CSS = css
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *Anchor) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *Anchor) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}
