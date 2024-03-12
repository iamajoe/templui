package theme

import (
	"fmt"

	"github.com/a-h/templ"
)

type ButtonProps struct {
	Variant Variant
	Size    Size
	Role    Role
}

// TODO: what about light scheme??

func Button(props ButtonProps) templ.CSSClass {
	id := fmt.Sprintf("templui_btn_%d_%d_%d", props.Variant, props.Size, props.Role)

	// TODO: each variant, each size, each role should have its differences

	background := "var(--templui-color-dark-default)"
	textColor := "var(--templui-color-dark-inverse)"

	switch props.Role {
	case RolePrimary:
		background = "var(--templui-color-dark-primary)"
		textColor = "var(--templui-color-dark-default)"
	case RoleSecondary:
		background = "var(--templui-color-dark-secondary)"
		textColor = "var(--templui-color-dark-default)"
	case RoleInfo:
		background = "var(--templui-color-dark-info)"
		textColor = "var(--templui-color-dark-default)"
	case RoleWarning:
		background = "var(--templui-color-dark-warning)"
		textColor = "var(--templui-color-dark-default)"
	case RoleDanger:
		background = "var(--templui-color-dark-error)"
		textColor = "var(--templui-color-dark-default)"
	case RoleSuccess:
		background = "var(--templui-color-dark-success)"
		textColor = "var(--templui-color-dark-default)"
	}

	switch props.Variant {
	case VariantBorder:
		if props.Role == RoleDefault {
			background = "var(--templui-color-dark-inverse)"
		}
		textColor = background
		background = "transparent"
		// VariantTransparent
		// VariantLink
	}

	padding := "8px 16px"
	radius := "var(--templui-radius)"
	border := fmt.Sprintf("1px solid %s", background)

	rules := fmt.Sprintf(`
    padding: %s;
    background: %s;
    color: %s;
    border: %s;
    border-radius: %s
  `,
		padding,
		background,
		textColor,
		border,
		radius,
	)

	return templ.ComponentCSSClass{
		ID: id,
		Class: templ.SafeCSS(fmt.Sprintf(`
.%s { %s }
  `, id, rules)),
	}
}
