/*
This is for the video: https://www.youtube.com/watch?v=Quman5pH2aY
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	router.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)

			// This for not continue the function
			return
		}
		fmt.Fprintf(w, "Hello, World! id: %d", id)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server is running on port:", server.Addr)

	///This for not close the server
	server.ListenAndServe()
}

// Codigo recomendado por CURSOR, se ve algo mejor.
/*
router := http.NewServeMux()

	router.HandleFunc("/notas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	router.HandleFunc("/notas/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Extraer el ID de la URL
		id := r.URL.Path[len("/notas/"):]
		fmt.Fprintf(w, "Hello, World! %s", id)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server is running on port:", server.Addr)

	///This for not close the server
	server.ListenAndServe()
*/
