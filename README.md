# TemplUI

This library aims to provide common web ui elements a functional and styling layer over [templ](https://github.com/a-h/templ).

## Install

```sh
go install github.com/iamajoe/templui
```

## Documentation

Components are nested under the `ui` package:
```go
import "github.com/iamajoe/templui/ui"
```

These components are built under `builder pattern`. As such, You construct the elements as you see fit without accessing the properties directly.
You can still construct the properties manually by providing the right custom function of the builder pattern:
```go
@ui.Anchor(
    func(p *ui.AnchorProps) error {
        p.ClassNames = append(p.ClassNames, "bar")
        return nil
    },
)
```

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
@ui.Button(
    ui.ButtonOpts.WithType(ui.ButtonTypeSubmit),
    ui.ButtonOpts.WithDisabled(),
    ui.ButtonOpts.WithClasses("foo"),
)
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
