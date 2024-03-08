package templui

type OptFn func(*Element)

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

func WithAttributes(attributes map[string]any) OptFn {
	return func(element *Element) {
		if element.Attributes == nil {
			element.Attributes = make(map[string]any)
		}

		for k, v := range attributes {
			element.Attributes[k] = v
		}
	}
}
