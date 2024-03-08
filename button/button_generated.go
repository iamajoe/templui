package button

// NOTE: Don't change this file. It is auto generated! You will lose any change

type OptFn func(*Button)

func WithID(id string) OptFn {
	return func(element *Button) {
		element.SetID(id)
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *Button) {
		element.SetClasses(classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *Button) {
		element.SetAttributes(attributes)
	}
}
