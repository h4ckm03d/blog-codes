package main

import (
	"fmt"
	"net/http"
)

// OmTeloletOm middleware
func OmTeloletOm(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("x-telolet", "OM telolet OM pake middleware")
		handler(w, r)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", OmTeloletOm(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}
