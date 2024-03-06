# TemplUI

- [Root](/)

## Input number

```go
import "github.com/iamajoe/templui/inputNumber"

@inputNumber.New(
    inputNumber.WithID("zing"),
    inputNumber.WithClasses("foo"),
    inputNumber.WithAttributes(map[string]any{"data-zed": "zung"}),
    inputNumber.WithDisabled(),
    inputNumber.WithRequired(),
    inputNumber.WithName("curr-salary"),
    inputNumber.WithValue(500),
    inputNumber.WithPlaceholder("Set your salary"),
    inputNumber.WithMin(100),
    inputNumber.WithMax(8000),
    inputNumber.WithStep(10),
    func(c *inputNumber.InputNumber) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
```


