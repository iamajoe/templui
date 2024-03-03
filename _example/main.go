package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

//go:generate go run github.com/a-h/templ/cmd/templ@latest generate

func main() {
	http.Handle("/", templ.Handler(web()))

	fmt.Println("Listening on :1234")
	if err := http.ListenAndServe(":1234", nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
