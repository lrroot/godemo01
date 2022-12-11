package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
