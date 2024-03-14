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

func buttonGetColors(props ButtonProps) (background string, textColor string) {
	background = "var(--templui-color-default)"
	textColor = "var(--templui-color-inverse)"

	switch props.Role {
	case RolePrimary:
		background = "var(--templui-color-primary)"
		textColor = "var(--templui-color-default)"
	case RoleSecondary:
		background = "var(--templui-color-secondary)"
		textColor = "var(--templui-color-default)"
	case RoleInfo:
		background = "var(--templui-color-info)"
		textColor = "var(--templui-color-default)"
	case RoleWarning:
		background = "var(--templui-color-warning)"
		textColor = "var(--templui-color-default)"
	case RoleDanger:
		background = "var(--templui-color-error)"
		textColor = "var(--templui-color-default)"
	case RoleSuccess:
		background = "var(--templui-color-success)"
		textColor = "var(--templui-color-default)"
	}

	return background, textColor
}

func buttonGetSizes(props ButtonProps) (fontSize string, padding string) {
	fontSize = "var(--templui-font-size)"
	padding = "var(--templui-btn-padding)"

	switch props.Size {
	case SizeXs:
		fontSize = "var(--templui-font-size-xs)"
		padding = "var(--templui-btn-padding-xs)"
	case SizeS:
		fontSize = "var(--templui-font-size-s)"
		padding = "var(--templui-btn-padding-s)"
	case SizeLg:
		fontSize = "var(--templui-font-size-lg)"
		padding = "var(--templui-btn-padding-lg)"
	}

	return fontSize, padding
}

func Button(props ButtonProps) templ.CSSClass {
	id := fmt.Sprintf("templui_btn_%d_%d_%d", props.Variant, props.Size, props.Role)

	background, textColor := buttonGetColors(props)
	fontSize, padding := buttonGetSizes(props)

	class := ""

	// TODO: consider using css nesting when fully supported
	// TODO: what could we do to simplify this?!

	switch props.Variant {
	case VariantTransparent:
		textColor = background
		class = fmt.Sprintf(`
    .%s {
      font-size: %s;
      padding: 0;
      background: transparent;
      color: %s;
    }`, id, fontSize, textColor)
	case VariantLink:
		textColor = background
		class = fmt.Sprintf(`
    .%s {
      font-size: %s;
      padding: 0;
      background: transparent;
      color: %s;
      text-decoration: underline;
    }`, id, fontSize, textColor)
	case VariantBorder:
		textColor = background
		class = fmt.Sprintf(`
    .%s {
      font-size: %s;
      padding: %s;
      background: transparent;
      color: %s;
      border: 1px solid %s;
      border-radius: var(--templui-btn-radius);
    }`, id, fontSize, padding, textColor, textColor)
	default:
		class = fmt.Sprintf(`
    .%s {
      font-size: %s;
      padding: %s;
      background: %s;
      color: %s;
      border: var(--templui-btn-border) solid %s;
      border-radius: var(--templui-btn-radius);
    }`, id, fontSize, padding, background, textColor, background)
	}

	if props.Role == RoleDefault && props.Variant != VariantDefault {
		class += fmt.Sprintf(`
    :root.dark .%s {
      color: var(--templui-color-inverse);
      border-color: var(--templui-color-inverse);
    }

    @media (prefers-color-scheme: dark)
      :root .%s {
        color: var(--templui-color-inverse);
        border-color: var(--templui-color-inverse);
      }
    }
    `, id, id)
	}

	return templ.ComponentCSSClass{
		ID:    id,
		Class: templ.SafeCSS(class),
	}
}
