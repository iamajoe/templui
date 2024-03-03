# TemplUI

This library aims to provide common web ui elements a functional and styling layer over [templ](https://github.com/a-h/templ).

## Install

```sh
go install github.com/iamajoe/templui
```

## Documentation

The components are built under `builder pattern`. As such, You construct the elements as you see fit without accessing the properties directly.

You can either use chaining:
```go
@button.New().WithID("zing")
```
Or spread the options under the arguments
```go
@button.New(WithID("zing"))
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
import "github.com/iamajoe/templui/button"

@button.New().
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

Whenever a `.templ` file is changed, the `.go` file according to that component needs to be generated. You can do so with:
```sh
make build
```

### Test

```sh
make test
```

## TODO

- [ ] Radio
- [ ] Checkbox
- [ ] Toggle
- [ ] Select
- [ ] Input
- [ ] Form group
- [ ] Alert
- [ ] Badge
- [ ] Card
- [ ] Dialog
- [ ] Progress
- [ ] Menubar
- [ ] Avatar
- [ ] Dropdown menu
- [ ] Tooltip
- [ ] Pagination
- [ ] Separator
- [ ] Slider
- [ ] Table
- [ ] Themes
- [ ] Styleguide
