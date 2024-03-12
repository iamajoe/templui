package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/iamajoe/templui/anchor"
	"github.com/iamajoe/templui/button"
	"github.com/iamajoe/templui/inputdate"
	"github.com/iamajoe/templui/inputnumber"
	"github.com/iamajoe/templui/inputtext"
	"github.com/iamajoe/templui/table"
)

//go:generate go run github.com/a-h/templ/cmd/templ@latest generate

type MenuItem struct {
	Label       string
	Description string
	Items       []MenuItem
	Route       string
	Component   func() templ.Component
}

var rootSideMenu = MenuItem{
	Items: []MenuItem{
		{Label: "Anchor", Route: "/anchor", Component: anchor.Styleguide},
		{Label: "Button", Route: "/button", Component: button.Styleguide},
		{Label: "Input date", Route: "/inputdate", Component: inputdate.Styleguide},
		{Label: "Input number", Route: "/inputnumber", Component: inputnumber.Styleguide},
		{Label: "Input text", Route: "/inputtext", Component: inputtext.Styleguide},
		{Label: "Table", Route: "/table", Component: table.Styleguide},
	},
}

func serveItem(item MenuItem) func(http.ResponseWriter, *http.Request) {
	contentType := "text/html; charset=utf-8"
	itemComponent := item.Component()

	return func(w http.ResponseWriter, r *http.Request) {
		activeUrlRaw := r.URL
		activeUrl := ""
		if activeUrlRaw != nil {
			activeUrl = activeUrlRaw.String()
		}

		component := renderItem(item.Label, item.Description, activeUrl, itemComponent)
		err := component.Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", contentType)
	}
}

func setHandleOnItems(item MenuItem, parentRoute string) {
	route := parentRoute + item.Route

	if item.Component != nil {
		http.HandleFunc(route, serveItem(item))
	}

	for _, child := range item.Items {
		setHandleOnItems(child, route)
	}
}

func main() {
	http.Handle("/", templ.Handler(renderRoot()))
	setHandleOnItems(rootSideMenu, "")

	fmt.Println("Listening on :1234")
	if err := http.ListenAndServe(":1234", nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
