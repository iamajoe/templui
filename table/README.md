# TemplUI

- [Root](/)

## Table

```go
import (
    "github.com/iamajoe/templui/table"
    "github.com/iamajoe/templui/table/tablerow"
    "github.com/iamajoe/templui/table/tablecell"
)

templ header() {
  @tablerow.New(
    tablerow.WithClasses("bg-slate-50 border-b border-slate-800"),
  ) {
    @tablecell.New(
      tablecell.WithClasses("p-2 text-center"),
      tablecell.WithIsHeader(),
    ) { ID }
  }
}

@table.New(
  table.WithID("zing"),
  table.WithClasses("border border-slate-200"),
  table.WithAttributes(map[string]any{"data-zed": "zung"}),
  table.WithHeader(header()),
) {
  @tablerow.New(
    tablerow.WithClasses("border-b border-slate-800"),
  ) {
    @tablecell.New(
      tablecell.WithClasses("p-2 text-center"),
    ) {
      <span>1</span>
    }
  }
}
```


