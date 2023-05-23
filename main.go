package main

import (
	"fmt"
	"net/http"

	"github.com/ChrisLo0751/gee/pkg/engine"
)

func main() {
	e := engine.New()
	e.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})
	e.Run(":9090")
}
