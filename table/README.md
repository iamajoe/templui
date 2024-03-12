# TemplUI

- [Root](/)

## Table

```go
import (
    "github.com/iamajoe/templui/table"
)

@table.New(
  table.WithID("zing"),
  table.WithClasses("border border-slate-200"),
  table.WithAttributes(map[string]any{"data-zed": "zung"}),
  table.WithHeader("ID", "Status"),
  table.WithHeaderClasses("bg-slate-50 border-b border-slate-800"),
  table.WithRows([]string{"1", "<span class=\"text-red-500\">Rejected</span>"}),
  table.WithRowClasses("border-b border-slate-800"),
) 
```


