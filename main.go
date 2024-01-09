package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Go")

	// Set response status code
	w.WriteHeader(http.StatusOK)

	// Your existing request handling logic
	switch r.URL.Path {
	case "/":
		if r.Method == http.MethodGet {
			fmt.Fprint(w, `{"message": "Ini adalah homepage"}`)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"message": "Halaman tidak dapat diakses dengan %s request"}`, r.Method)
		}
	case "/about":
		if r.Method == http.MethodGet {
			fmt.Fprint(w, `{"message": "Halo! Ini adalah halaman about"}`)
		} else if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, `{"error": "Error reading request body"}`)
				return
			}

			var data map[string]string
			if err := json.Unmarshal(body, &data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, `{"error": "Error decoding JSON"}`)
				return
			}

			name, ok := data["name"]
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, `{"error": "Missing 'name' field in JSON"}`)
				return
			}

			fmt.Fprintf(w, `{"message": "Halo, %s! Ini adalah halaman about"}`, name)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"message": "Halaman tidak dapat diakses menggunakan %s request"}`, r.Method)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"message": "Halaman tidak ditemukan!"}`)
	}
}

func main() {
	port := 5000
	host := "localhost"

	http.HandleFunc("/", handler)

	fmt.Printf("Server berjalan pada http://%s:%d\n", host, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}
