package selectbox

// NOTE: Don't change this file. It is auto generated! You will lose any change

import "github.com/a-h/templ"

type OptFn func(*SelectBox)

func WithDisabled() OptFn {
	return func(element *SelectBox) {
		element.Disabled = true
	}
}

func WithRequired() OptFn {
	return func(element *SelectBox) {
		element.Required = true
	}
}

func WithPlaceholder(placeholder string) OptFn {
	return func(element *SelectBox) {
		element.Placeholder = placeholder
	}
}

func WithName(name string) OptFn {
	return func(element *SelectBox) {
		element.Name = name
	}
}

///////////////////////////////
// NOTE: copied from ../element.go

func WithID(id string) OptFn {
	return func(element *SelectBox) {
		element.ID = id
	}
}

func WithCSS(css templ.CSSClass) OptFn {
	return func(element *SelectBox) {
		element.CSS = css
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *SelectBox) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *SelectBox) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}
