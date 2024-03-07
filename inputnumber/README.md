# TemplUI

- [Root](/)

## Input number

```go
import "github.com/iamajoe/templui/inputnumber"

@inputnumber.New(
    inputnumber.WithID("zing"),
    inputnumber.WithClasses("foo"),
    inputnumber.WithAttributes(map[string]any{"data-zed": "zung"}),
    inputnumber.WithDisabled(),
    inputnumber.WithRequired(),
    inputnumber.WithName("curr-salary"),
    inputnumber.WithValue(500),
    inputnumber.WithPlaceholder("Set your salary"),
    inputnumber.WithMin(100),
    inputnumber.WithMax(8000),
    inputnumber.WithStep(10),
    OptFn(func(c *inputnumber.InputNumber) {
        c.ClassNames = append(c.ClassNames, "bar")
    }),
)
```


