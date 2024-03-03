package ui

import (
	"github.com/a-h/templ"
)

type AnchorTarget string

const (
	AnchorTargetSelf   AnchorTarget = "_self"
	AnchorTargetParent AnchorTarget = "_parent"
	AnchorTargetTop    AnchorTarget = "_top"
	AnchorTargetBlank  AnchorTarget = "_blank"
)

type AnchorProps struct {
	Disabled   bool
	ClassNames []string

	Href   string
	Target AnchorTarget
}

type AnchorOptsFn func(*AnchorProps) error

var AnchorOpts = struct {
	WithDisabled func() AnchorOptsFn
	WithClasses  func(classes ...string) AnchorOptsFn
	WithHref     func(href string) AnchorOptsFn
	WithTarget   func(target AnchorTarget) AnchorOptsFn
}{
	WithHref: func(href string) AnchorOptsFn {
		return func(props *AnchorProps) error {
			props.Href = href
			return nil
		}
	},

	WithTarget: func(target AnchorTarget) AnchorOptsFn {
		return func(props *AnchorProps) error {
			props.Target = target
			return nil
		}
	},

	WithDisabled: func() AnchorOptsFn {
		return func(props *AnchorProps) error {
			props.Disabled = true
			return nil
		}
	},

	WithClasses: func(classes ...string) AnchorOptsFn {
		return func(props *AnchorProps) error {
			props.ClassNames = append(props.ClassNames, classes...)
			return nil
		}
	},
}

func Anchor(opts ...AnchorOptsFn) templ.Component {
	b := AnchorProps{}

	for _, opt := range opts {
		if err := opt(&b); err != nil {
			// TODO: need to figure how to handle errors
			return templ.NopComponent
		}
	}

	return anchorTmpl(b)
}
