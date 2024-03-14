package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Criando uma API REST com Go 1.22")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", func(w http.ResponseWriter, e *http.Request) {
		w.Write([]byte("books"))
	})

	mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		bookID := r.PathValue("id")
		w.Write([]byte("books - " + bookID))
	})

	http.ListenAndServe(":8080", mux)

}
