# TemplUI

This library aims to provide common web ui elements a functional and styling layer over [templ](https://github.com/a-h/templ).

## Install

```sh
go install github.com/iamajoe/templui
```

## Documentation

### Anchor

```go
@ui.Anchor(
    ui.AnchorOpts.WithTarget(ui.AnchorTargetBlank),
    ui.AnchorOpts.WithHref("https://google.com"),
    ui.AnchorOpts.WithDisabled(),
    ui.AnchorOpts.WithClasses("foo"),
)
```

### Button

```go
import "github.com/iamajoe/templui/button"

@button.New()
    WithID("zing").
    WithClasses("foo").
    WithAttributes(map[string]any{"data-zed": "zung"}).
    WithKind(button.KindSubmit).
    WithDisabled().
    WithFunction(func(el *button.Button) error {
        el.ClassNames = append(el.ClassNames, "bar")
        return nil
    })
```

## Development

### Build

Whenever a `.templ` file is changed, a build needs to happen so it can generate the `.go` file.

```sh
make build
```

### Test

```sh
make test
```

## TODO

- [ ]: Radio
- [ ]: Checkbox
- [ ]: Toggle
- [ ]: Select
- [ ]: Input
- [ ]: Form group
- [ ]: Alert
- [ ]: Badge
- [ ]: Card
- [ ]: Dialog
- [ ]: Progress
- [ ]: Menubar
- [ ]: Avatar
- [ ]: Dropdown menu
- [ ]: Tooltip
- [ ]: Pagination
- [ ]: Separator
- [ ]: Slider
- [ ]: Table
- [ ]: Themes
