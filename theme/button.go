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
	fontSize = "1rem"
	padding = "8px 16px"
	switch props.Size {
	case SizeXs:
		fontSize = "0.8rem"
		padding = "2px 6px"
	case SizeS:
		fontSize = "0.9rem"
		padding = "4px 10px"
	case SizeLg:
		fontSize = "1.2rem"
		padding = "12px 24px"
	}

	return fontSize, padding
}

func Button(props ButtonProps) templ.CSSClass {
	id := fmt.Sprintf("templui_btn_%d_%d_%d", props.Variant, props.Size, props.Role)

	background, textColor := buttonGetColors(props)
	fontSize, padding := buttonGetColors(props)

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
      border-radius: var(--templui-radius);
    }`, id, fontSize, padding, textColor, textColor)
	default:
		class = fmt.Sprintf(`
    .%s {
      font-size: %s;
      padding: %s;
      background: %s;
      color: %s;
      border: 1px solid %s;
      border-radius: var(--templui-radius);
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
