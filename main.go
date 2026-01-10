package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "eaea",
		Message: "Oi! templates",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	count := 0

	http.HandleFunc("/", handleTemplate)

	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		count = count + 1
		responseMessage := fmt.Sprintf("Clicked: %d time(s)", count)
		w.Write([]byte(responseMessage))
	})

	// Handle GET request
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "Hello, World!")
	})

	// Handle POST request
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		w.Header().Set("Content-Type", "text/plain")
		w.Write(body)
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
