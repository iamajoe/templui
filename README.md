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

### Extending

You can easily extend the component to be able to use them throughout your project by wrapping them in a function.

```go
func ProjectButton(id string, color string) templ.Component {
    return button.New().
        WithClasses("bg-" + color).
        WithAttributes(map[string]any{"data-id": id}).
        WithKind(button.KindSubmit)
}

@ProjectButton("indigo-500")
```

## Components

- [Anchor](./anchor)
- [Button](./button)

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
