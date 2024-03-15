# TemplUI

- [Root](/)

## Button

```go
import "github.com/iamajoe/templui/button"

css exampleCss() {
  border: 1px solid red;
}

@button.New(
    button.WithID("zing"),
    button.WithClasses("foo"),
    button.WithAttributes(map[string]any{"data-zed": "zung"}),
    button.WithKind(button.KindSubmit),
    button.WithCSS(exampleCss()),
    button.WithDisabled(),
    func(c *button.Button) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
) { Button }
```
