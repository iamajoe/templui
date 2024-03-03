# TemplUI

- [Root](/)

## Button

```go
import "github.com/iamajoe/templui/button"

@button.New().
    WithID("zing").
    WithClasses("foo").
    WithAttributes(map[string]any{"data-zed": "zung"}).
    WithKind(button.KindSubmit).
    WithDisabled().
    WithFunction(func(c *button.Button) error {
        c.ClassNames = append(c.ClassNames, "bar")
        return nil
    })
```


