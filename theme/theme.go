package theme

import (
	"fmt"

	"github.com/a-h/templ"
)

type Variant int

const (
	VariantDefault Variant = iota
	VariantBorder
	VariantTransparent
	VariantLink
)

type Size int

const (
	SizeNormal Size = iota
	SizeTiny
	SizeXs
	SizeS
	SizeLg
)

type Role int

const (
	RoleDefault Role = iota
	RolePrimary
	RoleSecondary
	RoleInfo
	RoleWarning
	RoleDanger
	RoleSuccess
)

func Synthwave() templ.CSSClass {
	tmpl := fmt.Sprintf(`
:root {
    color-scheme: light;

    --templui-color-inverse: oklch(0.979365 0.00819 301.358);
    --templui-color-default: oklab(0.168898 0.019427 -0.0603791);
    --templui-color-neutral: oklab(0.229999 0.0264779 -0.0893423);
    --templui-color-primary: oklch(0.722105 0.159514 342.009);
    --templui-color-secondary: oklch(0.782714 0.118101 227.382);
    --templui-color-info: oklch(0.765197 0.12273 231.832);
    --templui-color-success: oklch(0.860572 0.115038 178.625);
    --templui-color-warning: oklch(0.85531 0.122117 93.7222);
    --templui-color-error: oklch(0.737005 0.121339 32.6393);

    --templui-radius: 8px;
}

:root.dark {
    color-scheme: dark;

    --templui-color-default: var(--templui-color-neutral);
}

@media (prefers-color-scheme: dark)
  :root {
      color-scheme: dark;

      --templui-color-default: var(--templui-color-neutral);
  }
}
  `)

	return templ.ComponentCSSClass{
		ID:    "theme-synthwave",
		Class: templ.SafeCSS(tmpl),
	}
}
