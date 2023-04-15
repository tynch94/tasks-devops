package main

import (
	"io"
	"net/http"
)

func main() {
	a := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
	}
	http.HandleFunc("/", a)
	http.ListenAndServe(":8080", nil)
}
