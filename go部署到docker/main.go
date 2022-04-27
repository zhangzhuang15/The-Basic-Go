package main

import (
	"io"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8089", nil)
}
