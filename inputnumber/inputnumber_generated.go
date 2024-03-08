package inputnumber

// NOTE: Don't change this file. It is auto generated! You will lose any change

type OptFn func(*InputNumber)

func WithID(id string) OptFn {
	return func(element *InputNumber) {
		element.SetID(id)
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *InputNumber) {
		element.SetClasses(classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *InputNumber) {
		element.SetAttributes(attributes)
	}
}
