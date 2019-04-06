package main

import (
	"fmt"
	"net/http"

	"responder/processor"
)

func handler(rw http.ResponseWriter, request *http.Request) {
	var value processor.Export
	value.Process(rw, request)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Hello, world. This is a test!")
}
