package main

import (
	"fmt"
	"net/http"

	"github.com/joshbetz/bodyparser"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var name string

		if len(r.FormValue("name")) > 0 {
			name = r.FormValue("name")
		} else {
			name = "World"
		}

		body := fmt.Sprintf("Hello %s!", name)
		w.Write([]byte(body))
	})

	handler := bodyparser.Middleware(h)
	http.ListenAndServe(":3000", handler)
}
