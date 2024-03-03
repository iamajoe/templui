# TemplUI

- [Root](/)

## Anchor

```go
import "github.com/iamajoe/templui/anchor"

@anchor.New(
    WithID("zing"),
    WithClasses("foo"),
    WithAttributes(map[string]any{"data-zed": "zung"}),
    WithTarget(TargetBlank),
    WithHref("https://google.com"),
    func(c *Anchor) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
```

