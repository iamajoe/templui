# TemplUI

- [Root](/)

## Input text

```go
import "github.com/iamajoe/templui/inputText"

@inputtext.New(
    inputtext.WithID("zing"),
    inputtext.WithClasses("foo"),
    inputtext.WithAttributes(map[string]any{"data-zed": "zung"}),
    inputtext.WithDisabled(),
    inputtext.WithRequired(),
    inputtext.WithName("curr-salary"),
    inputtext.WithValue(500),
    inputtext.WithPlaceholder("Set your salary"),
    func(c *inputtext.InputText) {
        c.ClassNames = append(c.ClassNames, "bar")
    },
)
```


