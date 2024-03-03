# TemplUI

- [Root](/)

## Button

```go
import "github.com/iamajoe/templui/button"

@button.New().
    WithID("zing"),
    WithClasses("foo"),
    WithAttributes(map[string]any{"data-zed": "zung"}),
    WithKind(button.KindSubmit),
    WithDisabled(),
    func(c *button.Button) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
```


