package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>vascooooooooo</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running in http://localhost:8081...")
	http.ListenAndServe(":8081", nil)
}
