# TemplUI

- [Root](/)

## Button

```go
import "github.com/iamajoe/templui/button"

@button.New(
    button.WithID("zing"),
    button.WithClasses("foo"),
    button.WithAttributes(map[string]any{"data-zed": "zung"}),
    button.WithKind(button.KindSubmit),
    button.WithDisabled(),
    func(c *button.Button) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
```


