# TemplUI

- [Root](/)

## Input date

```go
import "github.com/iamajoe/templui/inputdate"

@inputdate.New(
    inputdate.WithID("zing"),
    inputdate.WithClasses("foo"),
    inputdate.WithAttributes(map[string]any{"data-zed": "zung"}),
    inputdate.WithDisabled(),
    inputdate.WithRequired(),
    inputdate.WithName("start-date"),
    inputdate.WithValue(time.Now()),
    inputdate.WithPlaceholder("Set your start date"),
    inputdate.WithMin(time.Now()),
    inputdate.WithMax(time.Now()),
    func(c *inputdate.InputDate) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
```
