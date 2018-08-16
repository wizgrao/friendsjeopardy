package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct {}

func (h* helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello")
}
func main() {
	http.Handle("/", &helloHandler{})
	fmt.Print(http.ListenAndServe(":8080", nil))
}
