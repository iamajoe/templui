package ui

import (
	"github.com/a-h/templ"
)

type ButtonType string

const (
	ButtonTypeSubmit ButtonType = "submit"
	ButtonTypeButton ButtonType = "button"
)

type ButtonProps struct {
	Disabled   bool
	ClassNames []string
	Type       ButtonType
}

type ButtonOptsFn func(*ButtonProps) error

var ButtonOpts = struct {
	WithDisabled func() ButtonOptsFn
	WithClasses  func(classes ...string) ButtonOptsFn
	WithType     func(kind ButtonType) ButtonOptsFn
}{
	WithType: func(kind ButtonType) ButtonOptsFn {
		return func(props *ButtonProps) error {
			props.Type = kind
			return nil
		}
	},

	WithDisabled: func() ButtonOptsFn {
		return func(props *ButtonProps) error {
			props.Disabled = true
			return nil
		}
	},

	WithClasses: func(classes ...string) ButtonOptsFn {
		return func(props *ButtonProps) error {
			props.ClassNames = append(props.ClassNames, classes...)
			return nil
		}
	},
}

func Button(opts ...ButtonOptsFn) templ.Component {
	b := ButtonProps{}

	for _, opt := range opts {
		if err := opt(&b); err != nil {
			// TODO: need to figure how to handle errors
			return templ.NopComponent
		}
	}

	return buttonTmpl(b)
}
