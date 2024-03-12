# TemplUI

- [Root](/)

## SelectBox

```go
import "github.com/iamajoe/templui/selectbox"

@selectbox.New(
  selectbox.WithID("zing"),
  selectbox.WithClasses("border border-slate-200"),
  selectbox.WithAttributes(map[string]any{"data-zed": "zung"}),
  // selectbox.WithDisabled(),
  selectbox.WithItems([]selectbox.SelectItem{
    {Value: "1", Label: "1", Selected: true},
    {Value: "2", Label: "2"},
  }),
  selectbox.WithValue("2"),
  func(c *selectbox.SelectBox) {
    c.ClassNames = append(c.ClassNames, "bar")
  },
)
```

