package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/iamajoe/templui/button"
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
		{Label: "Anchor", Route: "/anchor"},
		{Label: "Button", Route: "/button", Component: button.Styleguide},
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
