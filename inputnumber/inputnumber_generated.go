package inputnumber

// NOTE: Don't change this file. It is auto generated! You will lose any change

type OptFn func(*InputNumber)

func WithID(id string) OptFn {
	return func(element *InputNumber) {
		element.ID = id
	}
}

func WithClasses(classes ...string) OptFn {
	return func(element *InputNumber) {
		element.ClassNames = append(element.ClassNames, classes...)
	}
}

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *InputNumber) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}
