package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/thongthele/romanserver/numerals"
	"golang.org/x/net/html"
)

func main() {
	http.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")
		if paths[3] == "roman-number" {
			number, _ := strconv.Atoi(paths[4])
			if number > 0 && number <= 10 {
				fmt.Fprintf(w, "%q", html.EscapeString(numerals.Numerals[number]))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
