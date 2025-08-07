package main

import (
	"fmt"
	"net/http"

	"github.com/newton-miku/Goink"
)

func main() {
	engine := Goink.New()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	engine.Run(":9999")
}
