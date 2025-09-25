package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greeting)

	http.ListenAndServe(":8080", nil)
}

func greeting(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"Hello World!")
	}