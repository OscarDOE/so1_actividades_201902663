package main

import (
	"fmt"
	"net/http"
)

type datos struct {
	lnum   float32
	rnum   float32
	result float32
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {})
	fmt.Println("Hello World")
	http.ListenAndServe("localhost:5000", nil)
}
