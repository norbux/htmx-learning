package main

import (
	"fmt"
	"net/http"

	"github.com/norbux/htmx-learning/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.HandleTemplate)
	http.HandleFunc("/clicked", handler.HandleClicked)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
