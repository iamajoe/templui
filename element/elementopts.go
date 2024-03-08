package element

type OptFn func(*Element)

func WithID(id string) OptFn {
	return func(element *Element) {
		element.SetID(id)
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *Element) {
		element.SetClasses(classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *Element) {
		element.SetAttributes(attributes)
	}
}
