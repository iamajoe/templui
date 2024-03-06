# TemplUI

- [Root](/)

## Input date

```go
import "github.com/iamajoe/templui/inputDate"

@inputDate.New(
    inputDate.WithID("zing"),
    inputDate.WithClasses("foo"),
    inputDate.WithAttributes(map[string]any{"data-zed": "zung"}),
    inputDate.WithDisabled(),
    inputDate.WithRequired(),
    inputDate.WithName("start-date"),
    inputDate.WithValue(time.Now()),
    inputDate.WithPlaceholder("Set your start date"),
    inputDate.WithMin(time.Now()),
    inputDate.WithMax(time.Now()),
    func(c *inputDate.InputDate) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
```
