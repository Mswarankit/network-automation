package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Docker world")
	})
	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Docker blogging with my new site")
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
