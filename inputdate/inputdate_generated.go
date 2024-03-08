package inputdate

// NOTE: Don't change this file. It is auto generated! You will lose any change

type OptFn func(*InputDate)

func WithID(id string) OptFn {
	return func(element *InputDate) {
		element.SetID(id)
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *InputDate) {
		element.SetClasses(classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *InputDate) {
		element.SetAttributes(attributes)
	}
}
