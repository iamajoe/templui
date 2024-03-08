package anchor

// NOTE: Don't change this file. It is auto generated! You will lose any change

type OptFn func(*Anchor)

func WithID(id string) OptFn {
	return func(element *Anchor) {
		element.SetID(id)
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *Anchor) {
		element.SetClasses(classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *Anchor) {
		element.SetAttributes(attributes)
	}
}
